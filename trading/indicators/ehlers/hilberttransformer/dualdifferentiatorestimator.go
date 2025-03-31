package hilberttransformer

import (
	"math"
	"sync"
)

// DualDifferentiatorEstimator implements the Hilbert transformer of the
// WMA-smoothed and detrended data with the dual differentiator applied.
//
// John Ehlers, Rocket Science for Traders, Wiley, 2001, 0471405671, pp 70-74.
type DualDifferentiatorEstimator struct {
	mu                               sync.RWMutex
	smoothingLength                  int
	minPeriod                        int
	maxPeriod                        int
	alphaEmaQuadratureInPhase        float64
	alphaEmaPeriod                   float64
	warmUpPeriod                     int
	smoothingLengthPlusHtLengthMin1  int
	smoothingLengthPlus2HtLengthMin2 int
	smoothingLengthPlus3HtLengthMin3 int
	smoothingLengthPlus3HtLengthMin2 int
	smoothingLengthPlus3HtLengthMin1 int
	oneMinAlphaEmaQuadratureInPhase  float64
	oneMinAlphaEmaPeriod             float64
	rawValues                        []float64
	wmaFactors                       []float64
	wmaSmoothed                      []float64
	detrended                        []float64
	inPhase                          []float64
	quadrature                       []float64
	jInPhase                         []float64
	jQuadrature                      []float64
	count                            int
	smoothedInPhasePrevious          float64
	smoothedQuadraturePrevious       float64
	period                           float64
	isPrimed                         bool
	isWarmedUp                       bool
}

// NewDualDifferentiatorEstimator returns an instnce of the estimator using supplied parameters.
func NewDualDifferentiatorEstimator(
	p *CycleEstimatorParams,
) (*DualDifferentiatorEstimator, error) {
	err := verifyParameters(p)
	if err != nil {
		return nil, err
	}

	length := p.SmoothingLength
	alphaQuad := p.AlphaEmaQuadratureInPhase
	alphaPeriod := p.AlphaEmaPeriod

	smoothingLengthPlusHtLengthMin1 := length + htLength - 1
	smoothingLengthPlus2HtLengthMin2 := smoothingLengthPlusHtLengthMin1 + htLength - 1
	smoothingLengthPlus3HtLengthMin3 := smoothingLengthPlus2HtLengthMin2 + htLength - 1
	smoothingLengthPlus3HtLengthMin2 := smoothingLengthPlus3HtLengthMin3 + 1
	smoothingLengthPlus3HtLengthMin1 := smoothingLengthPlus3HtLengthMin2 + 1

	// These slices will be automatically filled with zeroes.
	wmaSmoothed := make([]float64, htLength)
	detrended := make([]float64, htLength)
	inPhase := make([]float64, htLength)
	quadrature := make([]float64, htLength)
	jInPhase := make([]float64, htLength)
	jQuadrature := make([]float64, htLength)
	rawValues := make([]float64, length)
	wmaFactors := make([]float64, length)

	fillWmaFactors(length, wmaFactors)

	return &DualDifferentiatorEstimator{
		smoothingLength:                  length,
		minPeriod:                        defaultMinPeriod,
		maxPeriod:                        defaultMaxPeriod,
		alphaEmaQuadratureInPhase:        alphaQuad,
		alphaEmaPeriod:                   alphaPeriod,
		warmUpPeriod:                     max(p.WarmUpPeriod, smoothingLengthPlus3HtLengthMin1),
		smoothingLengthPlusHtLengthMin1:  smoothingLengthPlusHtLengthMin1,
		smoothingLengthPlus2HtLengthMin2: smoothingLengthPlus2HtLengthMin2,
		smoothingLengthPlus3HtLengthMin3: smoothingLengthPlus3HtLengthMin3,
		smoothingLengthPlus3HtLengthMin2: smoothingLengthPlus3HtLengthMin2,
		smoothingLengthPlus3HtLengthMin1: smoothingLengthPlus3HtLengthMin1,
		oneMinAlphaEmaQuadratureInPhase:  1 - alphaQuad,
		oneMinAlphaEmaPeriod:             1 - alphaPeriod,
		rawValues:                        rawValues,
		wmaFactors:                       wmaFactors,
		wmaSmoothed:                      wmaSmoothed,
		detrended:                        detrended,
		inPhase:                          inPhase,
		quadrature:                       quadrature,
		jInPhase:                         jInPhase,
		jQuadrature:                      jQuadrature,
		count:                            0,
		smoothedInPhasePrevious:          0,
		smoothedQuadraturePrevious:       0,
		period:                           defaultMinPeriod,
		isPrimed:                         false,
		isWarmedUp:                       false,
	}, nil
}

