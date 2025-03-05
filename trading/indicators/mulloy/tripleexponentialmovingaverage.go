package mulloy

import (
	"fmt"
	"math"
	"sync"

	"mbg/trading/data"                        //nolint:depguard
	"mbg/trading/indicators/indicator"        //nolint:depguard
	"mbg/trading/indicators/indicator/output" //nolint:depguard
)

// https://store.traders.com/-v12-c01-smoothi-pdf.html
// https://store.traders.com/-v12-c02-smoothi-pdf.html

// TripleExponentialMovingAverage computes the Triple Exponential Moving Average (TEMA),
// a smoothing indicator with less lag than a straight exponential moving average.
//
// The TEMA was developed by Patrick G. Mulloy and is described in two articles:
//
//	❶ Technical Analysis of Stocks & Commodities v.12:1 (11-19), Smoothing Data With Faster Moving Averages.
//	❷ Technical Analysis of Stocks & Commodities v.12:2 (72-80), Smoothing Data With Less Lag.
//
// The calculation is as follows:
//
//	EMA¹ᵢ = EMA(Pᵢ) = αPᵢ + (1-α)EMA¹ᵢ₋₁ = EMA¹ᵢ₋₁ + α(Pᵢ - EMA¹ᵢ₋₁), 0 < α ≤ 1
//	EMA²ᵢ = EMA(EMA¹ᵢ) = αEMA¹ᵢ + (1-α)EMA²ᵢ₋₁ = EMA²ᵢ₋₁ + α(EMA¹ᵢ - EMA²ᵢ₋₁), 0 < α ≤ 1
//	EMA³ᵢ = EMA(EMA²ᵢ) = αEMA²ᵢ + (1-α)EMA³ᵢ₋₁ = EMA³ᵢ₋₁ + α(EMA²ᵢ - EMA³ᵢ₋₁), 0 < α ≤ 1
//	TEMAᵢ = 3(EMA¹ᵢ - EMA²ᵢ) + EMA³ᵢ
//
// The very first EMA value (the seed for subsequent values) is calculated differently.
// This implementation allows for two algorithms for this seed.
//
//	❶ Use a simple average of the first 'period'. This is the most widely documented approach.
//	❷ Use first sample value as a seed. This is used in Metastock.
type TripleExponentialMovingAverage struct {
	mu              sync.RWMutex
	name            string
	description     string
	smoothingFactor float64
	sum             float64
	ema1            float64
	ema2            float64
	ema3            float64
	length          int
	length2         int
	length3         int
	count           int
	firstIsAverage  bool
	primed          bool
	barFunc         data.BarFunc
	quoteFunc       data.QuoteFunc
	tradeFunc       data.TradeFunc
}

// NewTripleExponentialMovingAverageLength returns an instnce of the indicator
// created using supplied parameters based on length.
func NewTripleExponentialMovingAverageLength(
	p *TripleExponentialMovingAverageLengthParams,
) (*TripleExponentialMovingAverage, error) {
	return newTripleExponentialMovingAverage(p.Length, math.NaN(), p.FirstIsAverage,
		p.BarComponent, p.QuoteComponent, p.TradeComponent)
}

// NewTripleExponentialMovingAverageSmoothingFactor returns an instnce of the indicator
// created using supplied parameters based on smoothing factor.
func NewTripleExponentialMovingAverageSmoothingFactor(
	p *TripleExponentialMovingAverageSmoothingFactorParams,
) (*TripleExponentialMovingAverage, error) {
	return newTripleExponentialMovingAverage(0, p.SmoothingFactor, p.FirstIsAverage,
		p.BarComponent, p.QuoteComponent, p.TradeComponent)
}

//nolint:funlen
func newTripleExponentialMovingAverage(length int, alpha float64, firstIsAverage bool,
	bc data.BarComponent, qc data.QuoteComponent, tc data.TradeComponent,
) (*TripleExponentialMovingAverage, error) {
	const (
		invalid = "invalid triple exponential moving average parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtnl   = "tema(%d)"
		fmtna   = "tema(%d, %.8f)"
		minlen  = 2
		two     = 2
		three   = 3
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
			return nil, fmt.Errorf(fmts, invalid, "length should be greater than 1")
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

	desc := "Triple exponential moving average " + name

	return &TripleExponentialMovingAverage{
		name:            name,
		description:     desc,
		smoothingFactor: alpha,
		length:          length,
		length2:         two*length - 1,
		length3:         three*length - two,
		firstIsAverage:  firstIsAverage,
		barFunc:         barFunc,
		quoteFunc:       quoteFunc,
		tradeFunc:       tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *TripleExponentialMovingAverage) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the simple moving average.
func (s *TripleExponentialMovingAverage) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.TripleExponentialMovingAverage,
		Outputs: []output.Metadata{
			{
				Kind:        int(TripleExponentialMovingAverageValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
		},
	}
}

// Update updates the value of the indicator given the next sample.
func (s *TripleExponentialMovingAverage) Update(sample float64) float64 { //nolint:cyclop, funlen
	const three = 3.

	if math.IsNaN(sample) {
		return sample
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.primed {
		sf := s.smoothingFactor
		v1 := s.ema1
		v2 := s.ema2
		v3 := s.ema3
		v1 += (sample - v1) * sf
		v2 += (v1 - v2) * sf
		v3 += (v2 - v3) * sf
		s.ema1 = v1
		s.ema2 = v2
		s.ema3 = v3

		return three*(v1-v2) + v3
	}

	s.count++
	if s.firstIsAverage { //nolint:nestif
		if s.count == 1 {
			s.sum = sample
		} else if s.length >= s.count {
			s.sum += sample
			if s.length == s.count {
				s.ema1 = s.sum / float64(s.length)
				s.sum = s.ema1
			}
		} else if s.length2 >= s.count {
			s.ema1 += (sample - s.ema1) * s.smoothingFactor
			s.sum += s.ema1

			if s.length2 == s.count {
				s.ema2 = s.sum / float64(s.length)
				s.sum = s.ema2
			}
		} else { // if s.length3 >= s.count {
			s.ema1 += (sample - s.ema1) * s.smoothingFactor
			s.ema2 += (s.ema1 - s.ema2) * s.smoothingFactor
			s.sum += s.ema2

			if s.length3 == s.count {
				s.primed = true
				s.ema3 = s.sum / float64(s.length)

				return three*(s.ema1-s.ema2) + s.ema3
			}
		}
	} else { // Metastock
		if s.count == 1 {
			s.ema1 = sample
		} else if s.length >= s.count {
			s.ema1 += (sample - s.ema1) * s.smoothingFactor
			if s.length == s.count {
				s.ema2 = s.ema1
			}
		} else if s.length2 >= s.count {
			s.ema1 += (sample - s.ema1) * s.smoothingFactor
			s.ema2 += (s.ema1 - s.ema2) * s.smoothingFactor

			if s.length2 == s.count {
				s.ema3 = s.ema2
			}
		} else { // if s.length3 >= s.count {
			s.ema1 += (sample - s.ema1) * s.smoothingFactor
			s.ema2 += (s.ema1 - s.ema2) * s.smoothingFactor
			s.ema3 += (s.ema2 - s.ema3) * s.smoothingFactor

			if s.length3 == s.count {
				s.primed = true

				return three*(s.ema1-s.ema2) + s.ema3
			}
		}
	}

	return math.NaN()
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *TripleExponentialMovingAverage) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: s.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *TripleExponentialMovingAverage) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *TripleExponentialMovingAverage) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *TripleExponentialMovingAverage) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}
