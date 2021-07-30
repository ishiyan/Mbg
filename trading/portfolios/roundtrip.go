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
	maePrice   float64
	mfePrice   float64
	mae        float64
	mfe        float64
	entryEff   float64
	exitEff    float64
}

func newRoundtrip(instrument instruments.Instrument, entry, exit *Execution, qty float64) *Roundtrip {
	side := pos.Long
	if entry.side.IsSell() {
		side = pos.Short
	}

	high := math.Max(entry.roundtripPriceHigh, exit.roundtripPriceHigh)
	low := math.Min(entry.roundtripPriceLow, exit.roundtripPriceLow)

	var mae, mfe float64
	var entryEff, exitEff, totalEff float64

	switch side {
	case pos.Short:
		mae = high/entry.price - 1
		if high != low {
			entryEfficiency = (entry.price - low) / (high - low)
			exitEfficiency = (high - exit.price) / (high - low)
			totalEfficiency = (entry.price - exit.price) / (high - low)
		}
	case pos.Long:
		mae = 1 - low/entry.price
		if high != low {
			entryEfficiency = (high - entry.price) / (high - low)
			exitEfficiency = (exit.price - low) / (high - low)
			totalEfficiency = (exit.price - entry.price) / (high - low)
		}
	}

	return &Roundtrip{
		instrument: instrument,
		side:       side,
		quantity:   qty,
		entryTime:  entry.reportTime,
		entryPrice: entry.price,
		exitTime:   exit.reportTime,
		exitPrice:  exit.price,
		maePrice:   low,
		mfePrice:   high,
	}
}

// Duration returns a duration of this round-trip.
func (r *Roundtrip) Duration() time.Duration {
	return r.exitTime.Sub(r.entryTime)
}

// ExitDrawdown is the amount of profit given back before the position was closed.
func (r *Roundtrip) ExitDrawdown() float64 {
	return r.pnl - r.mfe
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
	return r.entryTime
}

// ExitPrice is the (average) price at which the position was closed.
func (r *Roundtrip) ExitPrice() float64 {
	return r.exitPrice
}

// PnL is the gross profit and loss of the round-trip in instrument's currency.
func (r *Roundtrip) PnL() float64 {
	return r.pnl
}

// Commission is the total commission associated with the round-trip in instrument's currency.
// This is always a positive value.
func (r *Roundtrip) Commission() float64 {
	return r.commission
}

// MaximumAdversePrice is the Maximum Adverse price in instrument's currency during the round-trip.
func (r *Roundtrip) MaximumAdversePrice() float64 {
	return r.maePrice
}

// MaximumFavorablePrice is the Maximum Favorable price in instrument's currency during the round-trip.
func (r *Roundtrip) MaximumFavorablePrice() float64 {
	return r.mfePrice
}

// MaximumAdverseExcursion is the Maximum Adverse Excursion (MAE) in instrument's currency.
// It means the maximum potential loss taken during the round-trip period.
func (r *Roundtrip) MaximumAdverseExcursion() float64 {
	return r.mae
}

// MaximumFavorableExcursion is the Maximum Favorable Excursion (MFE) in instrument's currency.
func (r *Roundtrip) MaximumFavorableExcursion() float64 {
	return r.mfe
}

/*
Maximum Favorable Excursion
Maximum Favorable Excursion is the peak profit before closing the trade. For example, you may have a closed trade which lost 25 pips but during the time the trade was open, it was making a 100 pips profit at some point - that was the Maximum Favorable Excursion for that particular trade.
This statistical concept originally created by John Sweeney to measure the distinctive characteristics of profitable trades, can be used as part of an analytical process to enable traders to distinguish between average trades and those that offer substantially greater profit potential.

Maximum Adverse Excursion
This is the maximum potential loss that the trade had before the trade closed in profit. For example, a trade closed with 25 points in profit but during the time it was open, at one point, it was losing 100 points - that was the Maximum Adverse Excursion for that trade.
*/

// EntryEfficiency is the Entry Efficiency which measures the percentage in range [0. 100] of
// the total round-trip potential taken by a round-trip given its entry and assuming
// the best possible exit during the round-trip period.
//
// It shows how close the entry price was to the best possible entry price during the round-trip.
//
// 100% is the perfect efficiency, 0% is the worst possible efficiency.
func (r *Roundtrip) EntryEfficiency() float64 {
	delta := r.mfePrice - r.maePrice
	if delta != 0 {
		switch r.side {
		case pos.Long:
			return 100 * (r.mfePrice - r.entryPrice) / delta
		case pos.Short:
			return 100 * (r.entryPrice - r.maePrice) / delta
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
	delta := r.mfePrice - r.maePrice
	if delta != 0 {
		switch r.side {
		case pos.Long:
			return 100 * (r.exitPrice - r.maePrice) / delta
		case pos.Short:
			return 100 * (r.mfePrice - r.exitPrice) / delta
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
	delta := r.mfePrice - r.maePrice
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
