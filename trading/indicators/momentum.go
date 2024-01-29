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

// Momentum is the absolute (not normalized) difference between today's sample and the sample ℓ periods ago.
//
// This implementation calculates the value of the MOM using the formula:
//
// MOMᵢ = Pᵢ - Pᵢ₋ℓ,
//
// where ℓ is the length.
//
// The indicator is not primed during the first ℓ updates.
type Momentum struct {
	mu           sync.RWMutex
	name         string
	description  string
	window       []float64
	windowSum    float64
	windowLength int
	windowCount  int
	lastIndex    int
	primed       bool
	barFunc      data.BarFunc
	quoteFunc    data.QuoteFunc
	tradeFunc    data.TradeFunc
}

// NewMomentum returns an instnce of the indicator created using supplied parameters.
func NewMomentum(p *MomentumParams) (*Momentum, error) {
	const (
		invalid = "invalid momentum parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtn    = "mom(%d)"
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
	desc := "Momentum " + name

	return &Momentum{
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
func (s *Momentum) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the momentum.
func (s *Momentum) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.Momentum,
		Outputs: []output.Metadata{
			{
				Kind:        int(MomentumValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
		},
	}
}

// Update updates the value of the momentum given the next sample.
//
// The indicator is not primed during the first ℓ updates.
func (s *Momentum) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.primed {
		for i := 0; i < s.lastIndex; i++ {
			s.window[i] = s.window[i+1]
		}

		s.window[s.lastIndex] = sample
		return sample - s.window[0]
	} else {
		s.window[s.windowCount] = sample
		s.windowCount++

		if s.windowLength == s.windowCount {
			s.primed = true
			return sample - s.window[0]
		}

		return math.NaN()
	}
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *Momentum) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: s.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *Momentum) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *Momentum) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *Momentum) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}
