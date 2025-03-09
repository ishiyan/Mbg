package tillson

import "mbg/trading/data" //nolint:depguard

// T2ExponentialMovingAverageLengthParams describes parameters to create an instance of the indicator
// based on length.
type T2ExponentialMovingAverageLengthParams struct {
	// Length is the length (the number of time periods, ℓ) of the moving window to calculate the average.
	//
	// The value should be greater than 1.
	Length int

	// VolumeFactor is the volume factor, v (0 ≤ ν ≤ 1), of the exponential moving average.
	// The default value is 0.7.
	// When ν=0, T2 is just an EMA, and when ν=1, T3 is DEMA.
	// In between, T2 is a cooler DEMA.
	VolumeFactor float64

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

// T2ExponentialMovingAverageLengthParams describes parameters to create an instance of the indicator
// based on smoothing factor.
type T2ExponentialMovingAverageSmoothingFactorParams struct {
	// SmoothingFactor is the smoothing factor, α (0 < α < 1), of the exponential moving average.
	//
	// The equivalent length ℓ is:
	//    ℓ = 2/α - 1, 0<α<1, 1≤ℓ.
	SmoothingFactor float64

	// VolumeFactor is the volume factor, v (0 ≤ ν ≤ 1), of the exponential moving average.
	// The default value is 0.7.
	// When ν=0, T2 is just an EMA, and when ν=1, T2 is DEMA.
	// In between, T2 is a cooler DEMA.
	VolumeFactor float64

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
