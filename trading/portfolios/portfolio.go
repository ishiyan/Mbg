package portfolios

//nolint:gofumpt
import (
	"math"
	"sync"
	"time"

	"mbg/trading/currencies"
	"mbg/trading/instruments"
	"mbg/trading/orders"
	"mbg/trading/orders/reports"
	"mbg/trading/portfolios/roundtrips/matchings"
)

// Portfolio is a portfolio.
type Portfolio struct {
	mu                sync.RWMutex
	roundtripMatching matchings.Matching
	currency          currencies.Currency
	converter         currencies.Converter
	initialCash       float64
	account           *Account
	positions         map[instruments.Instrument]*Position
	executions        []*Execution
	perf              *Performance
}

// NewPortfolio creates a new portfolio.
func NewPortfolio(holder string, cash float64, currency currencies.Currency, converter currencies.Converter,
	matching matchings.Matching,
) *Portfolio {
	p := &Portfolio{
		roundtripMatching: matching,
		currency:          currency,
		converter:         converter,
		initialCash:       cash,
		account:           newAccount(holder, currency, converter),
		positions:         make(map[instruments.Instrument]*Position),
		executions:        []*Execution{},
		perf:              newPerformance(),
	}

	p.Deposit(time.Time{}, cash, currency, "Initial deposit")

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

	p.mu.Lock()
	defer p.mu.Unlock()

	if pos, ok := p.positions[instr]; ok {
		p.addRoundtrips(pos, pos.add(exec, p.account))
		p.addPerformance(pos)
	} else {
		pos := newPosition(instr, exec, p.account, p.roundtripMatching)
		p.positions[instr] = pos
		p.addPerformance(pos)
	}

	p.executions = append(p.executions, exec)
}

func (p *Portfolio) portfolioAmount() float64 {
	var amount float64

	for _, pos := range p.positions {
		v := pos.Amount()
		if pos.Currency() != p.currency {
			v *= p.converter.ExchangeRate(pos.Currency(), p.currency)
		}

		amount += v
	}

	return amount
}

func (p *Portfolio) summatePerformance() (pnl, drawdown float64) {
	var amount float64

	for _, pos := range p.positions {
		v := pos.Amount()
		if pos.Currency() != p.currency {
			v *= p.converter.ExchangeRate(pos.Currency(), p.currency)
		}

		amount += v
	}

	return amount, 0
}

func (p *Portfolio) addPerformance(pos *Position) {
	switch {
	case pos.Currency() == p.currency:
		p.perf.addDrawdown(pos.perf.dd.amount.Time, pos.perf.dd.amount.Value)
	default:
		// rate := p.converter.ExchangeRate(pos.Currency(), p.currency)
	}
}

func (p *Portfolio) addRoundtrips(pos *Position, rts []*Roundtrip) {
	if len(rts) == 0 {
		return
	}

	switch {
	case pos.Currency() == p.currency:
		for _, rt := range rts {
			p.perf.addRoundtrip(*rt)
		}
	default:
		rate := p.converter.ExchangeRate(pos.Currency(), p.currency)

		for _, rt := range rts {
			r := *rt
			r.entryPrice *= rate
			r.exitPrice *= rate
			r.highPrice *= rate
			r.lowPrice *= rate
			r.commission *= rate
			r.pnl *= rate
			p.perf.addRoundtrip(r)
		}
	}
}
