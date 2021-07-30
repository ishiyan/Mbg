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
	pnl               *PnL
	drawdown          *Drawdown
	indexPnLExecution int
	executions        []*Execution
	mu                sync.RWMutex
}

// newPosition creates a new position in a given instrument.
// This is the only correct way to create a position instance.
func newPosition(instr instruments.Instrument, exec *Execution, account *Account) *Position {
	t := exec.reportTime
	p := Position{
		instrument:     instr,
		margin:         exec.margin,
		debt:           exec.debt,
		priceFactor:    1,
		price:          exec.price,
		quantityPnL:    exec.quantity,
		executions:     make([]*Execution, 0),
		entryAmount:    exec.amount,
		quantity:       exec.quantity,
		quantitySigned: scalarHistory{},
		cashFlow:       scalarHistory{},
		cashFlowNet:    scalarHistory{},
		amounts:        scalarHistory{},
		pnl:            newPnL(),
		drawdown:       &Drawdown{},
	}

	if instr.PriceFactor != 0 {
		p.priceFactor = instr.PriceFactor
	}

	exec.pnl = -exec.commissionConverted
	exec.realizedPnL = 0
	p.updateSideAndQuantities(exec, 0)
	p.executions = append(p.executions, exec)

	p.amounts.add(t, exec.amount)
	p.cashFlow.accumulate(t, exec.cashFlow)
	p.cashFlowNet.accumulate(t, exec.netCashFlow)
	p.pnl.add(t, exec.amount, exec.amount, exec.amount, exec.cashFlow, exec.netCashFlow)
	p.drawdown.add(t, exec.amount+exec.cashFlow)
	account.addExecution(exec)

	return &p
}

// add adds an execution to the existing position.
// This should only be called on position created by NewPosition.
// Instrument of the execution should match position instrument.
func (p *Position) add(exec *Execution, account *Account) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// debtIncrement := p.updateMarginAndDebt(exec)
	qtySigned := p.quantitySigned.Current()
	p.updateExecutionPnL(exec, qtySigned)
	p.updateSideAndQuantities(exec, qtySigned)
	p.executions = append(p.executions, exec)
	p.cashFlow.accumulate(exec.reportTime, exec.cashFlow)
	p.cashFlowNet.accumulate(exec.reportTime, exec.netCashFlow)

	account.addExecution(exec)

	p.updatePrice(exec.reportTime, exec.price)
}

// updateSideAndQuantities updates p.side, p.quantity, p.quantitySigned,
// p.quantityBought, p.quantitySold and p.quantitySoldShort
// for the given execution and the given current signed quantity.
func (p *Position) updateSideAndQuantities(exec *Execution, qtySigned float64) {
	switch exec.side {
	case sides.Buy, sides.BuyMinus:
		p.quantityBought += exec.quantity
		qtySigned += exec.quantity
	case sides.Sell, sides.SellPlus:
		p.quantitySold += exec.quantity
		qtySigned -= exec.quantity
	case sides.SellShort, sides.SellShortExempt:
		p.quantitySoldShort -= exec.quantity
		qtySigned -= exec.quantity
	}

	p.quantitySigned.add(exec.reportTime, qtySigned)
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
func (p *Position) updateMarginAndDebt(exec *Execution) float64 {
	if exec.margin == 0 {
		return 0
	}

	isLong := p.side == pos.Long
	isShort := p.side == pos.Short
	isBuy := exec.side.IsBuy()
	isSell := exec.side.IsSell()

	switch {
	case (isLong && isBuy) || (isShort && isSell):
		// Execution and updated position have the same directions.
		// Long position and buy execution or short position and sell execution.
		p.margin += exec.margin
		p.debt += exec.debt

		return exec.debt
	case (isLong && isSell) || (isShort && isBuy):
		// Execution and updated position have opposite directions.
		// Long position and sell execution or short position and buy execution.
		qtyDiff := p.quantity - exec.quantity

		switch {
		case qtyDiff > 0: // Executed less than updated position quantity.
			p.margin -= exec.margin
			delta := -p.debt * exec.quantity / p.quantity
			p.debt += delta

			return delta
		case qtyDiff < 0: // Executed more than updated position quantity.
			p.margin = -qtyDiff * p.instrument.Margin
			amtDiff := -qtyDiff*exec.price*p.priceFactor - p.margin
			delta := amtDiff - p.debt
			p.debt = amtDiff

			return delta
		default: // Executed exactly the updated position quantity.
			p.margin = 0
			p.debt = 0

			return -exec.debt
		}
	default:
		// Either order side or position side are unknown.
		return 0
	}
}

// updateExecutionPnL updates exec.PnL, exec.RealizedPnL, p.quantity and p.indexPnLExecution,
// assuming the execution has not been appended to the history yet.
func (p *Position) updateExecutionPnL(exec *Execution, qtySigned float64) {
	commission := 0.
	delta := qtySigned

	execSign := exec.quantitySign()
	qtyExec := exec.quantity

	if (delta >= 0 && execSign < 0) || (delta < 0 && execSign >= 0) {
		// Execution and updated position have opposite directions.
		// Long position and sell execution or short position and buy execution.
		qtyPnL := p.quantityPnL
		qtyMin := math.Min(qtyExec, qtyPnL)
		execCommPerQtyUnit := exec.commissionConverted / exec.quantity
		lenExecutions := len(p.executions)
		pnlExecution := p.executions[p.indexPnLExecution]
		pnlIndexNext := p.indexPnLExecution + 1

		commission = qtyMin * (execCommPerQtyUnit + pnlExecution.commissionConverted/pnlExecution.quantity)
		delta = -execSign * qtyMin * (exec.price - pnlExecution.price)

		for ; qtyExec > qtyPnL && pnlIndexNext < lenExecutions; pnlIndexNext++ {
			pnlExecution = p.executions[pnlIndexNext]
			if pnlExecution.quantitySign() != execSign {
				qtyMin = math.Min(qtyExec-qtyPnL, pnlExecution.quantity)
				commission += qtyMin * (execCommPerQtyUnit + pnlExecution.commissionConverted/pnlExecution.quantity)
				delta += -execSign * qtyMin * (exec.price - pnlExecution.price)
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
	exec.pnl = delta - exec.commissionConverted
	exec.realizedPnL = delta - commission
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
	p.drawdown.add(t, amt+p.cashFlow.Current())

	switch {
	case qty > 0:
		qty = p.quantityPnL
	case qty < 0:
		qty = -p.quantityPnL
	default:
		p.pnl.add(t, p.entryAmount, amt, 0, p.cashFlow.Current(), p.cashFlowNet.Current())

		return
	}

	var unr float64
	for i := p.indexPnLExecution; i < len(p.executions); i++ {
		unr += (price - p.executions[i].price) * qty
	}

	unr *= p.priceFactor
	p.pnl.add(t, p.entryAmount, amt, unr, p.cashFlow.Current(), p.cashFlowNet.Current())
}

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

// PnL (Profit and Loss) contains the last values and time series
// of the PnL, net PnL and unrealized PnL in instrument's currency,
// and their percentages.
func (p *Position) PnL() *PnL {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.pnl
}

// Drawdown contains the last values and time series of drawdown amount
// in instrument's currency, percentage and their maximal values.
func (p *Position) Drawdown() *Drawdown {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.drawdown
}
