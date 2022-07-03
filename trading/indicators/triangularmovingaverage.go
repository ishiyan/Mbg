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

// TriangularMovingAverage computes the triangular moving average (TRIMA) like a weighted moving average.
// Instead of the WMA who put more weight on the latest sample, the TRIMA puts more weight on the data
// in the middle of the window.
//
// Using algebra, it can be demonstrated that the TRIMA is equivalent to doing a SMA of a SMA.
// The following explain the rules.
//    ➊ When the period π is even, TRIMA(x,π) = SMA(SMA(x,π/2), (π/2)+1).
//    ➋ When the period π is odd, TRIMA(x,π) = SMA(SMA(x,(π+1)/2), (π+1)/2).
//
// The SMA of a SMA is the algorithm generally found in books.
//
// TradeStation deviate from the generally accepted implementation by making the TRIMA to be as follows:
//     TRIMA(x,π) = SMA(SMA(x, (int)(π/2)+1), (int)(π/2)+1).
// This formula is done regardless if the period is even or odd. In other words:
//    ➊ A period of 4 becomes TRIMA(x,4) = SMA(SMA(x,3), 3).
//    ➋ A period of 5 becomes TRIMA(x,5) = SMA(SMA(x,3), 3).
//    ➌ A period of 6 becomes TRIMA(x,6) = SMA(SMA(x,4), 4).
//    ➍ A period of 7 becomes TRIMA(x,7) = SMA(SMA(x,4), 4).
//
// The Metastock implementation is the same as the generally accepted one.
//
// To optimize speed, this implementation uses a better algorithm than the usual SMA of a SMA.
// The calculation from one TRIMA value to the next is done by doing 4 little adjustments.
//
// The following show a TRIMA 4-period:
//    TRIMA at time δ: ((1*α)+(2*β)+(2*γ)+(1*δ)) / 6
//    TRIMA at time ε: ((1*β)+(2*γ)+(2*δ)+(1*ε)) / 6
// To go from TRIMA δ to ε, the following is done:
//    ➊ α and β are subtract from the numerator.
//    ➋ δ is added to the numerator.
//    ➌ ε is added to the numerator.
//    ➍ TRIMA is calculated by doing numerator / 6.
//    ➎ Sequence is repeated for the next output.
type TriangularMovingAverage struct {
	mu               sync.RWMutex
	name             string
	description      string
	factor           float64
	numerator        float64
	numeratorSub     float64
	numeratorAdd     float64
	window           []float64
	windowLength     int
	windowLengthHalf int
	windowCount      int
	isOdd            bool
	primed           bool
	barFunc          data.BarFunc
	quoteFunc        data.QuoteFunc
	tradeFunc        data.TradeFunc
}

// NewTriangularMovingAverage returns an instnce of the indicator created using supplied parameters.
func NewTriangularMovingAverage(p *TriangularMovingAverageParams) (*TriangularMovingAverage, error) { //nolint:funlen
	const (
		invalid = "invalid triangular moving average parameters"
		fmts    = "%s: %s"
		fmtw    = "%s: %w"
		fmtn    = "trima(%d)"
		minlen  = 2
	)

	length := p.Length
	if length < minlen {
		return nil, fmt.Errorf(fmts, invalid, "length should be greater than 1")
	}

	var (
		factor    float64
		err       error
		barFunc   data.BarFunc
		quoteFunc data.QuoteFunc
		tradeFunc data.TradeFunc
	)

	lengthHalf := length >> 1
	l := 1 + lengthHalf
	isOdd := length%2 == 1 //nolint:gomnd

	if isOdd {
		// Let period = 5 and l=(int)(period/2), then the formula for a "triangular" series is:
		// 1+2+3+2+1 = l*(l+1) + l+1 = (l+1)*(l+1) = 3*3 = 9.
		factor = 1. / float64(l*l) //nolint:gomnd
	} else {
		// Let period = 6 and l=(int)(period/2), then  the formula for a "triangular" series is:
		// 1+2+3+3+2+1 = l*(l+1) = 3*4 = 12.
		factor = 1. / float64(lengthHalf*l) //nolint:gomnd
		lengthHalf--
	}

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
	desc := "Triangular moving average " + name

	return &TriangularMovingAverage{
		mu:               sync.RWMutex{},
		name:             name,
		description:      desc,
		factor:           factor,
		numerator:        0,
		numeratorSub:     0,
		numeratorAdd:     0,
		window:           make([]float64, length),
		windowLength:     length,
		windowLengthHalf: lengthHalf,
		windowCount:      0,
		isOdd:            isOdd,
		primed:           false,
		barFunc:          barFunc,
		quoteFunc:        quoteFunc,
		tradeFunc:        tradeFunc,
	}, nil
}

