package portfolios

//nolint:gci
import (
	"sync"
	"time"

	"mbg/trading/data"
)

const hundred = 100

// PnL (Profin and Loss) contains the last values and time series of
// PnL amount and percentage, unrealized PnL amount.
type PnL struct {
	mu sync.RWMutex

	amount           data.ScalarTimeSeries
	amountUnrealized data.ScalarTimeSeries
	percentage       data.ScalarTimeSeries
}

// newPnL creates a new PnL instance.
// This is the only correct way to create a PnL instance.
func newPnL() *PnL {
	return &PnL{
		amount:           data.ScalarTimeSeries{},
		amountUnrealized: data.ScalarTimeSeries{},
		percentage:       data.ScalarTimeSeries{},
	}
}

// Amount returns the current Profit and Loss amount
// (the value plus the cash flow) or zero if not initialized.
func (p *PnL) Amount() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amount.Current()
}

// AmountHistory returns the Profit and Loss amount time series
// or an empty slice if not initialized.
func (p *PnL) AmountHistory() []data.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amount.History()
}

// Percentage returns the Profit and Loss percentage
// (the PnL amount divided by the initial amount, expressed in %) or zero if not initialized.
func (p *PnL) Percentage() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.percentage.Current()
}

// PercentageHistory returns the Profit and Loss percentage time series
// or an empty slice if not initialized.
func (p *PnL) PercentageHistory() []data.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.percentage.History()
}

// UnrealizedAmount returns the unrealized Profit and Loss amount
// (the theoretical marked-to-market gain or loss on the open position(s)
// valued at current market price) or zero if not initialized.
// Unrealized gains and losses become PnL when position(s) is(are) closed.
func (p *PnL) UnrealizedAmount() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amountUnrealized.Current()
}

// UnrealizedAmountHistory returns the unrealized Profit and Loss amount time series
// or an empty slice if not initialized.
func (p *PnL) UnrealizedAmountHistory() []data.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amountUnrealized.History()
}

// add adds a new sample to the time series
// if the new sample time is later the last time of the time series.
// Otherwise (if the new sample time is less or equal to the last time),
// the time series will not be updated.
func (p *PnL) add(t time.Time, entryAmount, amount, unrealizedAmount, cashFlow float64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var pct float64

	if entryAmount != 0 {
		pct = (amount + cashFlow) / entryAmount * hundred
	}

	p.amount.Add(t, amount+cashFlow)
	p.amountUnrealized.Add(t, unrealizedAmount)
	p.percentage.Add(t, pct)
}

func (p *PnL) add2(t time.Time, entryAmount, amount, unrealizedAmount float64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var pct float64

	if entryAmount != 0 {
		pct = amount / entryAmount * hundred
	}

	p.amount.Add(t, amount)
	p.amountUnrealized.Add(t, unrealizedAmount)
	p.percentage.Add(t, pct)
}
