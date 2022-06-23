//nolint:testpackage
package statistics

//nolint: gofumpt
import (
	"math"
	"testing"
	"time"

	"mbg/trading/data"
)

func testStandardDeviationTime() time.Time {
	return time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})
}

// testStandardDeviationVarianceInput is variance input test data.
func testStandardDeviationVarianceInput() []float64 {
	return []float64{1, 2, 8, 4, 9, 6, 7, 13, 9, 10, 3, 12}
}

// testStandardDeviationVarianceExpectedLength3Population is the Excel (VAR.P) output of population variance of length 3.
func testStandardDeviationVarianceExpectedLength3Population() []float64 {
	return []float64{
		math.NaN(), math.NaN(),
		9.55555555555556000, 6.22222222222222000, 4.66666666666667000, 4.22222222222222000, 1.55555555555556000,
		9.55555555555556000, 6.22222222222222000, 2.88888888888889000, 9.55555555555556000, 14.88888888888890000,
	}
}

// testStandardDeviationVarianceExpectedLength3Sample is the Excel (VAR.S) output of population variance of length 3.
func testStandardDeviationVarianceExpectedLength3Sample() []float64 {
	return []float64{
		math.NaN(), math.NaN(),
		14.3333333333333000, 9.3333333333333400, 7.0000000000000000, 6.3333333333333400, 2.3333333333333300,
		14.3333333333333000, 9.3333333333333400, 4.3333333333333400, 14.3333333333333000, 22.3333333333333000,
	}
}

//nolint: funlen
func TestNewStandardDeviation(t *testing.T) {
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
	s, err := NewStandardDeviation(&params, []StandardDeviationOutput{StandardDeviationValue})
	check(cond, "err == nil", true, err == nil)
	check(cond, "name", "stdev.s(5)", s.name)
	check(cond, "description", "Standard deviation based on unbiased estimation of the sample variance stdev.s(5)", s.description)
	check(cond, "unbiased", true, s.unbiased)
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

func testStandardDeviationCreate(length int, unbiased bool) *StandardDeviation {
	params := VarianceParams{
		Length: length, IsUnbiased: unbiased, BarComponent: data.BarClosePrice,
		QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	s, _ := NewStandardDeviation(&params, []StandardDeviationOutput{StandardDeviationValue})

	return s
}
