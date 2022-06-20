package indicator

import (
	"mbg/trading/data"
)

// Indicator describes a common indicator functionality.
type Indicator interface {
	// IsPrimed indicates whether an indicator is primed.
	IsPrimed() bool

	// Metadata describes an output data of an indicator.
	Metadata() Metadata

	// UpdateScalar updates an indicator given the next scalar sample.
	UpdateScalar(sample *data.Scalar) Output

	// UpdateBar updates an indicator given the next bar sample.
	UpdateBar(sample *data.Bar) Output

	// UpdateQuote updates an indicator given the next quote sample.
	UpdateQuote(sample *data.Quote) Output

	// UpdateQuote updates an indicator given the next trade sample.
	UpdateTrade(sample *data.Trade) Output
}

// UpdateScalars updates the indicator given a slice of the next scalar samples.
func UpdateScalars(ind Indicator, samples []*data.Scalar) []Output {
	length := len(samples)
	output := make([]Output, length)

	for i, d := range samples {
		output[i] = ind.UpdateScalar(d)
	}

	return output
}

// UpdateBars updates the indicator given a slice of the next bar samples.
func UpdateBars(ind Indicator, samples []*data.Bar) []Output {
	length := len(samples)
	output := make([]Output, length)

	for i, d := range samples {
		output[i] = ind.UpdateBar(d)
	}

	return output
}

// UpdateQuotes updates the indicator given a slice of the next quote samples.
func UpdateQuotes(ind Indicator, samples []*data.Quote) []Output {
	length := len(samples)
	output := make([]Output, length)

	for i, d := range samples {
		output[i] = ind.UpdateQuote(d)
	}

	return output
}

// UpdateTrades updates the indicator given a slice of the next trade samples.
func UpdateTrades(ind Indicator, samples []*data.Trade) []Output {
	length := len(samples)
	output := make([]Output, length)

	for i, d := range samples {
		output[i] = ind.UpdateTrade(d)
	}

	return output
}
