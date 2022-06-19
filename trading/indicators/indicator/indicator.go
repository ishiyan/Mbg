package indicator

import (
	"mbg/trading/data"
)

// Indicator describes a common indicator functionality.
type Indicator interface {
	// IsPrimed indicates whether an indicator is primed.
	IsPrimed() bool

	// Reset resets an indicator. The indicator is not primed after this call.
	Reset()

	// Metadata describes an output data of an indicator.
	Metadata() Metadata

	// UpdateScalar updates an indicator given the next scalar sample.
	UpdateScalar(sample *data.Scalar) Output

	// UpdateQuote updates an indicator given the next quote sample.
	UpdateQuote(sample *data.Quote) Output

	// UpdateBar updates an indicator given the next bar sample.
	UpdateBar(sample *data.Bar) Output

	// UpdateScalars updates an indicator given a slice of the next scalar samples.
	UpdateScalars(samples []*data.Scalar) []Output

	// UpdateQuotes updates an indicator given a slice of the next quote samples.
	UpdateQuotes(samples []*data.Quote) []Output

	// UpdateBars updates an indicator given a slice of the next bar samples.
	UpdateBars(samples []*data.Bar) []Output
}
