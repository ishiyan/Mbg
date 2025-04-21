//nolint:testpackage
package ehlers

import (
	"math"
	"testing"
	"time"

	"mbg/trading/data"                                 //nolint:depguard
	"mbg/trading/indicators/ehlers/hilberttransformer" //nolint:depguard
	"mbg/trading/indicators/indicator"                 //nolint:depguard
	"mbg/trading/indicators/indicator/output"          //nolint:depguard
)

func testMesaAdaptiveMovingAverageTime() time.Time {
	return time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})
}

// Input data taken from TA-Lib (http://ta-lib.org/) tests, test_MAMA.xsl, Price, D5…D256, 252 entries.

func testMesaAdaptiveMovingAverageInput() []float64 {
	return []float64{
		92.0000, 93.1725, 95.3125, 94.8450, 94.4075, 94.1100, 93.5000, 91.7350, 90.9550, 91.6875,
		94.5000, 97.9700, 97.5775, 90.7825, 89.0325, 92.0950, 91.1550, 89.7175, 90.6100, 91.0000,
		88.9225, 87.5150, 86.4375, 83.8900, 83.0025, 82.8125, 82.8450, 86.7350, 86.8600, 87.5475,
		85.7800, 86.1725, 86.4375, 87.2500, 88.9375, 88.2050, 85.8125, 84.5950, 83.6575, 84.4550,
		83.5000, 86.7825, 88.1725, 89.2650, 90.8600, 90.7825, 91.8600, 90.3600, 89.8600, 90.9225,
		89.5000, 87.6725, 86.5000, 84.2825, 82.9075, 84.2500, 85.6875, 86.6100, 88.2825, 89.5325,
		89.5000, 88.0950, 90.6250, 92.2350, 91.6725, 92.5925, 93.0150, 91.1725, 90.9850, 90.3775,
		88.2500, 86.9075, 84.0925, 83.1875, 84.2525, 97.8600, 99.8750, 103.2650, 105.9375, 103.5000,
		103.1100, 103.6100, 104.6400, 106.8150, 104.9525, 105.5000, 107.1400, 109.7350, 109.8450, 110.9850,
		120.0000, 119.8750, 117.9075, 119.4075, 117.9525, 117.2200, 115.6425, 113.1100, 111.7500, 114.5175,
		114.7450, 115.4700, 112.5300, 112.0300, 113.4350, 114.2200, 119.5950, 117.9650, 118.7150, 115.0300,
		114.5300, 115.0000, 116.5300, 120.1850, 120.5000, 120.5950, 124.1850, 125.3750, 122.9700, 123.0000,
		124.4350, 123.4400, 124.0300, 128.1850, 129.6550, 130.8750, 132.3450, 132.0650, 133.8150, 135.6600,
		137.0350, 137.4700, 137.3450, 136.3150, 136.4400, 136.2850, 129.0950, 128.3100, 126.0000, 124.0300,
		123.9350, 125.0300, 127.2500, 125.6200, 125.5300, 123.9050, 120.6550, 119.9650, 120.7800, 124.0000,
		122.7800, 120.7200, 121.7800, 122.4050, 123.2500, 126.1850, 127.5600, 126.5650, 123.0600, 122.7150,
		123.5900, 122.3100, 122.4650, 123.9650, 123.9700, 124.1550, 124.4350, 127.0000, 125.5000, 128.8750,
		130.5350, 132.3150, 134.0650, 136.0350, 133.7800, 132.7500, 133.4700, 130.9700, 127.5950, 128.4400,
		127.9400, 125.8100, 124.6250, 122.7200, 124.0900, 123.2200, 121.4050, 120.9350, 118.2800, 118.3750,
		121.1550, 120.9050, 117.1250, 113.0600, 114.9050, 112.4350, 107.9350, 105.9700, 106.3700, 106.8450,
		106.9700, 110.0300, 91.0000, 93.5600, 93.6200, 95.3100, 94.1850, 94.7800, 97.6250, 97.5900,
		95.2500, 94.7200, 92.2200, 91.5650, 92.2200, 93.8100, 95.5900, 96.1850, 94.6250, 95.1200,
		94.0000, 93.7450, 95.9050, 101.7450, 106.4400, 107.9350, 103.4050, 105.0600, 104.1550, 103.3100,
		103.3450, 104.8400, 110.4050, 114.5000, 117.3150, 118.2500, 117.1850, 109.7500, 109.6550, 108.5300,
		106.2200, 107.7200, 109.8400, 109.0950, 109.0900, 109.1550, 109.3150, 109.0600, 109.9050, 109.6250,
		109.5300, 108.0600,
	}
}

