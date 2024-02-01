package wilder

import "mbg/trading/data"

// RelativeStrengthIndexParams describes parameters to create an instance of the indicator.
type RelativeStrengthIndexParams struct {
	// Length is the length (the number of time periods, ℓ) which defines today's sample and the sample ℓ periods ago.
	//
	// The value should be greater than 1.
	Length int

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}