// SmoothingLength returns the underlying WMA smoothing length in samples.
func (s *DualDifferentiatorEstimator) SmoothingLength() int {
	return s.smoothingLength
}

// MinPeriod returns the minimal cycle period.
func (s *DualDifferentiatorEstimator) MinPeriod() int {
	return s.minPeriod
}

// MaxPeriod returns the maximual cycle period.
func (s *DualDifferentiatorEstimator) MaxPeriod() int {
	return s.maxPeriod
}

// WarmUpPeriod returns the number of updates before the estimator
// is primed (MaxPeriod * 2 = 100).
func (s *DualDifferentiatorEstimator) WarmUpPeriod() int {
	return s.warmUpPeriod
}

// AlphaEmaQuadratureInPhase returns the value of α (0 < α ≤ 1)
// used in EMA to smooth the in-phase and quadrature components.
func (s *DualDifferentiatorEstimator) AlphaEmaQuadratureInPhase() float64 {
	return s.alphaEmaQuadratureInPhase
}

// AlphaEmaPeriod returns the value of α (0 < α ≤ 1)
// used in EMA to smooth the instantaneous period.
func (s *DualDifferentiatorEstimator) AlphaEmaPeriod() float64 {
	return s.alphaEmaPeriod
}

// Count returns the current count value.
func (s *DualDifferentiatorEstimator) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.count
}

// Primed indicates whether an instance is primed.
func (s *DualDifferentiatorEstimator) Primed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.isWarmedUp
}

// Period returns the current period value.
func (s *DualDifferentiatorEstimator) Period() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.period
}

// InPhase returns the current InPhase component value.
func (s *DualDifferentiatorEstimator) InPhase() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.inPhase[0]
}

// Quadrature returns the current Quadrature component value.
func (s *DualDifferentiatorEstimator) Quadrature() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.quadrature[0]
}

// Detrended returns the current detrended value.
func (s *DualDifferentiatorEstimator) Detrended() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.detrended[0]
}

// Smoothed returns the current WMA-smoothed value used by underlying Hilbert transformer.
//
// The linear-Weighted Moving Average has a window size of SmoothingLength.
func (s *DualDifferentiatorEstimator) Smoothed() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.wmaSmoothed[0]
}

