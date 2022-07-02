package statistics

//nolint: gofumpt
import (
	"fmt"
	"math"
	"sync"

	"mbg/trading/data"
	"mbg/trading/indicators/indicator"
	"mbg/trading/indicators/indicator/output"
)

// Variance computes the variance of the samples within a moving window of length ℓ:
//    σ² = (∑xᵢ² - (∑xᵢ)²/ℓ)/ℓ
// for the estimation of the population variance, or as:
//    σ² = (∑xᵢ² - (∑xᵢ)²/ℓ)/(ℓ-1)
// for the unbiased estimation of the sample variance, i={0,…,ℓ-1}.
type Variance struct {
	mu               sync.RWMutex
	name             string
	description      string
	window           []float64
	windowSum        float64
	windowSquaredSum float64
	windowLength     int
	windowCount      int
	lastIndex        int
	primed           bool
	unbiased         bool
	barFunc          data.BarFunc
	quoteFunc        data.QuoteFunc
	tradeFunc        data.TradeFunc
}

// NewVariance returns an instnce of the Variance indicator created using supplied parameters.
func NewVariance(p *VarianceParams) (*Variance, error) {
	const (
		invalid = "invalid variance parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtn    = "var.%c(%d)"
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

	var name, desc string
	if p.IsUnbiased {
		name = fmt.Sprintf(fmtn, 's', length)
		desc = "Unbiased estimation of the sample variance " + name
	} else {
		name = fmt.Sprintf(fmtn, 'p', length)
		desc = "Estimation of the population variance " + name
	}

	return &Variance{
		name:         name,
		description:  desc,
		window:       make([]float64, length),
		windowLength: length,
		lastIndex:    length - 1,
		unbiased:     p.IsUnbiased,
		barFunc:      barFunc,
		quoteFunc:    quoteFunc,
		tradeFunc:    tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (v *Variance) IsPrimed() bool {
	v.mu.RLock()
	defer v.mu.RUnlock()

	return v.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the variance.
func (v *Variance) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.Variance,
		Outputs: []output.Metadata{
			{
				Kind:        int(VarianceValue),
				Type:        output.Scalar,
				Name:        v.name,
				Description: v.description,
			},
		},
	}
}

// Update updates the value of the variance, σ², given the next sample.
//
// Depending on the isUnbiased, the value is the unbiased sample variance or the population variance.
//
// The indicator is not primed during the first ℓ-1 updates.
//nolint: funlen
func (v *Variance) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	var value float64

	temp := sample
	wlen := float64(v.windowLength)

	v.mu.Lock()
	defer v.mu.Unlock()

	//nolint: nestif
	if v.primed {
		v.windowSum += temp
		temp *= temp
		v.windowSquaredSum += temp
		temp = v.window[0]
		v.windowSum -= temp
		temp *= temp
		v.windowSquaredSum -= temp

		if v.unbiased {
			temp = v.windowSum
			temp *= temp
			temp /= wlen
			value = v.windowSquaredSum - temp
			value /= float64(v.lastIndex)
		} else {
			temp = v.windowSum / wlen
			temp *= temp
			value = v.windowSquaredSum/wlen - temp
		}

		for i := 0; i < v.lastIndex; i++ {
			v.window[i] = v.window[i+1]
		}

		v.window[v.lastIndex] = sample
	} else {
		v.windowSum += temp
		v.window[v.windowCount] = temp
		temp *= temp
		v.windowSquaredSum += temp

		v.windowCount++
		if v.windowLength == v.windowCount {
			v.primed = true
			if v.unbiased {
				temp = v.windowSum
				temp *= temp
				temp /= wlen
				value = v.windowSquaredSum - temp
				value /= float64(v.lastIndex)
			} else {
				temp = v.windowSum / wlen
				temp *= temp
				value = v.windowSquaredSum/wlen - temp
			}
		} else {
			return math.NaN()
		}
	}

	return value
}

// UpdateScalar updates the indicator given the next scalar sample.
func (v *Variance) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: v.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (v *Variance) UpdateBar(sample *data.Bar) indicator.Output {
	return v.UpdateScalar(&data.Scalar{Time: sample.Time, Value: v.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (v *Variance) UpdateQuote(sample *data.Quote) indicator.Output {
	return v.UpdateScalar(&data.Scalar{Time: sample.Time, Value: v.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (v *Variance) UpdateTrade(sample *data.Trade) indicator.Output {
	return v.UpdateScalar(&data.Scalar{Time: sample.Time, Value: v.tradeFunc(sample)})
}
