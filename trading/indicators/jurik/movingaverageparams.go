package jurik

import "mbg/trading/data"

// MovingAverageParams describes parameters to create an instance of the indicator.
type MovingAverageParams struct {
	// Length is the length (the number of time periods, ℓ) of the moving window to calculate the average.
	//
	// The value should be greater than 1. Typical values range from 5 to 20.
	//
	// Irrespective from the value, the indicator needs at 30 first values to be primed.
	Length int

	// Phase controls the trade-off between lag and overshoot.
	//
	// The phase values should be in [-100, 100].
	//
	// - The value of -100 results in maximum lag and no overshoot.
	//
	// - The value of 0 results in some lag and some overshoot.
	//
	// - The value of 100 results in minimum lag and maximum overshoot.
	Phase int

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}