// Update updates the estimator given the next sample value.
func (s *DualDifferentiatorEstimator) Update(sample float64) { //nolint:funlen, cyclop
	if math.IsNaN(sample) {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	const (
		twoPi = 2 * math.Pi
	)

	push(s.rawValues, sample)

	if s.isPrimed { //nolint:nestif
		if !s.isWarmedUp {
			s.count++
			if s.warmUpPeriod < s.count {
				s.isWarmedUp = true
			}
		}

		// The WMA is used to remove some high-frequency components before detrending the signal.
		push(s.wmaSmoothed, s.wma(s.rawValues))

		amplitudeCorrectionFactor := correctAmplitude(s.period)

		// Since we have an amplitude-corrected Hilbert transformer, and since we want to detrend
		// over its length, we simply use the Hilbert transformer itself as the detrender.
		push(s.detrended, ht(s.wmaSmoothed)*amplitudeCorrectionFactor)

		// Compute both the in-phase and quadrature components of the detrended signal.
		push(s.quadrature, ht(s.detrended)*amplitudeCorrectionFactor)
		push(s.inPhase, s.detrended[quadratureIndex])

		// Complex averaging: apply the Hilbert Transformer to both the in-phase and quadrature components.
		// This advances the phase of each component by 90°.
		push(s.jInPhase, ht(s.inPhase)*amplitudeCorrectionFactor)
		push(s.jQuadrature, ht(s.quadrature)*amplitudeCorrectionFactor)

		// Phasor addition for 3 bar averaging followed by exponential moving average smoothing.
		smoothedInPhase := s.emaQuadratureInPhase(s.inPhase[0]-s.jQuadrature[0], s.smoothedInPhasePrevious)
		smoothedQuadrature := s.emaQuadratureInPhase(s.quadrature[0]+s.jInPhase[0], s.smoothedQuadraturePrevious)

		// Dual Differential discriminator.
		discriminator := smoothedQuadrature*(smoothedInPhase-s.smoothedInPhasePrevious) -
			smoothedInPhase*(smoothedQuadrature-s.smoothedQuadraturePrevious)
		s.smoothedInPhasePrevious = smoothedInPhase
		s.smoothedQuadraturePrevious = smoothedQuadrature

		periodPrevious := s.period
		periodNew := twoPi * (smoothedInPhase*smoothedInPhase + smoothedQuadrature*smoothedQuadrature) / discriminator

		if !math.IsNaN(periodNew) && !math.IsInf(periodNew, 0) {
			s.period = periodNew
		}

		s.period = adjustPeriod(s.period, periodPrevious)

		// Exponential moving average smoothing of the period.
		s.period = s.emaPeriod(s.period, periodPrevious)
	} else { // Not primed.
		// On (smoothingLength)-th sample we calculate the first
		// WMA smoothed value and begin with the detrender.
		s.count++
		if s.smoothingLength > s.count { // count < 4
			return
		}

		push(s.wmaSmoothed, s.wma(s.rawValues)) // count >= 4

		if s.smoothingLengthPlusHtLengthMin1 > s.count { // count < 10
			return
		}

		amplitudeCorrectionFactor := correctAmplitude(s.period) // count >= 10
		push(s.detrended, ht(s.wmaSmoothed)*amplitudeCorrectionFactor)

		if s.smoothingLengthPlus2HtLengthMin2 > s.count { // count < 16
			return
		}

		push(s.quadrature, ht(s.detrended)*amplitudeCorrectionFactor) // count >= 16
		push(s.inPhase, s.detrended[quadratureIndex])

		if s.smoothingLengthPlus3HtLengthMin3 > s.count { // count < 22
			return
		}

		push(s.jInPhase, ht(s.inPhase)*amplitudeCorrectionFactor) // count >= 22
		push(s.jQuadrature, ht(s.quadrature)*amplitudeCorrectionFactor)

		if s.smoothingLengthPlus3HtLengthMin3 == s.count { // count == 22
			s.smoothedInPhasePrevious = s.inPhase[0] - s.jQuadrature[0]
			s.smoothedQuadraturePrevious = s.quadrature[0] + s.jInPhase[0]

			return
		}

		smoothedInPhase := s.emaQuadratureInPhase(s.inPhase[0]-s.jQuadrature[0], s.smoothedInPhasePrevious) // count >= 23
		smoothedQuadrature := s.emaQuadratureInPhase(s.quadrature[0]+s.jInPhase[0], s.smoothedQuadraturePrevious)

		discriminator := smoothedQuadrature*(smoothedInPhase-s.smoothedInPhasePrevious) -
			smoothedInPhase*(smoothedQuadrature-s.smoothedQuadraturePrevious)
		s.smoothedInPhasePrevious = smoothedInPhase
		s.smoothedQuadraturePrevious = smoothedQuadrature

		if s.smoothingLengthPlus3HtLengthMin2 == s.count { // count == 23
			s.period = twoPi * (smoothedInPhase*smoothedInPhase +
				smoothedQuadrature*smoothedQuadrature) / discriminator
			if math.IsNaN(s.period) || math.IsInf(s.period, 0) {
				s.period = float64(s.minPeriod)
			}

			s.period = adjustPeriod(s.period, s.period)

			return
		}

		periodPrevious := s.period
		periodNew := twoPi * (smoothedInPhase*smoothedInPhase + smoothedQuadrature*smoothedQuadrature) / discriminator

		if !math.IsNaN(periodNew) && !math.IsInf(periodNew, 0) {
			s.period = periodNew
		}

		s.period = adjustPeriod(s.period, periodPrevious)
		s.period = s.emaPeriod(s.period, periodPrevious)
		s.isPrimed = true
	}
}

func (s *DualDifferentiatorEstimator) wma(array []float64) float64 {
	value := 0.
	for i := range s.smoothingLength {
		value += s.wmaFactors[i] * array[i]
	}

	return value
}

func (s *DualDifferentiatorEstimator) emaQuadratureInPhase(value, valuePrevious float64) float64 {
	return s.alphaEmaQuadratureInPhase*value + s.oneMinAlphaEmaQuadratureInPhase*valuePrevious
}

func (s *DualDifferentiatorEstimator) emaPeriod(value, valuePrevious float64) float64 {
	return s.alphaEmaPeriod*value + s.oneMinAlphaEmaPeriod*valuePrevious
}
