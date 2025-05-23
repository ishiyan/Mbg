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

//nolint:lll
// Input data is taken from the TA-Lib (http://ta-lib.org/) tests,
//    test_data.c, TA_SREF_close_daily_ref_0_PRIV[252].
//
// Output data is taken from TA-Lib (http://ta-lib.org/) tests,
//    test_ma.c.
//
// /***************/
// /*  TRIMA TEST */
// /***************/
// { 1, TA_ANY_MA_TEST, 0, 0, 251, 10, TA_MAType_TRIMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      0,  93.6043, 9,  252-9  }, /* First Value */
// { 0, TA_ANY_MA_TEST, 0, 0, 251, 10, TA_MAType_TRIMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      1,  93.4252, 9,  252-9  },
// { 0, TA_ANY_MA_TEST, 0, 0, 251, 10, TA_MAType_TRIMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 252-11, 109.1850, 9,  252-9  },
// { 0, TA_ANY_MA_TEST, 0, 0, 251, 10, TA_MAType_TRIMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 252-10, 109.1407, 9,  252-9  }, /* Last Value */
//
// { 1, TA_ANY_MA_TEST, 0, 0, 251,  9, TA_MAType_TRIMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,     0,   93.8176,  8,  252-8  }, /* First Value */
// { 0, TA_ANY_MA_TEST, 0, 0, 251,  9, TA_MAType_TRIMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 252-9,  109.1312,  8,  252-8  }, /* Last Value */
//
// { 1, TA_ANY_MA_TEST, 0, 0, 251, 12, TA_MAType_TRIMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      0,  93.5329, 11,  252-11  }, /* First Value */
// { 0, TA_ANY_MA_TEST, 0, 0, 251, 12, TA_MAType_TRIMA, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 252-12, 109.1157, 11,  252-11  }, /* Last Value */

func testTriangularMovingAverageTime() time.Time {
	return time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})
}