// Expected data taken from TA-Lib (http://ta-lib.org/) tests, test_MAMA_new.xsl, MAMA, L5…AB256, 252 entries.

func testMesaAdaptiveMovingAverageExpected() []float64 {
	return []float64{
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(),
		math.NaN(), 82.81412500000000, 83.01016875000000, 83.20266031250000, 83.41990229687500,
		83.53790718203120, 83.66963682292970, 85.05356841146480, 85.16338999089160, 85.35209549134700,
		85.49474071677970, 85.51062868094070, 85.05281434047030, 84.98304862344680, 84.95664619227450,
		84.88381388266070, 85.83315694133040, 85.95012409426390, 86.11586788955070, 86.35307449507310,
		86.57454577031950, 89.21727288515970, 89.27440924090170, 89.30368877885670, 89.38462933991380,
		89.39039787291810, 89.30450297927220, 89.16427783030860, 86.72338891515430, 86.53259446939660,
		85.39129723469830, 85.40610737296340, 85.46630200431520, 85.60711190409940, 85.80338130889440,
		85.98821224344970, 86.09355163127720, 88.35927581563860, 88.55306202485670, 88.70903392361380,
		88.90320722743310, 89.10879686606150, 90.14064843303070, 90.18286601137920, 90.19259771081020,
		90.09546782526970, 89.93606943400620, 89.64389096230590, 86.41569548115300, 86.24460211882510,
		86.82537201288380, 87.47785341223960, 88.26721074162770, 89.15072520454630, 89.86818894431890,
		96.48909447215950, 96.84513974855150, 97.23488276112390, 97.71388862306770, 98.07581919191430,
		98.44702823231860, 101.82504408344900, 105.78002204172500, 107.81251102086200, 107.97113546981900,
		113.98556773491000, 116.93028386745500, 117.04756220351700, 117.16555909334200, 117.55902954667100,
		117.54046165237900, 117.44556356976000, 117.22878539127200, 116.95484612170900, 116.83297881562300,
		115.78898940781200, 115.77303993742100, 115.61088794055000, 115.43184354352200, 115.33200136634600,
		115.27640129802900, 117.43570064901500, 117.46216561656400, 118.08858280828200, 117.93565366786800,
		116.23282683393400, 116.17118549223700, 116.18912621762500, 116.38891990674400, 116.59447391140700,
		118.59473695570300, 118.87425010791800, 119.19928760252200, 121.08464380126100, 121.18041161119800,
		121.34314103063800, 122.39157051531900, 122.47349198955300, 122.75906739007500, 123.10386402057200,
		126.98943201028600, 129.66721600514300, 130.86610800257100, 131.01355260244300, 131.24587497232100,
		131.53533122370500, 131.83206466251900, 132.10771142939300, 134.21135571469700, 135.32567785734800,
		135.37364396448100, 135.05971176625700, 134.72222617794400, 134.28611486904700, 129.15805743452300,
		127.60796470846700, 127.47906647304400, 127.46761314939200, 127.37523249192200, 127.28297086732600,
		127.11407232396000, 126.79111870776200, 123.37805935388100, 123.24815638618700, 123.62407819309300,
		123.58187428343900, 123.43878056926700, 123.35584154080300, 123.30829946376300, 123.30538449057500,
		124.74519224528800, 124.88593263302300, 125.72546631651200, 124.39273315825600, 124.29395234553200,
		124.25875472825500, 124.16131699184200, 124.07650114225000, 124.02075057112500, 124.01821304256900,
		124.02505239044000, 124.04554977091800, 124.19327228237200, 124.84663614118600, 125.04805433412700,
		125.32240161742100, 128.81870080871000, 129.08101576827500, 129.42871497986100, 131.60435748993100,
		131.66163961543400, 131.75205763466200, 131.71295475292900, 131.50705701528300, 131.35370416451900,
		131.18301895629300, 128.49650947814600, 128.30293400423900, 128.02378730402700, 127.82709793882600,
		126.89727749566900, 124.15113874783400, 123.05638802581100, 122.68979938851700, 122.47405941909100,
		121.81452970954600, 121.35976485477300, 120.91051977170100, 119.99131672376500, 119.73700088757600,
		119.37190084319800, 118.49653563589600, 112.23326781794800, 111.94010442705100, 111.68534920569800,
		109.32767460284900, 109.36279087270700, 108.44465132907100, 107.70041876261800, 106.74281198325900,
		101.02640599163000, 97.60570299581490, 97.46441784602410, 97.47244695372290, 97.47832460603680,
		97.36690837573490, 97.23456295694820, 94.72728147847410, 94.56916740455040, 94.45170903432280,
		94.41962358260670, 95.00481179130340, 95.06382120173820, 95.04188014165130, 95.04578613456870,
		94.99349682784030, 94.36924841392010, 94.44603599322410, 94.81098419356290, 95.39243498388480,
		101.66371749194200, 101.75078161734500, 103.40539080867300, 103.44287126823900, 103.43622770482700,
		103.43166631958600, 103.50208300360600, 103.84722885342600, 109.17361442671300, 109.58068370537700,
		110.01414952010900, 113.59957476005400, 113.40709602205200, 113.21949122094900, 112.98501665990200,
		112.64676582690600, 112.40042753556100, 111.12021376778100, 111.01895307939200, 110.92250542542200,
		110.83413015415100, 110.75817364644300, 110.67326496412100, 110.28913248206100, 110.25592585795800,
		110.21962956506000, 110.11164808680700,
	}
}

