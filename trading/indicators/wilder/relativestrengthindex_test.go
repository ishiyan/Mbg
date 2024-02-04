//nolint:testpackage
package wilder

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
// Output data, length=14.
// Taken from TA-Lib (http://ta-lib.org/) tests, test_rsi.c.
//
// static TA_Test tableTest[] =
// {
// /**********************/
// /*      RSI TEST      */
// /**********************/
// { 1, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      0, 49.14,  14,  252-14 }, /* First Value */
// { 0, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      1, 52.32,  14,  252-14 },
// { 0, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      2, 46.07,  14,  252-14 },
// { 0, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 252-15, 49.63,  14,  252-14 },  /* Last Value */
//
// /* No output value. */
// { 0, TA_RSI_TEST, 0, 1, 1,  14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 0, 0, 0, 0},
//
// /* One value tests. */
// { 0, TA_RSI_TEST, 0, 14,  14, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 0, 49.14,     14, 1},
//
// /* Index too low test. */
// { 0, TA_RSI_TEST, 0, 0,  15, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 0, 49.14,     14, 2},
// { 0, TA_RSI_TEST, 0, 1,  15, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 0, 49.14,     14, 2},
// { 0, TA_RSI_TEST, 0, 2,  16, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 0, 49.14,     14, 3},
// { 0, TA_RSI_TEST, 0, 2,  16, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 1, 52.32,     14, 3},
// { 0, TA_RSI_TEST, 0, 2,  16, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 2, 46.07,     14, 3},
// { 0, TA_RSI_TEST, 0, 0,  14, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 0, 49.14,     14, 1},
// { 0, TA_RSI_TEST, 0, 0,  13, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 0, 49.14,     14, 0},
//
// /* Test with 1 unstable price bar. Test for period 1, 2, 14 */
// { 0, TA_RSI_TEST, 1, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      0,     52.32,  15,  252-(14+1) },
// { 0, TA_RSI_TEST, 1, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      1,     46.07,  15,  252-(14+1) },
// { 0, TA_RSI_TEST, 1, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 252-(15+1), 49.63,  15,  252-(14+1) },  /* Last Value */
//
// /* Test with 2 unstable price bar. Test for period 1, 2, 14 */
// { 0, TA_RSI_TEST, 2, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS,      0,     46.07,  16,  252-(14+2) },
// { 0, TA_RSI_TEST, 2, 0, 251, 14, TA_COMPATIBILITY_DEFAULT, TA_SUCCESS, 252-(15+2), 49.63,  16,  252-(14+2) },  /* Last Value */
//
//
// /**********************/
// /* RSI Metastock TEST */
// /**********************/
// { 1, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,      0, 47.11,  13,  252-13 }, /* First Value */
// { 0, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,      1, 49.14,  13,  252-13 },
// { 0, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,      2, 52.32,  13,  252-13 },
// { 0, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,      3, 46.07,  13,  252-13 },
// { 0, TA_RSI_TEST, 0, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 252-14, 49.63,  13,  252-13 }, /* Last Value */
//
// /* No output value. */
// { 0, TA_RSI_TEST, 0, 1, 1,  14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 0, 0, 0},
//
// /* One value tests. */
// { 0, TA_RSI_TEST, 0, 13, 13, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 47.11, 13, 1},
// { 0, TA_RSI_TEST, 0, 13, 13, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 47.11, 13, 1},
//
// /* Index too low test. */
// { 0, TA_RSI_TEST, 0, 0,  15, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 47.11,     13, 3},
// { 0, TA_RSI_TEST, 0, 1,  15, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 47.11,     13, 3},
// { 0, TA_RSI_TEST, 0, 2,  16, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 47.11,     13, 4},
// { 0, TA_RSI_TEST, 0, 2,  16, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 1, 49.14,     13, 4},
// { 0, TA_RSI_TEST, 0, 2,  16, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 2, 52.32,     13, 4},
// { 0, TA_RSI_TEST, 0, 0,  14, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 47.11,     13, 2},
// { 0, TA_RSI_TEST, 0, 0,  13, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 47.11,     13, 1},
// { 0, TA_RSI_TEST, 0, 0,  12, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 0, 47.11,     13, 0},
//
// /* Test with 1 unstable price bar. Test for period 1, 2, 14 */
// { 0, TA_RSI_TEST, 1, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,      0,     49.14,  14,  252-(13+1) },
// { 0, TA_RSI_TEST, 1, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,      1,     52.32,  14,  252-(13+1) },
// { 0, TA_RSI_TEST, 1, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 252-(14+1), 49.63,  14,  252-(13+1) },  /* Last Value */
//
// /* Test with 2 unstable price bar. Test for period 1, 2, 14 */
// { 0, TA_RSI_TEST, 2, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS,      0,     52.32,  15,  252-(13+2) },
// { 0, TA_RSI_TEST, 2, 0, 251, 14, TA_COMPATIBILITY_METASTOCK, TA_SUCCESS, 252-(14+2), 49.63,  15,  252-(13+2) },  /* Last Value */

