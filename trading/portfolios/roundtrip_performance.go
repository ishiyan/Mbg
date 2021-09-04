package portfolios

import (
	"sync"
	"time"
)

// RoundtripPerformance is a roundtrip performance.
type RoundtripPerformance struct {
	mu         sync.RWMutex
	roundtrips []*Roundtrip
}

// NewRoundtripPerformance creates a new round-trip portfolio performance.
// This is the only correct way to create a round-trip performance instance.
func NewRoundtripPerformance() *RoundtripPerformance {
	return &RoundtripPerformance{}
}

// Add adds a new round-trip.
func (rp *RoundtripPerformance) Add(r *Roundtrip) {
	rp.mu.Lock()
	defer rp.mu.Unlock()

	rp.roundtrips = append(rp.roundtrips, r)
}

// TotalCount is a total number of round-trips.
func (rp *RoundtripPerformance) TotalCount() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.totalCount()
}

// WinningCount is a number of winning round-trips.
func (rp *RoundtripPerformance) WinningCount() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.winningCount()
}

// NetWinningCount is a number of net winning round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetWinningCount() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netWinningCount()
}

// LoosingCount is a number of loosing round-trips.
func (rp *RoundtripPerformance) LoosingCount() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.loosingCount()
}

// NetLoosingCount is a number of net loosing round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetLoosingCount() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netLoosingCount()
}

// WinningPct is a percent of profitable round-trips.
func (rp *RoundtripPerformance) WinningPct() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divIntInt(rp.winningCount(), rp.totalCount())
}

// NetWinningPct is a percent of net profitable round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetWinningPct() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divIntInt(rp.netWinningCount(), rp.totalCount())
}

// LoosingPct is a percent of loosing round-trips.
func (rp *RoundtripPerformance) LoosingPct() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divIntInt(rp.loosingCount(), rp.totalCount())
}

// NetLoosingPct is a percent of net loosing round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetLoosingPct() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divIntInt(rp.netLoosingCount(), rp.totalCount())
}

// TotalPnL is a Profit and Loss of total round-trips.
func (rp *RoundtripPerformance) TotalPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.totalPnL()
}

// NetTotalPnL is a net Profit and Loss of total round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetTotalPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netTotalPnL()
}

// WinningPnL is a Profit and Loss of profitable round-trips.
func (rp *RoundtripPerformance) WinningPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.winningPnL()
}

// NetWinningPnL is a netProfit and Loss of profitable round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetWinningPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netWinningPnL()
}

// LoosingPnL is a Profit and Loss of loosing round-trips.
func (rp *RoundtripPerformance) LoosingPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.loosingPnL()
}

// NetLoosingPnL is a net Profit and Loss of loosing round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetLoosingPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netLoosingPnL()
}

// AvgTotalPnL is an average Profit and Loss of total round-trips.
func (rp *RoundtripPerformance) AvgTotalPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatInt(rp.totalPnL(), rp.totalCount())
}

// NetAvgTotalPnL is an average net Profit and Loss of total round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetAvgTotalPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatInt(rp.netTotalPnL(), rp.totalCount())
}

// AvgWinningPnL is an average Profit and Loss of profitable round-trips.
func (rp *RoundtripPerformance) AvgWinningPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatInt(rp.winningPnL(), rp.winningCount())
}

// NetAvgWinningPnL is an average net Profit and Loss of profitable round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetAvgWinningPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatInt(rp.netWinningPnL(), rp.netWinningCount())
}

// AvgLoosingPnL is an average Profit and Loss of loosing round-trips.
func (rp *RoundtripPerformance) AvgLoosingPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatInt(rp.loosingPnL(), rp.loosingCount())
}

// NetAvgLoosingPnL is an average net Profit and Loss of loosing round-trips (taking commission into account).
func (rp *RoundtripPerformance) NetAvgLoosingPnL() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatInt(rp.netLoosingPnL(), rp.netLoosingCount())
}

// AvgWinningLoosingPct is an average winning / average loosing Profit and Loss percentage.
func (rp *RoundtripPerformance) AvgWinningLoosingPct() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatFloat(
		100*rp.divFloatInt(rp.winningPnL(), rp.winningCount()), //nolint:gomnd
		-rp.divFloatInt(rp.loosingPnL(), rp.loosingCount()))
}

// NetAvgWinningLoosingPct is an average net winning / average net loosing Profit and Loss percentage
// (taking commission into account).
func (rp *RoundtripPerformance) NetAvgWinningLoosingPct() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatFloat(
		100*rp.divFloatInt(rp.netWinningPnL(), rp.netWinningCount()), //nolint:gomnd
		-rp.divFloatInt(rp.netLoosingPnL(), rp.netLoosingCount()))
}

