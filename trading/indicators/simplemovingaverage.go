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

// SimpleMovingAverage computes the simple, or arithmetic, moving average (SMA) by adding the samples
// for a number of time periods (length, ℓ) and then dividing this total by the number of time periods.
//
// In other words, this is an unweighted mean (gives equal weight to each sample) of the previous ℓ samples.
//
// This implementation updates the value of the SMA incrementally using the formula:
//    SMAᵢ = SMAᵢ₋₁ + (Pᵢ - Pᵢ₋ℓ) / ℓ,
// where ℓ is the length.
//
// The indicator is not primed during the first ℓ-1 updates.
type SimpleMovingAverage struct {
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

// NewSimpleMovingAverage returns an instnce of the indicator created using supplied parameters.
func NewSimpleMovingAverage(p *SimpleMovingAverageParams) (*SimpleMovingAverage, error) {
	const (
		invalid = "invalid simple moving average parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtn    = "sma(%d)"
		minlen  = 2
	)

	length := p.Length
	if length < minlen {
		return nil, fmt.Errorf(fmts, invalid, "length should be greater than 1")
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
	desc := "Simple moving average " + name

	return &SimpleMovingAverage{
		name:         name,
		description:  desc,
		window:       make([]float64, length),
		windowLength: length,
		lastIndex:    length - 1,
		barFunc:      barFunc,
		quoteFunc:    quoteFunc,
		tradeFunc:    tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *SimpleMovingAverage) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the simple moving average.
func (s *SimpleMovingAverage) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.SimpleMovingAverage,
		Outputs: []output.Metadata{
			{
				Kind:        int(SimpleMovingAverageValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
		},
	}
}

// Update updatess the value of the simple moving average given the next sample.
//
// The indicator is not primed during the first ℓ-1 updates.
func (s *SimpleMovingAverage) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	temp := sample

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.primed {
		s.windowSum += temp - s.window[0]

		for i := 0; i < s.lastIndex; i++ {
			s.window[i] = s.window[i+1]
		}

		s.window[s.lastIndex] = temp

		return temp
	} else {
		s.windowSum += temp
		s.window[s.windowCount] = temp
		s.windowCount++

		if s.windowLength != s.windowCount {
			return math.NaN()
		}

		s.primed = true
	}

	return s.windowSum / float64(s.windowLength)
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *SimpleMovingAverage) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: s.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *SimpleMovingAverage) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *SimpleMovingAverage) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *SimpleMovingAverage) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}
