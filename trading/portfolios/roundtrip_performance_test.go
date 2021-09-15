//nolint:testpackage
package portfolios

//nolint:gci
import (
	"math"
	"mbg/trading/instruments"
	"mbg/trading/portfolios/positions/sides"
	"testing"
	"time"
)

//nolint:funlen,gocognit,gocyclo
func TestRoundtripPerformanceEmpty(t *testing.T) {
	t.Parallel()

	const (
		fmtVal = "%v(): expected 0, actual %v"
	)

	rp := NewRoundtripPerformance()

	if rp.TotalCount() != 0 {
		t.Errorf(fmtVal, "TotalCount", rp.TotalCount())
	}

	if rp.WinningCount() != 0 {
		t.Errorf(fmtVal, "WinningCount", rp.WinningCount())
	}

	if rp.NetWinningCount() != 0 {
		t.Errorf(fmtVal, "NetWinningCount", rp.NetWinningCount())
	}

	if rp.LoosingCount() != 0 {
		t.Errorf(fmtVal, "LoosingCount", rp.LoosingCount())
	}

	if rp.NetLoosingCount() != 0 {
		t.Errorf(fmtVal, "NetLoosingCount", rp.NetLoosingCount())
	}

	if rp.WinningPct() != 0 {
		t.Errorf(fmtVal, "WinningPct", rp.WinningPct())
	}

	if rp.NetWinningPct() != 0 {
		t.Errorf(fmtVal, "NetWinningPct", rp.NetWinningPct())
	}

	if rp.LoosingPct() != 0 {
		t.Errorf(fmtVal, "LoosingPct", rp.LoosingPct())
	}

	if rp.NetLoosingPct() != 0 {
		t.Errorf(fmtVal, "NetLoosingPct", rp.NetLoosingPct())
	}

	if rp.TotalPnL() != 0 {
		t.Errorf(fmtVal, "TotalPnL", rp.TotalPnL())
	}

	if rp.NetTotalPnL() != 0 {
		t.Errorf(fmtVal, "NetTotalPnL", rp.NetTotalPnL())
	}

	if rp.WinningPnL() != 0 {
		t.Errorf(fmtVal, "WinningPnL", rp.WinningPnL())
	}

	if rp.NetWinningPnL() != 0 {
		t.Errorf(fmtVal, "NetWinningPnL", rp.NetWinningPnL())
	}

	if rp.LoosingPnL() != 0 {
		t.Errorf(fmtVal, "LoosingPnL", rp.LoosingPnL())
	}

	if rp.NetLoosingPnL() != 0 {
		t.Errorf(fmtVal, "NetLoosingPnL", rp.NetLoosingPnL())
	}

	if rp.AvgTotalPnL() != 0 {
		t.Errorf(fmtVal, "AvgTotalPnL", rp.AvgTotalPnL())
	}

	if rp.NetAvgTotalPnL() != 0 {
		t.Errorf(fmtVal, "NetAvgTotalPnL", rp.NetAvgTotalPnL())
	}

	if rp.AvgWinningPnL() != 0 {
		t.Errorf(fmtVal, "AvgWinningPnL", rp.AvgWinningPnL())
	}

	if rp.NetAvgWinningPnL() != 0 {
		t.Errorf(fmtVal, "NetAvgWinningPnL", rp.NetAvgWinningPnL())
	}

	if rp.AvgLoosingPnL() != 0 {
		t.Errorf(fmtVal, "AvgLoosingPnL", rp.AvgLoosingPnL())
	}

	if rp.NetAvgLoosingPnL() != 0 {
		t.Errorf(fmtVal, "NetAvgLoosingPnL", rp.NetAvgLoosingPnL())
	}

	if rp.AvgWinningLoosingPct() != 0 {
		t.Errorf(fmtVal, "AvgWinningLoosingPct", rp.AvgWinningLoosingPct())
	}

	if rp.NetAvgWinningLoosingPct() != 0 {
		t.Errorf(fmtVal, "NetAvgWinningLoosingPct", rp.NetAvgWinningLoosingPct())
	}

	if rp.ProfitPct() != 0 {
		t.Errorf(fmtVal, "ProfitPct", rp.ProfitPct())
	}

	if rp.NetProfitPct() != 0 {
		t.Errorf(fmtVal, "NetProfitPct", rp.NetProfitPct())
	}

	if rp.MaxConsecutiveWinners() != 0 {
		t.Errorf(fmtVal, "MaxConsecutiveWinners", rp.MaxConsecutiveWinners())
	}

	if rp.NetMaxConsecutiveWinners() != 0 {
		t.Errorf(fmtVal, "NetMaxConsecutiveWinners", rp.NetMaxConsecutiveWinners())
	}

	if rp.MaxConsecutiveLoosers() != 0 {
		t.Errorf(fmtVal, "MaxConsecutiveLoosers", rp.MaxConsecutiveLoosers())
	}

	if rp.NetMaxConsecutiveLoosers() != 0 {
		t.Errorf(fmtVal, "NetMaxConsecutiveLoosers", rp.NetMaxConsecutiveLoosers())
	}

	if rp.AvgDuration() != 0 {
		t.Errorf(fmtVal, "AvgDuration", rp.AvgDuration())
	}

	if rp.AvgWinningDuration() != 0 {
		t.Errorf(fmtVal, "AvgWinningDuration", rp.AvgWinningDuration())
	}

	if rp.NetAvgWinningDuration() != 0 {
		t.Errorf(fmtVal, "NetAvgWinningDuration", rp.NetAvgWinningDuration())
	}

	if rp.AvgLoosingDuration() != 0 {
		t.Errorf(fmtVal, "AvgLoosingDuration", rp.AvgLoosingDuration())
	}

	if rp.NetAvgLoosingDuration() != 0 {
		t.Errorf(fmtVal, "NetAvgLoosingDuration", rp.NetAvgLoosingDuration())
	}

	if rp.MinDuration() != 0 {
		t.Errorf(fmtVal, "MinDuration", rp.MinDuration())
	}

	if rp.MinWinningDuration() != 0 {
		t.Errorf(fmtVal, "MinWinningDuration", rp.MinWinningDuration())
	}

	if rp.NetMinWinningDuration() != 0 {
		t.Errorf(fmtVal, "NetMinWinningDuration", rp.NetMinWinningDuration())
	}

	if rp.MinLoosingDuration() != 0 {
		t.Errorf(fmtVal, "MinLoosingDuration", rp.MinLoosingDuration())
	}

	if rp.NetMinLoosingDuration() != 0 {
		t.Errorf(fmtVal, "NetMinLoosingDuration", rp.NetMinLoosingDuration())
	}

	if rp.MaxDuration() != 0 {
		t.Errorf(fmtVal, "MaxDuration", rp.MaxDuration())
	}

	if rp.MaxWinningDuration() != 0 {
		t.Errorf(fmtVal, "MaxWinningDuration", rp.MaxWinningDuration())
	}

	if rp.NetMaxWinningDuration() != 0 {
		t.Errorf(fmtVal, "NetMaxWinningDuration", rp.NetMaxWinningDuration())
	}

	if rp.MaxLoosingDuration() != 0 {
		t.Errorf(fmtVal, "MaxLoosingDuration", rp.MaxLoosingDuration())
	}

	if rp.NetMaxLoosingDuration() != 0 {
		t.Errorf(fmtVal, "NetMaxLoosingDuration", rp.NetMaxLoosingDuration())
	}

	if rp.AvgMAE() != 0 {
		t.Errorf(fmtVal, "AvgMAE", rp.AvgMAE())
	}

	if rp.AvgMFE() != 0 {
		t.Errorf(fmtVal, "AvgMFE", rp.AvgMFE())
	}

	if rp.AvgEntryEfficiency() != 0 {
		t.Errorf(fmtVal, "AvgEntryEfficiency", rp.AvgEntryEfficiency())
	}

	if rp.AvgExitEfficiency() != 0 {
		t.Errorf(fmtVal, "AvgExitEfficiency", rp.AvgExitEfficiency())
	}

	if rp.AvgTotalEfficiency() != 0 {
		t.Errorf(fmtVal, "AvgTotalEfficiency", rp.AvgTotalEfficiency())
	}
}