func testRelativeStrengthIndexTime() time.Time {
	return time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})
}

//nolint:dupl
func testRelativeStrengthIndexInput() []float64 {
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

//nolint:dupl
func testRelativeStrengthIndexLength9Output() []float64 {
	return []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(),
		53.58197932053180, 64.39471727507700, 67.05232851017130, 69.86404770177220, 42.02491601015560, 45.22756910292510,
		50.21882968739720, 41.61560247408060, 42.19395561081010, 49.01252917837510, 44.48189228865830, 40.84141245952300,
		39.04398236168950, 32.87048941206650, 29.57860354761250, 31.73890507799230, 27.67911650675580, 40.10971224636700,
		54.05268743674910, 46.73627112937570, 46.42887811500740, 43.83413208552280, 49.75704851163790, 45.94852006011590,
		55.23270950285110, 53.59153132366940, 48.42743797113670, 48.23075576524750, 42.02263869372570, 39.97852986218500,
		38.99717440722570, 37.35955423299950, 47.83355765059820, 60.67795922371490, 61.40231197146480, 66.05873660984950,
		64.24121878745430, 66.32584249167020, 53.83940672480780, 60.66750683371070, 58.24438584587450, 51.44787179990230,
		50.47867112141290, 34.88430505685490, 32.91666578055190, 30.88168929561010, 41.26986487163010, 45.33077140014450,
		47.73423225248430, 58.32837805203910, 59.47102733703560, 56.13506105425540, 55.46706692093850, 67.53283151601450,
		64.85526841473460, 69.86078994580260, 70.53523498917450, 68.16702266660970, 58.91151879407500, 49.79479569466790,
		48.27987313003640, 44.29419915450230, 31.19482365551140, 26.80270535389530, 35.27551156012400, 40.75332337605330,
		70.54050110525990, 73.95904732565060, 79.19876498767510, 80.13960925809880, 68.61190552802890, 68.31185131820730,
		71.53969720865650, 73.64933250885390, 73.15095438728780, 73.25682100808600, 66.70186428873520, 73.97815271099910,
		75.03675457631830, 76.85585802766400, 79.97320096459430, 88.15192061544870, 76.56874258621210, 73.74255533212520,
		74.35115499187500, 69.59648513548700, 64.29968057558690, 60.01740021198470, 50.06884724101550, 46.70580329150040,
		63.10623960986540, 57.48954473386070, 57.48954473386070, 47.43211280907020, 51.59963082814190, 49.55396420426640,
		56.82805632634870, 65.14340281635010, 54.88849410248800, 55.65759792059200, 51.14040364996210, 48.74952766857550,
		51.94892125456700, 52.97575218037450, 64.19054679341760, 62.45237026428300, 63.69118241763650, 71.30016815072380,
		65.93706793406470, 64.24314349701610, 62.64319513343330, 64.12446609473960, 61.38612510239970, 67.18527404483690,
		76.21674460383040, 78.70850061821470, 80.36183117224360, 73.90554199486880, 76.92541163796860, 78.74480234653510,
		83.02232633059410, 83.49763716604510, 83.58181674300740, 79.47697005694710, 73.42407113037200, 73.02470809100690,
		62.67066439895830, 38.49075120518990, 41.47689623459360, 30.19628536043190, 33.90204625499760, 30.40539186757520,
		42.40504170309070, 48.90114487671880, 41.46584657695990, 42.51128386200710, 34.95999631336530, 29.94937957476250,
		28.55836269864940, 44.25923005927580, 45.15551597945060, 41.94883579927610, 35.68176155664590, 47.85435226814970,
		42.85836705608540, 49.01107731347630, 58.04587582550450, 60.28251218726440, 48.39092739162940, 46.32589579835680,
		43.64480310546720, 50.87497508565820, 44.98549174965810, 46.05091336696410, 47.80284344330750, 51.12344517590550,
		48.47252144889760, 53.09993810940280, 60.86357330772370, 55.55593305281250, 63.34701846495520, 69.63568591658340,
		64.66675161504210, 71.88817672263370, 72.28643452081260, 61.94062123564820, 63.99824373140140, 58.73569358939160,
		51.93266237416960, 39.61362588908910, 52.61420760335070, 45.62624425414420, 41.63716997118120, 35.81473045195240,
		44.08691694257760, 40.20108523396490, 41.64760858559580, 35.07967607313700, 38.08147203419070, 32.27655864555370,
		39.09012863501310, 45.30946620643950, 39.34674507898310, 34.25695755518180, 29.83776928966220, 32.40144458502860,
		26.27987129912900, 21.07358886657860, 27.42565375634370, 30.25744020511560, 28.98509595453240, 29.44547109066120,
		29.23227584948250, 14.01321180142660, 22.36818266432260, 22.31011714761200, 27.25688548992740, 24.55872460536480,
		30.55964401838910, 39.75757239519880, 37.24259651380360, 34.10369033031380, 33.38572630887740, 28.99943906698160,
		27.13003133987030, 39.49075182691700, 38.84785922377950, 48.76007517850250, 44.01122602082110, 46.55832515418050,
		42.07797864021900, 43.95518354715070, 41.59886911288220, 54.88896553560580, 66.77671178115900, 72.22467867283780,
		66.56878756008010, 61.87092229131640, 62.81709656924280, 60.09909138869550, 56.27742654036670, 57.25169835018080,
		62.13022712593020, 74.02676854453240, 78.73051260042840, 79.36173872303730, 81.05524157612620, 63.67022410651560,
		52.41397722414160, 53.88319301379560, 52.70630339294140, 46.93940243540770, 52.61643968737750, 54.63598908385410,
		52.16558300983710, 54.80915051480850, 47.94075432094380, 50.00215267492850, 53.75665654988040, 53.96316225985600,
		50.53573892477520, 49.44529621603770, 45.55291315138690,
	}
}

