package tillson

import (
	"fmt"
	"math"
	"sync"

	"mbg/trading/data"                        //nolint:depguard
	"mbg/trading/indicators/indicator"        //nolint:depguard
	"mbg/trading/indicators/indicator/output" //nolint:depguard
)

// https://store.traders.com/-v16-c01-005smo-pdf.html

// T2ExponentialMovingAverage (T2 Exponential Moving Average, T2, T2EMA)
// is a smoothing indicator with less lag than a straight exponential moving average.
//
// The T2 was developed by Tim Tillson and is described in the article:
//
//	❶ Technical Analysis of Stocks & Commodities v.16:1 (33-37), Smoothing Techniques For More Accurate Signals.
//
// The calculation is as follows:
//
//	EMA¹ᵢ = EMA(Pᵢ) = αPᵢ + (1-α)EMA¹ᵢ₋₁ = EMA¹ᵢ₋₁ + α(Pᵢ - EMA¹ᵢ₋₁), 0 < α ≤ 1
//	EMA²ᵢ = EMA(EMA¹ᵢ) = αEMA¹ᵢ + (1-α)EMA²ᵢ₋₁ = EMA²ᵢ₋₁ + α(EMA¹ᵢ - EMA²ᵢ₋₁), 0 < α ≤ 1
//	GDᵛᵢ = (1+ν)EMA¹ᵢ - νEMA²ᵢ = EMA¹ᵢ + ν(EMA¹ᵢ - EMA²ᵢ), 0 < ν ≤ 1
//	T2ᵢ = GDᵛᵢ(GDᵛᵢ)
//
// where GD stands for 'Generalized DEMA' with 'volume' ν. The default value of ν is 0.7.
// When ν=0, GD is just an EMA, and when ν=1, GD is DEMA. In between, GD is a cooler DEMA.
//
// If x< stands for the action of running a time series through an EMA,
// ƒ is our formula for Generalized Dema with 'volume' ν:
//
//	ƒ = (1+ν)x -νx²
//
// Running the filter though itself three times is equivalent to cubing ƒ:
//
//	v²x⁴ - 2v(1+ν)x³ + (1+ν)²x²
//
// The Metastock code for T2 is:
//
//	e1=Mov(P,periods,E)
//	e2=Mov(e1,periods,E)
//	e3=Mov(e2,periods,E)
//	e4=Mov(e3,periods,E)
//	c1=v²
//	c2=-2v(1+ν)
//	c3=(1+ν)²
//	t2=c1*e4+c2*e3+c3*e2
//
// The very first EMA value (the seed for subsequent values) is calculated differently.
// This implementation allows for two algorithms for this seed.
//
//	❶ Use a simple average of the first 'period'. This is the most widely documented approach.
//	❷ Use first sample value as a seed. This is used in Metastock.
type T2ExponentialMovingAverage struct {
	mu              sync.RWMutex
	name            string
	description     string
	smoothingFactor float64
	c1              float64
	c2              float64
	c3              float64
	sum             float64
	ema1            float64
	ema2            float64
	ema3            float64
	ema4            float64
	length          int
	length2         int
	length3         int
	length4         int
	count           int
	firstIsAverage  bool
	primed          bool
	barFunc         data.BarFunc
	quoteFunc       data.QuoteFunc
	tradeFunc       data.TradeFunc
}

// NewT2ExponentialMovingAverageLength returns an instnce of the indicator
// created using supplied parameters based on length.
func NewT2ExponentialMovingAverageLength(
	p *T2ExponentialMovingAverageLengthParams,
) (*T2ExponentialMovingAverage, error) {
	return newT2ExponentialMovingAverage(p.Length, math.NaN(), p.VolumeFactor,
		p.FirstIsAverage, p.BarComponent, p.QuoteComponent, p.TradeComponent)
}

// NewT2ExponentialMovingAverageSmoothingFactor returns an instnce of the indicator
// created using supplied parameters based on smoothing factor.
func NewT2ExponentialMovingAverageSmoothingFactor(
	p *T2ExponentialMovingAverageSmoothingFactorParams,
) (*T2ExponentialMovingAverage, error) {
	return newT2ExponentialMovingAverage(0, p.SmoothingFactor, p.VolumeFactor,
		p.FirstIsAverage, p.BarComponent, p.QuoteComponent, p.TradeComponent)
}