//nolint:dupl
func testTriangularMovingAverageInput() []float64 {
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

// testTriangularMovingAverageExpectedXls is taken from TA-Lib (http://ta-lib.org/) tests,
// test_TRIMA.xsl, TRIMA(12), Q4…Q255, 252 entries.
func testTriangularMovingAverageExpectedXls12() []float64 {
	return []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(),
		93.5329761904762, 93.6096428571428, 93.5933333333333, 93.6425000000000, 93.7965476190476, 93.8471428571429,
		93.6536904761905, 93.2340476190476, 92.6722619047619, 92.1164285714286, 91.4207142857143, 90.6126190476190,
		89.8194047619047, 89.0209523809524, 88.1838095238095, 87.2529761904762, 86.4233333333333, 85.7552380952381,
		85.2686904761905, 84.9748809523809, 85.0114285714286, 85.2830952380952, 85.6417857142857, 86.0116666666667,
		86.3584523809524, 86.6651190476190, 86.8765476190476, 86.9123809523810, 86.7941666666667, 86.5613095238095,
		86.2458333333333, 85.9720238095238, 85.7696428571429, 85.7852380952381, 86.0032142857143, 86.5345238095238,
		87.2704761904762, 88.0822619047619, 88.8627380952381, 89.4853571428571, 89.8975000000000, 89.9754761904762,
		89.7304761904762, 89.2042857142857, 88.4980952380952, 87.6863095238095, 86.8611904761905, 86.1930952380952,
		85.8330952380952, 85.7453571428571, 85.9447619047619, 86.4314285714286, 87.1248809523809, 87.9834523809524,
		88.8315476190476, 89.6498809523810, 90.4035714285714, 91.0210714285714, 91.4451190476190, 91.6385714285714,
		91.5315476190476, 91.0911904761905, 90.3800000000000, 89.4954761904762, 88.8378571428571, 88.4852380952381,
		88.7070238095238, 89.6653571428571, 91.2761904761905, 93.4420238095238, 95.8794047619048, 98.2855952380953,
		100.4551190476190, 102.1559523809520, 103.3686904761900, 104.3098809523810, 104.9714285714290, 105.5622619047620,
		106.1650000000000, 107.1457142857140, 108.4820238095240, 110.0088095238100, 111.6240476190480, 113.3040476190480,
		114.9677380952380, 116.2847619047620, 117.0140476190480, 117.1920238095240, 117.1021428571430, 116.7295238095240,
		116.1692857142860, 115.4452380952380, 114.9517857142860, 114.6986904761900, 114.5891666666670, 114.6135714285710,
		114.6989285714290, 114.9138095238100, 115.2403571428570, 115.5548809523810, 115.8016666666670, 115.9888095238100,
		116.1657142857140, 116.4038095238090, 116.6538095238100, 117.1166666666670, 117.7342857142860, 118.5321428571430,
		119.4847619047620, 120.4102380952380, 121.3028571428570, 122.0614285714290, 122.7114285714290, 123.3659523809520,
		124.0828571428570, 124.9428571428570, 125.9771428571430, 127.1916666666670, 128.6028571428570, 130.0361904761900,
		131.4116666666670, 132.7052380952380, 133.8945238095240, 134.8933333333330, 135.6033333333330, 135.8921428571430,
		135.8073809523810, 135.2700000000000, 134.3100000000000, 132.9511904761900, 131.3392857142860, 129.7959523809520,
		128.3938095238100, 127.2464285714290, 126.3566666666670, 125.6542857142860, 125.0828571428570, 124.5873809523810,
		124.0442857142860, 123.5042857142860, 122.8509523809520, 122.3523809523810, 122.0026190476190, 121.8416666666670,
		121.8964285714290, 122.1459523809520, 122.5873809523810, 123.0900000000000, 123.5138095238100, 123.9007142857140,
		124.1554761904760, 124.1721428571430, 124.0164285714290, 123.7773809523810, 123.5814285714290, 123.3733333333330,
		123.2647619047620, 123.3673809523810, 123.7569047619050, 124.3590476190480, 125.1159523809520, 126.0811904761900,
		127.2280952380950, 128.4050000000000, 129.6045238095240, 130.6616666666670, 131.5104761904760, 131.9559523809520,
		132.0428571428570, 131.8200000000000, 131.2488095238090, 130.3350000000000, 129.3035714285710, 128.2335714285710,
		127.2290476190480, 126.1723809523810, 125.1411904761900, 124.2021428571430, 123.3776190476190, 122.6483333333330,
		121.8728571428570, 121.1673809523810, 120.4514285714290, 119.7519047619050, 118.9185714285710, 117.8040476190480,
		116.4230952380950, 114.9423809523810, 113.3947619047620, 111.8559523809520, 110.3290476190480, 108.7023809523810,
		107.1680952380950, 105.5907142857140, 103.9419047619050, 102.1116666666670, 100.1640476190480, 98.4604761904762,
		97.1585714285714, 96.1900000000000, 95.5278571428571, 95.1052380952381, 94.9071428571429, 94.8935714285714,
		94.6328571428571, 94.3573809523810, 94.0745238095238, 93.9211904761905, 93.8928571428571, 93.9923809523810,
		94.1976190476191, 94.5011904761905, 94.9654761904762, 95.7004761904762, 96.6185714285715, 97.6811904761905,
		98.9954761904762, 100.4540476190480, 101.8678571428570, 102.9628571428570, 103.7533333333330, 104.4335714285710,
		105.1404761904760, 105.8754761904760, 106.8254761904760, 108.0333333333330, 109.4359523809520, 110.8057142857140,
		111.8392857142860, 112.3819047619050, 112.4121428571430, 111.9997619047620, 111.3552380952380, 110.6319047619050,
		109.9304761904760, 109.4283333333330, 109.1685714285710, 109.1207142857140, 109.1483333333330, 109.1385714285710,
		109.1157142857140,
	}
}

