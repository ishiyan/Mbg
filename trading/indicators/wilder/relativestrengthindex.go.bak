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
	mu           sync.RWMutex
	name         string
	description  string
	prevSample   float64
	window       []float64
	windowGain   float64
	windowLoss   float64
	windowLength int
	windowCount  int
	primed       bool
	barFunc      data.BarFunc
	quoteFunc    data.QuoteFunc
	tradeFunc    data.TradeFunc
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
		name:         name,
		description:  desc,
		window:       make([]float64, length),
		windowLength: length,
		barFunc:      barFunc,
		quoteFunc:    quoteFunc,
		tradeFunc:    tradeFunc,
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
		d := s.window[0]
		if d < 0 {
			s.windowLoss += d
		} else if d > 0 {
			s.windowGain -= d
		}

		j := s.windowLength - 1
		for i := 0; i < j; i++ {
			s.window[i] = s.window[i+1]
		}

		d = sample - s.prevSample
		s.prevSample = sample
		s.window[j] = d

		if d < 0 {
			s.windowLoss -= d
		} else if d > 0 {
			s.windowGain += d
		}

		d = float64(s.windowLength)
		l := s.windowLoss / d
		g := s.windowGain / d

		d = g + l
		if d == 0 {
			return 0
		}

		return 100.0 * g / d
	} else {
		// Not primed.
		s.windowCount++

		if s.windowCount == 1 { // The very first sample.
			s.prevSample = sample
			return math.NaN()
		}

		d := sample - s.prevSample
		s.window[s.windowCount-2] = d
		s.prevSample = sample

		if d < 0 {
			s.windowLoss -= d
		} else if d > 0 {
			s.windowGain += d
		}

		if s.windowCount <= s.windowLength {
			return math.NaN()
		}

		s.primed = true

		d = float64(s.windowLength)
		l := s.windowLoss / d
		g := s.windowGain / d

		d = g + l
		if d == 0 {
			return 0
		}

		return 100.0 * g / d
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
