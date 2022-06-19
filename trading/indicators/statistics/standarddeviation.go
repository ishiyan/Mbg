package statistics

// StandardDeviation computes the standard deviation of the samples within a moving window of length ℓ
// as a square root of variance:
//
// σ² = (∑xᵢ² - (∑xᵢ)²/ℓ)/ℓ
//
// for the estimation of the population variance, or as:
//
// σ² = (∑xᵢ² - (∑xᵢ)²/ℓ)/(ℓ-1)
//
// for the unbiased estimation of the sample variance, i={0,…,ℓ-1}.
type StandardDeviation struct{}
