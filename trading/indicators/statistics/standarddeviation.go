package statistics

import (
	"fmt"
)

// StandardDeviation computes the standard deviation of the samples within a moving window of length ℓ
// as a square root of variance:
//    σ² = (∑xᵢ² - (∑xᵢ)²/ℓ)/ℓ
// for the estimation of the population variance, or as:
//    σ² = (∑xᵢ² - (∑xᵢ)²/ℓ)/(ℓ-1)
// for the unbiased estimation of the sample variance, i={0,…,ℓ-1}.
type StandardDeviation struct {
	name        string
	description string
	variance    *Variance
}

// NewStandardDeviation returns an instnce of the StandardDeviation indicator created using supplied parameters.
func NewStandardDeviation(p *VarianceParams) (*StandardDeviation, error) {
	const (
		fmtn = "stdev.%c(%d)"
	)

	variance, err := NewVariance(p)
	if err != nil {
		return nil, err
	}

	var name, desc string
	if p.IsUnbiased {
		name = fmt.Sprintf(fmtn, 's', p.Length)
		desc = "Standard deviation based on unbiased estimation of the sample variance " + name
	} else {
		name = fmt.Sprintf(fmtn, 'p', p.Length)
		desc = "Standard deviation based on estimation of the population variance " + name
	}

	return &StandardDeviation{
		name:        name,
		description: desc,
		variance:    variance,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *StandardDeviation) IsPrimed() bool {
	return s.variance.IsPrimed()
}
