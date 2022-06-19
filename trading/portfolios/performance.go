package portfolios

import (
	"sync"
	"time"
)

// Performance tracks performance of a portfolio or an individual position.
type Performance struct {
	mu  sync.RWMutex
	pnl PnL
	dd  Drawdown
	rt  RoundtripPerformance
}

func newPerformance() *Performance {
	p := &Performance{
		pnl: *newPnL(),
		dd:  Drawdown{},
		rt:  RoundtripPerformance{},
	}

	return p
}

func (p *Performance) addRoundtrip(r Roundtrip) {
	p.rt.add(r)
}

func (p *Performance) addPnL(t time.Time, entryAmount, amount, unrealizedAmount, cashFlow float64) {
	p.pnl.add(t, entryAmount, amount, amount, cashFlow)
}

func (p *Performance) addDrawdown(t time.Time, value float64) {
	p.dd.add(t, value)
}

// PnL (Profit and Loss) contains the last values and time series
// of the PnL, net PnL and unrealized PnL, and their percentages.
func (p *Performance) PnL() *PnL {
	return &p.pnl
}

// Drawdown contains the last values and time series of drawdown amount,
// percentage and their maximal values.
func (p *Performance) Drawdown() *Drawdown {
	return &p.dd
}

// Roundtrip contains the roundtrip performance statistics.
func (p *Performance) Roundtrip() *RoundtripPerformance {
	return &p.rt
}
