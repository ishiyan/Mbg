package indicators

//nolint: gofumpt
import (
	"fmt"
	"math"
	"sync"

	"mbg/trading/data"
	"mbg/trading/indicators/indicator"
	"mbg/trading/indicators/indicator/output"
)

// RateOfChangePercent is the difference between today's sample and the sample ℓ periods ago
// scaled by the old sample so as to represent the increase as a fraction.
//
// The values are centered at zero and can be positive and negative.
//
// ROC%ᵢ = (Pᵢ - Pᵢ₋ℓ) / Pᵢ₋ℓ = (Pᵢ/Pᵢ₋ℓ -1),
//
// where ℓ is the length.
//
// The indicator is not primed during the first ℓ updates.
type RateOfChangePercent struct {
	mu           sync.RWMutex
	name         string
	description  string
	window       []float64
	windowLength int
	windowCount  int
	lastIndex    int
	primed       bool
	barFunc      data.BarFunc
	quoteFunc    data.QuoteFunc
	tradeFunc    data.TradeFunc
}

// NewRateOfChangePercent returns an instnce of the indicator created using supplied parameters.
func NewRateOfChangePercent(p *RateOfChangePercentParams) (*RateOfChangePercent, error) {
	const (
		invalid = "invalid rate of change percent parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtn    = "rocp(%d)"
		minlen  = 1
	)

	length := p.Length
	if length < minlen {
		return nil, fmt.Errorf(fmts, invalid, "length should be positive")
	}

	var (
		err       error
		barFunc   data.BarFunc
		quoteFunc data.QuoteFunc
		tradeFunc data.TradeFunc
	)

	if barFunc, err = data.BarComponentFunc(p.BarComponent); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	if quoteFunc, err = data.QuoteComponentFunc(p.QuoteComponent); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	if tradeFunc, err = data.TradeComponentFunc(p.TradeComponent); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	name := fmt.Sprintf(fmtn, length)
	desc := "Rate of Change percent " + name

	return &RateOfChangePercent{
		name:         name,
		description:  desc,
		window:       make([]float64, length+1),
		windowLength: length + 1,
		lastIndex:    length,
		barFunc:      barFunc,
		quoteFunc:    quoteFunc,
		tradeFunc:    tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *RateOfChangePercent) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the rate of change percent.
func (s *RateOfChangePercent) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.RateOfChangePercent,
		Outputs: []output.Metadata{
			{
				Kind:        int(RateOfChangePercentValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
		},
	}
}

// Update updates the value of the indicator given the next sample.
//
// The indicator is not primed during the first ℓ updates.
func (s *RateOfChangePercent) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	const epsilon = 1e-13

	if s.primed {
		if s.lastIndex > 1 {
			for i := 0; i < s.lastIndex; i++ {
				s.window[i] = s.window[i+1]
			}
		}

		s.window[s.lastIndex] = sample
		previous := s.window[0]
		if math.Abs(previous) > epsilon {
			return sample/previous - 1
		}

		return 0
	} else { // Not primed.
		s.window[s.windowCount] = sample
		s.windowCount++

		if s.windowLength == s.windowCount {
			s.primed = true
			previous := s.window[0]
			if math.Abs(previous) > epsilon {
				return sample/previous - 1
			}

			return 0
		}

		return math.NaN()
	}
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *RateOfChangePercent) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: s.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *RateOfChangePercent) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *RateOfChangePercent) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *RateOfChangePercent) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}
