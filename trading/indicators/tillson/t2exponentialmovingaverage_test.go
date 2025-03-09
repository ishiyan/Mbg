package tillson //nolint:testpackage

import (
	"math"
	"testing"
	"time"

	"mbg/trading/data"                        //nolint:depguard
	"mbg/trading/indicators/indicator"        //nolint:depguard
	"mbg/trading/indicators/indicator/output" //nolint:depguard
)

func testT2ExponentialMovingAverageTime() time.Time {
	return time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})
}

func testT2ExponentialMovingAverageInput() []float64 { //nolint:dupl
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

// Expected data is taken from the modified TA-L:ib (http://ta-lib.org/) test_T3.xls: test_T2.xls, T2(5,0.7) column.

//nolint:lll
func testT2ExponentialMovingAverageExpected() []float64 {
	return []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		92.00445682439470, 90.91065523008910, 90.64635861170230, 90.30847892058210, 89.67203184711450, 88.90373302682710,
		87.58685794329080, 85.95483674752900, 84.70056892267930, 83.35230597480420, 83.01639757937170, 84.26606438522640,
		85.11761359114740, 85.64273951236740, 85.71656813204080, 86.07519575846890, 86.14040917330100, 86.84731118619480,
		87.49707083315330, 87.60487694128920, 87.48806230196600, 86.85037109765170, 86.00510559137550, 85.19270124163530,
		84.44671220986690, 84.39562430058820, 85.47846964605910, 86.77863391237980, 88.27480023486270, 89.43282744282170,
		90.38671380457610, 90.47144856898390, 90.70844123290330, 90.80974940198440, 90.47870378112050, 90.01014971788120,
		88.50241981733770, 86.77695943257320, 85.15371443592010, 84.41215046544750, 84.35591075095750, 84.68800316786790,
		85.75412809581360, 86.93319023556810, 87.74938215169110, 88.24027778265170, 89.34170209097540, 90.29228263445890,
		91.39771790246850, 92.36534535927510, 92.98185469287250, 92.98804953756390, 92.35873030252150, 91.56001444442250,
		90.66809450828110, 89.04851123746310, 87.09681376947570, 85.81489615554240, 85.29909582783180, 87.87278112697240,
		91.57930166956580, 96.09718145905370, 100.17004554905500, 102.36263587267900, 103.39235683035700, 104.28507155830600,
		105.25701353369700, 105.97310800537900, 106.42554279253100, 106.31201565835700, 106.95673692587200, 107.85537745407500,
		108.90779472362600, 110.30293857497500, 113.94246316874900, 116.72555020318500, 118.38770601306100, 119.39753212263300,
		119.61486726510600, 119.11850573798600, 118.15216910748100, 116.41234112714800, 114.45865964016000, 114.64061162124700,
		114.99246753582000, 115.33333018479700, 114.64625300711600, 114.18195589036400, 113.69963506110000, 114.05025810623800,
		115.68652710013500, 116.48211079384600, 116.91736166744200, 116.71412173592500, 116.10716021480600, 115.77116121989300,
		115.68159156874600, 116.84472084196100, 118.05979798379400, 119.16496881132300, 120.94593869721900, 122.23096010279300,
		122.95608471296700, 123.22656209567300, 123.39820574539500, 123.34529134741000, 123.70294192360900, 125.21356151217400,
		127.19355195747100, 129.19266295422300, 130.45262728011500, 131.60256664339700, 132.72482116981200, 134.37601221394000,
		135.92461267673900, 137.09424525891600, 137.70388483421700, 137.73781691122000, 137.52589963778700, 136.86260850470000,
		134.64367631336900, 132.54215730025200, 129.69729772115400, 127.44287309781500, 125.45301927369100, 124.78149677056900,
		125.23737028614000, 125.28834765521900, 125.29811473230200, 124.49118133074600, 122.91167204427900, 121.21791414741600,
		120.98626501620100, 121.39013089649900, 121.63114611836900, 121.09667781630300, 121.38604904393200, 121.37683602700200,
		121.81714963968400, 123.27680991970300, 125.04065301695500, 125.42231478992600, 125.02871976565900, 124.16184698358900,
		123.95348210659900, 123.42096848684700, 122.98054879392000, 122.80175403486300, 123.00506164290600, 123.10341273782800,
		123.47654621169500, 124.50686275946200, 125.20780013410700, 126.36056179556500, 128.16193338760800, 129.48687497042000,
		131.31008039649500, 132.94986405572100, 133.52752016483300, 133.82004244574000, 133.59810463804000, 132.76202894406400,
		130.77137039821000, 129.94273481747400, 128.95593356768700, 127.66988673048200, 125.80391233666600, 124.87150365869300,
		123.99029524198700, 123.45575776078900, 122.35139824579100, 121.53836332656800, 120.24081578533100, 119.58685232736100,
		119.81980247526100, 119.65181893967000, 118.73183613794600, 117.08643352117200, 115.74380263433200, 113.79159129198400,
		110.89385419602800, 108.78809510901800, 107.63783653194000, 106.88467024071700, 106.49446844969200, 106.31788640074800,
		102.45936340772700, 98.92286025711640, 96.26320146598560, 94.89830130673550, 93.71214241680890, 93.35493402958790,
		94.19091222655660, 94.93222890927330, 95.05352335355920, 94.85884233631840, 93.94374172767060, 92.68201046821450,
		92.46778442700140, 92.58855930337340, 93.61995692101500, 94.27141015939870, 94.87276841707030, 94.89831480519190,
		94.85776894321340, 94.59245060865460, 95.32163415260000, 97.64623243404460, 100.94595782878400, 103.42636980462600,
		104.70917474084900, 105.40458887459900, 105.52523492436600, 105.12339759255700, 104.68875644984300, 104.75482770270600,
		106.53547902116200, 109.48515731469700, 112.33128725513400, 114.90966925830500, 115.64161420440900, 114.51441441748600,
		113.13113208320500, 111.81663297883600, 110.24787839832700, 109.44340073972700, 109.26032936270900, 109.13994050792700,
		109.29451387663700, 109.00897104476500, 108.79138478114900, 108.91808951962300, 109.15266693557400, 109.18435708789600,
		109.08448284304800, 108.75304884708700,
	}
}

