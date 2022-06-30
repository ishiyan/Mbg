package kaufman

import "mbg/trading/data"

// KaufmanAdaptiveMovingAverageLengthParams describes parameters to create an instance of the indicator
// based on length.
type KaufmanAdaptiveMovingAverageLengthParams struct {
	// Length is the length (the number of time periods, ℓ) of the moving window to calculate the average.
	//
	// The value should be greater than 1.
	Length int

	// FirstIsAverage indicates whether the very first exponential moving average value is
	// a simple average of the first 'period' (the most widely documented approach) or
	// the first input value (used in Metastock).
	FirstIsAverage bool

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}

// KaufmanAdaptiveMovingAverageLengthParams describes parameters to create an instance of the indicator
// based on smoothing factor.
type KaufmanAdaptiveMovingAverageSmoothingFactorParams struct {
	// SmoothingFactor is the smoothing factor, α in (0,1), of the exponential moving average.
	//
	// The equivalent length ℓ is:
	//    ℓ = 2/α - 1, 0<α≤1, 1≤ℓ.
	SmoothingFactor float64

	// FirstIsAverage indicates whether the very first exponential moving average value is
	// a simple average of the first 'period' (the most widely documented approach) or
	// the first input value (used in Metastock).
	FirstIsAverage bool

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}
