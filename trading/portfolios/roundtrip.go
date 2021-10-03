package portfolios

//nolint:gci
import (
	"math"
	"mbg/trading/instruments"
	"mbg/trading/portfolios/positions/sides"
	pos "mbg/trading/portfolios/positions/sides"
	"time"
)

// Roundtrip is a position round-trip.
type Roundtrip struct {
	instrument instruments.Instrument
	side       sides.Side
	quantity   float64
	entryTime  time.Time
	entryPrice float64
	exitTime   time.Time
	exitPrice  float64
	pnl        float64
	commission float64
	highPrice  float64
	lowPrice   float64
}

func newRoundtrip(instr instruments.Instrument, entry, exit *Execution, qty float64) *Roundtrip {
	var pnl, pf float64

	if instr.PriceFactor != 0 {
		pf = instr.PriceFactor
	} else {
		pf = 1
	}

	side := pos.Long
	if entry.side.IsSell() {
		side = pos.Short
		pnl = qty * (entry.price - exit.price) * pf
	} else {
		pnl = qty * (exit.price - entry.price) * pf
	}

	commission := (entry.commissionConverted/entry.quantity + exit.commissionConverted/exit.quantity) * qty

	return &Roundtrip{
		instrument: instr,
		side:       side,
		quantity:   qty,
		entryTime:  entry.reportTime,
		entryPrice: entry.price,
		exitTime:   exit.reportTime,
		exitPrice:  exit.price,
		highPrice:  math.Max(entry.unrealizedPriceHigh, exit.unrealizedPriceHigh),
		lowPrice:   math.Min(entry.unrealizedPriceLow, exit.unrealizedPriceLow),
		commission: commission,
		pnl:        pnl,
	}
}

// Duration returns a duration of this round-trip.
func (r *Roundtrip) Duration() time.Duration {
	return r.exitTime.Sub(r.entryTime)
}

// Instrument is the traded instrument.
func (r *Roundtrip) Instrument() instruments.Instrument {
	return r.instrument
}

// Quantity is the total unsigned quantity of the position.
func (r *Roundtrip) Quantity() float64 {
	return r.quantity
}

// Side is the side of the roundtrip.
func (r *Roundtrip) Side() sides.Side {
	return r.side
}

// EntryTime is the date and time the position was opened.
func (r *Roundtrip) EntryTime() time.Time {
	return r.entryTime
}

// EntryPrice is the (average) price at which the position was opened.
func (r *Roundtrip) EntryPrice() float64 {
	return r.entryPrice
}

// ExitTime is the date and time the position was closed.
func (r *Roundtrip) ExitTime() time.Time {
	return r.exitTime
}

// ExitPrice is the (average) price at which the position was closed.
func (r *Roundtrip) ExitPrice() float64 {
	return r.exitPrice
}

// PnL is the gross profit and loss of the round-trip in instrument's currency.
func (r *Roundtrip) PnL() float64 {
	return r.pnl
}

// NetPnL is the net profit and loss of the round-trip in instrument's currency.
func (r *Roundtrip) NetPnL() float64 {
	return r.pnl - r.commission
}

// Commission is the total commission associated with the round-trip in instrument's currency.
// This is always a positive value.
func (r *Roundtrip) Commission() float64 {
	return r.commission
}

// HighestPrice is the highest price in instrument's currency during the round-trip.
func (r *Roundtrip) HighestPrice() float64 {
	return r.highPrice
}

// LowestPrice is the lowest price in instrument's currency during the round-trip.
func (r *Roundtrip) LowestPrice() float64 {
	return r.lowPrice
}

// MaximumAdversePrice is the Maximum Adverse price in instrument's currency during the round-trip.
func (r *Roundtrip) MaximumAdversePrice() float64 {
	switch r.side {
	case pos.Long:
		return r.lowPrice
	default:
		return r.highPrice
	}
}

