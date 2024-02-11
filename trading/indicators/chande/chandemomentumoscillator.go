package chande

//nolint: gofumpt
import (
	"fmt"
	"math"
	"sync"

	"mbg/trading/data"
	"mbg/trading/indicators/indicator"
	"mbg/trading/indicators/indicator/output"
)

// ChandeMomentumOscillator is a momentum indicator based on the average
// of up samples and down samples over a specified period ℓ.
//
// The calculation formula is:
//
// CMOᵢ = 100 (SUᵢ-SDᵢ) / (SUᵢ + SDᵢ),
//
// where SUᵢ (sum up) is the sum of gains and SDᵢ (sum down)
// is the sum of losses over the chosen period [i-ℓ, i].
//
// The indicator is not primed during the first ℓ updates.
type ChandeMomentumOscillator struct {
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

// NewChandeMomentumOscillator returns an instnce of the indicator created using supplied parameters.
func NewChandeMomentumOscillator(p *ChandeMomentumOscillatorParams) (*ChandeMomentumOscillator, error) {
	const (
		invalid = "invalid Chande momentum oscillator parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtn    = "cmo(%d)"
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
	desc := "Chande Momentum Oscillator " + name

	return &ChandeMomentumOscillator{
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
func (s *ChandeMomentumOscillator) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the indicator.
func (s *ChandeMomentumOscillator) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.ChandeMomentumOscillator,
		Outputs: []output.Metadata{
			{
				Kind:        int(ChandeMomentumOscillatorValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
		},
	}
}

// Update updates the value of the iChande momentum oscillator given the next sample.
//
// The indicator is not primed during the first ℓ updates.
func (s *ChandeMomentumOscillator) Update(sample float64) float64 {
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

		l := s.windowLoss
		g := s.windowGain

		d = g + l
		if d == 0 {
			return 0
		}

		return 100.0 * (g - l) / d
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

		l := s.windowLoss
		g := s.windowGain

		d = g + l
		if d == 0 {
			return 0
		}

		return 100.0 * (g - l) / d
	}
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *ChandeMomentumOscillator) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: s.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *ChandeMomentumOscillator) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *ChandeMomentumOscillator) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *ChandeMomentumOscillator) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}
