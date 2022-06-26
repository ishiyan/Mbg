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

// ExponentialMovingAverage computes the exponential, or exponentially weighted, moving average (EMA).
//
// Given a constant smoothing percentage factor 0 < α ≤ 1, EMA is calculated by applying a constant
// smoothing factor α to a difference of today's sample and yesterday's EMA value:
//    EMAᵢ = αPᵢ + (1-α)EMAᵢ₋₁ = EMAᵢ₋₁ + α(Pᵢ - EMAᵢ₋₁), 0 < α ≤ 1.
// Thus, the weighting for each older sample is given by the geometric progression 1, α, α², α³, …,
// giving much more importance to recent observations while not discarding older ones: all data
// previously used are always part of the new EMA value.
//
// α may be expressed as a percentage, so a smoothing factor of 10% is equivalent to α = 0.1. A higher α
// discounts older observations faster. Alternatively, α may be expressed in terms of ℓ time periods (length),
// where:
//    α = 2 / (ℓ + 1) and ℓ = 2/α - 1.
// The indicator is not primed during the first ℓ-1 updates.
//
// The 12- and 26-day EMAs are the most popular short-term averages, and they are used to create indicators
// like MACD and PPO. In general, the 50- and 200-day EMAs are used as signals of long-term trends.
//
// The very first EMA value (the seed for subsequent values) is calculated differently. This implementation
// allows for two algorithms for this seed.
// ❶ Use a simple average of the first 'period'. This is the most widely documented approach.
// ❷ Use first sample value as a seed. This is used in Metastock.
type ExponentialMovingAverage struct {
	mu              sync.RWMutex
	name            string
	description     string
	value           float64
	sum             float64
	smoothingFactor float64
	length          int
	count           int
	firstIsAverage  bool
	primed          bool
	barFunc         data.BarFunc
	quoteFunc       data.QuoteFunc
	tradeFunc       data.TradeFunc
}

// NewExponentialMovingAverageLength returns an instnce of the indicator
// created using supplied parameters with nased on length.
func NewExponentialMovingAverageLength(
	p *ExponentialMovingAverageLengthParams,
) (*ExponentialMovingAverage, error) {
	return newExponentialMovingAverage(p.Length, math.NaN(), p.FirstIsAverage,
		p.BarComponent, p.QuoteComponent, p.TradeComponent)
}

// NewExponentialMovingAverageSmoothingFactor returns an instnce of the indicator
// created using supplied parameters based on smoothing factor.
func NewExponentialMovingAverageSmoothingFactor(
	p *ExponentialMovingAverageSmoothingFactorParams,
) (*ExponentialMovingAverage, error) {
	return newExponentialMovingAverage(0, p.SmoothingFactor, p.FirstIsAverage,
		p.BarComponent, p.QuoteComponent, p.TradeComponent)
}

//nolint:funlen
func newExponentialMovingAverage(length int, alpha float64, firstIsAverage bool,
	bc data.BarComponent, qc data.QuoteComponent, tc data.TradeComponent,
) (*ExponentialMovingAverage, error) {
	const (
		invalid = "invalid exponential moving average parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtnl   = "ema(%d)"
		fmtna   = "ema(%d, %.8f)"
		minlen  = 1
		two     = 2.
		epsilon = 0.00000001
	)

	var (
		name      string
		err       error
		barFunc   data.BarFunc
		quoteFunc data.QuoteFunc
		tradeFunc data.TradeFunc
	)

	if math.IsNaN(alpha) {
		if length < minlen {
			return nil, fmt.Errorf(fmts, invalid, "length should be positive")
		}

		alpha = two / float64(1+length)
		name = fmt.Sprintf(fmtnl, length)
	} else {
		if alpha < 0. || alpha > 1. {
			return nil, fmt.Errorf(fmts, invalid, "smoothing factor should be in range [0, 1]")
		}

		if alpha < epsilon {
			alpha = epsilon
		}

		length = int(math.Round(two/alpha)) - 1
		name = fmt.Sprintf(fmtna, length, alpha)
	}

	if barFunc, err = data.BarComponentFunc(bc); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	if quoteFunc, err = data.QuoteComponentFunc(qc); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	if tradeFunc, err = data.TradeComponentFunc(tc); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	desc := "Exponential moving average " + name

	return &ExponentialMovingAverage{
		name:            name,
		description:     desc,
		smoothingFactor: alpha,
		length:          length,
		firstIsAverage:  firstIsAverage,
		barFunc:         barFunc,
		quoteFunc:       quoteFunc,
		tradeFunc:       tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *ExponentialMovingAverage) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the simple moving average.
func (s *ExponentialMovingAverage) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.ExponentialMovingAverage,
		Outputs: []output.Metadata{
			{
				Kind:        int(ExponentialMovingAverageValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
		},
	}
}

// Update updatess the value of the exponential moving average given the next sample.
//
// The indicator is not primed during the first ℓ-1 updates.
func (s *ExponentialMovingAverage) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	temp := sample

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.primed { //nolint:nestif
		s.value += (temp - s.value) * s.smoothingFactor
	} else {
		s.count++
		if s.firstIsAverage {
			s.sum += temp
			if s.length != s.count {
				return math.NaN()
			}

			s.value = s.sum / float64(s.length)
		} else {
			if s.count == 1 {
				s.value = temp
			} else {
				s.value += (temp - s.value) * s.smoothingFactor
			}

			if s.count != s.length {
				return math.NaN()
			}
		}

		s.primed = true
	}

	return s.value
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *ExponentialMovingAverage) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: s.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *ExponentialMovingAverage) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *ExponentialMovingAverage) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *ExponentialMovingAverage) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}
