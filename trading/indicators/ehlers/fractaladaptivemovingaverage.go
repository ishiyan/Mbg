package ehlers

import (
	"fmt"
	"math"
	"sync"
	"time"

	"mbg/trading/data"                        //nolint:depguard
	"mbg/trading/indicators/indicator"        //nolint:depguard
	"mbg/trading/indicators/indicator/output" //nolint:depguard
)

// FractalAdaptiveMovingAverage (Ehler's fractal adaptive moving average, FRAMA)
// is an EMA with the smoothing factor, α, being changed with each new sample:
//
// FRAMAᵢ = αᵢPᵢ + (1 - αᵢ)*FRAMAᵢ₋₁,  αs ≤ αᵢ ≤ 1
//
// Here the αs is the slowest α (default suggested value is 0.01 or equivalent
// length of 199 samples).
//
// The concept of FRAMA is to relate the fractal dimension FDᵢ, calculated on a window
// samples, to the EMA smoothing factor αᵢ, thus making the EMA adaptive.
//
// This dependency is defined as follows:
//
//	αᵢ = exp(-w(FDᵢ - 1)),  1 ≤ FDᵢ ≤ 2,
//
//	w = ln(αs)
//
//	or, given the length ℓs = 2/αs - 1,
//
//	w = ln(2/(ℓs + 1))
//
// The fractal dimension varies over the range from 1 to 2.
//
// When FDᵢ = 1 (series forms a straight line), the exponent is zero – which means that
// αᵢ = 1, and the output of the exponential moving average is equal to the input.
//
// When FDᵢ = 2 (series fills all plane, excibiting extreme volatility), the exponent
// is -w, which means that αᵢ = αs, and the output of the exponential moving average
// is equal to the output of the slowest moving average with αs.
//
// The fractal dimension is estimated by using a "box count" method.
// Since price samples are typically uniformly spaced, the box count is approximated
// by the average slope of the price curve. This is calculated as the highest price
// minus the lowest price within an interval, divided by the length of that interval.
//
// FDᵢ = (ln(N1+N2) − ln(N3)) / ln(2)
//
// N1 is calculated over the first half of the total lookback period ℓ as the
// (highest price - lowest price) during the first ℓ/2 bars, divided by ℓ/2.
//
// N2 is calculated over the second half of the total lookback period ℓ as the
// (lighest price - lowest price) during the second ℓ/2 bars (from ℓ/2 to ℓ-1 bars ago),
// divided by ℓ/2.
//
// N3 is calculated over the entire lookback period ℓ as the
// (highest price - lowest price) during the full ℓ bars, divided by ℓ.
//
// The box-counting method itself, often used to determine fractal dimension,
// is also known as the Minkowski–Bouligand dimension, named after Hermann Minkowski
// and Georges Bouligand. Here's how the box-counting method generally works.
//
//  1. Cover the pattern by overlaying it with a grid of square boxes,
//     all of the same side length (s).
//  2. Count the number of boxes (N) in the grid that contain at least
//     some part of the pattern.
//  3. Repeat steps 1 and 2 using grids with progressively smaller box sizes (s).
//  4. Calculate dimension by examining the relationship between N(s) and s,
//     typically as s approaches zero. Mathematically, it's defined as the limit:
//     FD = lim(s→0) [ln(N(s)) / ln(1/s)]
//
// Reference:
//
//	Falconer, K. (2014). Fractal geometry: Mathematical foundations and applications (3rd ed) Wiley.
//	Ehlers, John F. (2005). Fractal Adaptive Moving Average. Technical Analysis of Stocks & Commodities, 23(10), 81–82.
//	Ehlers, John F. (2006). FRAMA – Fractal Adaptive Moving Average, https://www.mesasoftware.com/papers/FRAMA.pdf.
type FractalAdaptiveMovingAverage struct {
	mu               sync.RWMutex
	name             string
	description      string
	nameFdim         string
	descriptionFdim  string
	alphaSlowest     float64
	scalingFactor    float64
	fractalDimension float64
	value            float64
	length           int
	lengthMinOne     int
	halfLength       int
	windowCount      int
	windowHigh       []float64
	windowLow        []float64
	primed           bool
	barFunc          data.BarFunc
	quoteFunc        data.QuoteFunc
	tradeFunc        data.TradeFunc
	isLogging        bool
}