// MaximumFavorablePrice is the Maximum Favorable price in instrument's currency during the round-trip.
func (r *Roundtrip) MaximumFavorablePrice() float64 {
	switch r.side {
	case pos.Long:
		return r.highPrice
	default:
		return r.lowPrice
	}
}

// MAE is the percentage of the Maximum Adverse Excursion (MAE)
// which measures the maximum potential loss per unit of quantity
// taken during the round-trip period.
//
// 0% is the perfect MAE, 100% and higher is the worst possible MAE.
func (r *Roundtrip) MaximumAdverseExcursion() float64 {
	switch r.side {
	case pos.Long:
		return 100 * (1 - r.lowPrice/r.entryPrice) //nolint:gomnd
	default:
		return 100 * (r.highPrice/r.entryPrice - 1) //nolint:gomnd
	}
}

// MFE is the percentage of the Maximum Favorable Excursion (MFE)
// which measures the peak potential profit per unit of quantity
// taken during the round-trip period.
//
// This statistical concept originally created by John Sweeney to measure
// the distinctive characteristics of profitable trades, can be used as part of
// an analytical process to distinguish between average trades and those
// that offer substantially greater profit potential.
//
// 0% is the perfect MFE, 100% and higher is the worst possible MFE.
func (r *Roundtrip) MaximumFavorableExcursion() float64 {
	switch r.side {
	case pos.Long:
		return 100 * (r.highPrice/r.exitPrice - 1) //nolint:gomnd
	default:
		return 100 * (1 - r.lowPrice/r.exitPrice) //nolint:gomnd
	}
}

// EntryEfficiency is the Entry Efficiency which measures the percentage in range [0. 100] of
// the total round-trip potential taken by a round-trip given its entry and assuming
// the best possible exit during the round-trip period.
//
// It shows how close the entry price was to the best possible entry price during the round-trip.
//
// 100% is the perfect efficiency, 0% is the worst possible efficiency.
func (r *Roundtrip) EntryEfficiency() float64 {
	delta := r.highPrice - r.lowPrice
	if delta != 0 {
		switch r.side {
		case pos.Long:
			return 100 * (r.highPrice - r.entryPrice) / delta
		case pos.Short:
			return 100 * (r.entryPrice - r.lowPrice) / delta
		}
	}

	// Undefined, but we return a fixed value.
	return 0
}

// ExitEfficiency is the Exit Efficiency which measures the percentage in range [0. 100] of
// the total round-trip potential taken by a round-trip given its exit and assuming
// the best possible entry during the round-trip period.
//
// It shows how close the exit price was to the best possible exit price during the round-trip.
//
// 100% is the perfect efficiency, 0% is the worst possible efficiency.
func (r *Roundtrip) ExitEfficiency() float64 {
	delta := r.highPrice - r.lowPrice
	if delta != 0 {
		switch r.side {
		case pos.Long:
			return 100 * (r.exitPrice - r.lowPrice) / delta
		case pos.Short:
			return 100 * (r.highPrice - r.exitPrice) / delta
		}
	}

	// Undefined, but we return a fixed value.
	return 0
}

// TotalEfficiency is the Total Efficiency which measures the percentage in range [0. 100] of
// the total round-trip potential taken by a round-trip during the round-trip period.
//
// It shows how close the entry and exit prices were to the best possible entry and exit prices
// during the round-trip,
// or the ability to capture the maximum profit potential during the round-trip period.
//
// 100% is the perfect efficiency, 0% is the worst possible efficiency.
func (r *Roundtrip) TotalEfficiency() float64 {
	delta := r.highPrice - r.lowPrice
	if delta != 0 {
		switch r.side {
		case pos.Long:
			return 100 * (r.exitPrice - r.entryPrice) / delta
		case pos.Short:
			return 100 * (r.entryPrice - r.exitPrice) / delta
		}
	}

	// Undefined, but we return a fixed value.
	return 0
}
