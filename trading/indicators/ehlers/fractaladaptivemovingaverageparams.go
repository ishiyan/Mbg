package ehlers

import (
	"mbg/trading/data" //nolint:depguard
)

// FractalAdaptiveMovingAverageLengthParams describes parameters to create an instance of the indicator.
type FractalAdaptiveMovingAverageParams struct {
	// Length is the length, ℓ, (the number of time periods) of the Fractal Adaptive Moving Average.
	//
	// The value should be an even integer, greater or equal to 2.
	// The default value is 16.
	Length int

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	//
	// The original FRAMA indicator uses the median price (high+low)/2, which is the default.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}