func TestT2ExponentialMovingAverageUpdate(t *testing.T) { //nolint: funlen
	t.Parallel()

	check := func(index int, exp, act float64) {
		t.Helper()

		if math.Abs(exp-act) > 1e-8 {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	checkNaN := func(index int, act float64) {
		t.Helper()

		if !math.IsNaN(act) {
			t.Errorf("[%v] is incorrect: expected NaN, actual %v", index, act)
		}
	}

	input := testT2ExponentialMovingAverageInput()

	const (
		l       = 5
		lprimed = 4*l - 4
	)

	t.Run("length = 5, firstIsAverage = false (Metastock)", func(t *testing.T) {
		t.Parallel()

		const (
			firstCheck = lprimed + 43
		)

		t2 := testT2ExponentialMovingAverageCreateLength(l, false, 0.7)

		exp := testT2ExponentialMovingAverageExpected()

		for i := 0; i < lprimed; i++ {
			checkNaN(i, t2.Update(input[i]))
		}

		for i := lprimed; i < len(input); i++ {
			act := t2.Update(input[i])

			if i >= firstCheck {
				check(i, exp[i], act)
			}
		}

		checkNaN(0, t2.Update(math.NaN()))
	})

	t.Run("length = 5, firstIsAverage = true (t2.xls)", func(t *testing.T) {
		t.Parallel()

		const (
			firstCheck = lprimed + 1
		)

		t2 := testT2ExponentialMovingAverageCreateLength(l, true, 0.7)

		exp := testT2ExponentialMovingAverageExpected()

		for i := 0; i < lprimed; i++ {
			checkNaN(i, t2.Update(input[i]))
		}

		for i := lprimed; i < len(input); i++ {
			act := t2.Update(input[i])

			if i >= firstCheck {
				check(i, exp[i], act)
			}
		}

		checkNaN(0, t2.Update(math.NaN()))
	})
}

func TestT2ExponentialMovingAverageUpdateEntity(t *testing.T) { //nolint: funlen, dupl
	t.Parallel()

	const (
		l        = 2
		lprimed  = 4*l - 4
		alpha    = 2. / float64(l+1)
		inp      = 3.
		expFalse = 0.7437037037037035
		expTrue  = 0.6711111111111108
	)

	time := testT2ExponentialMovingAverageTime()
	check := func(exp float64, act indicator.Output) {
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
		t2 := testT2ExponentialMovingAverageCreateLength(l, false, 0.7)

		for i := 0; i < lprimed; i++ {
			t2.Update(0.)
		}

		check(expFalse, t2.UpdateScalar(&s))
	})

	t.Run("update bar", func(t *testing.T) {
		t.Parallel()

		b := data.Bar{Time: time, Close: inp}
		t2 := testT2ExponentialMovingAverageCreateLength(l, true, 0.7)

		for i := 0; i < lprimed; i++ {
			t2.Update(0.)
		}

		check(expTrue, t2.UpdateBar(&b))
	})

	t.Run("update quote", func(t *testing.T) {
		t.Parallel()

		q := data.Quote{Time: time, Bid: inp}
		t2 := testT2ExponentialMovingAverageCreateLength(l, false, 0.7)

		for i := 0; i < lprimed; i++ {
			t2.Update(0.)
		}

		check(expFalse, t2.UpdateQuote(&q))
	})

	t.Run("update trade", func(t *testing.T) {
		t.Parallel()

		r := data.Trade{Time: time, Price: inp}
		t2 := testT2ExponentialMovingAverageCreateLength(l, true, 0.7)

		for i := 0; i < lprimed; i++ {
			t2.Update(0.)
		}

		check(expTrue, t2.UpdateTrade(&r))
	})
}

func TestT2ExponentialMovingAverageIsPrimed(t *testing.T) {
	t.Parallel()

	input := testT2ExponentialMovingAverageInput()
	check := func(index int, exp, act bool) {
		t.Helper()

		if exp != act {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	const (
		l       = 5
		lprimed = 4*l - 4
	)

	t.Run("length = 5, firstIsAverage = true", func(t *testing.T) {
		t.Parallel()

		t2 := testT2ExponentialMovingAverageCreateLength(l, true, 0.7)

		check(0, false, t2.IsPrimed())

		for i := 0; i < lprimed; i++ {
			t2.Update(input[i])
			check(i+1, false, t2.IsPrimed())
		}

		for i := lprimed; i < len(input); i++ {
			t2.Update(input[i])
			check(i+1, true, t2.IsPrimed())
		}
	})

	t.Run("length = 5, firstIsAverage = false (Metastock)", func(t *testing.T) {
		t.Parallel()

		t2 := testT2ExponentialMovingAverageCreateLength(l, false, 0.7)

		check(0, false, t2.IsPrimed())

		for i := 0; i < lprimed; i++ {
			t2.Update(input[i])
			check(i+1, false, t2.IsPrimed())
		}

		for i := lprimed; i < len(input); i++ {
			t2.Update(input[i])
			check(i+1, true, t2.IsPrimed())
		}
	})
}

func TestT2ExponentialMovingAverageMetadata(t *testing.T) {
	t.Parallel()

	check := func(what string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", what, exp, act)
		}
	}

	checkInstance := func(act indicator.Metadata, name string) {
		check("Type", indicator.T2ExponentialMovingAverage, act.Type)
		check("len(Outputs)", 1, len(act.Outputs))
		check("Outputs[0].Kind", int(T2ExponentialMovingAverageValue), act.Outputs[0].Kind)
		check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
		check("Outputs[0].Name", name, act.Outputs[0].Name)
		check("Outputs[0].Description", "T2 exponential moving average "+name, act.Outputs[0].Description)
	}

	t.Run("length = 10, v=0.3333, firstIsAverage = true", func(t *testing.T) {
		t.Parallel()

		t2 := testT2ExponentialMovingAverageCreateLength(10, true, 0.3333)
		act := t2.Metadata()
		checkInstance(act, "t2(10, 0.33)")
	})

	t.Run("alpha = 2/11 = 0.18181818..., v=0.3333333, firstIsAverage = false", func(t *testing.T) {
		t.Parallel()

		// α = 2 / (ℓ + 1) = 2/11 = 0.18181818...
		const alpha = 2. / 11.

		t2 := testT2ExponentialMovingAverageCreateAlpha(alpha, false, 0.3333333)
		act := t2.Metadata()
		checkInstance(act, "t2(0.1818 (10), 0.33)")
	})
}

func TestNewT2ExponentialMovingAverage(t *testing.T) { //nolint: funlen, maintidx
	t.Parallel()

	const (
		bc     data.BarComponent   = data.BarMedianPrice
		qc     data.QuoteComponent = data.QuoteMidPrice
		tc     data.TradeComponent = data.TradePrice
		length                     = 10
		alpha                      = 2. / 11.

		errlen   = "invalid t2 exponential moving average parameters: length should be greater than 1"
		erralpha = "invalid t2 exponential moving average parameters: smoothing factor should be in range [0, 1]"
		errvol   = "invalid t2 exponential moving average parameters: volume factor should be in range [0, 1]"
		errbc    = "invalid t2 exponential moving average parameters: 9999: unknown bar component"
		errqc    = "invalid t2 exponential moving average parameters: 9999: unknown quote component"
		errtc    = "invalid t2 exponential moving average parameters: 9999: unknown trade component"
	)

	check := func(name string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", name, exp, act)
		}
	}

	checkInstance := func(
		t2 *T2ExponentialMovingAverage, name string, length int, alpha float64, firstIsAverage bool,
	) {
		check("name", name, t2.name)
		check("description", "T2 exponential moving average "+name, t2.description)
		check("firstIsAverage", firstIsAverage, t2.firstIsAverage)
		check("primed", false, t2.primed)
		check("length", length, t2.length)
		check("length2", length+length-1, t2.length2)
		check("length3", length+length+length-2, t2.length3)
		check("length4", length+length+length+length-3, t2.length4)
		check("smoothingFactor", alpha, t2.smoothingFactor)
		check("count", 0, t2.count)
		check("sum", 0., t2.sum)
		check("ema1", 0., t2.ema1)
		check("ema2", 0., t2.ema2)
		check("ema3", 0., t2.ema3)
		check("ema4", 0., t2.ema4)
		check("barFunc == nil", false, t2.barFunc == nil)
		check("quoteFunc == nil", false, t2.quoteFunc == nil)
		check("tradeFunc == nil", false, t2.tradeFunc == nil)
	}

	t.Run("length > 1, firstIsAverage = false", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: length, VolumeFactor: 0.7, FirstIsAverage: false,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(t2, "t2(10, 0.70)", length, alpha, false)
	})

	t.Run("length = 1, firstIsAverage = true", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: 1, VolumeFactor: 0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length = 0", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: 0, VolumeFactor: 0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length < 0", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: -1, VolumeFactor: 0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", errlen, err.Error())
	})

	t.Run("epsilon < α ≤ 1", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: alpha, VolumeFactor: 0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(t2, "t2(0.1818 (10), 0.70)", length, alpha, true)
	})

	t.Run("0 < α < epsilon", func(t *testing.T) {
		t.Parallel()

		const (
			alpha  = 0.00000001
			length = 199999999 // 2./0.00000001 - 1.
		)

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: alpha, VolumeFactor: 0.7, FirstIsAverage: false,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(t2, "t2(0.0000 (199999999), 0.70)", length, alpha, false)
	})

	t.Run("α = 0", func(t *testing.T) {
		t.Parallel()

		const (
			alpha  = 0.00000001
			length = 199999999 // 2./0.00000001 - 1.
		)

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 0, VolumeFactor: 0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(t2, "t2(0.0000 (199999999), 0.70)", length, alpha, true)
	})

	t.Run("α = 1", func(t *testing.T) {
		t.Parallel()

		const (
			alpha  = 1
			length = 1 // 2./1 - 1.
		)

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: alpha, VolumeFactor: 0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(t2, "t2(1.0000 (1), 0.70)", length, alpha, true)
	})

	t.Run("α < 0", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: -1, VolumeFactor: 0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", erralpha, err.Error())
	})

	t.Run("α > 1", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 2, VolumeFactor: 0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", erralpha, err.Error())
	})

	t.Run("volume factor = 0.5", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: length, VolumeFactor: 0.5, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(t2, "t2(10, 0.50)", length, alpha, true)
	})

	t.Run("volume factor = 0", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: length, VolumeFactor: 0, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(t2, "t2(10, 0.00)", length, alpha, true)
	})

	t.Run("volume factor = 1", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: length, VolumeFactor: 1, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(t2, "t2(10, 1.00)", length, alpha, true)
	})

	t.Run("volume factor < 0", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: 3, VolumeFactor: -0.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", errvol, err.Error())
	})

	t.Run("volume factor > 1", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageLengthParams{
			Length: 3, VolumeFactor: 1.7, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageLength(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", errvol, err.Error())
	})

	t.Run("invalid bar component", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 1, VolumeFactor: 0.5, FirstIsAverage: true,
			BarComponent: data.BarComponent(9999), QuoteComponent: qc, TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", errbc, err.Error())
	})

	t.Run("invalid quote component", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 1, VolumeFactor: 0.5, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: data.QuoteComponent(9999), TradeComponent: tc,
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", errqc, err.Error())
	})

	t.Run("invalid trade component", func(t *testing.T) {
		t.Parallel()

		params := T2ExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 1, VolumeFactor: 0.5, FirstIsAverage: true,
			BarComponent: bc, QuoteComponent: qc, TradeComponent: data.TradeComponent(9999),
		}

		t2, err := NewT2ExponentialMovingAverageSmoothingFactor(&params)
		check("t2 == nil", true, t2 == nil)
		check("err", errtc, err.Error())
	})
}

func testT2ExponentialMovingAverageCreateLength(
	length int, firstIsAverage bool, volume float64,
) *T2ExponentialMovingAverage {
	params := T2ExponentialMovingAverageLengthParams{
		Length: length, VolumeFactor: volume, FirstIsAverage: firstIsAverage,
		BarComponent: data.BarClosePrice, QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	t2, _ := NewT2ExponentialMovingAverageLength(&params)

	return t2
}

func testT2ExponentialMovingAverageCreateAlpha(
	alpha float64, firstIsAverage bool, volume float64,
) *T2ExponentialMovingAverage {
	params := T2ExponentialMovingAverageSmoothingFactorParams{
		SmoothingFactor: alpha, VolumeFactor: volume, FirstIsAverage: firstIsAverage,
		BarComponent: data.BarClosePrice, QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	t2, _ := NewT2ExponentialMovingAverageSmoothingFactor(&params)

	return t2
}
