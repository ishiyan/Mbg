package mulloy

import (
	"fmt"
	"math"
	"sync"

	"mbg/trading/data"                        //nolint:depguard
	"mbg/trading/indicators/indicator"        //nolint:depguard
	"mbg/trading/indicators/indicator/output" //nolint:depguard
)

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
	sum1            float64
	sum2            float64
	sum3            float64
	value1          float64
	value2          float64
	value3          float64
	length          int
	length2         int
	length3         int
	count1          int
	count2          int
	count3          int
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
		length2:         two * length,
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
		v1 := s.value1
		v2 := s.value2
		v3 := s.value3
		v1 += (sample - v1) * sf
		v2 += (v1 - v2) * sf
		v3 += (v2 - v3) * sf
		// fmt.Printf("4; value1: %.4f -> %.4f\n", s.value1, v1)
		// fmt.Printf("4; value2: %.4f -> %.4f\n", s.value2, v2)
		// fmt.Printf("4; value3: %.4f -> %.4f\n", s.value3, v3)
		s.value1 = v1
		s.value2 = v2
		s.value3 = v3

		// fmt.Printf("4; -> %.4f\n", three*(v1-v2)+v3)
		return three*(v1-v2) + v3
	}

	if s.firstIsAverage { //nolint:nestif
		if s.length > s.count1 {
			s.sum1 += sample
			s.count1++

			if s.length == s.count1 {
				s.value1 = s.sum1 / float64(s.length)
				s.sum2 += s.value1
				s.count2++
			}
		} else if s.length > s.count2 {
			s.value1 += (sample - s.value1) * s.smoothingFactor
			s.sum2 += s.value1
			s.count2++

			if s.length == s.count2 {
				s.value2 = s.sum2 / float64(s.length)
				s.sum3 += s.value2
				s.count3++
			}
		} else {
			s.value1 += (sample - s.value1) * s.smoothingFactor
			s.value2 += (s.value1 - s.value2) * s.smoothingFactor
			s.sum3 += s.value2
			s.count3++

			if s.length == s.count3 {
				s.value3 = s.sum3 / float64(s.length)
				s.primed = true

				return three*(s.value1-s.value2) + s.value3
			}
		}
	} else { // Metastock
		if s.length > s.count1 {
			s.count1++
			// fmt.Printf("1; count1: %d -> %d\n", s.count1-1, s.count1)
			if s.count1 == 1 {
				s.value1 = sample
				// fmt.Printf("1; value1: 0 -> %.4f\n", s.value1)
			} else {
				// v := s.value1
				s.value1 += (sample - s.value1) * s.smoothingFactor
				// fmt.Printf("1; value1: %.4f -> %.4f\n", v, s.value1)
			}

			if s.length == s.count1 {
				s.value2 = s.value1
				// fmt.Printf("1; value2: 0 -> %.4f\n", s.value2)
			}
		} else if s.length2 > s.count1 {
			// fmt.Printf("2; count1: %d -> %d\n", s.count1, s.count1+1)
			// v := s.value1
			s.value1 += (sample - s.value1) * s.smoothingFactor
			// fmt.Printf("2; value1: %.4f -> %.4f\n", v, s.value1)
			// v = s.value2
			s.value2 += (s.value1 - s.value2) * s.smoothingFactor
			// fmt.Printf("2; value2: %.4f -> %.4f\n", v, s.value2)
			s.count1++

			if s.length2 == s.count1 {
				s.value3 = s.value2
				// fmt.Printf("2; value2: 0 -> %.4f\n", s.value2)
			}
		} else {
			// fmt.Printf("3; count1: %d -> %d\n", s.count1, s.count1+1)
			// v := s.value1
			s.value1 += (sample - s.value1) * s.smoothingFactor
			// fmt.Printf("3; value1: %.4f -> %.4f\n", v, s.value1)
			// v = s.value2
			s.value2 += (s.value1 - s.value2) * s.smoothingFactor
			// fmt.Printf("3; value2: %.4f -> %.4f\n", v, s.value2)
			// v = s.value3
			s.value3 += (s.value2 - s.value3) * s.smoothingFactor
			// fmt.Printf("3; value3: %.4f -> %.4f\n", v, s.value3)
			s.count1++

			if s.length3 == s.count1 {
				// fmt.Printf("3; primed -> %.4f\n", three*(s.value1-s.value2)+s.value3)
				s.primed = true

				return three*(s.value1-s.value2) + s.value3
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
