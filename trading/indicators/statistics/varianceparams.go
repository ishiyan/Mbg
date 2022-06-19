package statistics

import "mbg/trading/data"

// VarianceParams describes parameters to create an instance of the indicator.
type VarianceParams struct {
	// Length is the length (the number of time periods, ℓ) of the moving window to calculate the variance.
	//
	// The value should be greater than 1.
	Length int

	// IsUnbiased indicates whether the estimate of the variance is the unbiased sample variance or the population variance.
	//
	// When in doubt, use the unbiased sample variance (value is true).
	IsUnbiased bool

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}
