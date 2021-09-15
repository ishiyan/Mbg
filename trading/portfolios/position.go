package portfolios

//nolint:gci
import (
	"math"
	"mbg/trading/currencies"
	"mbg/trading/data/entities"
	"mbg/trading/instruments"
	"mbg/trading/orders/sides"
	pos "mbg/trading/portfolios/positions/sides"
	"sync"
	"time"
)

// Position is a portfolio position.
type Position struct {
	// Notifies when this position has been executed.
	// event Action<PortfolioPosition, PortfolioExecution> Executed
	// Notifies when this position has been changed.
	// event Action<PortfolioPosition, DateTime> Changed
	// PortfolioMonitors Monitors

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
	quantitySigned    scalarHistory
	quantityPnL       float64
	side              pos.Side
	cashFlow          scalarHistory
	cashFlowNet       scalarHistory
	amounts           scalarHistory
	indexPnLExecution int
	executions        []*Execution
	perf              *Performance
	mu                sync.RWMutex
}

// newPosition creates a new position in a given instrument.
// This is the only correct way to create a position instance.
func newPosition(instr instruments.Instrument, ex *Execution, account *Account) *Position {
	t := ex.reportTime
	p := Position{
		instrument:     instr,
		margin:         ex.margin,
		debt:           ex.debt,
		priceFactor:    1,
		price:          ex.price,
		quantityPnL:    ex.quantity,
		executions:     make([]*Execution, 0),
		entryAmount:    ex.amount,
		quantity:       ex.quantity,
		quantitySigned: scalarHistory{},
		cashFlow:       scalarHistory{},
		cashFlowNet:    scalarHistory{},
		amounts:        scalarHistory{},
		perf:           newPerformance(),
	}

	if instr.PriceFactor != 0 {
		p.priceFactor = instr.PriceFactor
	}

	ex.pnl = -ex.commissionConverted
	ex.realizedPnL = 0
	p.updateSideAndQuantities(ex, 0)
	p.executions = append(p.executions, ex)

	p.amounts.add(t, ex.amount)
	p.cashFlow.accumulate(t, ex.cashFlow)
	p.cashFlowNet.accumulate(t, ex.netCashFlow)
	p.perf.addPnL(t, ex.amount, ex.amount, ex.amount, ex.cashFlow, ex.netCashFlow)
	p.perf.addDrawdown(t, ex.amount+ex.cashFlow)
	account.addExecution(ex)

	return &p
}

// add adds an execution to the existing position.
// This should only be called on position created by newPosition.
// Instrument of the execution should match position instrument.
func (p *Position) add(ex *Execution, account *Account) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// debtIncrement := p.updateMarginAndDebt(exec)
	qtySigned := p.quantitySigned.Current()
	p.updateExecutionPnL(ex, qtySigned)
	p.updateSideAndQuantities(ex, qtySigned)
	p.executions = append(p.executions, ex)
	p.cashFlow.accumulate(ex.reportTime, ex.cashFlow)
	p.cashFlowNet.accumulate(ex.reportTime, ex.netCashFlow)

	account.addExecution(ex)

	p.updatePrice(ex.reportTime, ex.price)
}

// updateSideAndQuantities updates p.side, p.quantity, p.quantitySigned,
// p.quantityBought, p.quantitySold and p.quantitySoldShort
// for the given execution and the given current signed quantity.
func (p *Position) updateSideAndQuantities(ex *Execution, qtySigned float64) {
	switch ex.side {
	case sides.Buy, sides.BuyMinus:
		p.quantityBought += ex.quantity
		qtySigned += ex.quantity
	case sides.Sell, sides.SellPlus:
		p.quantitySold += ex.quantity
		qtySigned -= ex.quantity
	case sides.SellShort, sides.SellShortExempt:
		p.quantitySoldShort -= ex.quantity
		qtySigned -= ex.quantity
	}

	p.quantitySigned.add(ex.reportTime, qtySigned)
	p.quantity = math.Abs(qtySigned)

	if qtySigned < 0 {
		p.side = pos.Short
	} else {
		p.side = pos.Long
	}
}

// updateMarginAndDebt updates p.margin and p.debt and returns
// a signed increment in debt caused by this execution.
// Uses p.side and p.quantity updated by updateSideAndQuantities.
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
			p.margin = -qtyDiff * p.instrument.Margin
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

// updateExecutionPnL updates exec.PnL, exec.RealizedPnL, p.quantity and p.indexPnLExecution,
// assuming the execution has not been appended to the history yet.
func (p *Position) updateExecutionPnL(ex *Execution, qtySigned float64) {
	commission := 0.
	delta := qtySigned

	execSign := ex.quantitySign()
	qtyExec := ex.quantity

	if (delta >= 0 && execSign < 0) || (delta < 0 && execSign >= 0) {
		// Execution and updated position have opposite directions.
		// Long position and sell execution or short position and buy execution.
		qtyPnL := p.quantityPnL
		qtyMin := math.Min(qtyExec, qtyPnL)
		execCommPerQtyUnit := ex.commissionConverted / ex.quantity
		lenExecutions := len(p.executions)
		pnlExecution := p.executions[p.indexPnLExecution]
		pnlIndexNext := p.indexPnLExecution + 1

		commission = qtyMin * (execCommPerQtyUnit + pnlExecution.commissionConverted/pnlExecution.quantity)
		delta = -execSign * qtyMin * (ex.price - pnlExecution.price)

		for ; qtyExec > qtyPnL && pnlIndexNext < lenExecutions; pnlIndexNext++ {
			pnlExecution = p.executions[pnlIndexNext]
			if pnlExecution.quantitySign() != execSign {
				qtyMin = math.Min(qtyExec-qtyPnL, pnlExecution.quantity)
				commission += qtyMin * (execCommPerQtyUnit + pnlExecution.commissionConverted/pnlExecution.quantity)
				delta += -execSign * qtyMin * (ex.price - pnlExecution.price)
				qtyPnL += pnlExecution.quantity
			}
		}

		p.quantity = math.Abs(qtyExec - qtyPnL)

		if (qtyExec == qtyPnL && pnlIndexNext == lenExecutions) || (qtyExec > qtyPnL) {
			p.indexPnLExecution = lenExecutions
		} else {
			p.indexPnLExecution = pnlIndexNext - 1
		}
	}

	delta *= p.priceFactor
	ex.pnl = delta - ex.commissionConverted
	ex.realizedPnL = delta - commission
}

