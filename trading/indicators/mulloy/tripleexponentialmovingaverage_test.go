package mulloy //nolint:testpackage, gci, gofmt, gofumpt, goimports

import (
	"math"
	"testing"
	"time"

	"mbg/trading/data"                        //nolint:depguard
	"mbg/trading/indicators/indicator"        //nolint:depguard
	"mbg/trading/indicators/indicator/output" //nolint:depguard
)

//nolint:lll
// Input data is taken from the TA-Lib (http://ta-lib.org/) tests,
//    test_data.c, TA_SREF_close_daily_ref_0_PRIV[252].
//
// Output data is taken from TA-Lib (http://ta-lib.org/) tests,
//    test_ma.c.
//
//   /*******************************/
//   /*  TEMA TEST - Metastock      */
//   /*******************************/
//   /* No output value. */
//   { 0, TA_ANY_MA_TEST, 0, 1, 1,  14, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 0, 0, 0},
//#ifndef TA_FUNC_NO_RANGE_CHECK
//   { 0, TA_ANY_MA_TEST, 0, 0, 251,  0, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_BAD_PARAM, 0, 0, 0, 0 },
//#endif
//
//   /* Test with period 14 */
//   { 1, TA_ANY_MA_TEST, 0, 0, 251, 14, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   0,  84.721, 39, 252-39 }, /* First Value */
//   { 0, TA_ANY_MA_TEST, 0, 0, 251, 14, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,   1,  84.089, 39, 252-39 },
//   { 0, TA_ANY_MA_TEST, 0, 0, 251, 14, TA_MAType_TEMA, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 252-40, 108.418, 39, 252-39 }, /* Last Value */

func testTripleExponentialMovingAverageTime() time.Time {
	return time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})
}

func testTripleExponentialMovingAverageInput() []float64 { //nolint:dupl
	return []float64{
		91.500000, 94.815000, 94.375000, 95.095000, 93.780000, 94.625000, 92.530000, 92.750000, 90.315000, 92.470000,
		96.125000, 97.250000, 98.500000, 89.875000, 91.000000, 92.815000, 89.155000, 89.345000, 91.625000, 89.875000,
		88.375000, 87.625000, 84.780000, 83.000000, 83.500000, 81.375000, 84.440000, 89.250000, 86.375000, 86.250000,
		85.250000, 87.125000, 85.815000, 88.970000, 88.470000, 86.875000, 86.815000, 84.875000, 84.190000, 83.875000,
		83.375000, 85.500000, 89.190000, 89.440000, 91.095000, 90.750000, 91.440000, 89.000000, 91.000000, 90.500000,
		89.030000, 88.815000, 84.280000, 83.500000, 82.690000, 84.750000, 85.655000, 86.190000, 88.940000, 89.280000,
		88.625000, 88.500000, 91.970000, 91.500000, 93.250000, 93.500000, 93.155000, 91.720000, 90.000000, 89.690000,
		88.875000, 85.190000, 83.375000, 84.875000, 85.940000, 97.250000, 99.875000, 104.940000, 106.000000, 102.500000,
		102.405000, 104.595000, 106.125000, 106.000000, 106.065000, 104.625000, 108.625000, 109.315000, 110.500000,
		112.750000, 123.000000, 119.625000, 118.750000, 119.250000, 117.940000, 116.440000, 115.190000, 111.875000,
		110.595000, 118.125000, 116.000000, 116.000000, 112.000000, 113.750000, 112.940000, 116.000000, 120.500000,
		116.620000, 117.000000, 115.250000, 114.310000, 115.500000, 115.870000, 120.690000, 120.190000, 120.750000,
		124.750000, 123.370000, 122.940000, 122.560000, 123.120000, 122.560000, 124.620000, 129.250000, 131.000000,
		132.250000, 131.000000, 132.810000, 134.000000, 137.380000, 137.810000, 137.880000, 137.250000, 136.310000,
		136.250000, 134.630000, 128.250000, 129.000000, 123.870000, 124.810000, 123.000000, 126.250000, 128.380000,
		125.370000, 125.690000, 122.250000, 119.370000, 118.500000, 123.190000, 123.500000, 122.190000, 119.310000,
		123.310000, 121.120000, 123.370000, 127.370000, 128.500000, 123.870000, 122.940000, 121.750000, 124.440000,
		122.000000, 122.370000, 122.940000, 124.000000, 123.190000, 124.560000, 127.250000, 125.870000, 128.860000,
		132.000000, 130.750000, 134.750000, 135.000000, 132.380000, 133.310000, 131.940000, 130.000000, 125.370000,
		130.130000, 127.120000, 125.190000, 122.000000, 125.000000, 123.000000, 123.500000, 120.060000, 121.000000,
		117.750000, 119.870000, 122.000000, 119.190000, 116.370000, 113.500000, 114.250000, 110.000000, 105.060000,
		107.000000, 107.870000, 107.000000, 107.120000, 107.000000, 91.000000, 93.940000, 93.870000, 95.500000, 93.000000,
		94.940000, 98.250000, 96.750000, 94.810000, 94.370000, 91.560000, 90.250000, 93.940000, 93.620000, 97.000000,
		95.000000, 95.870000, 94.060000, 94.620000, 93.750000, 98.000000, 103.940000, 107.870000, 106.060000, 104.500000,
		105.000000, 104.190000, 103.060000, 103.420000, 105.270000, 111.870000, 116.000000, 116.620000, 118.280000,
		113.370000, 109.000000, 109.700000, 109.250000, 107.000000, 109.190000, 110.000000, 109.200000, 110.120000,
		108.000000, 108.620000, 109.750000, 109.810000, 109.000000, 108.750000, 107.870000,
	}
}