// IsPrimed indicates whether an indicator is primed.
func (t *TriangularMovingAverage) IsPrimed() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.primed
}

// Metadata describes an output data of the indicator.
// It always has a single scalar output -- the calculated value of the simple moving average.
func (t *TriangularMovingAverage) Metadata() indicator.Metadata {
	return indicator.Metadata{
		Type: indicator.TriangularMovingAverage,
		Outputs: []output.Metadata{
			{
				Kind:        int(TriangularMovingAverageValue),
				Type:        output.Scalar,
				Name:        t.name,
				Description: t.description,
			},
		},
	}
}

// Update updatess the value of the moving average given the next sample.
//
// The indicator is not primed during the first ℓ-1 updates.
func (t *TriangularMovingAverage) Update(sample float64) float64 {
	if math.IsNaN(sample) {
		return sample
	}

	temp := sample

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.primed {
		t.numerator -= t.numeratorSub
		t.numeratorSub -= t.window[0]

		j := t.windowLength - 1
		for i := 0; i < j; i++ {
			t.window[i] = t.window[i+1]
		}

		t.window[j] = temp
		temp = t.window[t.windowLengthHalf]
		t.numeratorSub += temp

		if t.isOdd { // The logic for an odd length.
			t.numerator += t.numeratorAdd
			t.numeratorAdd -= temp
		} else { // The logic for an even length.
			t.numeratorAdd -= temp
			t.numerator += t.numeratorAdd
		}

		temp = sample
		t.numeratorAdd += temp
		t.numerator += temp
	} else {
		t.window[t.windowCount] = temp
		t.windowCount++

		if t.windowLength > t.windowCount {
			return math.NaN()
		}

		for i := t.windowLengthHalf; i >= 0; i-- {
			t.numeratorSub += t.window[i]
			t.numerator += t.numeratorSub
		}

		for i := t.windowLengthHalf + 1; i < t.windowLength; i++ {
			t.numeratorAdd += t.window[i]
			t.numerator += t.numeratorAdd
		}

		t.primed = true
	}

	return t.numerator * t.factor
}

// UpdateScalar updates the indicator given the next scalar sample.
func (t *TriangularMovingAverage) UpdateScalar(sample *data.Scalar) indicator.Output {
	output := make([]any, 1)
	output[0] = data.Scalar{Time: sample.Time, Value: t.Update(sample.Value)}

	return output
}

// UpdateBar updates the indicator given the next bar sample.
func (t *TriangularMovingAverage) UpdateBar(sample *data.Bar) indicator.Output {
	return t.UpdateScalar(&data.Scalar{Time: sample.Time, Value: t.barFunc(sample)})
}

// UpdateQuote updates the indicator given the next quote sample.
func (t *TriangularMovingAverage) UpdateQuote(sample *data.Quote) indicator.Output {
	return t.UpdateScalar(&data.Scalar{Time: sample.Time, Value: t.quoteFunc(sample)})
}

// UpdateTrade updates the indicator given the next trade sample.
func (t *TriangularMovingAverage) UpdateTrade(sample *data.Trade) indicator.Output {
	return t.UpdateScalar(&data.Scalar{Time: sample.Time, Value: t.tradeFunc(sample)})
}