// ProfitPct is a winning / loosing Profit and Loss percentage.
func (rp *RoundtripPerformance) ProfitPct() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatFloat(100*rp.winningPnL(), -rp.loosingPnL()) //nolint:gomnd
}

// NetProfitPct is a winning / loosing net Profit and Loss percentage
// (taking commission into account).
func (rp *RoundtripPerformance) NetProfitPct() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.divFloatFloat(100*rp.netWinningPnL(), -rp.netLoosingPnL()) //nolint:gomnd
}

// MaxConsecutiveWinners is a maximal number of consecutive winning round-trips.
func (rp *RoundtripPerformance) MaxConsecutiveWinners() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	m := 0
	cnt := 0

	for _, r := range rp.roundtrips {
		if r.PnL() > 0 {
			cnt++
		} else {
			m = max(m, cnt)
			cnt = 0
		}
	}

	return max(m, cnt)
}

// NetMaxConsecutiveWinners is a maximal number of consecutive net winning round-trips
// (taking commission into account).
func (rp *RoundtripPerformance) NetMaxConsecutiveWinners() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	m := 0
	cnt := 0

	for _, r := range rp.roundtrips {
		if r.NetPnL() > 0 {
			cnt++
		} else {
			m = max(m, cnt)
			cnt = 0
		}
	}

	return max(m, cnt)
}

// MaxConsecutiveLoosers is a maximal number of consecutive loosing round-trips.
func (rp *RoundtripPerformance) MaxConsecutiveLoosers() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	m := 0
	cnt := 0

	for _, r := range rp.roundtrips {
		if r.PnL() < 0 {
			cnt++
		} else {
			m = max(m, cnt)
			cnt = 0
		}
	}

	return max(m, cnt)
}

// NetMaxConsecutiveLoosers is a maximal number of consecutive net loosing round-trips
// (taking commission into account).
func (rp *RoundtripPerformance) NetMaxConsecutiveLoosers() int {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	m := 0
	cnt := 0

	for _, r := range rp.roundtrips {
		if r.NetPnL() < 0 {
			cnt++
		} else {
			m = max(m, cnt)
			cnt = 0
		}
	}

	return max(m, cnt)
}

// AvgDuration is an average duration of a round-trip.
func (rp *RoundtripPerformance) AvgDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.avgDuration()
}

// AvgWinningDuration is an average duration of a profitable round-trip.
func (rp *RoundtripPerformance) AvgWinningDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.avgWinningDuration()
}

// NetAvgWinningDuration is an average duration of a net profitable round-trip
// (taking commission into account).
func (rp *RoundtripPerformance) NetAvgWinningDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netAvgWinningDuration()
}

// AvgLoosingDuration is an average duration of a loosing round-trip.
func (rp *RoundtripPerformance) AvgLoosingDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.avgLoosingDuration()
}

// NetAvgLoosingDuration is an average duration of a net loosing round-trip
// (taking commission into account).
func (rp *RoundtripPerformance) NetAvgLoosingDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netAvgLoosingDuration()
}

// MinDuration is a minimal duration of a round-trip.
func (rp *RoundtripPerformance) MinDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.minDuration()
}

// MinWinningDuration is a minimal profitable duration of a round-trip.
func (rp *RoundtripPerformance) MinWinningDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.minWinningDuration()
}

// NetMinWinningDuration is a minimal net profitable duration of a round-trip
// (taking commission into account).
func (rp *RoundtripPerformance) NetMinWinningDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netMinWinningDuration()
}

// MinLoosingDuration is a minimal loosing duration of a round-trip.
func (rp *RoundtripPerformance) MinLoosingDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.minLoosingDuration()
}

// NetMinLoosingDuration is a minimal net loosing duration of a round-trip
// (taking commission into account).
func (rp *RoundtripPerformance) NetMinLoosingDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netMinLoosingDuration()
}

// MaxDuration is a maximal duration of a round-trip.
func (rp *RoundtripPerformance) MaxDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.maxDuration()
}

// MaxWinningDuration is a maximal profitable duration of a round-trip.
func (rp *RoundtripPerformance) MaxWinningDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.maxWinningDuration()
}

// NetMaxWinningDuration is a maximal net profitable duration of a round-trip
// (taking commission into account).
func (rp *RoundtripPerformance) NetMaxWinningDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netMaxWinningDuration()
}

// MaxLoosingDuration is a maximal loosing duration of a round-trip.
func (rp *RoundtripPerformance) MaxLoosingDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.maxLoosingDuration()
}

