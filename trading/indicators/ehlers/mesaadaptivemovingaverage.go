package ehlers

import (
	"fmt"
	"math"
	"sync"

	"mbg/trading/data"                                        //nolint:depguard
	"mbg/trading/indicators/ehlers/hilberttransformer"        //nolint:depguard
	"mbg/trading/indicators/indicator"                        //nolint:depguard
	"mbg/trading/indicators/indicator/output"                 //nolint:depguard
	outputdata "mbg/trading/indicators/indicator/output/data" //nolint:depguard
)

// MesaAdaptiveMovingAverage (Ehler's Mesa adaptive moving average, or Mother of All Moving Averages, MAMA)
// is an EMA with the smoothing factor, α, being changed with each new sample within the fast and the slow
// limit boundaries which are the constant parameters of MAMA:
//
// MAMAᵢ = αᵢPᵢ + (1 - αᵢ)*MAMAᵢ₋₁,  αs ≤ αᵢ ≤ αf
//
// The αf is the α of the fast (shortest, default suggested value 0.5 or 3 samples) limit boundary.
//
// The αs is the α of the slow (longest, default suggested value 0.05 or 39 samples) limit boundary.
//
// The concept of MAMA is to relate the phase rate of change, as measured by a Hilbert Transformer
// estimator, to the EMA smoothing factor α, thus making the EMA adaptive.
//
// The cycle phase is computed from the arctangent of the ratio of the Quadrature component to the
// InPhase component. The rate of change is obtained by taking the difference of successive phase
// measurements. The α is computed as the fast limit αf divided by the phase rate of change.
// Any time there is a negative phase rate of change the value of α is set to the fast limit αf;
// if the phase rate of change is large, the α is bounded at the slow limit αs.
//
// The Following Adaptive Moving Average (FAMA) is produced by applying the MAMA to the first
// MAMA indicator.
//
// By using an α in FAMA that is the half the value of the α in MAMA, the FAMA has steps in
// time synchronization with MAMA, but the vertical movement is not as great.
//
// As a result, MAMA and FAMA do not cross unless there has been a major change in the
// market direction. This suggests an adaptive moving average crossover system that is
// virtually free of whipsaw trades.
//
// Reference:
// John Ehlers, Rocket Science for Traders, Wiley, 2001, 0471405671, pp 177-184.
type MesaAdaptiveMovingAverage struct {
	mu              sync.RWMutex
	name            string
	description     string
	nameFama        string
	descriptionFama string
	nameBand        string
	descriptionBand string
	alphaFastLimit  float64
	alphaSlowLimit  float64
	previousPhase   float64
	mama            float64
	fama            float64
	htce            hilberttransformer.CycleEstimator
	isPhaseCached   bool
	primed          bool
	barFunc         data.BarFunc
	quoteFunc       data.QuoteFunc
	tradeFunc       data.TradeFunc
}

// NewMesadaptiveMovingAverageDefault returns an instnce of the indicator
// created using default values of the parameters.
func NewMesaAdaptiveMovingAverageDefault() (*MesaAdaptiveMovingAverage, error) {
	const (
		fastLimitLength           = 3
		slowLimitLength           = 39
		smoothingLength           = 4
		alphaEmaQuadratureInPhase = 0.2
		AlphaEmaPeriod            = 0.2
		warmUpPeriod              = 0
	)

	return newMesaAdaptiveMovingAverage(
		hilberttransformer.HomodyneDiscriminator,
		&hilberttransformer.CycleEstimatorParams{
			SmoothingLength:           smoothingLength,
			AlphaEmaQuadratureInPhase: alphaEmaQuadratureInPhase,
			AlphaEmaPeriod:            AlphaEmaPeriod,
			WarmUpPeriod:              warmUpPeriod,
		},
		fastLimitLength, slowLimitLength,
		math.NaN(), math.NaN(),
		data.BarMedianPrice, data.QuoteMidPrice, data.TradePrice)
}

// NewMesadaptiveMovingAverageLength returns an instnce of the indicator
// created using supplied parameters based on length.
func NewMesaAdaptiveMovingAverageLength(
	p *MesaAdaptiveMovingAverageLengthParams,
) (*MesaAdaptiveMovingAverage, error) {
	return newMesaAdaptiveMovingAverage(
		p.EstimatorType, &p.EstimatorParams,
		p.FastLimitLength, p.SlowLimitLength,
		math.NaN(), math.NaN(),
		p.BarComponent, p.QuoteComponent, p.TradeComponent)
}

