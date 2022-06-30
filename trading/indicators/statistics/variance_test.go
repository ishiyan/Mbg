//nolint:testpackage
package statistics

//nolint: gofumpt
import (
	"math"
	"testing"
	"time"

	"mbg/trading/data"
	"mbg/trading/indicators/indicator"
	"mbg/trading/indicators/indicator/output"
)

func testVarianceTime() time.Time {
	return time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})
}

// testVarianceInput is variance input test data.
func testVarianceInput() []float64 { return []float64{1, 2, 8, 4, 9, 6, 7, 13, 9, 10, 3, 12} }

// testVarianceExpectedLength3Population is the Excel (VAR.P) output of population variance of length 3.
func testVarianceExpectedLength3Population() []float64 {
	return []float64{
		math.NaN(), math.NaN(),
		9.55555555555556000, 6.22222222222222000, 4.66666666666667000, 4.22222222222222000, 1.55555555555556000,
		9.55555555555556000, 6.22222222222222000, 2.88888888888889000, 9.55555555555556000, 14.88888888888890000,
	}
}

// testVarianceExpectedLength5Population is the Excel (VAR.P) output of population variance of length 5.
func testVarianceExpectedLength5Population() []float64 {
	return []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		10.16000, 6.56000, 2.96000, 9.36000, 5.76000, 6.00000, 11.04000, 12.24000,
	}
}

// testVarianceExpectedLength3Sample is the Excel (VAR.S) output of population variance of length 3.
func testVarianceExpectedLength3Sample() []float64 {
	return []float64{
		math.NaN(), math.NaN(),
		14.3333333333333000, 9.3333333333333400, 7.0000000000000000, 6.3333333333333400, 2.3333333333333300,
		14.3333333333333000, 9.3333333333333400, 4.3333333333333400, 14.3333333333333000, 22.3333333333333000,
	}
}

// testVarianceExpectedLength5Sample is the Excel (VAR.S) output of population variance of length 5.
func testVarianceExpectedLength5Sample() []float64 {
	return []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		12.7000, 8.2000, 3.7000, 11.7000, 7.2000, 7.5000, 13.8000, 15.3000,
	}
}

func TestVarianceUpdate(t *testing.T) { //nolint: funlen
	t.Parallel()

	check := func(index int, exp, act float64) {
		t.Helper()

		if math.Abs(exp-act) > 1e-13 {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	checkNaN := func(index int, act float64) {
		t.Helper()

		if !math.IsNaN(act) {
			t.Errorf("[%v] is incorrect: expected NaN, actual %v", index, act)
		}
	}

	input := testVarianceInput()

	t.Run("population variance length of 3", func(t *testing.T) {
		t.Parallel()
		v := testVarianceCreate(3, false)
		expected := testVarianceExpectedLength3Population()

		for i := 0; i < 2; i++ {
			checkNaN(i, v.Update(input[i]))
		}

		for i := 2; i < len(input); i++ {
			exp := expected[i]
			act := v.Update(input[i])
			check(i, exp, act)
		}

		checkNaN(0, v.Update(math.NaN()))
	})

	t.Run("population variance length of 5", func(t *testing.T) {
		t.Parallel()
		v := testVarianceCreate(5, false)
		expected := testVarianceExpectedLength5Population()

		for i := 0; i < 4; i++ {
			checkNaN(i, v.Update(input[i]))
		}

		for i := 4; i < len(input); i++ {
			exp := expected[i]
			act := v.Update(input[i])
			check(i, exp, act)
		}
	})

	t.Run("sample variance length of 3", func(t *testing.T) {
		t.Parallel()
		v := testVarianceCreate(3, true)
		expected := testVarianceExpectedLength3Sample()

		for i := 0; i < 2; i++ {
			checkNaN(i, v.Update(input[i]))
		}

		for i := 2; i < len(input); i++ {
			exp := expected[i]
			act := v.Update(input[i])
			check(i, exp, act)
		}

		checkNaN(0, v.Update(math.NaN()))
	})
}

func TestVarianceUpdateEntity(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		l   = 3
		inp = 3.
		exp = inp * inp / float64(l)
	)

	time := testVarianceTime()
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
		v := testVarianceCreate(l, true)
		v.Update(0.)
		v.Update(0.)
		check(v.UpdateScalar(&s))
	})

	t.Run("update bar", func(t *testing.T) {
		t.Parallel()

		b := data.Bar{Time: time, Close: inp}
		v := testVarianceCreate(l, true)
		v.Update(0.)
		v.Update(0.)
		check(v.UpdateBar(&b))
	})

	t.Run("update quote", func(t *testing.T) {
		t.Parallel()

		q := data.Quote{Time: time, Bid: inp}
		v := testVarianceCreate(l, true)
		v.Update(0.)
		v.Update(0.)
		check(v.UpdateQuote(&q))
	})

	t.Run("update trade", func(t *testing.T) {
		t.Parallel()

		r := data.Trade{Time: time, Price: inp}
		v := testVarianceCreate(l, true)
		v.Update(0.)
		v.Update(0.)
		check(v.UpdateTrade(&r))
	})
}

