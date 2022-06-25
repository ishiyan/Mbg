package statistics

//nolint:gofumpt
import (
	"fmt"
	"math"
	"time"

	"mbg/trading/data"
	"mbg/trading/indicators/indicator"
	"mbg/trading/indicators/indicator/output"
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
	outputs     []StandardDeviationOutput
}

// NewStandardDeviation returns an instnce of the StandardDeviation indicator created using supplied parameters.
func NewStandardDeviation(p *VarianceParams, outputs []StandardDeviationOutput) (*StandardDeviation, error) {
	const (
		fmtn = "stdev.%c(%d)"
	)

	for i, o := range outputs {
		if !o.IsKnown() {
			return nil, fmt.Errorf("unknown standard deviation output[%d]: %d", i, int(o))
		}
	}

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
		outputs:     outputs,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *StandardDeviation) IsPrimed() bool {
	return s.variance.IsPrimed()
}

// Metadata describes an output data of the indicator.
func (s *StandardDeviation) Metadata() indicator.Metadata {
	length := len(s.outputs)
	outputs := make([]output.Metadata, length)

	for i, o := range s.outputs {
		outputs[i].Kind = int(o)
		outputs[i].Type = output.Scalar

		switch o {
		case StandardDeviationValue:
			outputs[i].Name = s.name
			outputs[i].Description = s.description
		default: // StandardDeviationVarianceValue
			m := s.variance.Metadata()
			outputs[i].Name = m.Outputs[0].Name
			outputs[i].Description = m.Outputs[0].Description
		}
	}

	return indicator.Metadata{
		Type:    indicator.StandardDeviation,
		Outputs: outputs,
	}
}

func (s *StandardDeviation) Update(time time.Time, sample float64) indicator.Output {
	length := len(s.outputs)
	outputs := make([]any, length)

	v := s.variance.Update(sample)

	for i, o := range s.outputs {
		switch o {
		case StandardDeviationValue:
			if math.IsNaN(v) {
				outputs[i] = data.Scalar{Time: time, Value: math.NaN()}
			} else {
				outputs[i] = data.Scalar{Time: time, Value: math.Sqrt(v)}
			}
		default: // StandardDeviationVarianceValue
			outputs[i] = data.Scalar{Time: time, Value: v}
		}
	}

	return outputs
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *StandardDeviation) UpdateScalar(sample *data.Scalar) indicator.Output {
	return s.Update(sample.Time, sample.Value)
}

// UpdateBar updates the indicator given the next bar sample.
func (s *StandardDeviation) UpdateBar(sample *data.Bar) indicator.Output {
	return s.Update(sample.Time, s.variance.barFunc(sample))
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *StandardDeviation) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.Update(sample.Time, s.variance.quoteFunc(sample))
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *StandardDeviation) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.Update(sample.Time, s.variance.tradeFunc(sample))
}