// NewMesaAdaptiveMovingAverageSmoothingFactor returns an instnce of the indicator
// created using supplied parameters based on smoothing factor.
func NewMesaAdaptiveMovingAverageSmoothingFactor(
	p *MesaAdaptiveMovingAverageSmoothingFactorParams,
) (*MesaAdaptiveMovingAverage, error) {
	return newMesaAdaptiveMovingAverage(
		p.EstimatorType, &p.EstimatorParams,
		0, 0,
		p.FastLimitSmoothingFactor, p.SlowLimitSmoothingFactor,
		p.BarComponent, p.QuoteComponent, p.TradeComponent)
}

//nolint:funlen,cyclop
func newMesaAdaptiveMovingAverage(
	estimatorType hilberttransformer.CycleEstimatorType,
	estimatorParams *hilberttransformer.CycleEstimatorParams,
	fastLimitLength int, slowLimitLength int,
	fastLimitSmoothingFactor float64, slowLimitSmoothingFactor float64,
	bc data.BarComponent, qc data.QuoteComponent, tc data.TradeComponent,
) (*MesaAdaptiveMovingAverage, error) {
	const (
		invalid = "invalid mesa adaptive moving average parameters"
		fmtl    = "%s: %s length should be larger than 1"
		fmta    = "%s: %s smoothing factor should be in range [0, 1]"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtnl   = "mama(%d, %d%s)"
		fmtna   = "mama(%.4f, %.4f%s)"
		fmtnlf  = "fama(%d, %d%s)"
		fmtnaf  = "fama(%.4f, %.4f%s)"
		fmtnlb  = "mama-fama(%d, %d%s)"
		fmtnab  = "mama-fama(%.4f, %.4f%s)"
		two     = 2
		four    = 4
		alpha   = 0.2
		epsilon = 0.00000001
		flim    = "fast limit"
		slim    = "slow limit"
		descr   = "Mesa adaptive moving average "
	)

	var (
		name      string
		nameFama  string
		nameBand  string
		err       error
		barFunc   data.BarFunc
		quoteFunc data.QuoteFunc
		tradeFunc data.TradeFunc
	)

	estimator, err := hilberttransformer.NewCycleEstimator(estimatorType, estimatorParams)
	if err != nil {
		return nil, err
	}

	estimmatorMoniker := ""
	if estimatorType != hilberttransformer.HomodyneDiscriminator ||
		estimatorParams.SmoothingLength != four ||
		estimatorParams.AlphaEmaQuadratureInPhase != alpha ||
		estimatorParams.AlphaEmaPeriod != alpha {
		estimmatorMoniker = hilberttransformer.EstimatorMoniker(estimatorType, estimator)
		if len(estimmatorMoniker) > 0 {
			estimmatorMoniker = fmt.Sprintf(", %s", estimmatorMoniker)
		}
	}

	if math.IsNaN(fastLimitSmoothingFactor) { //nolint:nestif
		if fastLimitLength < two {
			return nil, fmt.Errorf(fmtl, invalid, flim)
		}

		if slowLimitLength < two {
			return nil, fmt.Errorf(fmtl, invalid, slim)
		}

		fastLimitSmoothingFactor = two / float64(1+fastLimitLength)
		slowLimitSmoothingFactor = two / float64(1+slowLimitLength)

		name = fmt.Sprintf(fmtnl,
			fastLimitLength, slowLimitLength, estimmatorMoniker)
		nameFama = fmt.Sprintf(fmtnlf,
			fastLimitLength, slowLimitLength, estimmatorMoniker)
		nameBand = fmt.Sprintf(fmtnlb,
			fastLimitLength, slowLimitLength, estimmatorMoniker)
	} else {
		if fastLimitSmoothingFactor < 0. || fastLimitSmoothingFactor > 1. {
			return nil, fmt.Errorf(fmta, invalid, flim)
		}

		if slowLimitSmoothingFactor < 0. || slowLimitSmoothingFactor > 1. {
			return nil, fmt.Errorf(fmta, invalid, slim)
		}

		if fastLimitSmoothingFactor < epsilon {
			fastLimitSmoothingFactor = epsilon
		}

		if slowLimitSmoothingFactor < epsilon {
			slowLimitSmoothingFactor = epsilon
		}

		name = fmt.Sprintf(fmtna,
			fastLimitSmoothingFactor, slowLimitSmoothingFactor, estimmatorMoniker)
		nameFama = fmt.Sprintf(fmtnaf,
			fastLimitSmoothingFactor, slowLimitSmoothingFactor, estimmatorMoniker)
		nameBand = fmt.Sprintf(fmtnab,
			fastLimitSmoothingFactor, slowLimitSmoothingFactor, estimmatorMoniker)
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

	return &MesaAdaptiveMovingAverage{
		name:            name,
		description:     descr + name,
		nameFama:        nameFama,
		descriptionFama: descr + nameFama,
		nameBand:        nameBand,
		descriptionBand: descr + nameBand,
		alphaFastLimit:  fastLimitSmoothingFactor,
		alphaSlowLimit:  slowLimitSmoothingFactor,
		htce:            estimator,
		barFunc:         barFunc,
		quoteFunc:       quoteFunc,
		tradeFunc:       tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (s *MesaAdaptiveMovingAverage) IsPrimed() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the moving average.
func (s *MesaAdaptiveMovingAverage) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.MesaAdaptiveMovingAverage,
		Outputs: []output.Metadata{
			{
				Kind:        int(MesaAdaptiveMovingAverageValue),
				Type:        output.Scalar,
				Name:        s.name,
				Description: s.description,
			},
			{
				Kind:        int(MesaAdaptiveMovingAverageValueFama),
				Type:        output.Scalar,
				Name:        s.nameFama,
				Description: s.descriptionFama,
			},
			{
				Kind:        int(MesaAdaptiveMovingAverageBand),
				Type:        output.Band,
				Name:        s.nameBand,
				Description: s.descriptionBand,
			},
		},
	}
}

// Update updates the value of the moving average given the next sample.
func (s *MesaAdaptiveMovingAverage) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.htce.Update(sample)

	if s.primed {
		return s.calculate(sample)
	}

	if s.htce.Primed() {
		if s.isPhaseCached {
			s.primed = true

			return s.calculate(sample)
		}

		s.isPhaseCached = true
		s.previousPhase = s.calculatePhase()
		s.mama = sample
		s.fama = sample
	}

	return math.NaN()
}

