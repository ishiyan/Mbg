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

// WeightedMovingAverage computes the weighted moving average (WMA) that has multiplying factors
// to give arithmetically decreasing weights to the samples in the look back window.
//    WMAᵢ = (ℓPᵢ + (ℓ-1)Pᵢ₋₁ + ... + Pᵢ₋ℓ) / (ℓ + (ℓ-1) + ... + 2 + 1),
// where ℓ is the length.
//
// The denominator is a triangle number and can be computed as
//    ½ℓ(ℓ-1).
//
// When calculating the WMA across successive values,
//    WMAᵢ₊₁ - WMAᵢ = ℓPᵢ₊₁ - Pᵢ - ... - Pᵢ₋ℓ₊₁
// If we denote the sum
//    Totalᵢ = Pᵢ + ... + Pᵢ₋ℓ₊₁
// then
//    Totalᵢ₊₁ = Totalᵢ + Pᵢ₊₁ - Pᵢ₋ℓ₊₁
//    Numeratorᵢ₊₁ = Numeratorᵢ + ℓPᵢ₊₁ - Totalᵢ
//    WMAᵢ₊₁ = Numeratorᵢ₊₁ / ½ℓ(ℓ-1)
//
// The WMA indicator is not primed during the first ℓ-1 updates.
type WeightedMovingAverage struct {
	mu           sync.RWMutex
	name         string
	description  string
	window       []float64
	windowSum    float64
	windowSub    float64
	divider      float64
	windowLength int
	windowCount  int
	lastIndex    int
	primed       bool
	barFunc      data.BarFunc
	quoteFunc    data.QuoteFunc
	tradeFunc    data.TradeFunc
}

// NewWeightedMovingAverage returns an instnce of the indicator created using supplied parameters.
func NewWeightedMovingAverage(p *WeightedMovingAverageParams) (*WeightedMovingAverage, error) {
	const (
		invalid = "invalid weighted moving average parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtn    = "wma(%d)"
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
	desc := "Weighted moving average " + name
	divider := float64(length) * float64(length+1) / 2.

	return &WeightedMovingAverage{
		mu:           sync.RWMutex{},
		name:         name,
		description:  desc,
		window:       make([]float64, length),
		windowSum:    0,
		windowSub:    0,
		divider:      divider,
		windowLength: length,
		windowCount:  0,
		lastIndex:    length - 1,
		primed:       false,
		barFunc:      barFunc,
		quoteFunc:    quoteFunc,
		tradeFunc:    tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (w *WeightedMovingAverage) IsPrimed() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return w.primed
}

// Metadata describes an output data of the indicator.
//
// It always has a single scalar output -- the calculated value of the indicator.
func (w *WeightedMovingAverage) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.WeightedMovingAverage,
		Outputs: []output.Metadata{
			{
				Kind:        int(WeightedMovingAverageValue),
				Type:        output.Scalar,
				Name:        w.name,
				Description: w.description,
			},
		},
	}
}

// Update updates the value of the moving average given the next sample.
//
// The indicator is not primed during the first ℓ-1 updates.
func (w *WeightedMovingAverage) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	temp := sample

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.primed {
		w.windowSum -= w.windowSub
		w.windowSum += temp * float64(w.windowLength)
		w.windowSub -= w.window[0]
		w.windowSub += temp

		for i := 0; i < w.lastIndex; i++ {
			w.window[i] = w.window[i+1]
		}

		w.window[w.lastIndex] = temp
	} else { // Not primed.
		w.window[w.windowCount] = temp
		w.windowSub += temp
		w.windowCount++
		w.windowSum += temp * float64(w.windowCount)

		if w.windowLength > w.windowCount {
			return math.NaN()
		}

		w.primed = true
	}

	return w.windowSum / w.divider
}

// UpdateScalar updates the indicator given the next scalar sample.
func (w *WeightedMovingAverage) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: w.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (w *WeightedMovingAverage) UpdateBar(sample *data.Bar) indicator.Output {
	return w.UpdateScalar(&data.Scalar{Time: sample.Time, Value: w.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (w *WeightedMovingAverage) UpdateQuote(sample *data.Quote) indicator.Output {
	return w.UpdateScalar(&data.Scalar{Time: sample.Time, Value: w.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (w *WeightedMovingAverage) UpdateTrade(sample *data.Trade) indicator.Output {
	return w.UpdateScalar(&data.Scalar{Time: sample.Time, Value: w.tradeFunc(sample)})
}
