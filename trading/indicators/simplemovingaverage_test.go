//nolint:testpackage
package indicators

//nolint: gofumpt
import (
	"math"
	"testing"
	"time"

	"mbg/trading/data"
	"mbg/trading/indicators/indicator"
	"mbg/trading/indicators/indicator/output"
)

// Test data taken from:
// Perry Kaufman, Trading Systems an Methods, 3rd edition, page 72.

func testSimpleMovingAverageTime() time.Time {
	return time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})
}

func testSimpleMovingAverageInput() []float64 {
	return []float64{
		64.59, 64.23, 65.26, 65.24, 65.07, 65.14, 64.98, 64.76, 65.11, 65.46,
		65.94, 66.10, 66.87, 66.56, 66.71, 66.19, 66.14, 66.64, 67.33, 68.18,
		67.48, 67.19, 66.46, 67.20, 67.62, 67.66, 67.89, 69.19, 69.68, 69.31,
		69.11, 69.27, 68.97, 69.11, 69.50, 69.70, 69.94, 69.11, 67.64, 67.75,
		67.47, 67.50, 68.18, 67.35, 66.74, 67.00, 67.46, 67.36, 67.37, 67.78,
		67.96,
	}
}

func testSimpleMovingAverageExpected3() []float64 {
	return []float64{
		math.NaN(), math.NaN(), 64.69, 64.91, 65.19, 65.15, 65.06, 64.96, 64.95, 65.11,
		65.50, 65.83, 66.30, 66.51, 66.71, 66.49, 66.35, 66.32, 66.70, 67.38,
		67.66, 67.62, 67.04, 66.95, 67.09, 67.49, 67.72, 68.25, 68.92, 69.39,
		69.37, 69.23, 69.12, 69.12, 69.19, 69.44, 69.71, 69.58, 68.90, 68.17,
		67.62, 67.57, 67.72, 67.68, 67.42, 67.03, 67.07, 67.27, 67.40, 67.50,
		67.70,
	}
}

func testSimpleMovingAverageExpected5() []float64 {
	return []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), 64.88, 64.99, 65.14, 65.04, 65.01, 65.09,
		65.25, 65.47, 65.90, 66.19, 66.44, 66.49, 66.49, 66.45, 66.60, 66.90,
		67.15, 67.36, 67.33, 67.30, 67.19, 67.23, 67.37, 67.91, 68.41, 68.75,
		69.04, 69.31, 69.27, 69.15, 69.19, 69.31, 69.44, 69.47, 69.18, 68.83,
		68.38, 67.89, 67.71, 67.65, 67.45, 67.35, 67.35, 67.18, 67.19, 67.39,
		67.59,
	}
}

func testSimpleMovingAverageExpected10() []float64 {
	return []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), 64.98,
		65.12, 65.31, 65.47, 65.60, 65.76, 65.87, 65.98, 66.17, 66.39, 66.67,
		66.82, 66.93, 66.89, 66.95, 67.04, 67.19, 67.37, 67.62, 67.86, 67.97,
		68.13, 68.34, 68.59, 68.78, 68.97, 69.17, 69.38, 69.37, 69.17, 69.01,
		68.85, 68.67, 68.59, 68.41, 68.14, 67.87, 67.62, 67.45, 67.42, 67.42,
		67.47,
	}
}