func TestMesaAdaptiveMovingAverageUpdate(t *testing.T) {
	t.Parallel()

	check := func(index int, exp, act float64) {
		t.Helper()

		if math.Abs(exp-act) > 1e-9 {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	checkNaN := func(index int, act float64) {
		t.Helper()

		if !math.IsNaN(act) {
			t.Errorf("[%v] is incorrect: expected NaN, actual %v", index, act)
		}
	}

	input := testMesaAdaptiveMovingAverageInput()

	const (
		lprimed = 26
		f       = 3
		s       = 39
	)

	t.Run("reference implementation test_mama_new.xls", func(t *testing.T) {
		t.Parallel()

		mama := testMesaAdaptiveMovingAverageCreateLength(f, s)
		exp := testMesaAdaptiveMovingAverageExpected()

		for i := 0; i < lprimed; i++ {
			checkNaN(i, mama.Update(input[i]))
		}

		for i := lprimed; i < len(input); i++ {
			act := mama.Update(input[i])
			check(i, exp[i], act)
		}

		checkNaN(0, mama.Update(math.NaN()))
	})
}

func TestMesaAdaptiveMovingAverageUpdateEntity(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		lprimed  = 26
		fast     = 3
		slow     = 39
		inp      = 3.
		expected = 1.5
	)

	time := testMesaAdaptiveMovingAverageTime()
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
		mama := testMesaAdaptiveMovingAverageCreateLength(fast, slow)

		for i := 0; i < lprimed; i++ {
			mama.Update(0.)
		}

		check(expected, mama.UpdateScalar(&s))
	})

	t.Run("update bar", func(t *testing.T) {
		t.Parallel()

		b := data.Bar{Time: time, Close: inp}
		mama := testMesaAdaptiveMovingAverageCreateLength(fast, slow)

		for i := 0; i < lprimed; i++ {
			mama.Update(0.)
		}

		check(expected, mama.UpdateBar(&b))
	})

	t.Run("update quote", func(t *testing.T) {
		t.Parallel()

		q := data.Quote{Time: time, Bid: inp}
		mama := testMesaAdaptiveMovingAverageCreateLength(fast, slow)

		for i := 0; i < lprimed; i++ {
			mama.Update(0.)
		}

		check(expected, mama.UpdateQuote(&q))
	})

	t.Run("update trade", func(t *testing.T) {
		t.Parallel()

		r := data.Trade{Time: time, Price: inp}
		mama := testMesaAdaptiveMovingAverageCreateLength(fast, slow)

		for i := 0; i < lprimed; i++ {
			mama.Update(0.)
		}

		check(expected, mama.UpdateTrade(&r))
	})
}

