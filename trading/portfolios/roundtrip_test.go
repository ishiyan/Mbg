//nolint:testpackage
package portfolios

//nolint:gci
import (
	"math"
	"mbg/trading/instruments"
	"mbg/trading/orders/sides"
	pos "mbg/trading/portfolios/positions/sides"
	"testing"
	"time"
)

//nolint:funlen,gocognit
func TestRoundtrip(t *testing.T) {
	t.Parallel()

	const (
		fmtVal            = "%v(): expected %v, actual %v"
		equalityThreshold = 1e-13
	)

	notEqual := func(a, b float64) bool {
		return math.Abs(a-b) > equalityThreshold
	}

	t0 := time.Now()
	tests := []struct {
		factor  float64
		enSide  sides.Side
		enTime  time.Time
		enPrice float64
		enQty   float64
		enComm  float64
		enHigh  float64
		enLow   float64
		exTime  time.Time
		exPrice float64
		exQty   float64
		exComm  float64
		exHigh  float64
		exLow   float64
		qty     float64
		side    pos.Side
		pnl     float64
		npnl    float64
		comm    float64
		high    float64
		low     float64
		dur     time.Duration
		maep    float64
		mfep    float64
		mae     float64
		mfe     float64
		enef    float64
		exef    float64
		toef    float64
	}{
		{
			factor: 0, enSide: sides.Sell,
			enTime: t0.AddDate(0, 0, 1), enPrice: 20, enQty: 5, enComm: 10, enHigh: 24, enLow: 14,
			exTime: t0.AddDate(0, 0, 2), exPrice: 16, exQty: 8, exComm: 16, exHigh: 16, exLow: 16, qty: 5,
			side: pos.Short, pnl: 20, npnl: 0, comm: 20, high: 24, low: 14, dur: time.Duration(time.Hour * 24),
			maep: 24, mfep: 14, mae: 20, mfe: 12.5, enef: 60, exef: 80, toef: 40,
		},
		{
			factor: 2, enSide: sides.Buy,
			enTime: t0.AddDate(0, 0, 1), enPrice: 16, enQty: 8, enComm: 16, enHigh: 24, enLow: 14,
			exTime: t0.AddDate(0, 0, 2), exPrice: 20, exQty: 5, exComm: 10, exHigh: 20, exLow: 20, qty: 5,
			side: pos.Long, pnl: 40, npnl: 20, comm: 20, high: 24, low: 14, dur: time.Duration(time.Hour * 24),
			maep: 14, mfep: 24, mae: 12.5, mfe: 20, enef: 80, exef: 60, toef: 40,
		},
		{
			factor: 1, enSide: sides.Buy,
			enTime: t0.AddDate(0, 0, 1), enPrice: 20, enQty: 5, enComm: 10, enHigh: 20, enLow: 20,
			exTime: t0.AddDate(0, 0, 2), exPrice: 20, exQty: 8, exComm: 16, exHigh: 20, exLow: 20, qty: 5,
			side: pos.Long, pnl: 0, npnl: -20, comm: 20, high: 20, low: 20, dur: time.Duration(time.Hour * 24),
			maep: 20, mfep: 20, mae: 0, mfe: 0, enef: 0, exef: 0, toef: 0,
		},
	}

	for _, tt := range tests {
		instr := instruments.Instrument{PriceFactor: tt.factor}
		en := &Execution{
			reportTime:          tt.enTime,
			side:                tt.enSide,
			quantity:            tt.enQty,
			price:               tt.enPrice,
			commissionConverted: tt.enComm,
			unrealizedPriceHigh: tt.enHigh,
			unrealizedPriceLow:  tt.enLow,
		}
		ex := &Execution{
			reportTime:          tt.exTime,
			quantity:            tt.exQty,
			price:               tt.exPrice,
			commissionConverted: tt.exComm,
			unrealizedPriceHigh: tt.exHigh,
			unrealizedPriceLow:  tt.exLow,
		}

		r := newRoundtrip(instr, en, ex, tt.qty)

		if r.Instrument() != instr {
			t.Errorf(fmtVal, "Instrument", instr, r.Instrument())
		}

		if r.Duration() != tt.dur {
			t.Errorf(fmtVal, "Duration", tt.dur, r.Duration())
		}

		if notEqual(r.Quantity(), tt.qty) {
			t.Errorf(fmtVal, "Quantity", tt.qty, r.Quantity())
		}

		if r.Side() != tt.side {
			t.Errorf(fmtVal, "Side", tt.side, r.Side())
		}

		if r.EntryTime() != tt.enTime {
			t.Errorf(fmtVal, "EntryTime", tt.enTime, r.EntryTime())
		}

		if notEqual(r.EntryPrice(), tt.enPrice) {
			t.Errorf(fmtVal, "EntryPrice", tt.enPrice, r.EntryPrice())
		}

		if r.ExitTime() != tt.exTime {
			t.Errorf(fmtVal, "ExitTime", tt.exTime, r.ExitTime())
		}

		if notEqual(r.ExitPrice(), tt.exPrice) {
			t.Errorf(fmtVal, "ExitPrice", tt.exPrice, r.ExitPrice())
		}

		if notEqual(r.PnL(), tt.pnl) {
			t.Errorf(fmtVal, "PnL", tt.pnl, r.PnL())
		}

		if notEqual(r.NetPnL(), tt.npnl) {
			t.Errorf(fmtVal, "NetPnL", tt.npnl, r.NetPnL())
		}

		if notEqual(r.Commission(), tt.comm) {
			t.Errorf(fmtVal, "Commission", tt.comm, r.Commission())
		}

		if notEqual(r.HighestPrice(), tt.high) {
			t.Errorf(fmtVal, "HighestPrice", tt.high, r.HighestPrice())
		}

		if notEqual(r.LowestPrice(), tt.low) {
			t.Errorf(fmtVal, "LowestPrice", tt.low, r.LowestPrice())
		}

		if notEqual(r.MaximumAdversePrice(), tt.maep) {
			t.Errorf(fmtVal, "MaximumAdversePrice", tt.maep, r.MaximumAdversePrice())
		}

		if notEqual(r.MaximumFavorablePrice(), tt.mfep) {
			t.Errorf(fmtVal, "MaximumFavorablePrice", tt.mfep, r.MaximumFavorablePrice())
		}

		if notEqual(r.MaximumAdverseExcursion(), tt.mae) {
			t.Errorf(fmtVal, "MaximumAdverseExcursion", tt.mae, r.MaximumAdverseExcursion())
		}

		if notEqual(r.MaximumFavorableExcursion(), tt.mfe) {
			t.Errorf(fmtVal, "MaximumFavorableExcursion", tt.mfe, r.MaximumFavorableExcursion())
		}

		if notEqual(r.EntryEfficiency(), tt.enef) {
			t.Errorf(fmtVal, "EntryEfficiency", tt.enef, r.EntryEfficiency())
		}

		if notEqual(r.ExitEfficiency(), tt.exef) {
			t.Errorf(fmtVal, "ExitEfficiency", tt.exef, r.ExitEfficiency())
		}

		if notEqual(r.TotalEfficiency(), tt.toef) {
			t.Errorf(fmtVal, "TotalEfficiency", tt.toef, r.TotalEfficiency())
		}
	}
}
