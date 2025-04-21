package ehlers

import (
	"mbg/trading/data"                                 //nolint:depguard
	"mbg/trading/indicators/ehlers/hilberttransformer" //nolint:depguard
)

// MesaAdaptiveMovingAverageLengthParams describes parameters to create an instance of the indicator
// based on lengths.
type MesaAdaptiveMovingAverageLengthParams struct {
	// EstimatorType is the type of cycle estimator to use.
	// The default value is hilberttransformer.HomodyneDiscriminator.
	EstimatorType hilberttransformer.CycleEstimatorType

	// CycleEstimatorParams describes parameters to create an instance
	// of the Hilbert transformer cycle estimator.
	EstimatorParams hilberttransformer.CycleEstimatorParams

	// FastLimitLength is the fastest boundary length, ℓf.
	// The equivalent smoothing factor αf is
	//
	//   αf = 2/(ℓf + 1), 2 ≤ ℓ
	//
	// The value should be greater than 1.
	// The default value is 3 (αf=0.5).
	FastLimitLength int

	// SlowLimitLength is the slowest boundary length, ℓs.
	// The equivalent smoothing factor αs is
	//
	//   αs = 2/(ℓs + 1), 2 ≤ ℓ
	//
	// The value should be greater than 1.
	// The default value is 39 (αs=0.05).
	SlowLimitLength int

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}

// MesaAdaptiveMovingAverageSmoothingFactorParams describes parameters to create an instance of the indicator
// based on smoothing factors.
type MesaAdaptiveMovingAverageSmoothingFactorParams struct {
	// EstimatorType is the type of cycle estimator to use.
	// The default value is hilberttransformer.HomodyneDiscriminator.
	EstimatorType hilberttransformer.CycleEstimatorType

	// CycleEstimatorParams describes parameters to create an instance
	// of the Hilbert transformer cycle estimator.
	EstimatorParams hilberttransformer.CycleEstimatorParams

	// FastLimitSmoothingFactor is the fastest boundary smoothing factor, αf in (0,1).
	// The equivalent length ℓf is
	//
	//   ℓf = 2/αf - 1, 0 < αf ≤ 1, 1 ≤ ℓf
	//
	// The default value is 0.5 (ℓf=3).
	FastLimitSmoothingFactor float64

	// SlowLimitSmoothingFactor is the slowest boundary smoothing factor, αs in (0,1).
	// The equivalent length ℓs is
	//
	//   ℓs = 2/αs - 1, 0 < αs ≤ 1, 1 ≤ ℓs
	//
	// The default value is 0.05 (ℓs=39).
	SlowLimitSmoothingFactor float64

	// BarComponent indicates the component of a bar to use when updating the indicator with a bar sample.
	BarComponent data.BarComponent

	// QuoteComponent indicates the component of a quote to use when updating the indicator with a quote sample.
	QuoteComponent data.QuoteComponent

	// TradeComponent indicates the component of a trade to use when updating the indicator with a trade sample.
	TradeComponent data.TradeComponent
}