// NetMaxLoosingDuration is a maximal net loosing duration of a round-trip
// (taking commission into account).
func (rp *RoundtripPerformance) NetMaxLoosingDuration() time.Duration {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.netMaxLoosingDuration()
}

// AvgMAE is an average Maximum Adverse Excursion of all round-trips.
func (rp *RoundtripPerformance) AvgMAE() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.avgMAE()
}

// AvgMFE is an average Maximum Favorable Excursion of all round-trips.
func (rp *RoundtripPerformance) AvgMFE() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.avgMFE()
}

// AvgEntryEfficiency is an average Entry Efficiency of all round-trips.
func (rp *RoundtripPerformance) AvgEntryEfficiency() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.avgEntryEfficiency()
}

// AvgExitEfficiency is an average Exit Efficiency of all round-trips.
func (rp *RoundtripPerformance) AvgExitEfficiency() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.avgExitEfficiency()
}

// AvgTotalEfficiency is an average Total Efficiency of all round-trips.
func (rp *RoundtripPerformance) AvgTotalEfficiency() float64 {
	rp.mu.RLock()
	defer rp.mu.RUnlock()

	return rp.avgTotalEfficiency()
}

func (rp *RoundtripPerformance) totalCount() int {
	return len(rp.roundtrips)
}

func (rp *RoundtripPerformance) winningCount() int {
	cnt := 0

	for _, r := range rp.roundtrips {
		if r.PnL() > 0 {
			cnt++
		}
	}

	return cnt
}

func (rp *RoundtripPerformance) loosingCount() int {
	cnt := 0

	for _, r := range rp.roundtrips {
		if r.PnL() < 0 {
			cnt++
		}
	}

	return cnt
}

func (rp *RoundtripPerformance) netWinningCount() int {
	cnt := 0

	for _, r := range rp.roundtrips {
		if r.NetPnL() > 0 {
			cnt++
		}
	}

	return cnt
}

func (rp *RoundtripPerformance) netLoosingCount() int {
	cnt := 0

	for _, r := range rp.roundtrips {
		if r.NetPnL() < 0 {
			cnt++
		}
	}

	return cnt
}

func (rp *RoundtripPerformance) divIntInt(numerator, denominator int) float64 {
	switch {
	case denominator == 0:
		return 0
	default:
		return 100 * float64(numerator) / float64(denominator)
	}
}

func (rp *RoundtripPerformance) divFloatInt(numerator float64, denominator int) float64 {
	switch {
	case denominator == 0:
		return 0
	default:
		return numerator / float64(denominator)
	}
}

func (rp *RoundtripPerformance) divFloatFloat(numerator, denominator float64) float64 {
	switch {
	case denominator == 0:
		return 0
	default:
		return numerator / denominator
	}
}

func (rp *RoundtripPerformance) divDurInt(numerator time.Duration, denominator int) time.Duration {
	switch {
	case denominator == 0:
		return time.Duration(0)
	default:
		return time.Duration(int64(numerator) / int64(denominator))
	}
}

func (rp *RoundtripPerformance) totalPnL() float64 {
	pnl := 0.

	for _, r := range rp.roundtrips {
		pnl += r.PnL()
	}

	return pnl
}

func (rp *RoundtripPerformance) netTotalPnL() float64 {
	pnl := 0.

	for _, r := range rp.roundtrips {
		pnl += r.NetPnL()
	}

	return pnl
}

func (rp *RoundtripPerformance) winningPnL() float64 {
	pnl := 0.

	for _, r := range rp.roundtrips {
		v := r.PnL()
		if v > 0 {
			pnl += v
		}
	}

	return pnl
}

func (rp *RoundtripPerformance) netWinningPnL() float64 {
	pnl := 0.

	for _, r := range rp.roundtrips {
		v := r.NetPnL()
		if v > 0 {
			pnl += v
		}
	}

	return pnl
}

func (rp *RoundtripPerformance) loosingPnL() float64 {
	pnl := 0.

	for _, r := range rp.roundtrips {
		v := r.PnL()
		if v < 0 {
			pnl += v
		}
	}

	return pnl
}

func (rp *RoundtripPerformance) netLoosingPnL() float64 {
	pnl := 0.

	for _, r := range rp.roundtrips {
		v := r.NetPnL()
		if v < 0 {
			pnl += v
		}
	}

	return pnl
}

func (rp *RoundtripPerformance) avgDuration() time.Duration {
	var d time.Duration

	cnt := 0

	for _, r := range rp.roundtrips {
		d += r.Duration()
		cnt++
	}

	return rp.divDurInt(d, cnt)
}