func TestRelativeStrengthIndexUpdate(t *testing.T) { //nolint: funlen
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

	input := testRelativeStrengthIndexInput()

	t.Run("length = 9 (excel reference implementation)", func(t *testing.T) {
		output := testRelativeStrengthIndexLength9Output()

		t.Parallel()
		rsi := testRelativeStrengthIndexCreate(9)

		for i := 0; i < len(input); i++ {
			act := rsi.Update(input[i])

			if i < 9 {
				checkNaN(i, act)
			} else {
				check(i, output[i], act)
			}
		}

		checkNaN(0, rsi.Update(math.NaN()))
	})

	t.Run("length = 14", func(t *testing.T) {
		const ( // Values from index=0 to index=13 are NaN.
			i14value  = 49.14733969986360 // Index=14 value, 49.14.
			i15value  = 52.32555279533660 // Index=15 value, 52.32.
			i16value  = 46.07239657691370 // Index=16 value, 46.07.
			i251value = 49.63210207086760 // Index=251 (last) value, 49.63.
		)

		t.Parallel()
		rsi := testRelativeStrengthIndexCreate(14)

		for i := 0; i < 13; i++ {
			checkNaN(i, rsi.Update(input[i]))
		}

		for i := 13; i < len(input); i++ {
			act := rsi.Update(input[i])

			switch i {
			case 14:
				check(i, i14value, act)
			case 15:
				check(i, i15value, act)
			case 16:
				check(i, i16value, act)
			case 251:
				check(i, i251value, act)
			}
		}

		checkNaN(0, rsi.Update(math.NaN()))
	})
}

