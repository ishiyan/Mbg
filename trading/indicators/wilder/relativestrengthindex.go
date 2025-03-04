package wilder

//nolint: gofumpt
import (
	"fmt"
	"math"
	"sync"

	"mbg/trading/data"
	"mbg/trading/indicators/indicator"
	"mbg/trading/indicators/indicator/output"
)

// RelativeStrengthIndex is a momentum indicator based on the average of up samples
// and down samples over a specified period ℓ (commonly 14 samples).
//
// The calculation formula is:
//
// RSIᵢ = 100 - 100 / (1 + RSᵢ),
//
// where RSᵢ (Relative Strength) is the average gain divided by the average loss
// over the chosen period [i-ℓ, i].
//
// The indicator is not primed during the first ℓ updates.
type RelativeStrengthIndex struct {
	mu          sync.RWMutex
	name        string
	description string
	prevSample  float64
	gain        float64
	loss        float64
	length      int
	count       int
	primed      bool
	barFunc     data.BarFunc
	quoteFunc   data.QuoteFunc
	tradeFunc   data.TradeFunc
}

// NewRelativeStrengthIndex returns an instnce of the indicator created using supplied parameters.
func NewRelativeStrengthIndex(p *RelativeStrengthIndexParams) (*RelativeStrengthIndex, error) {
	const (
		invalid = "invalid relative strength index parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtn    = "rsi(%d)"
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
	desc := "Relative Strength Index " + name

	return &RelativeStrengthIndex{
		name:        name,
		description: desc,
		length:      length,
		barFunc:     barFunc,
		quoteFunc:   quoteFunc,
		tradeFunc:   tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *RelativeStrengthIndex) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the indicator.
func (s *RelativeStrengthIndex) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.RelativeStrengthIndex,
		Outputs: []output.Metadata{
			{
				Kind:        int(RelativeStrengthIndexValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
		},
	}
}

// Update updates the value of the relative strength index given the next sample.
//
// The indicator is not primed during the first ℓ updates.
func (s *RelativeStrengthIndex) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.primed {
		d := sample - s.prevSample
		s.prevSample = sample

		var g, l float64
		if d < 0 {
			l = -d
		} else if d > 0 {
			g = d
		}

		l0 := float64(s.length)
		l1 := float64(s.length - 1)

		s.gain = (s.gain*l1 + g) / l0
		s.loss = (s.loss*l1 + l) / l0

		d = s.gain + s.loss
		if d == 0 {
			return 0
		}

		return 100.0 * s.gain / d
	} else {
		// Not primed.
		s.count++

		if s.count == 1 { // The very first sample.
			s.prevSample = sample
			return math.NaN()
		}

		d := sample - s.prevSample
		s.prevSample = sample

		if d < 0 {
			s.loss -= d
		} else if d > 0 {
			s.gain += d
		}

		if s.count <= s.length {
			return math.NaN()
		}

		s.primed = true

		d = float64(s.length)
		s.gain /= d
		s.loss /= d

		d = s.gain + s.loss
		if d == 0 {
			return 0
		}

		return 100.0 * s.gain / d
	}
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *RelativeStrengthIndex) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: s.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *RelativeStrengthIndex) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *RelativeStrengthIndex) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *RelativeStrengthIndex) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}