// NewFractaldaptiveMovingAverageLength returns an instnce of the indicator
// created using supplied parameters.
func NewFractalAdaptiveMovingAverage( //nolint:funlen
	params *FractalAdaptiveMovingAverageParams,
) (*FractalAdaptiveMovingAverage, error) {
	const (
		invalid = "invalid fractal adaptive moving average parameters"
		fmtl    = "%s: length should be an even integer larger than 1"
		fmta    = "%s: slowest smoothing factor should be in range [0, 1]"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtnl   = "frama(%d, %.3f)"
		fmtnld  = "framaDim(%d, %.3f)"
		two     = 2
		descr   = "Fractal adaptive moving average "
	)

	var (
		name      string
		nameFdim  string
		err       error
		barFunc   data.BarFunc
		quoteFunc data.QuoteFunc
		tradeFunc data.TradeFunc
	)

	if params.Length < two {
		return nil, fmt.Errorf(fmtl, invalid)
	}

	if params.SlowestSmoothingFactor < 0. || params.SlowestSmoothingFactor > 1. {
		return nil, fmt.Errorf(fmta, invalid)
	}

	length := params.Length
	if length%2 != 0 {
		length++
	}

	halfLength := length / two

	name = fmt.Sprintf(fmtnl, length, params.SlowestSmoothingFactor)
	nameFdim = fmt.Sprintf(fmtnld, length, params.SlowestSmoothingFactor)

	if barFunc, err = data.BarComponentFunc(params.BarComponent); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	if quoteFunc, err = data.QuoteComponentFunc(params.QuoteComponent); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	if tradeFunc, err = data.TradeComponentFunc(params.TradeComponent); err != nil {
		return nil, fmt.Errorf(fmtw, invalid, err)
	}

	return &FractalAdaptiveMovingAverage{
		name:             name,
		description:      descr + name,
		nameFdim:         nameFdim,
		descriptionFdim:  descr + nameFdim,
		length:           length,
		lengthMinOne:     length - 1,
		halfLength:       halfLength,
		windowHigh:       make([]float64, length),
		windowLow:        make([]float64, length),
		windowCount:      0,
		alphaSlowest:     params.SlowestSmoothingFactor,
		scalingFactor:    math.Log(params.SlowestSmoothingFactor),
		fractalDimension: math.NaN(),
		value:            math.NaN(),
		primed:           false,
		barFunc:          barFunc,
		quoteFunc:        quoteFunc,
		tradeFunc:        tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *FractalAdaptiveMovingAverage) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the moving average.
func (s *FractalAdaptiveMovingAverage) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.FractalAdaptiveMovingAverage,
		Outputs: []output.Metadata{
			{
				Kind:        int(FractalAdaptiveMovingAverageValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
			{
				Kind:        int(FractalAdaptiveMovingAverageValueFdim),
				Type:        output.Scalar,
				Name:        s.nameFdim,
				Description: s.descriptionFdim,
			},
		},
	}
}

// Update updates the value of the moving average given the next sample.
func (s *FractalAdaptiveMovingAverage) Update(sample, sampleHigh, sampleLow float64) float64 {
	if math.IsNaN(sampleHigh) || math.IsNaN(sampleLow) || math.IsNaN(sample) {
		return math.NaN()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.primed {
		s.windowCount++
		if s.isLogging {
			fmt.Println("windowCount", s.windowCount, "lows", s.windowLow)
			fmt.Println("windowCount", s.windowCount, "highs", s.windowHigh)
		}

		for i := range s.lengthMinOne { // - 1
			j := i + 1
			if s.isLogging {
				fmt.Println("i", i, "j", j, "low", s.windowLow[i], "<-", s.windowLow[j], "high", s.windowHigh[i], "<-", s.windowHigh[j])
			}
			s.windowHigh[i] = s.windowHigh[j]
			s.windowLow[i] = s.windowLow[j]
		}

		if s.isLogging {
			fmt.Println("i", s.lengthMinOne, "low", s.windowLow[s.lengthMinOne], "<-", sampleLow, "high", s.windowHigh[s.lengthMinOne], "<-", sampleHigh)
		}
		s.windowHigh[s.lengthMinOne] = sampleHigh
		s.windowLow[s.lengthMinOne] = sampleLow

		if s.isLogging {
			fmt.Println("lows ->", s.windowLow)
			fmt.Println("highs ->", s.windowHigh)
		}

		s.fractalDimension = s.estimateFractalDimension()
		alpha := s.estimateAlpha()
		// s.value += (sample - s.value) * alpha

		if s.isLogging {
			fmt.Println("windowCount", s.windowCount, "fdim", s.fractalDimension, "alpha", alpha, "sample", sample, "value", s.value, "-->", s.value+(sample-s.value)*alpha)
		}

		s.value += (sample - s.value) * alpha

		return s.value
	} else {
		s.windowHigh[s.windowCount] = sampleHigh
		s.windowLow[s.windowCount] = sampleLow

		s.windowCount++
		if s.windowCount == s.lengthMinOne {
			s.value = sample
		} else if s.windowCount == s.length {
			s.fractalDimension = s.estimateFractalDimension()
			alpha := s.estimateAlpha()
			// s.value += (sample - s.value) * alpha
			s.primed = true

			if s.isLogging {
				fmt.Println("windowCount", s.windowCount, "fdim", s.fractalDimension, "alpha", alpha, "sample", sample, "value", s.value, "-->", s.value+(sample-s.value)*alpha)
			}

			s.value += (sample - s.value) * alpha

			return s.value
		}
	}

	return math.NaN()
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *FractalAdaptiveMovingAverage) UpdateScalar(sample *data.Scalar) indicator.Output {
	v := sample.Value

	return s.updateEntity(sample.Time, v, v, v)
}

// UpdateBar updates the indicator given the next bar sample.
func (s *FractalAdaptiveMovingAverage) UpdateBar(sample *data.Bar) indicator.Output {
	v := s.barFunc(sample)

	return s.updateEntity(sample.Time, v, sample.High, sample.Low)
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *FractalAdaptiveMovingAverage) UpdateQuote(sample *data.Quote) indicator.Output {
	v := s.quoteFunc(sample)

	return s.updateEntity(sample.Time, v, sample.Ask, sample.Bid)
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *FractalAdaptiveMovingAverage) UpdateTrade(sample *data.Trade) indicator.Output {
	v := s.tradeFunc(sample)

	return s.updateEntity(sample.Time, v, v, v)
}

func (s *FractalAdaptiveMovingAverage) updateEntity(
	time time.Time, sample, sampleHigh, sampleLow float64,
) indicator.Output {
	const length = 2

	output := make([]any, length)
	frama := s.Update(sample, sampleHigh, sampleLow)

	fdim := s.fractalDimension
	if math.IsNaN(frama) {
		fdim = math.NaN()
	}

	i := 0
	output[i] = data.Scalar{Time: time, Value: frama}
	i++
	output[i] = data.Scalar{Time: time, Value: fdim}

	return output
}

//nolint:cyclop
func (s *FractalAdaptiveMovingAverage) estimateFractalDimension() float64 {
	minLowHalf := math.MaxFloat64
	maxHighHalf := math.SmallestNonzeroFloat64

	for i := range s.halfLength {
		l := s.windowLow[i]
		if minLowHalf > l {
			minLowHalf = l
		}

		h := s.windowHigh[i]
		if maxHighHalf < h {
			maxHighHalf = h
		}

		if s.isLogging {
			fmt.Println("i", i, "low", l, "high", h, "N1 minLow", minLowHalf, "maxHigh", maxHighHalf)
		}
	}

	rangeN1 := maxHighHalf - minLowHalf
	minLowFull := minLowHalf
	maxHighFull := maxHighHalf
	minLowHalf = math.MaxFloat64
	maxHighHalf = math.SmallestNonzeroFloat64

	// for i := s.halfLength; i < s.length; i++ {
	for j := range s.halfLength {
		i := j + s.halfLength
		l := s.windowLow[i]
		if minLowFull > l {
			minLowFull = l
		}

		if minLowHalf > l {
			minLowHalf = l
		}

		h := s.windowHigh[i]
		if maxHighFull < h {
			maxHighFull = h
		}

		if maxHighHalf < h {
			maxHighHalf = h
		}

		if s.isLogging {
			fmt.Println("i", i, "low", l, s.windowLow[i], "high", h, s.windowHigh[i], "N2 minLow", minLowHalf, "maxHigh", maxHighHalf, "N3 minLow", minLowFull, "maxHigh", maxHighFull)
		}
	}

	rangeN2 := maxHighHalf - minLowHalf
	rangeN3 := maxHighFull - minLowFull

	if s.isLogging {
		fmt.Println("range N1", rangeN1, "range N2", rangeN2, "range N3", rangeN3)
	}

	fdim := (math.Log((rangeN1+rangeN2)/float64(s.halfLength)) -
		math.Log(rangeN3/float64(s.length))) * math.Log2E

	const two = 2

	return math.Min(math.Max(fdim, 1), two)
}

func (s *FractalAdaptiveMovingAverage) estimateAlpha() float64 {
	factor := s.scalingFactor

	// We use the fractal dimension to dynamically change the alpha of an exponential moving average.
	// The fractal dimension varies over the range from 1 to 2.
	// Since the prices are log-normal, it seems reasonable to use an exponential function to relate
	// the fractal dimension to alpha.

	// An empirically chosen scaling in Ehlers’s method to map fractal dimension (1–2)
	// to the exponential α.
	alpha := math.Exp(factor * (s.fractalDimension - 1))

	// When the fractal dimension is 1, the exponent is zero – which means that alpha is 1, and
	// the output of the exponential moving average is equal to the input.

	// Limit alpha to vary only from αs to 1.
	/*alpha = s.alphaSlowest
	} else if alpha > 1 {
		alpha = 1
	}
	return alpha*/

	return math.Min(math.Max(alpha, s.alphaSlowest), 1)
}