func TestMesaAdaptiveMovingAverageIsPrimed(t *testing.T) {
	t.Parallel()

	input := testMesaAdaptiveMovingAverageInput()
	check := func(index int, exp, act bool) {
		t.Helper()

		if exp != act {
			t.Errorf("[%v] is incorrect: expected %v, actual %v", index, exp, act)
		}
	}

	const (
		lprimed = 26
		f       = 3
		s       = 39
	)

	t.Run("fast len = 3, slow len = 39 (mama.xls)", func(t *testing.T) {
		t.Parallel()

		mama := testMesaAdaptiveMovingAverageCreateLength(f, s)

		check(0, false, mama.IsPrimed())

		for i := 0; i < lprimed; i++ {
			mama.Update(input[i])
			check(i+1, false, mama.IsPrimed())
		}

		for i := lprimed; i < len(input); i++ {
			mama.Update(input[i])
			check(i+1, true, mama.IsPrimed())
		}
	})
}

func TestMesaAdaptiveMovingAverageMetadata(t *testing.T) {
	t.Parallel()

	check := func(what string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", what, exp, act)
		}
	}

	checkInstance := func(act indicator.Metadata, name string) {
		check("Type", indicator.MesaAdaptiveMovingAverage, act.Type)
		check("len(Outputs)", 1, len(act.Outputs))
		check("Outputs[0].Kind", int(MesaAdaptiveMovingAverageValue), act.Outputs[0].Kind)
		check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
		check("Outputs[0].Name", name, act.Outputs[0].Name)
		check("Outputs[0].Description", "Mesa adaptive moving average "+name, act.Outputs[0].Description)
	}

	t.Run("(fast, slow) limit length = (2, 40)", func(t *testing.T) {
		t.Parallel()

		const (
			f = 2
			s = 40
		)

		mama := testMesaAdaptiveMovingAverageCreateLength(f, s)
		act := mama.Metadata()
		checkInstance(act, "mama(2, 40)")
	})

	t.Run("(fast, slow) α = (0.666666666, 0.064516129)", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.666666666
			s = 0.064516129
		)

		mama := testMesaAdaptiveMovingAverageCreateAlpha(f, s)
		act := mama.Metadata()
		checkInstance(act, "mama(0.6667, 0.0645)")
	})
}