func testTripleExponentialMovingAverageTascInput() []float64 { //nolint:dupl
	return []float64{
		451.61, 455.20, 453.29, 446.48, 446.17, 440.86, 441.88, 451.61, 438.43, 406.33,
		328.45, 323.30, 326.39, 322.97, 312.49, 316.47, 292.92, 302.57, 326.91, 333.19,
		330.47, 338.47, 340.14, 337.59, 344.66, 345.75, 353.27, 357.12, 363.40, 373.37,
		375.48, 381.58, 372.54, 374.64, 381.83, 373.90, 374.04, 379.23, 379.42, 372.48,
		366.03, 366.66, 376.86, 386.25, 386.92, 391.62, 394.69, 394.33, 394.59, 387.35,
		387.33, 387.71, 378.95, 377.42, 374.43, 376.51, 381.60, 383.91, 384.98, 387.71,
		385.67, 384.59, 388.59, 382.79, 381.02, 373.76, 367.58, 366.38, 373.91, 375.21,
		375.80, 377.34, 381.38, 384.74, 387.09, 391.66, 397.96, 406.35, 402.37, 407.19,
		399.96, 403.99, 405.90, 402.19, 400.94, 406.73, 410.71, 417.68, 423.76, 427.55,
		430.74, 434.83, 442.05, 445.21, 451.63, 453.65, 447.21, 448.36, 435.29, 442.42,
		448.90, 449.29, 452.82, 457.42, 462.48, 461.97, 466.75, 471.34, 471.31, 467.57,
		468.07, 472.92, 483.64, 467.29, 470.67, 452.76, 452.97, 456.19, 456.72, 456.63,
		457.10, 456.22, 443.84, 444.57, 454.82, 458.22, 439.72, 440.88, 421.33, 422.21,
		428.84, 429.01, 419.52, 431.02, 436.76, 442.16, 437.25, 435.54, 430.90, 436.31,
		425.79, 417.98, 428.61, 438.10, 448.31, 453.69, 462.13, 460.87, 467.55, 459.33,
		462.29, 460.53, 468.44, 455.27, 442.59, 417.46, 408.03, 393.49, 367.33, 381.21,
		380.38, 374.42, 362.25, 344.51, 347.36, 327.55, 337.36, 334.36, 336.45, 341.95,
		350.85, 349.04, 359.06, 371.54, 368.83, 373.60, 371.20, 367.24, 361.80, 376.99,
		394.28, 417.69, 436.80, 448.71, 448.95, 456.73, 475.11, 466.29, 464.15, 482.30,
		495.79, 501.62, 501.19, 494.64, 492.10, 493.42, 481.38, 492.67, 506.11, 498.54,
		495.07, 485.82, 475.92, 474.05, 492.71, 497.55, 492.69, 505.67, 508.31, 512.47,
		521.06, 525.68, 516.94, 516.71, 527.19, 524.48, 520.40, 519.05, 538.90, 525.13,
		540.93, 548.08, 531.29, 523.47, 523.90, 536.30, 540.90, 535.76, 565.71, 592.65,
		615.70, 626.85, 624.68, 620.21, 634.95, 636.43, 629.75, 633.47, 615.95, 618.62,
		624.28, 604.67, 590.01, 584.24, 591.81, 572.89, 578.14, 585.76, 574.43, 580.30,
		585.31, 585.43, 569.52, 554.20, 547.84, 563.35, 567.80, 570.52, 565.61, 580.83,
		573.74, 573.18, 563.70, 563.56, 573.44, 583.01, 589.12, 577.20, 571.63, 570.52,
		582.61, 597.30, 605.17, 616.82, 637.16, 642.60, 649.49, 661.60, 655.79, 661.29,
		665.88, 676.95, 677.21, 697.15, 701.64, 696.34, 700.98, 690.54, 663.61, 670.77,
		681.37, 692.78, 682.72, 681.54, 669.85, 665.26, 666.78, 658.41, 661.42, 681.44,
		676.37, 694.29, 700.53, 702.01, 693.19, 689.59, 694.81, 704.49, 705.81, 699.73,
		700.24, 704.70, 718.08, 718.26, 730.96, 734.07,
	}
}