// updatePrice adds p.price, p.amounts, p.drawdown, p.pnl, p.pnlNet,
// p.pnlPercent, p.pnlNetPercent, p.pnlUnrealized based on new price
// assuming updatePnl() has been called.
func (p *Position) updatePrice(t time.Time, price float64) {
	for _, e := range p.executions {
		if e.quantity != e.roundtripQuantity {
			e.roundtripPriceHigh = math.Max(e.roundtripPriceHigh, price)
			e.roundtripPriceLow = math.Min(e.roundtripPriceLow, price)
		}
	}

	p.price = price
	qty := p.quantitySigned.Current()
	amt := price * p.priceFactor * qty
	p.amounts.add(t, amt)
	p.perf.addDrawdown(t, amt+p.cashFlow.Current())

	switch {
	case qty > 0:
		qty = p.quantityPnL
	case qty < 0:
		qty = -p.quantityPnL
	default:
		p.perf.addPnL(t, p.entryAmount, amt, 0, p.cashFlow.Current(), p.cashFlowNet.Current())

		return
	}

	var unr float64
	for i := p.indexPnLExecution; i < len(p.executions); i++ {
		unr += (price - p.executions[i].price) * qty
	}

	unr *= p.priceFactor
	p.perf.addPnL(t, p.entryAmount, amt, unr, p.cashFlow.Current(), p.cashFlowNet.Current())
}

/*
// matchRoundtrips matches LIFO roundtrips
// assuming the execution has not been appended to the history yet.
func (p *Position) matchRoundtrips(exec *Execution) []*Roundtrip {
	eq := exec.quantity
	es := exec.quantitySign()

	for _, e := range p.executions {
		q := e.quantity
		qr := e.roundtripQuantity

		// Skip if the full quantity has already been matched or execution directions are the same.
		if qr == q || es == e.quantitySign() {
			continue
		}

		// Execution directions are opposite and there is an unmatched quantity.
		qty := math.Min(eq, q-qr)

	}

	return nil
}
*/
func (p *Position) newRoundtrip(entry, exit *Execution, qty float64) *Roundtrip {
	entry.roundtripQuantity += qty
	exit.roundtripQuantity += qty

	rt := newRoundtrip(p.instrument, entry, exit, qty)

	return rt
}

// Instrument is a financial instrument associated with this position.
func (p *Position) Instrument() instruments.Instrument {
	return p.instrument
}

// Currency is the instrument's currency.
func (p *Position) Currency() currencies.Currency {
	return p.instrument.Currency
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

// EntryAmount is the current round-trip entry amount of this position in instrument's currency.
func (p *Position) EntryAmount() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.entryAmount
}

// Debt is the position debt in instrument's currency.
func (p *Position) Debt() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.debt
}

// Margin is the position margin.
func (p *Position) Margin() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.margin
}

// Leverage is the current position leverage in instrument's currency.
func (p *Position) Leverage() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	v := p.margin // ??? * p.cashFlow.Current()
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

// Side is the position side (short or long).
func (p *Position) Side() pos.Side {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.side
}

// Quantity is the unsigned position quantity.
func (p *Position) Quantity() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.quantity
}

// QuantitySigned is the signed position quantity
// (bought minus sold minus sold short).
func (p *Position) QuantitySigned() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.quantitySigned.Current()
}

// QuantitySignedHistory is a time series of the signed position quantities
// (bought minus sold minus sold short).
func (p *Position) QuantitySignedHistory() []entities.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.quantitySigned.History()
}

// CashFlow is the cash flow
// (the sum of cash flows af all order executions) in instrument's currency.
func (p *Position) CashFlow() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.cashFlow.Current()
}

// CashFlowHistory is a time series of the cash flow
// (the sum of cash flows af all order executions) in instrument's currency.
func (p *Position) CashFlowHistory() []entities.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.cashFlow.History()
}

// NetCashFlow is the net cash flow
// (the sum of net cash flows af all executions) in instrument's currency.
func (p *Position) NetCashFlow() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.cashFlowNet.Current()
}

// NetCashFlowHistory is a time series of the net cash flow
// (the sum of net cash flows af all executions) in instrument's currency.
func (p *Position) NetCashFlowHistory() []entities.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.cashFlowNet.History()
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
func (p *Position) AmountHistory() []entities.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amounts.History()
}

// Performance tracks performance of this position in instrument's currency.
func (p *Position) Performance() *Performance {
	return p.perf
}