//nolint:funlen,gocognit,gocyclo
func TestRoundtripPerformance(t *testing.T) {
	t.Parallel()

	const (
		fmtVal                   = "%v(): expected %v, actual %v"
		equalityThreshold        = 1e-13
		totalCount               = 6
		winningCount             = 5
		loosingCount             = 1
		winningPct               = (100. * winningCount) / totalCount
		loosingPct               = (100. * loosingCount) / totalCount
		netWinningCount          = 4
		netLoosingCount          = 2
		netWinningPct            = (100. * netWinningCount) / totalCount
		netLoosingPct            = (100. * netLoosingCount) / totalCount
		totalPnL                 = 159.
		winningPnL               = 229.
		loosingPnL               = -70.
		netTotalPnL              = 99.
		netWinningPnL            = 180.
		netLoosingPnL            = -81.
		avgTotalPnL              = totalPnL / totalCount
		avgWinningPnL            = winningPnL / winningCount
		avgLoosingPnL            = loosingPnL / loosingCount
		netAvgTotalPnL           = netTotalPnL / totalCount
		netAvgWinningPnL         = netWinningPnL / netWinningCount
		netAvgLoosingPnL         = netLoosingPnL / netLoosingCount
		avgWinningLoosingPct     = (100. * avgWinningPnL) / -avgLoosingPnL
		netAvgWinningLoosingPct  = (100. * netAvgWinningPnL) / -netAvgLoosingPnL
		profitPct                = (100. * winningPnL) / -loosingPnL
		netProfitPct             = (100. * netWinningPnL) / -netLoosingPnL
		maxConsecutiveWinners    = 5
		maxConsecutiveLoosers    = 1
		netMaxConsecutiveWinners = 4
		netMaxConsecutiveLoosers = 2
		avgDuration              = time.Duration(int64(time.Duration(time.Hour*24*52)) / int64(totalCount))
		avgWinningDuration       = time.Duration(int64(time.Duration(time.Hour*24*47)) / int64(winningCount))
		avgLoosingDuration       = time.Duration(int64(time.Duration(time.Hour*24*5)) / int64(loosingCount))
		netAvgWinningDuration    = time.Duration(int64(time.Duration(time.Hour*24*36)) / int64(netWinningCount))
		netAvgLoosingDuration    = time.Duration(int64(time.Duration(time.Hour*24*16)) / int64(netLoosingCount))
		minDuration              = time.Duration(time.Hour * 24)
		minWinningDuration       = time.Duration(time.Hour * 24)
		minLoosingDuration       = time.Duration(time.Hour * 24 * 5)
		netMinWinningDuration    = time.Duration(time.Hour * 24)
		netMinLoosingDuration    = time.Duration(time.Hour * 24 * 5)
		maxDuration              = time.Duration(time.Hour * 24 * 19)
		maxWinningDuration       = time.Duration(time.Hour * 24 * 19)
		maxLoosingDuration       = time.Duration(time.Hour * 24 * 5)
		netMaxWinningDuration    = time.Duration(time.Hour * 24 * 19)
		netMaxLoosingDuration    = time.Duration(time.Hour * 24 * 11)
		avgMAE                   = 7.61072912235703
		avgMFE                   = 7.841795144426725
		avgEntryEfficiency       = 70.97883597883598
		avgExitEfficiency        = 70.97883597883598
		avgTotalEfficiency       = 41.95767195767196
	)

	notEqual := func(a, b float64) bool {
		return math.Abs(a-b) > equalityThreshold
	}

	t0 := time.Now()
	tests := []struct {
		side    sides.Side
		qty     float64
		enTime  time.Time
		enPrice float64
		exTime  time.Time
		exPrice float64
		comm    float64
		high    float64
		low     float64
	}{
		{
			// PnL: 100, NetPnL: 90, Duration: 19, MAE: 10, MFE: 6.66666666666667,
			// EnE: 85.71428571428571, ExE: 85.71428571428571, TE: 71.42857142857143
			side: sides.Long, qty: 10, comm: 10, high: 32, low: 18,
			enTime: t0.AddDate(0, 0, 1), enPrice: 20, exTime: t0.AddDate(0, 0, 20), exPrice: 30,
		},
		{
			// PnL: 30, NetPnL: 20, Duration: 1, MAE: 6.66666666666667, MFE: 5.55555555555556,
			// EnE: 80, ExE: 80, TE: 60
			side: sides.Long, qty: 5, comm: 10, high: 38, low: 28,
			enTime: t0.AddDate(0, 1, 1), enPrice: 30, exTime: t0.AddDate(0, 1, 2), exPrice: 36,
		},
		{
			// PnL: 40, NetPnL: 30, Duration: 9, MAE: 2.27272727272727, MFE: 2.5,
			// EnE: 83.33333333333333, ExE: 83.33333333333333, TE: 66.66666666666667
			side: sides.Short, qty: 10, comm: 10, high: 45, low: 39,
			enTime: t0.AddDate(0, 2, 1), enPrice: 44, exTime: t0.AddDate(0, 2, 10), exPrice: 40,
		},
		{
			// PnL: 50, NetPnL: 40, Duration: 7, MAE: 2.32558139534884, MFE: 2.63157894736842,
			// EnE: 85.71428571428571, ExE: 85.71428571428571, TE: 71.42857142857143
			side: sides.Short, qty: 10, comm: 10, high: 44, low: 37,
			enTime: t0.AddDate(0, 3, 1), enPrice: 43, exTime: t0.AddDate(0, 3, 8), exPrice: 38,
		},
		{
			// PnL: 9, NetPnL: -1, Duration: 11, MAE: 2.77777777777778, MFE: 3.03030303030303,
			// EnE: 80, ExE: 80, TE: 60
			side: sides.Short, qty: 3, comm: 10, high: 37, low: 32,
			enTime: t0.AddDate(0, 4, 1), enPrice: 36, exTime: t0.AddDate(0, 4, 12), exPrice: 33,
		},
		{
			// PnL: -70, NetPnL: -80, Duration: 5, MAE: 21.62162162162162, MFE: 26.66666666666667,
			// EnE: 11.11111111111111, ExE: 11.11111111111111, TE: -77.77777777777777
			side: sides.Long, qty: 10, comm: 10, high: 38, low: 29,
			enTime: t0.AddDate(0, 5, 1), enPrice: 37, exTime: t0.AddDate(0, 5, 6), exPrice: 30,
		},
	}

	rp := NewRoundtripPerformance()

	for _, tt := range tests {
		var pnl float64
		if tt.side == sides.Long {
			pnl = tt.qty * (tt.exPrice - tt.enPrice)
		} else {
			pnl = tt.qty * (tt.enPrice - tt.exPrice)
		}

		r := Roundtrip{
			instrument: instruments.Instrument{PriceFactor: 1},
			side:       tt.side,
			quantity:   tt.qty,
			entryTime:  tt.enTime,
			entryPrice: tt.enPrice,
			exitTime:   tt.exTime,
			exitPrice:  tt.exPrice,
			highPrice:  tt.high,
			lowPrice:   tt.low,
			commission: tt.comm,
			pnl:        pnl,
		}

		rp.Add(r)
	}

	if rp.TotalCount() != totalCount {
		t.Errorf(fmtVal, "TotalCount", totalCount, rp.TotalCount())
	}

	if rp.WinningCount() != winningCount {
		t.Errorf(fmtVal, "WinningCount", winningCount, rp.WinningCount())
	}

	if rp.NetWinningCount() != netWinningCount {
		t.Errorf(fmtVal, "NetWinningCount", netWinningCount, rp.NetWinningCount())
	}

	if rp.LoosingCount() != loosingCount {
		t.Errorf(fmtVal, "LoosingCount", loosingCount, rp.LoosingCount())
	}

	if rp.NetLoosingCount() != netLoosingCount {
		t.Errorf(fmtVal, "NetLoosingCount", netLoosingCount, rp.NetLoosingCount())
	}

	if notEqual(rp.WinningPct(), winningPct) {
		t.Errorf(fmtVal, "WinningPct", winningPct, rp.WinningPct())
	}

	if notEqual(rp.NetWinningPct(), netWinningPct) {
		t.Errorf(fmtVal, "NetWinningPct", netWinningPct, rp.NetWinningPct())
	}

	if notEqual(rp.LoosingPct(), loosingPct) {
		t.Errorf(fmtVal, "LoosingPct", loosingPct, rp.LoosingPct())
	}

	if notEqual(rp.NetLoosingPct(), netLoosingPct) {
		t.Errorf(fmtVal, "NetLoosingPct", netLoosingPct, rp.NetLoosingPct())
	}

	if notEqual(rp.TotalPnL(), totalPnL) {
		t.Errorf(fmtVal, "TotalPnL", totalPnL, rp.TotalPnL())
	}

	if notEqual(rp.NetTotalPnL(), netTotalPnL) {
		t.Errorf(fmtVal, "NetTotalPnL", netTotalPnL, rp.NetTotalPnL())
	}

	if notEqual(rp.WinningPnL(), winningPnL) {
		t.Errorf(fmtVal, "WinningPnL", winningPnL, rp.WinningPnL())
	}

	if notEqual(rp.NetWinningPnL(), netWinningPnL) {
		t.Errorf(fmtVal, "NetWinningPnL", netWinningPnL, rp.NetWinningPnL())
	}

	if notEqual(rp.LoosingPnL(), loosingPnL) {
		t.Errorf(fmtVal, "LoosingPnL", loosingPnL, rp.LoosingPnL())
	}

	if notEqual(rp.NetLoosingPnL(), netLoosingPnL) {
		t.Errorf(fmtVal, "NetLoosingPnL", netLoosingPnL, rp.NetLoosingPnL())
	}

	if notEqual(rp.AvgTotalPnL(), avgTotalPnL) {
		t.Errorf(fmtVal, "AvgTotalPnL", avgTotalPnL, rp.AvgTotalPnL())
	}

	if notEqual(rp.NetAvgTotalPnL(), netAvgTotalPnL) {
		t.Errorf(fmtVal, "NetAvgTotalPnL", netAvgTotalPnL, rp.NetAvgTotalPnL())
	}

	if notEqual(rp.AvgWinningPnL(), avgWinningPnL) {
		t.Errorf(fmtVal, "AvgWinningPnL", avgWinningPnL, rp.AvgWinningPnL())
	}

	if notEqual(rp.NetAvgWinningPnL(), netAvgWinningPnL) {
		t.Errorf(fmtVal, "NetAvgWinningPnL", netAvgWinningPnL, rp.NetAvgWinningPnL())
	}

	if notEqual(rp.AvgLoosingPnL(), avgLoosingPnL) {
		t.Errorf(fmtVal, "AvgLoosingPnL", avgLoosingPnL, rp.AvgLoosingPnL())
	}

	if notEqual(rp.NetAvgLoosingPnL(), netAvgLoosingPnL) {
		t.Errorf(fmtVal, "NetAvgLoosingPnL", netAvgLoosingPnL, rp.NetAvgLoosingPnL())
	}

	if notEqual(rp.AvgWinningLoosingPct(), avgWinningLoosingPct) {
		t.Errorf(fmtVal, "AvgWinningLoosingPct", avgWinningLoosingPct, rp.AvgWinningLoosingPct())
	}

	if notEqual(rp.NetAvgWinningLoosingPct(), netAvgWinningLoosingPct) {
		t.Errorf(fmtVal, "NetAvgWinningLoosingPct", netAvgWinningLoosingPct, rp.NetAvgWinningLoosingPct())
	}

	if notEqual(rp.ProfitPct(), profitPct) {
		t.Errorf(fmtVal, "ProfitPct", profitPct, rp.ProfitPct())
	}

	if notEqual(rp.NetProfitPct(), netProfitPct) {
		t.Errorf(fmtVal, "NetProfitPct", netProfitPct, rp.NetProfitPct())
	}

	if rp.MaxConsecutiveWinners() != maxConsecutiveWinners {
		t.Errorf(fmtVal, "MaxConsecutiveWinners", maxConsecutiveWinners, rp.MaxConsecutiveWinners())
	}

	if rp.NetMaxConsecutiveWinners() != netMaxConsecutiveWinners {
		t.Errorf(fmtVal, "NetMaxConsecutiveWinners", netMaxConsecutiveWinners, rp.NetMaxConsecutiveWinners())
	}

	if rp.MaxConsecutiveLoosers() != maxConsecutiveLoosers {
		t.Errorf(fmtVal, "MaxConsecutiveLoosers", maxConsecutiveLoosers, rp.MaxConsecutiveLoosers())
	}

	if rp.NetMaxConsecutiveLoosers() != netMaxConsecutiveLoosers {
		t.Errorf(fmtVal, "NetMaxConsecutiveLoosers", netMaxConsecutiveLoosers, rp.NetMaxConsecutiveLoosers())
	}

	if rp.AvgDuration() != avgDuration {
		t.Errorf(fmtVal, "AvgDuration", avgDuration, rp.AvgDuration())
	}

	if rp.AvgWinningDuration() != avgWinningDuration {
		t.Errorf(fmtVal, "AvgWinningDuration", avgWinningDuration, rp.AvgWinningDuration())
	}

	if rp.NetAvgWinningDuration() != netAvgWinningDuration {
		t.Errorf(fmtVal, "NetAvgWinningDuration", netAvgWinningDuration, rp.NetAvgWinningDuration())
	}

	if rp.AvgLoosingDuration() != avgLoosingDuration {
		t.Errorf(fmtVal, "AvgLoosingDuration", avgLoosingDuration, rp.AvgLoosingDuration())
	}

	if rp.NetAvgLoosingDuration() != netAvgLoosingDuration {
		t.Errorf(fmtVal, "NetAvgLoosingDuration", netAvgLoosingDuration, rp.NetAvgLoosingDuration())
	}

	if rp.MinDuration() != minDuration {
		t.Errorf(fmtVal, "MinDuration", minDuration, rp.MinDuration())
	}

	if rp.MinWinningDuration() != minWinningDuration {
		t.Errorf(fmtVal, "MinWinningDuration", minWinningDuration, rp.MinWinningDuration())
	}

	if rp.NetMinWinningDuration() != netMinWinningDuration {
		t.Errorf(fmtVal, "NetMinWinningDuration", netMinWinningDuration, rp.NetMinWinningDuration())
	}

	if rp.MinLoosingDuration() != minLoosingDuration {
		t.Errorf(fmtVal, "MinLoosingDuration", minLoosingDuration, rp.MinLoosingDuration())
	}

	if rp.NetMinLoosingDuration() != netMinLoosingDuration {
		t.Errorf(fmtVal, "NetMinLoosingDuration", netMinLoosingDuration, rp.NetMinLoosingDuration())
	}

	if rp.MaxDuration() != maxDuration {
		t.Errorf(fmtVal, "MaxDuration", maxDuration, rp.MaxDuration())
	}

	if rp.MaxWinningDuration() != maxWinningDuration {
		t.Errorf(fmtVal, "MaxWinningDuration", maxWinningDuration, rp.MaxWinningDuration())
	}

	if rp.NetMaxWinningDuration() != netMaxWinningDuration {
		t.Errorf(fmtVal, "NetMaxWinningDuration", netMaxWinningDuration, rp.NetMaxWinningDuration())
	}

	if rp.MaxLoosingDuration() != maxLoosingDuration {
		t.Errorf(fmtVal, "MaxLoosingDuration", maxLoosingDuration, rp.MaxLoosingDuration())
	}

	if rp.NetMaxLoosingDuration() != netMaxLoosingDuration {
		t.Errorf(fmtVal, "NetMaxLoosingDuration", netMaxLoosingDuration, rp.NetMaxLoosingDuration())
	}

	if notEqual(rp.AvgMAE(), avgMAE) {
		t.Errorf(fmtVal, "AvgMAE", avgMAE, rp.AvgMAE())
	}

	if notEqual(rp.AvgMFE(), avgMFE) {
		t.Errorf(fmtVal, "AvgMFE", avgMFE, rp.AvgMFE())
	}

	if notEqual(rp.AvgEntryEfficiency(), avgEntryEfficiency) {
		t.Errorf(fmtVal, "AvgEntryEfficiency", avgEntryEfficiency, rp.AvgEntryEfficiency())
	}

	if notEqual(rp.AvgExitEfficiency(), avgExitEfficiency) {
		t.Errorf(fmtVal, "AvgExitEfficiency", avgExitEfficiency, rp.AvgExitEfficiency())
	}

	if notEqual(rp.AvgTotalEfficiency(), avgTotalEfficiency) {
		t.Errorf(fmtVal, "AvgTotalEfficiency", avgTotalEfficiency, rp.AvgTotalEfficiency())
	}
}