func testTripleExponentialMovingAverageTascExpected() []float64 { //nolint:dupl
	return []float64{722.4268577073440}
}

func TestTripleExponentialMovingAverageUpdate(t *testing.T) { //nolint: funlen, cyclop
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

	input := testTripleExponentialMovingAverageInput()

	const (
		l       = 14
		lprimed = 3*l - 3
	)

	t.Run("length = 14, firstIsAverage = true", func(t *testing.T) {
		t.Parallel()

		const (
			i39value  = 84.8629 // Index=39 value.
			i40value  = 84.2246 // Index=40 value.
			i251value = 108.418 // Index=251 (last) value.
		)

		tema := testTripleExponentialMovingAverageCreateLength(l, true)

		for i := 0; i < lprimed; i++ {
			checkNaN(i, tema.Update(input[i]))
		}

		for i := lprimed; i < len(input); i++ {
			act := tema.Update(input[i])

			switch i {
			case 39:
				check(i, i39value, act)
			case 40:
				check(i, i40value, act)
			case 251:
				check(i, i251value, act)
			}
		}

		checkNaN(0, tema.Update(math.NaN()))
	})

	/*
		t.Run("length = 14, firstIsAverage = false (Metastock)", func(t *testing.T) {
			t.Parallel()

			const (
				i39value  = 84.721  // Index=39 value.
				i40value  = 84.089  // Index=40 value.
				i251value = 108.418 // Index=251 (last) value.
			)

			tema := testTripleExponentialMovingAverageCreateLength(l, false)
			t.Logf("length=%d, length2=%d, length3=%d", tema.length, tema.length2, tema.length3)

			for i := 0; i < lprimed; i++ {
				t.Logf("i=%d, primed=%v, count1=%d, count2=%d, count3=%d, sum1=%v, sum2=%v, sum3=%v, value1=%v, value2=%v, value3=%v", i, tema.primed, tema.count1, tema.count2, tema.count3, tema.sum1, tema.sum2, tema.sum3, tema.value1, tema.value2, tema.value3)
				checkNaN(i, tema.Update(input[i]))
			}

			for i := lprimed; i < len(input); i++ {
				t.Logf("i=%d, primed=%v, count1=%d, count2=%d, count3=%d, sum1=%v, sum2=%v, sum3=%v, value1=%v, value2=%v, value3=%v", i, tema.primed, tema.count1, tema.count2, tema.count3, tema.sum1, tema.sum2, tema.sum3, tema.value1, tema.value2, tema.value3)
				act := tema.Update(input[i])
				t.Logf("i=%d, act=%v", i, act)

				switch i {
				case 39:
					check(i, i39value, act)
				case 40:
					check(i, i40value, act)
				case 251:
					check(i, i251value, act)
				}
			}

			checkNaN(0, tema.Update(math.NaN()))
		})
	*/
	t.Run("length = 26, firstIsAverage = false (Metastock)", func(t *testing.T) {
		t.Parallel()

		const (
			lastValue = 722.4268577073440
		)

		tema := testTripleExponentialMovingAverageCreateLength(26, false)
		t.Logf("length=%d, length2=%d, length3=%d", tema.length, tema.length2, tema.length3)

		in := testTripleExponentialMovingAverageTascInput()
		inlen := len(in)
		t.Logf("inlen=%d", inlen)

		for i := 0; i < inlen; i++ {
			t.Logf("i=%d, act=%v primed=%v, count1=%d, count2=%d, count3=%d, sum1=%v, sum2=%v, sum3=%v, value1=%v, value2=%v, value3=%v", i, 3.*(tema.value1-tema.value2)+tema.value3, tema.primed, tema.count1, tema.count2, tema.count3, tema.sum1, tema.sum2, tema.sum3, tema.value1, tema.value2, tema.value3)
			act := tema.Update(in[i])
			t.Logf("i=%d, -> act=%v", i, act)

			switch i {
			case inlen - 1:
				check(i, lastValue, act)
			}
		}

		checkNaN(0, tema.Update(math.NaN()))
	})
}

