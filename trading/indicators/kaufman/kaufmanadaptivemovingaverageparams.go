package kaufman

import "mbg/trading/data"

// KaufmanAdaptiveMovingAverageLengthParams describes parameters to create an instance of the indicator
// based on length.
//
// Kaufman's adaptive moving average (KAMA) is an EMA with the smoothing factor, α,
// being changed with each new sample within the fastest and the slowest boundaries.
//
// KAMAᵢ = αPᵢ + (1 - α)*KAMAᵢ₋₁,  α = (αs + (αf - αs)ε)²
//
// where the αf is the α of the fastest (shortest, default 2 samples) period boundary,
// the αs is the α of the slowest (longest, default 30 samples) period boundary,
// and ε is the efficiency ratio:
//
// ε = |P - Pℓ| / ∑|Pᵢ - Pᵢ₊₁|,  i ≤ ℓ-1
//
// where ℓ is a number of samples used to calculate the ε.
// The recommended values of ℓ are in the range of 8 to 10.
//
// The efficiency ratio has the value of 1 when samples move in the same direction for
// the full ℓ periods, and a value of 0 when samples are unchanged over the ℓ periods.
// When samples move in wide swings within the interval, the sum of the denominator
// becomes very large compared with the numerator and the ε approaches 0.
// Smaller values of ε result in a smaller smoothing constant and a slower trend.
//
// The indicator is not primed during the first ℓ updates.
//
// See
// Perry J. Kaufman, Smarter Trading, McGraw-Hill, Ney York, 1995, pp. 129-153
// for a complete discussion.
type KaufmanAdaptiveMovingAverageLengthParams struct {
	// EfficiencyRatioLength is the number of last samples used to calculate the efficiency ratio.
	//
	// The value should be greater than 1.
	// The default value is 10.
	EfficiencyRatioLength int

	// FastestLength is the fastest boundary length, ℓf.
	// The equivalent smoothing factor αf is
	//
	//   αf = 2/(ℓf + 1), 2 ≤ ℓ
	//
	// The value should be greater than 1.
	// The default value is 2.
	FastestLength int

	// SlowestLength is the slowest boundary length, ℓs.
	// The equivalent smoothing factor αs is
	//
	//   αs = 2/(ℓs + 1), 2 ≤ ℓ
	//
	// The value should be greater than 1.
	// The default value is 30.
	SlowestLength int

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
	// EfficiencyRatioLength is the number of last samples used to calculate the efficiency ratio.
	//
	// The value should be greater than 1.
	// The default value is 10.
	EfficiencyRatioLength int

	// FastestSmoothingFactor is the fastest boundary smoothing factor, αf in (0,1).
	// The equivalent length ℓf is
	//
	//   ℓf = 2/αf - 1, 0 < αf ≤ 1, 1 ≤ ℓf
	//
	// The default value is 2/3 (0.6666...).
	FastestSmoothingFactor float64

	// SlowestSmoothingFactor is the slowest boundary smoothing factor, αs in (0,1).
	// The equivalent length ℓs is
	//
	//   ℓs = 2/αs - 1, 0 < αs ≤ 1, 1 ≤ ℓs
	//
	// The default value is 2/31 (0.06451612903225806451612903225806).
	SlowestSmoothingFactor float64

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}