func (rp *RoundtripPerformance) avgWinningDuration() time.Duration {
	var d time.Duration

	cnt := 0

	for _, r := range rp.roundtrips {
		if r.PnL() > 0 {
			d += r.Duration()
			cnt++
		}
	}

	return rp.divDurInt(d, cnt)
}

func (rp *RoundtripPerformance) netAvgWinningDuration() time.Duration {
	var d time.Duration

	cnt := 0

	for _, r := range rp.roundtrips {
		if r.NetPnL() > 0 {
			d += r.Duration()
			cnt++
		}
	}

	return rp.divDurInt(d, cnt)
}

func (rp *RoundtripPerformance) avgLoosingDuration() time.Duration {
	var d time.Duration

	cnt := 0

	for _, r := range rp.roundtrips {
		if r.PnL() < 0 {
			d += r.Duration()
			cnt++
		}
	}

	return rp.divDurInt(d, cnt)
}

func (rp *RoundtripPerformance) netAvgLoosingDuration() time.Duration {
	var d time.Duration

	cnt := 0

	for _, r := range rp.roundtrips {
		if r.NetPnL() < 0 {
			d += r.Duration()
			cnt++
		}
	}

	return rp.divDurInt(d, cnt)
}

func (rp *RoundtripPerformance) minDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		d := r.Duration()

		switch {
		case m == 0, m > d:
			m = d
		}
	}

	return m
}

func (rp *RoundtripPerformance) minWinningDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		if r.PnL() > 0 {
			d := r.Duration()

			switch {
			case m == 0, m > d:
				m = d
			}
		}
	}

	return m
}

func (rp *RoundtripPerformance) netMinWinningDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		if r.NetPnL() > 0 {
			d := r.Duration()

			switch {
			case m == 0, m > d:
				m = d
			}
		}
	}

	return m
}

func (rp *RoundtripPerformance) minLoosingDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		if r.PnL() < 0 {
			d := r.Duration()

			switch {
			case m == 0, m > d:
				m = d
			}
		}
	}

	return m
}

func (rp *RoundtripPerformance) netMinLoosingDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		if r.NetPnL() < 0 {
			d := r.Duration()

			switch {
			case m == 0, m > d:
				m = d
			}
		}
	}

	return m
}

func (rp *RoundtripPerformance) maxDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		d := r.Duration()

		switch {
		case m == 0, m < d:
			m = d
		}
	}

	return m
}

func (rp *RoundtripPerformance) maxWinningDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		if r.PnL() > 0 {
			d := r.Duration()

			switch {
			case m == 0, m < d:
				m = d
			}
		}
	}

	return m
}

func (rp *RoundtripPerformance) netMaxWinningDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		if r.NetPnL() > 0 {
			d := r.Duration()

			switch {
			case m == 0, m < d:
				m = d
			}
		}
	}

	return m
}

func (rp *RoundtripPerformance) maxLoosingDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		if r.PnL() < 0 {
			d := r.Duration()

			switch {
			case m == 0, m < d:
				m = d
			}
		}
	}

	return m
}

func (rp *RoundtripPerformance) netMaxLoosingDuration() time.Duration {
	var m time.Duration

	for _, r := range rp.roundtrips {
		if r.NetPnL() < 0 {
			d := r.Duration()

			switch {
			case m == 0, m < d:
				m = d
			}
		}
	}

	return m
}

func (rp *RoundtripPerformance) avgMAE() float64 {
	v := 0.

	for _, r := range rp.roundtrips {
		v += r.MaximumAdverseExcursion()
	}

	return rp.divFloatFloat(v, float64(rp.totalCount()))
}

func (rp *RoundtripPerformance) avgMFE() float64 {
	v := 0.

	for _, r := range rp.roundtrips {
		v += r.MaximumFavorableExcursion()
	}

	return rp.divFloatFloat(v, float64(rp.totalCount()))
}

func (rp *RoundtripPerformance) avgEntryEfficiency() float64 {
	v := 0.

	for _, r := range rp.roundtrips {
		v += r.EntryEfficiency()
	}

	return rp.divFloatFloat(v, float64(rp.totalCount()))
}

func (rp *RoundtripPerformance) avgExitEfficiency() float64 {
	v := 0.

	for _, r := range rp.roundtrips {
		v += r.ExitEfficiency()
	}

	return rp.divFloatFloat(v, float64(rp.totalCount()))
}

func (rp *RoundtripPerformance) avgTotalEfficiency() float64 {
	v := 0.

	for _, r := range rp.roundtrips {
		v += r.TotalEfficiency()
	}

	return rp.divFloatFloat(v, float64(rp.totalCount()))
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