func TestTripleExponentialMovingAverageUpdateEntity(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		l       = 2
		lprimed = 3*l - 3
		alpha   = 2. / float64(l+1)
		inp     = 3.
		exp     = 2.6666666666666665
	)

	time := testTripleExponentialMovingAverageTime()
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
		tema := testTripleExponentialMovingAverageCreateLength(l, false)

		for i := 0; i < lprimed; i++ {
			tema.Update(0.)
		}

		check(tema.UpdateScalar(&s))
	})

	t.Run("update bar", func(t *testing.T) {
		t.Parallel()

		b := data.Bar{Time: time, Close: inp}
		tema := testTripleExponentialMovingAverageCreateLength(l, true)

		for i := 0; i < lprimed; i++ {
			tema.Update(0.)
		}

		check(tema.UpdateBar(&b))
	})

	t.Run("update quote", func(t *testing.T) {
		t.Parallel()

		q := data.Quote{Time: time, Bid: inp}
		tema := testTripleExponentialMovingAverageCreateLength(l, true)

		for i := 0; i < lprimed; i++ {
			tema.Update(0.)
		}

		check(tema.UpdateQuote(&q))
	})

	t.Run("update trade", func(t *testing.T) {
		t.Parallel()

		r := data.Trade{Time: time, Price: inp}
		tema := testTripleExponentialMovingAverageCreateLength(l, true)

		for i := 0; i < lprimed; i++ {
			tema.Update(0.)
		}

		check(tema.UpdateTrade(&r))
	})
}