func TestTriangularMovingAverageUpdate(t *testing.T) { //nolint: funlen
	t.Parallel()

	check := func(index int, exp, act float64) {
		t.Helper()

		if math.Abs(exp-act) > 1e-4 {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	checkNaN := func(index int, act float64) {
		t.Helper()

		if !math.IsNaN(act) {
			t.Errorf("[%v] is incorrect: expected NaN, actual %v", index, act)
		}
	}

	input := testTriangularMovingAverageInput()

	t.Run("length = 9", func(t *testing.T) {
		const ( // Values from index=0 to index=7 are NaN.
			i8value   = 93.8176  // Index=8 value.
			i251value = 109.1312 // Index=251 (last) value.
		)

		t.Parallel()
		trima := testTriangularMovingAverageCreate(9)

		for i := 0; i < 8; i++ {
			checkNaN(i, trima.Update(input[i]))
		}

		for i := 8; i < len(input); i++ {
			act := trima.Update(input[i])

			switch i {
			case 8:
				check(i, i8value, act)
			case 251:
				check(i, i251value, act)
			}
		}

		checkNaN(0, trima.Update(math.NaN()))
	})

	t.Run("length = 10", func(t *testing.T) {
		const ( // Values from index=0 to index=8 are NaN.
			i9value   = 93.6043  // Index=9 value.
			i10value  = 93.4252  // Index=10 value.
			i250value = 109.1850 // Index=250 value.
			i251value = 109.1407 // Index=251 (last) value.
		)

		t.Parallel()
		trima := testTriangularMovingAverageCreate(10)

		for i := 0; i < 9; i++ {
			checkNaN(i, trima.Update(input[i]))
		}

		for i := 9; i < len(input); i++ {
			act := trima.Update(input[i])

			switch i {
			case 9:
				check(i, i9value, act)
			case 10:
				check(i, i10value, act)
			case 250:
				check(i, i250value, act)
			case 251:
				check(i, i251value, act)
			}
		}

		checkNaN(0, trima.Update(math.NaN()))
	})

	t.Run("length = 12", func(t *testing.T) {
		const ( // Values from index=0 to index=10 are NaN.
			i11value  = 93.5329  // Index=11 value.
			i251value = 109.1157 // Index=251 (last) value.
		)

		t.Parallel()
		trima := testTriangularMovingAverageCreate(12)

		for i := 0; i < 10; i++ {
			checkNaN(i, trima.Update(input[i]))
		}

		for i := 10; i < len(input); i++ {
			act := trima.Update(input[i])

			switch i {
			case 11:
				check(i, i11value, act)
			case 251:
				check(i, i251value, act)
			}
		}

		checkNaN(0, trima.Update(math.NaN()))
	})
}

func TestTriangularMovingAverageUpdateXls(t *testing.T) {
	t.Parallel()

	check := func(index int, exp, act float64) {
		t.Helper()

		if math.Abs(exp-act) > 1e-12 {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	checkNaN := func(index int, act float64) {
		t.Helper()

		if !math.IsNaN(act) {
			t.Errorf("[%v] is incorrect: expected NaN, actual %v", index, act)
		}
	}

	input := testTriangularMovingAverageInput()
	expected := testTriangularMovingAverageExpectedXls12()
	trima := testTriangularMovingAverageCreate(12)

	for i := 0; i < 11; i++ {
		checkNaN(i, trima.Update(input[i]))
	}

	for i := 11; i < len(input); i++ {
		act := trima.Update(input[i])
		check(i, expected[i], act)
	}
}

func TestTriangularMovingAverageUpdateEntity(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		l   = 12
		l1  = l - 1
		exp = 93.5329761904762
	)

	input := testTriangularMovingAverageInput()
	time := testTriangularMovingAverageTime()
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

		if math.Abs(exp-s.Value) > 1e-12 {
			t.Errorf("value is incorrect: expected %v, actual %v", exp, s.Value)
		}
	}

	t.Run("update scalar", func(t *testing.T) {
		t.Parallel()

		s := data.Scalar{Time: time, Value: input[l1]}
		trima := testTriangularMovingAverageCreate(l)

		for i := 0; i < l1; i++ {
			trima.Update(input[i])
		}

		check(trima.UpdateScalar(&s))
	})

	t.Run("update bar", func(t *testing.T) {
		t.Parallel()

		b := data.Bar{Time: time, Close: input[l1]}
		trima := testTriangularMovingAverageCreate(l)

		for i := 0; i < l1; i++ {
			trima.Update(input[i])
		}

		check(trima.UpdateBar(&b))
	})

	t.Run("update quote", func(t *testing.T) {
		t.Parallel()

		q := data.Quote{Time: time, Bid: input[l1]}
		trima := testTriangularMovingAverageCreate(l)

		for i := 0; i < l1; i++ {
			trima.Update(input[i])
		}

		check(trima.UpdateQuote(&q))
	})

	t.Run("update trade", func(t *testing.T) {
		t.Parallel()

		r := data.Trade{Time: time, Price: input[l1]}
		trima := testTriangularMovingAverageCreate(l)

		for i := 0; i < l1; i++ {
			trima.Update(input[i])
		}

		check(trima.UpdateTrade(&r))
	})
}

func TestTriangularMovingAverageIsPrimed(t *testing.T) { //nolint:dupl
	t.Parallel()

	input := testTriangularMovingAverageInput()
	check := func(index int, exp, act bool) {
		t.Helper()

		if exp != act {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	t.Run("length = 9", func(t *testing.T) {
		t.Parallel()
		trima := testTriangularMovingAverageCreate(9)

		check(-1, false, trima.IsPrimed())

		for i := 0; i < 8; i++ {
			trima.Update(input[i])
			check(i, false, trima.IsPrimed())
		}

		for i := 8; i < len(input); i++ {
			trima.Update(input[i])
			check(i, true, trima.IsPrimed())
		}
	})

	t.Run("length = 12", func(t *testing.T) {
		t.Parallel()
		trima := testTriangularMovingAverageCreate(12)

		check(-1, false, trima.IsPrimed())

		for i := 0; i < 11; i++ {
			trima.Update(input[i])
			check(i, false, trima.IsPrimed())
		}

		for i := 11; i < len(input); i++ {
			trima.Update(input[i])
			check(i, true, trima.IsPrimed())
		}
	})
}

func TestTriangularMovingAverageMetadata(t *testing.T) {
	t.Parallel()

	trima := testTriangularMovingAverageCreate(5)
	act := trima.Metadata()

	check := func(what string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", what, exp, act)
		}
	}

	check("Type", indicator.TriangularMovingAverage, act.Type)
	check("len(Outputs)", 1, len(act.Outputs))
	check("Outputs[0].Kind", int(TriangularMovingAverageValue), act.Outputs[0].Kind)
	check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
	check("Outputs[0].Name", "trima(5)", act.Outputs[0].Name)
	check("Outputs[0].Description", "Triangular moving average trima(5)", act.Outputs[0].Description)
}

func TestNewTriangularMovingAverage(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		bc data.BarComponent   = data.BarMedianPrice
		qc data.QuoteComponent = data.QuoteMidPrice
		tc data.TradeComponent = data.TradePrice

		errlen = "invalid triangular moving average parameters: length should be greater than 1"
		errbc  = "invalid triangular moving average parameters: 9999: unknown bar component"
		errqc  = "invalid triangular moving average parameters: 9999: unknown quote component"
		errtc  = "invalid triangular moving average parameters: 9999: unknown trade component"
	)

	check := func(name string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", name, exp, act)
		}
	}

	t.Run("length > 1, even", func(t *testing.T) {
		t.Parallel()

		const (
			length     = 6
			lengthHalf = 3
			factor     = 1. / 12.
		)

		params := TriangularMovingAverageParams{
			Length: length, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		trima, err := NewTriangularMovingAverage(&params)
		check("err == nil", true, err == nil)
		check("name", "trima(6)", trima.name)
		check("description", "Triangular moving average trima(6)", trima.description)
		check("numerator", 0., trima.numerator)
		check("numeratorSub", 0., trima.numeratorSub)
		check("numeratorAdd", 0., trima.numeratorAdd)
		check("len(window)", length, len(trima.window))
		check("windowLength", length, trima.windowLength)
		check("windowLengthHalf", lengthHalf-1, trima.windowLengthHalf)
		check("windowCount", 0, trima.windowCount)
		check("factor", factor, trima.factor)
		check("isOdd", false, trima.isOdd)
		check("primed", false, trima.primed)
		check("barFunc == nil", false, trima.barFunc == nil)
		check("quoteFunc == nil", false, trima.quoteFunc == nil)
		check("tradeFunc == nil", false, trima.tradeFunc == nil)
	})

	t.Run("length > 1, odd", func(t *testing.T) {
		t.Parallel()

		const (
			length     = 5
			lengthHalf = 2
			factor     = 1. / 9.
		)

		params := TriangularMovingAverageParams{
			Length: length, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		trima, err := NewTriangularMovingAverage(&params)
		check("err == nil", true, err == nil)
		check("name", "trima(5)", trima.name)
		check("description", "Triangular moving average trima(5)", trima.description)
		check("numerator", 0., trima.numerator)
		check("numeratorSub", 0., trima.numeratorSub)
		check("numeratorAdd", 0., trima.numeratorAdd)
		check("len(window)", length, len(trima.window))
		check("windowLength", length, trima.windowLength)
		check("windowLengthHalf", lengthHalf, trima.windowLengthHalf)
		check("windowCount", 0, trima.windowCount)
		check("factor", factor, trima.factor)
		check("isOdd", true, trima.isOdd)
		check("primed", false, trima.primed)
		check("barFunc == nil", false, trima.barFunc == nil)
		check("quoteFunc == nil", false, trima.quoteFunc == nil)
		check("tradeFunc == nil", false, trima.tradeFunc == nil)
	})

	t.Run("length = 1", func(t *testing.T) {
		t.Parallel()
		params := TriangularMovingAverageParams{
			Length: 1, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		trima, err := NewTriangularMovingAverage(&params)
		check("trima == nil", true, trima == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length = 0", func(t *testing.T) {
		t.Parallel()
		params := TriangularMovingAverageParams{
			Length: 0, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		trima, err := NewTriangularMovingAverage(&params)
		check("trima == nil", true, trima == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length < 0", func(t *testing.T) {
		t.Parallel()
		params := TriangularMovingAverageParams{
			Length: -1, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		trima, err := NewTriangularMovingAverage(&params)
		check("trima == nil", true, trima == nil)
		check("err", errlen, err.Error())
	})

	t.Run("invalid bar component", func(t *testing.T) {
		t.Parallel()
		params := TriangularMovingAverageParams{
			Length: 5, BarComponent: data.BarComponent(9999), QuoteComponent: qc, TradeComponent: tc,
		}

		trima, err := NewTriangularMovingAverage(&params)
		check("trima == nil", true, trima == nil)
		check("err", errbc, err.Error())
	})

	t.Run("invalid quote component", func(t *testing.T) {
		t.Parallel()
		params := TriangularMovingAverageParams{
			Length: 5, BarComponent: bc, QuoteComponent: data.QuoteComponent(9999), TradeComponent: tc,
		}

		trima, err := NewTriangularMovingAverage(&params)
		check("trima == nil", true, trima == nil)
		check("err", errqc, err.Error())
	})

	t.Run("invalid trade component", func(t *testing.T) {
		t.Parallel()
		params := TriangularMovingAverageParams{
			Length: 5, BarComponent: bc, QuoteComponent: qc, TradeComponent: data.TradeComponent(9999),
		}

		trima, err := NewTriangularMovingAverage(&params)
		check("trima == nil", true, trima == nil)
		check("err", errtc, err.Error())
	})
}

func testTriangularMovingAverageCreate(length int) *TriangularMovingAverage {
	params := TriangularMovingAverageParams{
		Length: length, BarComponent: data.BarClosePrice, QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	trima, _ := NewTriangularMovingAverage(&params)

	return trima
}