//nolint:funlen, cyclop
func newT2ExponentialMovingAverage(length int, alpha float64, v float64, firstIsAverage bool,
	bc data.BarComponent, qc data.QuoteComponent, tc data.TradeComponent,
) (*T2ExponentialMovingAverage, error) {
	const (
		invalid = "invalid t2 exponential moving average parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtnl   = "t2(%d, %.2f)"
		fmtna   = "t2(%.4f (%d), %.2f)"
		minlen  = 2
		two     = 2
		three   = 3
		four    = 4
		epsilon = 0.00000001
	)

	var (
		name      string
		err       error
		barFunc   data.BarFunc
		quoteFunc data.QuoteFunc
		tradeFunc data.TradeFunc
	)

	if v < 0. || v > 1. {
		return nil, fmt.Errorf(fmts, invalid, "volume factor should be in range [0, 1]")
	}

	if math.IsNaN(alpha) {
		if length < minlen {
			return nil, fmt.Errorf(fmts, invalid, "length should be greater than 1")
		}

		alpha = two / float64(1+length)
		name = fmt.Sprintf(fmtnl, length, v)
	} else {
		if alpha < 0. || alpha > 1. {
			return nil, fmt.Errorf(fmts, invalid, "smoothing factor should be in range [0, 1]")
		}

		if alpha < epsilon {
			alpha = epsilon
		}

		length = int(math.Round(two/alpha)) - 1
		name = fmt.Sprintf(fmtna, alpha, length, v)
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

	desc := "T2 exponential moving average " + name

	v1 := v + 1
	c1 := v * v
	c2 := -two * v * v1
	c3 := v1 * v1

	return &T2ExponentialMovingAverage{
		name:            name,
		description:     desc,
		smoothingFactor: alpha,
		c1:              c1,
		c2:              c2,
		c3:              c3,
		length:          length,
		length2:         two*length - 1,
		length3:         three*length - two,
		length4:         four*length - three,
		firstIsAverage:  firstIsAverage,
		barFunc:         barFunc,
		quoteFunc:       quoteFunc,
		tradeFunc:       tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *T2ExponentialMovingAverage) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the simple moving average.
func (s *T2ExponentialMovingAverage) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.T2ExponentialMovingAverage,
		Outputs: []output.Metadata{
			{
				Kind:        int(T2ExponentialMovingAverageValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
		},
	}
}

// Update updates the value of the indicator given the next sample.
func (s *T2ExponentialMovingAverage) Update(sample float64) float64 { //nolint:cyclop, funlen, gocognit
	if math.IsNaN(sample) {
		return sample
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	sf := s.smoothingFactor

	if s.primed {
		v1 := s.ema1
		v2 := s.ema2
		v3 := s.ema3
		v4 := s.ema4
		v1 += (sample - v1) * sf
		v2 += (v1 - v2) * sf
		v3 += (v2 - v3) * sf
		v4 += (v3 - v4) * sf
		s.ema1 = v1
		s.ema2 = v2
		s.ema3 = v3
		s.ema4 = v4

		return s.c1*v4 + s.c2*v3 + s.c3*v2
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
			s.ema1 += (sample - s.ema1) * sf
			s.sum += s.ema1

			if s.length2 == s.count {
				s.ema2 = s.sum / float64(s.length)
				s.sum = s.ema2
			}
		} else if s.length3 >= s.count {
			s.ema1 += (sample - s.ema1) * sf
			s.ema2 += (s.ema1 - s.ema2) * sf
			s.sum += s.ema2

			if s.length3 == s.count {
				s.ema3 = s.sum / float64(s.length)
				s.sum = s.ema3
			}
		} else { // if s.length4 >= s.count {
			s.ema1 += (sample - s.ema1) * sf
			s.ema2 += (s.ema1 - s.ema2) * sf
			s.ema3 += (s.ema2 - s.ema3) * sf
			s.sum += s.ema3

			if s.length4 == s.count {
				s.primed = true
				s.ema4 = s.sum / float64(s.length)

				return s.c1*s.ema4 + s.c2*s.ema3 + s.c3*s.ema3
			}
		}
	} else { // Metastock
		if s.count == 1 {
			s.ema1 = sample
		} else if s.length >= s.count {
			s.ema1 += (sample - s.ema1) * sf
			if s.length == s.count {
				s.ema2 = s.ema1
			}
		} else if s.length2 >= s.count {
			s.ema1 += (sample - s.ema1) * sf
			s.ema2 += (s.ema1 - s.ema2) * sf

			if s.length2 == s.count {
				s.ema3 = s.ema2
			}
		} else if s.length3 >= s.count {
			s.ema1 += (sample - s.ema1) * sf
			s.ema2 += (s.ema1 - s.ema2) * sf
			s.ema3 += (s.ema2 - s.ema3) * sf

			if s.length3 == s.count {
				s.ema4 = s.ema3
			}
		} else { // if s.length4 >= s.count {
			s.ema1 += (sample - s.ema1) * sf
			s.ema2 += (s.ema1 - s.ema2) * sf
			s.ema3 += (s.ema2 - s.ema3) * sf
			s.ema4 += (s.ema3 - s.ema4) * sf

			if s.length4 == s.count {
				s.primed = true

				return s.c1*s.ema4 + s.c2*s.ema3 + s.c3*s.ema3
			}
		}
	}

	return math.NaN()
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *T2ExponentialMovingAverage) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: s.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *T2ExponentialMovingAverage) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *T2ExponentialMovingAverage) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *T2ExponentialMovingAverage) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}