func TestNewMesaAdaptiveMovingAverage(t *testing.T) { //nolint: funlen, maintidx
	t.Parallel()

	const (
		bc  data.BarComponent   = data.BarMedianPrice
		qc  data.QuoteComponent = data.QuoteMidPrice
		tc  data.TradeComponent = data.TradePrice
		two                     = 2

		errfl = "invalid mesa adaptive moving average parameters: fast limit length should be larger than 1"
		errsl = "invalid mesa adaptive moving average parameters: slow limit length should be larger than 1"
		errfa = "invalid mesa adaptive moving average parameters: fast limit smoothing factor should be in range [0, 1]"
		errsa = "invalid mesa adaptive moving average parameters: slow limit smoothing factor should be in range [0, 1]"
		errbc = "invalid mesa adaptive moving average parameters: 9999: unknown bar component"
		errqc = "invalid mesa adaptive moving average parameters: 9999: unknown quote component"
		errtc = "invalid mesa adaptive moving average parameters: 9999: unknown trade component"
	)

	check := func(name string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", name, exp, act)
		}
	}

	checkInstance := func(mama *MesaAdaptiveMovingAverage,
		name string, lenFast int, lenSlow int, aFast float64, aSlow float64,
	) {
		if math.IsNaN(aFast) {
			aFast = two / float64(lenFast+1)
		}

		if math.IsNaN(aSlow) {
			aSlow = two / float64(lenSlow+1)
		}

		check("name", name, mama.name)
		check("description", "Mesa adaptive moving average "+name, mama.description)
		check("primed", false, mama.primed)
		check("alphaFastLimit", aFast, mama.alphaFastLimit)
		check("alphaSlowLimit", aSlow, mama.alphaSlowLimit)
		check("previousPhase", 0., mama.previousPhase)
		check("value", 0., mama.value)
		check("htce != nil", true, mama.htce != nil)
		check("isPhaseCached", false, mama.isPhaseCached)
		check("primed", false, mama.primed)
		check("barFunc == nil", false, mama.barFunc == nil)
		check("quoteFunc == nil", false, mama.quoteFunc == nil)
		check("tradeFunc == nil", false, mama.tradeFunc == nil)
	}

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		const (
			f = 3
			s = 39
		)

		mama, err := NewMesaAdaptiveMovingAverageDefault()
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(3, 39)", f, s, math.NaN(), math.NaN())
	})

	t.Run("both limit lengths > 1, default estimator", func(t *testing.T) {
		t.Parallel()

		const (
			f = 3
			s = 39
		)

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(3, 39)", f, s, math.NaN(), math.NaN())
	})

	t.Run("both limit lengths > 1, default estimator (different length)", func(t *testing.T) {
		t.Parallel()

		const (
			l = 3
			f = 2
			s = 40
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()
		cep.SmoothingLength = l

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(2, 40, hd(3, 0.200, 0.200))", f, s, math.NaN(), math.NaN())
	})

	t.Run("both limit lengths > 1, default estimator (different α quad)", func(t *testing.T) {
		t.Parallel()

		const (
			f = 2
			s = 40
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()
		cep.AlphaEmaQuadratureInPhase = 0.567

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(2, 40, hd(4, 0.567, 0.200))", f, s, math.NaN(), math.NaN())
	})

	t.Run("both limit lengths > 1, default estimator (different α period)", func(t *testing.T) {
		t.Parallel()

		const (
			f = 2
			s = 40
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()
		cep.AlphaEmaPeriod = 0.567

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(2, 40, hd(4, 0.200, 0.567))", f, s, math.NaN(), math.NaN())
	})

	t.Run("both limit lengths > 1, homodyne discriminator unrolled", func(t *testing.T) {
		t.Parallel()

		const (
			f = 2
			s = 40
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminatorUnrolled,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(2, 40, hdu(4, 0.200, 0.200))", f, s, math.NaN(), math.NaN())
	})

	t.Run("both limit lengths > 1, phase accumulator", func(t *testing.T) {
		t.Parallel()

		const (
			f = 2
			s = 40
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.PhaseAccumulator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(2, 40, pa(4, 0.200, 0.200))", f, s, math.NaN(), math.NaN())
	})

	t.Run("both limit lengths > 1, dual differentiator", func(t *testing.T) {
		t.Parallel()

		const (
			f = 2
			s = 40
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.DualDifferentiator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(2, 40, dd(4, 0.200, 0.200))", f, s, math.NaN(), math.NaN())
	})

	t.Run("fast limit length = 1, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = 1
			s = 39
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errfl, err.Error())
	})

	t.Run("fast limit length = 0, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0
			s = 39
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errfl, err.Error())
	})

	t.Run("fast limit length < 0, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = -1
			s = 39
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errfl, err.Error())
	})

	t.Run("slow limit length = 1, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = 3
			s = 1
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errsl, err.Error())
	})

	t.Run("slow limit length = 0, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = 3
			s = 0
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errsl, err.Error())
	})

	t.Run("slow limit length < 0, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = 3
			s = -1
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errsl, err.Error())
	})

	t.Run("both smoothing factors 0 ≤ α ≤ 1, default estimator", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.66666666
			s = 0.33333333
		)

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(0.6667, 0.3333)", 0, 0, f, s)
	})

	t.Run("both smoothing factors 0 ≤ α ≤ 1, default estimator (different length)", func(t *testing.T) {
		t.Parallel()

		const (
			l = 3
			f = 0.66666666
			s = 0.33333333
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()
		cep.SmoothingLength = l

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(0.6667, 0.3333, hd(3, 0.200, 0.200))", 0, 0, f, s)
	})

	t.Run("both smoothing factors 0 ≤ α ≤ 1, default estimator (different α quad)", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.66666666
			s = 0.33333333
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()
		cep.AlphaEmaQuadratureInPhase = 0.567

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(0.6667, 0.3333, hd(4, 0.567, 0.200))", 0, 0, f, s)
	})

	t.Run("both smoothing factors 0 ≤ α ≤ 1, default estimator (different α period)", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.66666666
			s = 0.33333333
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()
		cep.AlphaEmaPeriod = 0.567

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(0.6667, 0.3333, hd(4, 0.200, 0.567))", 0, 0, f, s)
	})

	t.Run("both smoothing factors 0 ≤ α ≤ 1, homodyne discriminator unrolled", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.66666666
			s = 0.33333333
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminatorUnrolled,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(0.6667, 0.3333, hdu(4, 0.200, 0.200))", 0, 0, f, s)
	})

	t.Run("both smoothing factors 0 ≤ α ≤ 1, phase accumulator", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.66666666
			s = 0.33333333
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.PhaseAccumulator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(0.6667, 0.3333, pa(4, 0.200, 0.200))", 0, 0, f, s)
	})

	t.Run("both smoothing factors 0 ≤ α ≤ 1, dual differentiator", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.66666666
			s = 0.33333333
		)

		cep := testMesaAdaptiveMovingAverageCreateCycleEstimatorParams()

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.DualDifferentiator,
			EstimatorParams: cep,
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(mama, "mama(0.6667, 0.3333, dd(4, 0.200, 0.200))", 0, 0, f, s)
	})

	t.Run("α < 0, fastest, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = -0.00000001
			s = 0.33333333
		)

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.DualDifferentiator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("mama == nil", true, mama == nil)
		check("err", errfa, err.Error())
	})

	t.Run("α > 1, fastest, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = 1.00000001
			s = 0.33333333
		)

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.DualDifferentiator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("mama == nil", true, mama == nil)
		check("err", errfa, err.Error())
	})

	t.Run("α < 0, slowest, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.66666666
			s = -0.00000001
		)

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.DualDifferentiator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("mama == nil", true, mama == nil)
		check("err", errsa, err.Error())
	})

	t.Run("α > 1, slowest, error", func(t *testing.T) {
		t.Parallel()

		const (
			f = 0.66666666
			s = 1.00000001
		)

		params := MesaAdaptiveMovingAverageSmoothingFactorParams{
			FastLimitSmoothingFactor: f, SlowLimitSmoothingFactor: s,
			EstimatorType:   hilberttransformer.DualDifferentiator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)
		check("mama == nil", true, mama == nil)
		check("err", errsa, err.Error())
	})

	t.Run("invalid bar component", func(t *testing.T) {
		t.Parallel()

		const (
			f = 3
			s = 39
		)

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    data.BarComponent(9999), QuoteComponent: qc, TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errbc, err.Error())
	})

	t.Run("invalid quote component", func(t *testing.T) {
		t.Parallel()

		const (
			f = 3
			s = 39
		)

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    bc, QuoteComponent: data.QuoteComponent(9999), TradeComponent: tc,
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errqc, err.Error())
	})

	t.Run("invalid trade component", func(t *testing.T) {
		t.Parallel()

		const (
			f = 3
			s = 39
		)

		params := MesaAdaptiveMovingAverageLengthParams{
			FastLimitLength: f, SlowLimitLength: s,
			EstimatorType:   hilberttransformer.HomodyneDiscriminator,
			EstimatorParams: testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
			BarComponent:    bc, QuoteComponent: qc, TradeComponent: data.TradeComponent(9999),
		}

		mama, err := NewMesaAdaptiveMovingAverageLength(&params)
		check("mama == nil", true, mama == nil)
		check("err", errtc, err.Error())
	})
}