func TestRelativeStrengthIndexUpdateEntity(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		l   = 2
		inp = 3.
		exp = 100.
	)

	time := testRelativeStrengthIndexTime()
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
		rsi := testRelativeStrengthIndexCreate(l)
		rsi.Update(0.)
		rsi.Update(0.)
		check(rsi.UpdateScalar(&s))
	})

	t.Run("update bar", func(t *testing.T) {
		t.Parallel()

		b := data.Bar{Time: time, Close: inp}
		rsi := testRelativeStrengthIndexCreate(l)
		rsi.Update(0.)
		rsi.Update(0.)
		check(rsi.UpdateBar(&b))
	})

	t.Run("update quote", func(t *testing.T) {
		t.Parallel()

		q := data.Quote{Time: time, Bid: inp}
		rsi := testRelativeStrengthIndexCreate(l)
		rsi.Update(0.)
		rsi.Update(0.)
		check(rsi.UpdateQuote(&q))
	})

	t.Run("update trade", func(t *testing.T) {
		t.Parallel()

		r := data.Trade{Time: time, Price: inp}
		rsi := testRelativeStrengthIndexCreate(l)
		rsi.Update(0.)
		rsi.Update(0.)
		check(rsi.UpdateTrade(&r))
	})
}

func TestRelativeStrengthIndexIsPrimed(t *testing.T) { //nolint:funlen
	t.Parallel()

	input := testRelativeStrengthIndexInput()
	check := func(index int, exp, act bool) {
		t.Helper()

		if exp != act {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	t.Run("length = 1", func(t *testing.T) {
		t.Parallel()
		rsi := testRelativeStrengthIndexCreate(1)

		check(-1, false, rsi.IsPrimed())

		for i := 0; i < 1; i++ {
			rsi.Update(input[i])
			check(i, false, rsi.IsPrimed())
		}

		for i := 1; i < len(input); i++ {
			rsi.Update(input[i])
			check(i, true, rsi.IsPrimed())
		}
	})

	t.Run("length = 2", func(t *testing.T) {
		t.Parallel()
		rsi := testRelativeStrengthIndexCreate(2)

		check(-1, false, rsi.IsPrimed())

		for i := 0; i < 2; i++ {
			rsi.Update(input[i])
			check(i, false, rsi.IsPrimed())
		}

		for i := 2; i < len(input); i++ {
			rsi.Update(input[i])
			check(i, true, rsi.IsPrimed())
		}
	})

	t.Run("length = 3", func(t *testing.T) {
		t.Parallel()
		rsi := testRelativeStrengthIndexCreate(3)

		check(-1, false, rsi.IsPrimed())

		for i := 0; i < 3; i++ {
			rsi.Update(input[i])
			check(i, false, rsi.IsPrimed())
		}

		for i := 3; i < len(input); i++ {
			rsi.Update(input[i])
			check(i, true, rsi.IsPrimed())
		}
	})

	t.Run("length = 5", func(t *testing.T) {
		t.Parallel()
		rsi := testRelativeStrengthIndexCreate(5)

		check(-1, false, rsi.IsPrimed())

		for i := 0; i < 5; i++ {
			rsi.Update(input[i])
			check(i, false, rsi.IsPrimed())
		}

		for i := 5; i < len(input); i++ {
			rsi.Update(input[i])
			check(i, true, rsi.IsPrimed())
		}
	})

	t.Run("length = 10", func(t *testing.T) {
		t.Parallel()
		rsi := testRelativeStrengthIndexCreate(10)

		check(-1, false, rsi.IsPrimed())

		for i := 0; i < 10; i++ {
			rsi.Update(input[i])
			check(i, false, rsi.IsPrimed())
		}

		for i := 10; i < len(input); i++ {
			rsi.Update(input[i])
			check(i, true, rsi.IsPrimed())
		}
	})
}

func TestRelativeStrengthIndexMetadata(t *testing.T) {
	t.Parallel()

	rsi := testRelativeStrengthIndexCreate(5)
	act := rsi.Metadata()

	check := func(what string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", what, exp, act)
		}
	}

	check("Type", indicator.RelativeStrengthIndex, act.Type)
	check("len(Outputs)", 1, len(act.Outputs))
	check("Outputs[0].Kind", int(RelativeStrengthIndexValue), act.Outputs[0].Kind)
	check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
	check("Outputs[0].Name", "rsi(5)", act.Outputs[0].Name)
	check("Outputs[0].Description", "Relative Strength Index rsi(5)", act.Outputs[0].Description)
}

func TestNewRelativeStrengthIndex(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		bc data.BarComponent   = data.BarMedianPrice
		qc data.QuoteComponent = data.QuoteMidPrice
		tc data.TradeComponent = data.TradePrice

		length = 5
		errlen = "invalid relative strength index parameters: length should be positive"
		errbc  = "invalid relative strength index parameters: 9999: unknown bar component"
		errqc  = "invalid relative strength index parameters: 9999: unknown quote component"
		errtc  = "invalid relative strength index parameters: 9999: unknown trade component"
	)

	check := func(name string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", name, exp, act)
		}
	}

	t.Run("length > 0", func(t *testing.T) {
		t.Parallel()
		params := RelativeStrengthIndexParams{
			Length: length, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		rsi, err := NewRelativeStrengthIndex(&params)
		check("err == nil", true, err == nil)
		check("name", "rsi(5)", rsi.name)
		check("description", "Relative Strength Index rsi(5)", rsi.description)
		check("primed", false, rsi.primed)
		check("len(window)", length, len(rsi.window))
		check("windowLength", length, rsi.windowLength)
		check("windowCount", 0, rsi.windowCount)
		check("barFunc == nil", false, rsi.barFunc == nil)
		check("quoteFunc == nil", false, rsi.quoteFunc == nil)
		check("tradeFunc == nil", false, rsi.tradeFunc == nil)
	})

	t.Run("length = 1", func(t *testing.T) {
		t.Parallel()
		params := RelativeStrengthIndexParams{
			Length: 1, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		rsi, err := NewRelativeStrengthIndex(&params)
		check("err == nil", true, err == nil)
		check("name", "rsi(1)", rsi.name)
		check("description", "Relative Strength Index rsi(1)", rsi.description)
		check("primed", false, rsi.primed)
		check("len(window)", 1, len(rsi.window))
		check("windowLength", 1, rsi.windowLength)
		check("windowCount", 0, rsi.windowCount)
		check("barFunc == nil", false, rsi.barFunc == nil)
		check("quoteFunc == nil", false, rsi.quoteFunc == nil)
		check("tradeFunc == nil", false, rsi.tradeFunc == nil)
	})

	t.Run("length = 0", func(t *testing.T) {
		t.Parallel()
		params := RelativeStrengthIndexParams{
			Length: 0, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		rsi, err := NewRelativeStrengthIndex(&params)
		check("rsi == nil", true, rsi == nil)
		check("err", errlen, err.Error())
	})

	t.Run("length < 0", func(t *testing.T) {
		t.Parallel()
		params := RelativeStrengthIndexParams{
			Length: -1, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		rsi, err := NewRelativeStrengthIndex(&params)
		check("rsi == nil", true, rsi == nil)
		check("err", errlen, err.Error())
	})

	t.Run("invalid bar component", func(t *testing.T) {
		t.Parallel()
		params := RelativeStrengthIndexParams{
			Length: length, BarComponent: data.BarComponent(9999), QuoteComponent: qc, TradeComponent: tc,
		}

		rsi, err := NewRelativeStrengthIndex(&params)
		check("rsi == nil", true, rsi == nil)
		check("err", errbc, err.Error())
	})

	t.Run("invalid quote component", func(t *testing.T) {
		t.Parallel()
		params := RelativeStrengthIndexParams{
			Length: length, BarComponent: bc, QuoteComponent: data.QuoteComponent(9999), TradeComponent: tc,
		}

		rsi, err := NewRelativeStrengthIndex(&params)
		check("rsi == nil", true, rsi == nil)
		check("err", errqc, err.Error())
	})

	t.Run("invalid trade component", func(t *testing.T) {
		t.Parallel()
		params := RelativeStrengthIndexParams{
			Length: length, BarComponent: bc, QuoteComponent: qc, TradeComponent: data.TradeComponent(9999),
		}

		rsi, err := NewRelativeStrengthIndex(&params)
		check("rsi == nil", true, rsi == nil)
		check("err", errtc, err.Error())
	})
}

func testRelativeStrengthIndexCreate(length int) *RelativeStrengthIndex {
	params := RelativeStrengthIndexParams{
		Length: length, BarComponent: data.BarClosePrice, QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	rsi, _ := NewRelativeStrengthIndex(&params)

	return rsi
}