func TestTripleExponentialMovingAverageIsPrimed(t *testing.T) {
	t.Parallel()

	input := testTripleExponentialMovingAverageInput()
	check := func(index int, exp, act bool) {
		t.Helper()

		if exp != act {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	const (
		l       = 14
		lprimed = 3*l - 3
	)

	t.Run("length = 14, firstIsAverage = true", func(t *testing.T) {
		t.Parallel()

		tema := testTripleExponentialMovingAverageCreateLength(l, true)

		check(0, false, tema.IsPrimed())

		for i := 0; i < lprimed; i++ {
			tema.Update(input[i])
			check(i+1, false, tema.IsPrimed())
		}

		for i := lprimed; i < len(input); i++ {
			tema.Update(input[i])
			check(i+1, true, tema.IsPrimed())
		}
	})

	t.Run("length = 14, firstIsAverage = false (Metastock)", func(t *testing.T) {
		t.Parallel()

		tema := testTripleExponentialMovingAverageCreateLength(l, false)

		check(0, false, tema.IsPrimed())

		for i := 0; i < lprimed; i++ {
			tema.Update(input[i])
			check(i+1, false, tema.IsPrimed())
		}

		for i := lprimed; i < len(input); i++ {
			tema.Update(input[i])
			check(i+1, true, tema.IsPrimed())
		}
	})
}

func TestTripleExponentialMovingAverageMetadata(t *testing.T) { //nolint:dupl
	t.Parallel()

	check := func(what string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", what, exp, act)
		}
	}

	t.Run("length = 10, firstIsAverage = true", func(t *testing.T) {
		t.Parallel()

		tema := testTripleExponentialMovingAverageCreateLength(10, true)
		act := tema.Metadata()

		check("Type", indicator.TripleExponentialMovingAverage, act.Type)
		check("len(Outputs)", 1, len(act.Outputs))
		check("Outputs[0].Kind", int(TripleExponentialMovingAverageValue), act.Outputs[0].Kind)
		check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
		check("Outputs[0].Name", "tema(10)", act.Outputs[0].Name)
		check("Outputs[0].Description", "Triple exponential moving average tema(10)", act.Outputs[0].Description)
	})

	t.Run("alpha = 2/11 = 0.18181818..., firstIsAverage = false", func(t *testing.T) {
		t.Parallel()

		// α = 2 / (ℓ + 1) = 2/11 = 0.18181818...
		const alpha = 2. / 11.

		tema := testTripleExponentialMovingAverageCreateAlpha(alpha, false)
		act := tema.Metadata()

		check("Type", indicator.TripleExponentialMovingAverage, act.Type)
		check("len(Outputs)", 1, len(act.Outputs))
		check("Outputs[0].Kind", int(TripleExponentialMovingAverageValue), act.Outputs[0].Kind)
		check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
		check("Outputs[0].Name", "tema(10, 0.18181818)", act.Outputs[0].Name)
		check("Outputs[0].Description", "Triple exponential moving average tema(10, 0.18181818)", act.Outputs[0].Description)
	})
}

func TestNewTripleExponentialMovingAverage(t *testing.T) { //nolint: funlen, maintidx
	t.Parallel()

	const (
		bc     data.BarComponent   = data.BarMedianPrice
		qc     data.QuoteComponent = data.QuoteMidPrice
		tc     data.TradeComponent = data.TradePrice
		length                     = 10
		alpha                      = 2. / 11.

		errlen   = "invalid triple exponential moving average parameters: length should be greater than 1"
		erralpha = "invalid triple exponential moving average parameters: smoothing factor should be in range [0, 1]"
		errbc    = "invalid triple exponential moving average parameters: 9999: unknown bar component"
		errqc    = "invalid triple exponential moving average parameters: 9999: unknown quote component"
		errtc    = "invalid triple exponential moving average parameters: 9999: unknown trade component"
	)

	check := func(name string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", name, exp, act)
		}
	}

	t.Run("length > 1, firstIsAverage = false", func(t *testing.T) { //nolint:dupl
		t.Parallel()

		params := TripleExponentialMovingAverageLengthParams{
			Length: length, FirstIsAverage: false, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		check("name", "tema(10)", tema.name)
		check("description", "Triple exponential moving average tema(10)", tema.description)
		check("firstIsAverage", false, tema.firstIsAverage)
		check("primed", false, tema.primed)
		check("length", length, tema.length)
		check("length2", tema.length+tema.length, tema.length2)
		check("length3", tema.length2+tema.length-2, tema.length3)
		check("smoothingFactor", alpha, tema.smoothingFactor)
		check("count1", 0, tema.count1)
		check("count2", 0, tema.count2)
		check("count3", 0, tema.count3)
		check("sum1", 0., tema.sum1)
		check("sum2", 0., tema.sum2)
		check("sum3", 0., tema.sum3)
		check("value1", 0., tema.value1)
		check("value2", 0., tema.value2)
		check("value3", 0., tema.value3)
		check("barFunc == nil", false, tema.barFunc == nil)
		check("quoteFunc == nil", false, tema.quoteFunc == nil)
		check("tradeFunc == nil", false, tema.tradeFunc == nil)
	})

	t.Run("length = 1, firstIsAverage = true", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageLengthParams{
			Length: 1, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageLength(&params)
		check("tema == nil", true, tema == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length = 0", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageLengthParams{
			Length: 0, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageLength(&params)
		check("tema == nil", true, tema == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length < 0", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageLengthParams{
			Length: -1, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageLength(&params)
		check("tema == nil", true, tema == nil)
		check("err", errlen, err.Error())
	})

	t.Run("epsilon < α ≤ 1", func(t *testing.T) { //nolint:dupl
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: alpha, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		check("name", "tema(10, 0.18181818)", tema.name)
		check("description", "Triple exponential moving average tema(10, 0.18181818)", tema.description)
		check("firstIsAverage", true, tema.firstIsAverage)
		check("primed", false, tema.primed)
		check("length", length, tema.length)
		check("length2", tema.length+tema.length, tema.length2)
		check("length3", tema.length2+tema.length-2, tema.length3)
		check("smoothingFactor", alpha, tema.smoothingFactor)
		check("count1", 0, tema.count1)
		check("count2", 0, tema.count2)
		check("count3", 0, tema.count3)
		check("sum1", 0., tema.sum1)
		check("sum2", 0., tema.sum2)
		check("sum3", 0., tema.sum3)
		check("value1", 0., tema.value1)
		check("value2", 0., tema.value2)
		check("value3", 0., tema.value3)
		check("barFunc == nil", false, tema.barFunc == nil)
		check("quoteFunc == nil", false, tema.quoteFunc == nil)
		check("tradeFunc == nil", false, tema.tradeFunc == nil)
	})

	t.Run("0 < α < epsilon", func(t *testing.T) { //nolint:dupl
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 0.000000001, FirstIsAverage: false, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		check("name", "tema(199999999, 0.00000001)", tema.name)
		check("description", "Triple exponential moving average tema(199999999, 0.00000001)", tema.description)
		check("firstIsAverage", false, tema.firstIsAverage)
		check("primed", false, tema.primed)
		check("length", 199999999, tema.length) // 2./0.00000001 - 1.
		check("length2", tema.length+tema.length, tema.length2)
		check("length3", tema.length2+tema.length-2, tema.length3)
		check("smoothingFactor", 0.00000001, tema.smoothingFactor)
		check("count1", 0, tema.count1)
		check("count2", 0, tema.count2)
		check("count3", 0, tema.count3)
		check("sum1", 0., tema.sum1)
		check("sum2", 0., tema.sum2)
		check("sum3", 0., tema.sum3)
		check("value1", 0., tema.value1)
		check("value2", 0., tema.value2)
		check("value3", 0., tema.value3)
		check("barFunc == nil", false, tema.barFunc == nil)
		check("quoteFunc == nil", false, tema.quoteFunc == nil)
		check("tradeFunc == nil", false, tema.tradeFunc == nil)
	})

	t.Run("α = 0", func(t *testing.T) { //nolint:dupl
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 0, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		check("name", "tema(199999999, 0.00000001)", tema.name)
		check("description", "Triple exponential moving average tema(199999999, 0.00000001)", tema.description)
		check("firstIsAverage", true, tema.firstIsAverage)
		check("primed", false, tema.primed)
		check("length", 199999999, tema.length) // 2./0.00000001 - 1.
		check("length2", tema.length+tema.length, tema.length2)
		check("length3", tema.length2+tema.length-2, tema.length3)
		check("smoothingFactor", 0.00000001, tema.smoothingFactor)
		check("count1", 0, tema.count1)
		check("count2", 0, tema.count2)
		check("count3", 0, tema.count3)
		check("sum1", 0., tema.sum1)
		check("sum2", 0., tema.sum2)
		check("sum3", 0., tema.sum3)
		check("value1", 0., tema.value1)
		check("value2", 0., tema.value2)
		check("value3", 0., tema.value3)
		check("barFunc == nil", false, tema.barFunc == nil)
		check("quoteFunc == nil", false, tema.quoteFunc == nil)
		check("tradeFunc == nil", false, tema.tradeFunc == nil)
	})

	t.Run("α = 1", func(t *testing.T) { //nolint:dupl
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 1, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		check("name", "tema(1, 1.00000000)", tema.name)
		check("description", "Triple exponential moving average tema(1, 1.00000000)", tema.description)
		check("firstIsAverage", true, tema.firstIsAverage)
		check("primed", false, tema.primed)
		check("length", 1, tema.length) // 2./1 - 1.
		check("length2", tema.length+tema.length, tema.length2)
		check("length3", tema.length2+tema.length-2, tema.length3)
		check("smoothingFactor", 1., tema.smoothingFactor)
		check("count1", 0, tema.count1)
		check("count2", 0, tema.count2)
		check("count3", 0, tema.count3)
		check("sum1", 0., tema.sum1)
		check("sum2", 0., tema.sum2)
		check("sum3", 0., tema.sum3)
		check("value1", 0., tema.value1)
		check("value2", 0., tema.value2)
		check("value3", 0., tema.value3)
		check("barFunc == nil", false, tema.barFunc == nil)
		check("quoteFunc == nil", false, tema.quoteFunc == nil)
		check("tradeFunc == nil", false, tema.tradeFunc == nil)
	})

	t.Run("α < 0", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: -1, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("tema == nil", true, tema == nil)
		check("err", erralpha, err.Error())
	})

	t.Run("α > 1", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 2, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("tema == nil", true, tema == nil)
		check("err", erralpha, err.Error())
	})

	t.Run("invalid bar component", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 1, FirstIsAverage: true,
			BarComponent: data.BarComponent(9999), QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("tema == nil", true, tema == nil)
		check("err", errbc, err.Error())
	})

	t.Run("invalid quote component", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 1, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: data.QuoteComponent(9999), TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("tema == nil", true, tema == nil)
		check("err", errqc, err.Error())
	})

	t.Run("invalid trade component", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 1, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: data.TradeComponent(9999),
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("tema == nil", true, tema == nil)
		check("err", errtc, err.Error())
	})
}

func testTripleExponentialMovingAverageCreateLength(length int, firstIsAverage bool) *TripleExponentialMovingAverage {
	params := TripleExponentialMovingAverageLengthParams{
		Length: length, FirstIsAverage: firstIsAverage, BarComponent: data.BarClosePrice,
		QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	tema, _ := NewTripleExponentialMovingAverageLength(&params)

	return tema
}

func testTripleExponentialMovingAverageCreateAlpha(alpha float64, firstIsAverage bool) *TripleExponentialMovingAverage {
	params := TripleExponentialMovingAverageSmoothingFactorParams{
		SmoothingFactor: alpha, FirstIsAverage: firstIsAverage, BarComponent: data.BarClosePrice,
		QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	tema, _ := NewTripleExponentialMovingAverageSmoothingFactor(&params)

	return tema
}