func testMesaAdaptiveMovingAverageCreateCycleEstimatorParams() hilberttransformer.CycleEstimatorParams {
	const (
		smoothingLength           = 4
		alphaEmaQuadratureInPhase = 0.2
		alphaEmaPeriod            = 0.2
		warmUpPeriod              = 0
	)

	return hilberttransformer.CycleEstimatorParams{
		SmoothingLength:           smoothingLength,
		AlphaEmaQuadratureInPhase: alphaEmaQuadratureInPhase,
		AlphaEmaPeriod:            alphaEmaPeriod,
		WarmUpPeriod:              warmUpPeriod,
	}
}

func testMesaAdaptiveMovingAverageCreateLength(
	fastLimitLength int, slowLimitLength int,
) *MesaAdaptiveMovingAverage {
	return testMesaAdaptiveMovingAverageCreateLengthEx(
		fastLimitLength,
		slowLimitLength,
		hilberttransformer.HomodyneDiscriminator,
		testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
	)
}

func testMesaAdaptiveMovingAverageCreateLengthEx(
	fastLimitLength int, slowLimitLength int,
	estimatorType hilberttransformer.CycleEstimatorType,
	estimatorParams hilberttransformer.CycleEstimatorParams,
) *MesaAdaptiveMovingAverage {
	params := MesaAdaptiveMovingAverageLengthParams{
		EstimatorType: estimatorType, EstimatorParams: estimatorParams,
		FastLimitLength: fastLimitLength, SlowLimitLength: slowLimitLength,
		BarComponent: data.BarClosePrice, QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	mama, _ := NewMesaAdaptiveMovingAverageLength(&params)

	return mama
}

func testMesaAdaptiveMovingAverageCreateAlpha(
	fastLimitSmoothingFactor float64, slowLimitSmoothingFactor float64,
) *MesaAdaptiveMovingAverage {
	return testMesaAdaptiveMovingAverageCreateAlphaEx(
		fastLimitSmoothingFactor, slowLimitSmoothingFactor,
		hilberttransformer.HomodyneDiscriminator,
		testMesaAdaptiveMovingAverageCreateCycleEstimatorParams(),
	)
}

func testMesaAdaptiveMovingAverageCreateAlphaEx(
	fastLimitSmoothingFactor float64, slowLimitSmoothingFactor float64,
	estimatorType hilberttransformer.CycleEstimatorType,
	estimatorParams hilberttransformer.CycleEstimatorParams,
) *MesaAdaptiveMovingAverage {
	params := MesaAdaptiveMovingAverageSmoothingFactorParams{
		EstimatorType: estimatorType, EstimatorParams: estimatorParams,
		FastLimitSmoothingFactor: fastLimitSmoothingFactor,
		SlowLimitSmoothingFactor: slowLimitSmoothingFactor,
		BarComponent:             data.BarClosePrice, QuoteComponent: data.QuoteBidPrice, TradeComponent: data.TradePrice,
	}

	mama, _ := NewMesaAdaptiveMovingAverageSmoothingFactor(&params)

	return mama
}