func TestVarianceIsPrimed(t *testing.T) {
	t.Parallel()

	check := func(index int, exp, act bool) {
		t.Helper()

		if exp != act {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	input := testVarianceInput()
	v := testVarianceCreate(3, false)

	check(0, false, v.IsPrimed())

	for i := 0; i < 2; i++ {
		v.Update(input[i])
		check(i+1, false, v.IsPrimed())
	}

	for i := 2; i < len(input); i++ {
		v.Update(input[i])
		check(i+1, true, v.IsPrimed())
	}
}

func TestVarianceMetadata(t *testing.T) {
	t.Parallel()

	check := func(what string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", what, exp, act)
		}
	}

	t.Run("population variance", func(t *testing.T) {
		t.Parallel()
		v := testVarianceCreate(7, false)
		act := v.Metadata()

		check("Type", indicator.Variance, act.Type)
		check("len(Outputs)", 1, len(act.Outputs))
		check("Outputs[0].Kind", int(VarianceValue), act.Outputs[0].Kind)
		check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
		check("Outputs[0].Name", "var.p(7)", act.Outputs[0].Name)
		check("Outputs[0].Description", "Estimation of the population variance var.p(7)", act.Outputs[0].Description)
	})

	t.Run("sample variance", func(t *testing.T) {
		t.Parallel()
		v := testVarianceCreate(7, true)
		act := v.Metadata()

		check("Type", indicator.Variance, act.Type)
		check("len(Outputs)", 1, len(act.Outputs))
		check("Outputs[0].Kind", int(VarianceValue), act.Outputs[0].Kind)
		check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
		check("Outputs[0].Name", "var.s(7)", act.Outputs[0].Name)
		check("Outputs[0].Description", "Unbiased estimation of the sample variance var.s(7)", act.Outputs[0].Description)
	})
}

func TestNewVariance(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		bc     data.BarComponent   = data.BarMedianPrice
		qc     data.QuoteComponent = data.QuoteMidPrice
		tc     data.TradeComponent = data.TradePrice
		length                     = 5
		errlen                     = "invalid variance parameters: length should be greater than 1"
		errbc                      = "invalid variance parameters: 9999: unknown bar component"
		errqc                      = "invalid variance parameters: 9999: unknown quote component"
		errtc                      = "invalid variance parameters: 9999: unknown trade component"
	)

	check := func(condition, name string, exp, act any) {
		if exp != act {
			t.Errorf("(%s): %s is incorrect: expected %v, actual %v", condition, name, exp, act)
		}
	}

	params := VarianceParams{Length: length, IsUnbiased: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc}

	cond := "length > 1, unbiased"
	v, err := NewVariance(&params)
	check(cond, "err == nil", true, err == nil)
	check(cond, "name", "var.s(5)", v.name)
	check(cond, "description", "Unbiased estimation of the sample variance var.s(5)", v.description)
	check(cond, "unbiased", true, v.unbiased)
	check(cond, "primed", false, v.primed)
	check(cond, "windowLength", length, v.windowLength)
	check(cond, "lastIndex", length-1, v.lastIndex)
	check(cond, "len(window)", length, len(v.window))
	check(cond, "windowSum", 0., v.windowSum)
	check(cond, "windowSquaredSum", 0., v.windowSquaredSum)
	check(cond, "windowCount", 0, v.windowCount)
	check(cond, "barFunc == nil", false, v.barFunc == nil)
	check(cond, "quoteFunc == nil", false, v.quoteFunc == nil)
	check(cond, "tradeFunc == nil", false, v.tradeFunc == nil)

	cond = "length > 1, biased"
	params.IsUnbiased = false
	v, err = NewVariance(&params)
	check(cond, "err == nil", true, err == nil)
	check(cond, "name", "var.p(5)", v.name)
	check(cond, "description", "Estimation of the population variance var.p(5)", v.description)
	check(cond, "unbiased", false, v.unbiased)
	check(cond, "primed", false, v.primed)
	check(cond, "windowLength", length, v.windowLength)
	check(cond, "lastIndex", length-1, v.lastIndex)
	check(cond, "len(window)", length, len(v.window))
	check(cond, "windowSum", 0., v.windowSum)
	check(cond, "windowSquaredSum", 0., v.windowSquaredSum)
	check(cond, "windowCount", 0, v.windowCount)
	check(cond, "barFunc == nil", false, v.barFunc == nil)
	check(cond, "quoteFunc == nil", false, v.quoteFunc == nil)
	check(cond, "tradeFunc == nil", false, v.tradeFunc == nil)

	cond = "length = 1"
	params.Length = 1
	v, err = NewVariance(&params)
	check(cond, "v == nil", true, v == nil)
	check(cond, "err", errlen, err.Error())

	cond = "length = 0"
	params.Length = 0
	v, err = NewVariance(&params)
	check(cond, "v == nil", true, v == nil)
	check(cond, "err", errlen, err.Error())

	cond = "length = -1"
	params.Length = -1
	v, err = NewVariance(&params)
	check(cond, "v == nil", true, v == nil)
	check(cond, "err", errlen, err.Error())

	cond = "invalid bar component"
	params.BarComponent = data.BarComponent(9999)
	params.Length = length
	v, err = NewVariance(&params)
	check(cond, "v == nil", true, v == nil)
	check(cond, "err", errbc, err.Error())

	cond = "invalid quote component"
	params.BarComponent = bc
	params.QuoteComponent = data.QuoteComponent(9999)
	v, err = NewVariance(&params)
	check(cond, "v == nil", true, v == nil)
	check(cond, "err", errqc, err.Error())

	cond = "invalid trade component"
	params.QuoteComponent = qc
	params.TradeComponent = data.TradeComponent(9999)
	v, err = NewVariance(&params)
	check(cond, "v == nil", true, v == nil)
	check(cond, "err", errtc, err.Error())
}

func testVarianceCreate(length int, unbiased bool) *Variance {
	params := VarianceParams{
		Length: length, IsUnbiased: unbiased, BarComponent: data.BarClosePrice,
		QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	v, _ := NewVariance(&params)

	return v
}