// UpdateScalar updates the indicator given the next scalar sample.
func (s *MesaAdaptiveMovingAverage) UpdateScalar(sample *data.Scalar) indicator.Output {
	const length = 3

	output := make([]any, length)
	mama := s.Update(sample.Value)

	fama := s.fama
	if math.IsNaN(mama) {
		fama = math.NaN()
	}

	i := 0
	output[i] = data.Scalar{Time: sample.Time, Value: mama}
	i++
	output[i] = data.Scalar{Time: sample.Time, Value: fama}
	i++
	output[i] = outputdata.Band{Time: sample.Time, Upper: mama, Lower: fama}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (s *MesaAdaptiveMovingAverage) UpdateBar(sample *data.Bar) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (s *MesaAdaptiveMovingAverage) UpdateQuote(sample *data.Quote) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (s *MesaAdaptiveMovingAverage) UpdateTrade(sample *data.Trade) indicator.Output {
	return s.UpdateScalar(&data.Scalar{Time: sample.Time, Value: s.tradeFunc(sample)})
}

func (s *MesaAdaptiveMovingAverage) calculatePhase() float64 {
	if s.htce.InPhase() == 0 {
		return s.previousPhase
	}

	const rad2deg = 180.0 / math.Pi

	// The cycle phase is computed from the arctangent of the ratio
	// of the Quadrature component to the InPhase component.
	// phase := math.Atan2(s.htce.InPhase(), s.htce.Quadrature()) * rad2deg
	phase := math.Atan(s.htce.Quadrature()/s.htce.InPhase()) * rad2deg
	if !math.IsNaN(phase) && !math.IsInf(phase, 0) {
		return phase
	}

	return s.previousPhase
}

func (s *MesaAdaptiveMovingAverage) calculateMama(sample float64) float64 {
	phase := s.calculatePhase()

	// The phase rate of change is obtained by taking the
	// difference of successive previousPhase measurements.
	phaseRateOfChange := s.previousPhase - phase
	s.previousPhase = phase

	// Any negative rate change is theoretically impossible
	// because phase must advance as the time increases.
	// We therefore limit all rate changes of phase to be
	// no less than unity.
	if phaseRateOfChange < 1 {
		phaseRateOfChange = 1
	}

	// The α is computed as the fast limit divided
	// by the phase rate of change.
	alpha := s.alphaFastLimit / phaseRateOfChange
	if alpha < s.alphaSlowLimit {
		alpha = s.alphaSlowLimit
	}

	if alpha > s.alphaFastLimit {
		alpha = s.alphaFastLimit
	}

	s.mama = alpha*sample + (1.0-alpha)*s.mama

	return alpha
}

func (s *MesaAdaptiveMovingAverage) calculate(sample float64) float64 {
	const two = 2

	alpha := s.calculateMama(sample) / two
	s.fama = alpha*s.mama + (1.0-alpha)*s.fama

	return s.mama
}
