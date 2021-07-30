package portfolios

//nolint:gci
import (
	"mbg/trading/data/entities"
	"sync"
	"time"
)

const hundred = 100

// PnL (Profin and Loss) contains the last values and time series of
// PnL amount and percentage, net PnL amount and percentage, unrealized PnL amount.
type PnL struct {
	mu sync.RWMutex

	amount           scalarHistory
	amountNet        scalarHistory
	amountUnrealized scalarHistory
	percentage       scalarHistory
	percentageNet    scalarHistory
}

// newPnL creates a new PnL instance.
// This is the only correct way to create a PnL instance.
func newPnL() *PnL {
	return &PnL{
		amount:           scalarHistory{},
		amountNet:        scalarHistory{},
		amountUnrealized: scalarHistory{},
		percentage:       scalarHistory{},
		percentageNet:    scalarHistory{},
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
func (p *PnL) AmountHistory() []entities.Scalar {
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
func (p *PnL) PercentageHistory() []entities.Scalar {
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
func (p *PnL) UnrealizedAmountHistory() []entities.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amountUnrealized.History()
}

// NetAmount returns the current net Profit and Loss amount
// (the value plus the net cash flow) or zero if not initialized.
func (p *PnL) NetAmount() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amountNet.Current()
}

// NetAmountHistory returns the net Profit and Loss amount time series
// or an empty slice if not initialized.
func (p *PnL) NetAmountHistory() []entities.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.amountNet.History()
}

// NetPercentage returns the net Profit and Loss percentage
// (the net PnL amount divided by the initial amount, expressed in %) or zero if not initialized.
func (p *PnL) NetPercentage() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.percentageNet.Current()
}

// NetPercentageHistory returns the net Profit and Loss percentage time series
// or an empty slice if not initialized.
func (p *PnL) NetPercentageHistory() []entities.Scalar {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.percentageNet.History()
}

// add adds a new sample to the time series
// if the new sample time is later the last time of the time series.
// Otherwise (if the new sample time is less or equal to the last time),
// the time series will not be updated.
func (p *PnL) add(t time.Time, entryAmount, amount, unrealizedAmount, cashFlow, netCashFlow float64) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var pct, netPct float64

	if entryAmount != 0 {
		pct = (amount + cashFlow) / entryAmount * hundred
		netPct = (amount + netCashFlow) / entryAmount * hundred
	}

	p.amount.add(t, amount+cashFlow)
	p.amountNet.add(t, amount+netCashFlow)
	p.amountUnrealized.add(t, unrealizedAmount)
	p.percentage.add(t, pct)
	p.percentageNet.add(t, netPct)
}
