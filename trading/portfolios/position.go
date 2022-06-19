package portfolios

//nolint:gofumpt
import (
	"math"
	"sync"
	"time"

	"mbg/trading/currencies"
	"mbg/trading/data"
	"mbg/trading/instruments"
	"mbg/trading/orders/sides"
	pos "mbg/trading/portfolios/positions/sides"
	"mbg/trading/portfolios/roundtrips/matchings"
)

// Position is a portfolio position.
type Position struct {
	mu                sync.RWMutex
	roundtripMatching matchings.Matching
	instrument        instruments.Instrument
	entryAmount       float64
	debt              float64
	margin            float64
	price             float64
	priceFactor       float64
	quantityBought    float64
	quantitySold      float64
	quantitySoldShort float64
	quantity          float64
	quantitySigned    float64
	side              pos.Side
	cashFlow          float64
	amounts           data.ScalarTimeSeries
	executions        []*Execution
	perf              *Performance
}

// newPosition creates a new position in a given instrument.
func newPosition(instr instruments.Instrument, ex *Execution, account *Account, matching matchings.Matching) *Position {
	p := Position{
		roundtripMatching: matching,
		instrument:        instr,
		priceFactor:       1,
		executions:        make([]*Execution, 0),
		amounts:           data.ScalarTimeSeries{},
		perf:              newPerformance(),
	}

	if instr.PriceFactor() != 0 {
		p.priceFactor = instr.PriceFactor()
	}

	p.init(ex, account)

	return &p
}

// add adds an execution to the existing position.
// This should only be called on position created by newPosition.
// Instrument of the execution should match position instrument.
func (p *Position) add(ex *Execution, account *Account) []*Roundtrip {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.quantity == 0 {
		p.init(ex, account)

		return make([]*Roundtrip, 0)
	}

	qtySigned := p.quantitySigned
	rts := p.updateExecutionPnLAndMatchRoundtrips(ex, qtySigned)
	p.updateSideAndQuantities(ex, qtySigned)
	p.updateMarginAndDebt(ex)
	p.executions = append(p.executions, ex)
	p.cashFlow += ex.cashFlow - ex.commissionConverted

	for _, rt := range rts {
		p.perf.addRoundtrip(*rt)
	}

	account.addExecution(ex)
	p.updatePrice(ex.reportTime, ex.price)

	return rts
}

// init re-initializes the closed position.
func (p *Position) init(ex *Execution, account *Account) {
	p.margin = ex.margin
	p.debt = ex.debt
	p.price = ex.price
	p.entryAmount = ex.amount
	p.quantity = ex.quantity

	for _, e := range p.executions {
		e.unrealizedQuantity = 0
	}

	p.updateSideAndQuantities(ex, 0)
	p.executions = append(p.executions, ex)

	cf := ex.cashFlow - ex.commissionConverted
	p.cashFlow = cf
	t := ex.reportTime
	p.amounts.Add(t, ex.amount)
	p.perf.addPnL(t, ex.amount, ex.amount, ex.amount, cf)
	p.perf.addDrawdown(t, ex.amount+cf)
	account.addExecution(ex)
}

// updateSideAndQuantities updates p.side, p.quantity, p.quantitySigned,
// p.quantityBought, p.quantitySold and p.quantitySoldShort
// for the given execution and the given current signed quantity.
func (p *Position) updateSideAndQuantities(ex *Execution, qtySigned float64) {
	switch ex.side {
	case sides.Buy, sides.BuyMinus:
		p.quantityBought += ex.quantity
	case sides.Sell, sides.SellPlus:
		p.quantitySold += ex.quantity
	case sides.SellShort, sides.SellShortExempt:
		p.quantitySoldShort += ex.quantity
	}

	qtySigned += ex.quantitySign * ex.quantity
	p.quantitySigned = qtySigned
	p.quantity = math.Abs(qtySigned)

	if qtySigned < 0 {
		p.side = pos.Short
	} else {
		p.side = pos.Long
	}
}

// updateMarginAndDebt updates p.margin and p.debt and returns
// a signed increment in debt caused by this execution.
// Uses p.side and p.quantity updated by p.updateSideAndQuantities().
func (p *Position) updateMarginAndDebt(ex *Execution) float64 {
	if ex.margin == 0 {
		return 0
	}

	isLong := p.side == pos.Long
	isShort := p.side == pos.Short
	isBuy := ex.side.IsBuy()
	isSell := ex.side.IsSell()

	switch {
	case (isLong && isBuy) || (isShort && isSell):
		// Execution and updated position have the same directions.
		// Long position and buy execution or short position and sell execution.
		p.margin += ex.margin
		p.debt += ex.debt

		return ex.debt
	case (isLong && isSell) || (isShort && isBuy):
		// Execution and updated position have opposite directions.
		// Long position and sell execution or short position and buy execution.
		qtyDiff := p.quantity - ex.quantity

		switch {
		case qtyDiff > 0: // Executed less than updated position quantity.
			p.margin -= ex.margin
			delta := -p.debt * ex.quantity / p.quantity
			p.debt += delta

			return delta
		case qtyDiff < 0: // Executed more than updated position quantity.
			p.margin = -qtyDiff * p.instrument.Margin()
			amtDiff := -qtyDiff*ex.price*p.priceFactor - p.margin
			delta := amtDiff - p.debt
			p.debt = amtDiff

			return delta
		default: // Executed exactly the updated position quantity.
			p.margin = 0
			p.debt = 0

			return -ex.debt
		}
	default:
		// Either order side or position side are unknown.
		return 0
	}
}

