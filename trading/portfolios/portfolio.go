package portfolios

//nolint:gci
import (
	"math"
	"mbg/trading/currencies"
	"mbg/trading/instruments"
	"mbg/trading/orders"
	"mbg/trading/orders/reports"
	"sync"
	"time"
)

// Portfolio is a portfolio.
type Portfolio struct {
	mu         sync.RWMutex
	currency   currencies.Currency
	converter  currencies.Converter
	account    *Account
	positions  map[instruments.Instrument]*Position
	executions []*Execution
}

// NewPortfolio creates a new portfolio.
// This is the only correct way to create a portfolio instance.
func NewPortfolio(holder string, currency currencies.Currency, converter currencies.Converter) *Portfolio {
	p := &Portfolio{
		currency:   currency,
		converter:  converter,
		account:    newAccount(holder, currency, converter),
		positions:  make(map[instruments.Instrument]*Position),
		executions: []*Execution{},
	}

	return p
}

// Deposit deposits an amount of money in the specified currency into an account associated with this portfolio.
//
// The amount will be converted into the home currency if the indicated currency differs from the home one.
func (p *Portfolio) Deposit(time time.Time, amount float64, currency currencies.Currency, note string) {
	p.account.add(time, math.Abs(amount), currency, note)
}

// Withdraw withdraws an amount of money in the specified currency from an account associated with this portfolio.
//
// It does not check if the balance becomes negative.
//
// The amount will be converted into the home currency if the indicated currency differs from the home one.
func (p *Portfolio) Withdraw(time time.Time, amount float64, currency currencies.Currency, note string) {
	p.account.add(time, -math.Abs(amount), currency, note)
}

// OrderSingleExecution adds an order execution to the related portfolio position.
func (p *Portfolio) OrderSingleExecution(report orders.OrderSingleExecutionReport) {
	switch report.ReportType() {
	case reports.Filled, reports.PartiallyFilled:
		break
	default:
		return
	}

	exec := newExecutionOrderSingle(report, p.converter)
	instr := report.Order().Instrument

	var pos *Position

	p.mu.Lock()
	defer p.mu.Unlock()

	if po, ok := p.positions[instr]; ok {
		pos = po
		pos.add(exec, p.account)
	} else {
		pos := newPosition(instr, exec, p.account)
		p.positions[instr] = pos
	}

	p.executions = append(p.executions, exec)
}