func TestSimpleMovingAverageUpdate(t *testing.T) { //nolint: funlen
	t.Parallel()

	check := func(index int, exp, act float64) {
		t.Helper()

		if math.Abs(exp-act) > 1e-2 {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	checkNaN := func(index int, act float64) {
		t.Helper()

		if !math.IsNaN(act) {
			t.Errorf("[%v] is incorrect: expected NaN, actual %v", index, act)
		}
	}

	input := testSimpleMovingAverageInput()

	t.Run("length = 3", func(t *testing.T) {
		t.Parallel()
		sma := testSimpleMovingAverageCreate(3)
		expected := testSimpleMovingAverageExpected3()

		for i := 0; i < 2; i++ {
			checkNaN(i, sma.Update(input[i]))
		}

		for i := 2; i < len(input); i++ {
			exp := expected[i]
			act := sma.Update(input[i])
			check(i, exp, act)
		}

		checkNaN(0, sma.Update(math.NaN()))
	})

	t.Run("length = 5", func(t *testing.T) {
		t.Parallel()
		sma := testSimpleMovingAverageCreate(5)
		expected := testSimpleMovingAverageExpected5()

		for i := 0; i < 4; i++ {
			checkNaN(i, sma.Update(input[i]))
		}

		for i := 4; i < len(input); i++ {
			exp := expected[i]
			act := sma.Update(input[i])
			check(i, exp, act)
		}

		checkNaN(0, sma.Update(math.NaN()))
	})

	t.Run("length = 10", func(t *testing.T) {
		t.Parallel()
		sma := testSimpleMovingAverageCreate(10)
		expected := testSimpleMovingAverageExpected10()

		for i := 0; i < 8; i++ {
			checkNaN(i, sma.Update(input[i]))
		}

		for i := 8; i < len(input); i++ {
			exp := expected[i]
			act := sma.Update(input[i])
			check(i, exp, act)
		}

		checkNaN(0, sma.Update(math.NaN()))
	})
}

func TestSimpleMovingAverageUpdateEntity(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		l   = 2
		inp = 3.
		exp = inp / float64(l)
	)

	time := testSimpleMovingAverageTime()
	check := func(act indicator.Output) {
		t.Helper()

		if len(act) != 1 {
			t.Errorf("len(output) is incorrect: expected 1, actual %v", len(act))
		}

		s, ok := act[0].(data.Scalar)
		if !ok {
			t.Error("output is not scalar")
		}

		if s.Time != time {
			t.Errorf("time is incorrect: expected %v, actual %v", time, s.Time)
		}

		if s.Value != exp {
			t.Errorf("value is incorrect: expected %v, actual %v", exp, s.Value)
		}
	}

	t.Run("update scalar", func(t *testing.T) {
		t.Parallel()

		s := data.Scalar{Time: time, Value: inp}
		sma := testSimpleMovingAverageCreate(l)
		sma.Update(0.)
		check(sma.UpdateScalar(&s))
	})

	t.Run("update bar", func(t *testing.T) {
		t.Parallel()

		b := data.Bar{Time: time, Close: inp}
		sma := testSimpleMovingAverageCreate(l)
		sma.Update(0.)
		check(sma.UpdateBar(&b))
	})

	t.Run("update quote", func(t *testing.T) {
		t.Parallel()

		q := data.Quote{Time: time, Bid: inp}
		sma := testSimpleMovingAverageCreate(l)
		sma.Update(0.)
		check(sma.UpdateQuote(&q))
	})

	t.Run("update trade", func(t *testing.T) {
		t.Parallel()

		r := data.Trade{Time: time, Price: inp}
		sma := testSimpleMovingAverageCreate(l)
		sma.Update(0.)
		check(sma.UpdateTrade(&r))
	})
}

func TestSimpleMovingAverageIsPrimed(t *testing.T) { //nolint:funlen
	t.Parallel()

	input := testSimpleMovingAverageInput()
	check := func(index int, exp, act bool) {
		t.Helper()

		if exp != act {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	t.Run("length = 3", func(t *testing.T) {
		t.Parallel()
		sma := testSimpleMovingAverageCreate(3)

		check(-1, false, sma.IsPrimed())

		for i := 0; i < 2; i++ {
			sma.Update(input[i])
			check(i, false, sma.IsPrimed())
		}

		for i := 2; i < len(input); i++ {
			sma.Update(input[i])
			check(i, true, sma.IsPrimed())
		}
	})

	t.Run("length = 5", func(t *testing.T) {
		t.Parallel()
		sma := testSimpleMovingAverageCreate(5)

		check(-1, false, sma.IsPrimed())

		for i := 0; i < 4; i++ {
			sma.Update(input[i])
			check(i, false, sma.IsPrimed())
		}

		for i := 4; i < len(input); i++ {
			sma.Update(input[i])
			check(i, true, sma.IsPrimed())
		}
	})

	t.Run("length = 10", func(t *testing.T) {
		t.Parallel()
		sma := testSimpleMovingAverageCreate(10)

		check(-1, false, sma.IsPrimed())

		for i := 0; i < 9; i++ {
			sma.Update(input[i])
			check(i, false, sma.IsPrimed())
		}

		for i := 9; i < len(input); i++ {
			sma.Update(input[i])
			check(i, true, sma.IsPrimed())
		}
	})
}

func TestSimpleMovingAverageMetadata(t *testing.T) {
	t.Parallel()

	sma := testSimpleMovingAverageCreate(5)
	act := sma.Metadata()

	check := func(what string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", what, exp, act)
		}
	}

	check("Type", indicator.SimpleMovingAverage, act.Type)
	check("len(Outputs)", 1, len(act.Outputs))
	check("Outputs[0].Kind", int(SimpleMovingAverageValue), act.Outputs[0].Kind)
	check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
	check("Outputs[0].Name", "sma(5)", act.Outputs[0].Name)
	check("Outputs[0].Description", "Simple moving average sma(5)", act.Outputs[0].Description)
}

func TestNewSimpleMovingAverage(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		bc data.BarComponent   = data.BarMedianPrice
		qc data.QuoteComponent = data.QuoteMidPrice
		tc data.TradeComponent = data.TradePrice

		length = 5
		errlen = "invalid simple moving average parameters: length should be greater than 1"
		errbc  = "invalid simple moving average parameters: 9999: unknown bar component"
		errqc  = "invalid simple moving average parameters: 9999: unknown quote component"
		errtc  = "invalid simple moving average parameters: 9999: unknown trade component"
	)

	check := func(name string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", name, exp, act)
		}
	}

	t.Run("length > 1", func(t *testing.T) {
		t.Parallel()
		params := SimpleMovingAverageParams{
			Length: length, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		sma, err := NewSimpleMovingAverage(&params)
		check("err == nil", true, err == nil)
		check("name", "sma(5)", sma.name)
		check("description", "Simple moving average sma(5)", sma.description)
		check("primed", false, sma.primed)
		check("lastIndex", length-1, sma.lastIndex)
		check("len(window)", length, len(sma.window))
		check("windowLength", length, sma.windowLength)
		check("windowCount", 0, sma.windowCount)
		check("windowSum", 0., sma.windowSum)
		check("barFunc == nil", false, sma.barFunc == nil)
		check("quoteFunc == nil", false, sma.quoteFunc == nil)
		check("tradeFunc == nil", false, sma.tradeFunc == nil)
	})

	t.Run("length = 1", func(t *testing.T) {
		t.Parallel()
		params := SimpleMovingAverageParams{
			Length: 1, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		sma, err := NewSimpleMovingAverage(&params)
		check("sma == nil", true, sma == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length = 0", func(t *testing.T) {
		t.Parallel()
		params := SimpleMovingAverageParams{
			Length: 0, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		sma, err := NewSimpleMovingAverage(&params)
		check("sma == nil", true, sma == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length < 0", func(t *testing.T) {
		t.Parallel()
		params := SimpleMovingAverageParams{
			Length: -1, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		sma, err := NewSimpleMovingAverage(&params)
		check("sma == nil", true, sma == nil)
		check("err", errlen, err.Error())
	})

	t.Run("invalid bar component", func(t *testing.T) {
		t.Parallel()
		params := SimpleMovingAverageParams{
			Length: length, BarComponent: data.BarComponent(9999), QuoteComponent: qc, TradeComponent: tc,
		}

		sma, err := NewSimpleMovingAverage(&params)
		check("sma == nil", true, sma == nil)
		check("err", errbc, err.Error())
	})

	t.Run("invalid quote component", func(t *testing.T) {
		t.Parallel()
		params := SimpleMovingAverageParams{
			Length: length, BarComponent: bc, QuoteComponent: data.QuoteComponent(9999), TradeComponent: tc,
		}

		sma, err := NewSimpleMovingAverage(&params)
		check("sma == nil", true, sma == nil)
		check("err", errqc, err.Error())
	})

	t.Run("invalid trade component", func(t *testing.T) {
		t.Parallel()
		params := SimpleMovingAverageParams{
			Length: length, BarComponent: bc, QuoteComponent: qc, TradeComponent: data.TradeComponent(9999),
		}

		sma, err := NewSimpleMovingAverage(&params)
		check("sma == nil", true, sma == nil)
		check("err", errtc, err.Error())
	})
}

func testSimpleMovingAverageCreate(length int) *SimpleMovingAverage {
	params := SimpleMovingAverageParams{
		Length: length, BarComponent: data.BarClosePrice, QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	sma, _ := NewSimpleMovingAverage(&params)

	return sma
}