// updateExecutionPnLAndMatchRoundtrips updates ex.PnL and ex.RealizedPnL, and matches roundtrips
// assuming the execution has not been appended to the history yet.
//nolint:gocognit
func (p *Position) updateExecutionPnLAndMatchRoundtrips(ex *Execution, qtySigned float64) []*Roundtrip {
	var commissionMatched, amountMatched float64

	rts := make([]*Roundtrip, 0)
	exSign := ex.quantitySign
	exQtyLeft := ex.quantity

	if (qtySigned >= 0 && exSign < 0) || (qtySigned < 0 && exSign >= 0) {
		// Execution and previous position have opposite sides.
		// Long position and sell execution or short position and buy execution.
		switch p.roundtripMatching {
		case matchings.FirstInFirstOut:
			for _, x := range p.executions {
				if exQtyLeft <= 0 {
					break
				}

				// Skip if the full quantity has already been matched or execution sides are the same.
				if x.unrealizedQuantity > 0 && exSign != x.quantitySign {
					// Execution sides are opposite and there is an unmatched quantity.
					minQty := math.Min(exQtyLeft, x.unrealizedQuantity)
					commissionMatched += minQty * (ex.commissionConvertedPerUnit + x.commissionConvertedPerUnit)
					amountMatched += -exSign * minQty * (ex.price - x.price)
					exQtyLeft -= minQty
					rts = append(rts, p.newRoundtrip(x, ex, minQty))
				}
			}
		case matchings.LastInFirstOut:
			for i := len(p.executions) - 1; i >= 0; i-- {
				if exQtyLeft <= 0 {
					break
				}

				// Skip if the full quantity has already been matched or execution sides are the same.
				x := p.executions[i]
				if x.unrealizedQuantity > 0 && exSign != x.quantitySign {
					// Execution sides are opposite and there is an unmatched quantity.
					minQty := math.Min(exQtyLeft, x.unrealizedQuantity)
					commissionMatched += minQty * (ex.commissionConvertedPerUnit + x.commissionConvertedPerUnit)
					amountMatched += -exSign * minQty * (ex.price - x.price)
					exQtyLeft -= minQty
					rts = append(rts, p.newRoundtrip(x, ex, minQty))
				}
			}
		}
	}

	amountMatched *= p.priceFactor
	ex.pnl += amountMatched
	ex.realizedPnL = amountMatched - commissionMatched

	return rts
}

// updatePrice updates p.price, p.amounts and p.performance based on new price
// assuming p.updateExecutionPnLAndMatchRoundtrips() has been called
// and the execution has been appended.
func (p *Position) updatePrice(t time.Time, price float64) {
	var unrealizedAmt float64

	for _, e := range p.executions {
		qty := e.unrealizedQuantity
		if qty > 0 {
			// Delta increments for the changed price.
			unrealizedAmt += (price - e.price) * qty * e.quantitySign
			e.unrealizedPriceHigh = math.Max(e.unrealizedPriceHigh, price)
			e.unrealizedPriceLow = math.Min(e.unrealizedPriceLow, price)
		}
	}

	p.price = price
	amt := price * p.priceFactor * p.quantitySigned
	p.amounts.Add(t, amt)
	p.perf.addDrawdown(t, amt+p.cashFlow)
	p.perf.addPnL(t, p.entryAmount, amt, unrealizedAmt*p.priceFactor, p.cashFlow)
}

// newRoundtrip creates a new roundtrip updating unrealized quantities in both entry and exit executions.
func (p *Position) newRoundtrip(entry, exit *Execution, qty float64) *Roundtrip {
	entry.unrealizedQuantity -= qty
	exit.unrealizedQuantity -= qty

	return newRoundtrip(p.instrument, entry, exit, qty)
}

// Instrument is a financial instrument associated with this position.
func (p *Position) Instrument() instruments.Instrument {
	return p.instrument
}

// Currency is the instrument's currency.
func (p *Position) Currency() currencies.Currency {
	return p.instrument.Currency()
}

// ExecutionHistory is a collection of order executions
// related to this position in chronological order.
func (p *Position) ExecutionHistory() []Execution {
	p.mu.RLock()
	defer p.mu.RUnlock()

	v := make([]Execution, len(p.executions))
	for i, e := range p.executions {
		v[i] = *e
	}

	return v
}

// Debt is the position debt in instrument's currency.
func (p *Position) Debt() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.debt
}

// Margin is the position margin in instrument's currency.
func (p *Position) Margin() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.margin
}

// Leverage is the current position leverage ratio.
func (p *Position) Leverage() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	v := p.margin
	if v == 0 {
		return 0
	}

	return p.amounts.Current() / v
}

// Price is the current price in instrument's currency.
func (p *Position) Price() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.price
}

// QuantityBought is the unsigned quantity bought in this position.
func (p *Position) QuantityBought() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.quantityBought
}

// QuantitySold is the unsigned quantity sold in this position.
func (p *Position) QuantitySold() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.quantitySold
}

// QuantitySoldShort is the unsigned quantity sold short in this position.
func (p *Position) QuantitySoldShort() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.quantitySoldShort
}

// Side is the position side (long or short).
func (p *Position) Side() pos.Side {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.side
}

// Quantity is the unsigned position quantity
// (bought minus sold minus sold short).
func (p *Position) Quantity() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.quantity
}

// CashFlow is the cash flow
// (the sum of cash flows af all order executions) in instrument's currency.
func (p *Position) CashFlow() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.cashFlow
}

// Amount is the current position amount
// (factored price times signed quantity) in instrument's currency.
func (p *Position) Amount() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amounts.Current()
}

// AmountHistory is a time series of the position amounts
// (factored price times signed quantity) in instrument's currency.
func (p *Position) AmountHistory() []data.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amounts.History()
}

// Performance tracks performance of this position in instrument's currency.
func (p *Position) Performance() *Performance {
	return p.perf
}
