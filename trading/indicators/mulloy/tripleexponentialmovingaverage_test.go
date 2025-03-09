package mulloy //nolint:testpackage

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

// Input and output data taken from Excel file describing TEMA calculations in
// Technical Analysis of Stocks & Commodities v.12:2 (72-80), Smoothing Data With Less Lag.

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
	return []float64{
		447.7532551201420, 443.0945245466050, 439.3127472726470, 435.2518723371600, 432.2710763296830, 429.1642847834070,
		427.2327936092430, 428.0821531817640, 426.5642655329860, 419.0888204711490, 397.0552029730080, 377.5576505182240,
		361.9745888044020, 348.4345573260620, 335.2088133377440, 325.2474782643930, 312.4163246000170, 303.9761331490760,
		302.3327376381150, 302.7780461091920, 303.1295873873310, 305.5798603729110, 308.4913425220430, 310.8923560402970,
		314.8048544252180, 318.7458780707020, 324.0015674119110, 329.5915729975590, 335.9319183187670, 343.6388864445580,
		350.8577411754070, 358.4417891703670, 363.1805414104080, 367.7455088517980, 373.1861445591210, 376.2355447688490,
		378.8777238094130, 382.1966643343470, 385.0500737510610, 386.0199880574620, 385.4664676938710, 385.0630954688660,
		386.7606315681500, 390.0879632471660, 393.0046051974930, 396.3884559283380, 399.8257956356140, 402.5906056605140,
		404.8971093997830, 405.2573236638770, 405.4348352082500, 405.5393293145810, 403.6985510392290, 401.6858147461310,
		399.2332589251670, 397.4583396940800, 396.8937227605320, 396.7990039945650, 396.8540479790120, 397.3812728832470,
		397.3304482735930, 396.9828044158080, 397.4303712027860, 396.5407532508720, 395.3380643625140, 392.7402919137640,
		389.1770493104540, 385.8243757851450, 384.4644853865600, 383.5369247504470, 382.8376548127300, 382.5327828867360,
		383.0831742449620, 384.2262203247070, 385.6664051571090, 387.8149761522220, 390.9220104972630, 395.2736186969200,
		398.1306093181220, 401.5101382897160, 402.8450026113750, 404.7429994134900, 406.6840088659490, 407.4996964507750,
		407.8562638034450, 409.2718370758190, 411.2223377592010, 414.2454720799570, 417.9997648337040, 421.9010844750510,
		425.7959721697150, 429.8622843243570, 434.7123735845690, 439.3876760131870, 444.5749888670540, 449.2837705328210,
		451.8281999793430, 454.0768445711660, 453.1394648960810, 453.6403988952490, 455.2482749407150, 456.5529745116560,
		458.2471043687100, 460.4974979756810, 463.3198247167110, 465.4824898200690, 468.1688721484940, 471.2640105053360,
		473.7545600416750, 474.9598499093730, 475.9386229655560, 477.6237194345840, 481.1273356342330, 480.6075335702210,
		480.7083781757940, 476.9598328026220, 473.6598427764300, 471.3830105758960, 469.4436422405790, 467.6778088950180,
		466.1869203451790, 464.6629560824800, 460.7484512177790, 457.4997769166250, 456.7993806231350, 456.8801879512060,
		453.1185493323270, 450.1238394529140, 443.5294639473960, 438.0778124441250, 434.8113851799720, 432.1002481282950,
		427.8821658851730, 426.7093894078130, 426.9666044271660, 428.3786709465310, 428.6499108049290, 428.5976817947370,
		427.6614820488370, 428.0389029359660, 426.2569774599580, 423.1826844740630, 422.8079829336740, 424.5158328614040,
		428.1526165009020, 432.4359102736830, 437.8890978161300, 442.3277477172340, 447.5163606932100, 450.2598986990470,
		453.1995417039280, 455.3257606490070, 458.7412031649510, 458.9083294879620, 456.3876814850100, 449.0020212527150,
		440.7007074664930, 430.5880606622500, 416.5574924193970, 407.4618991600910, 399.5999731084520, 391.7619551784840,
		382.6849532665060, 371.4193465810620, 362.5491004860270, 351.0858326066390, 343.5306536601080, 336.7060553904140,
		331.5697033687270, 328.5916293748200, 328.1659309303330, 327.7134822766880, 329.6661880359370, 334.1735396735250,
		337.7171067877820, 341.9525151507030, 345.2801347422690, 347.4864096473240, 348.4089353944940, 352.4715022936610,
		359.6437750494890, 370.7187792893960, 384.2174856020850, 398.2649742882960, 410.3291818830410, 422.2041735834390,
		436.0642920845510, 445.9791022193630, 453.8542925013730, 464.1409773143540, 475.5116683919040, 486.2050385165940,
		495.0002824311040, 500.8851448275950, 505.0892659755070, 508.6441565462490, 508.8881063398530, 511.1102184547510,
		515.4803199000990, 517.3620973848650, 517.9582991886640, 516.2693891414160, 512.5017825798550, 508.6287386116080,
		508.9219085413370, 509.9539303000130, 509.6271950441720, 511.8220468709190, 514.0531069701780, 516.6303502169080,
		520.4188040492950, 524.4256651970370, 525.8615874608350, 526.8482836504300, 529.6632272156770, 531.3291171208500,
		531.7291245148990, 531.6123636566610, 535.4317283632810, 535.6965931458210, 539.0122889799320, 543.1625967914890,
		543.0887282359920, 541.2445272085140, 539.5951294541200, 540.5943601920610, 542.2647031712210, 542.5049907577120,
		548.7575882556430, 559.5400748269270, 573.3828311898380, 587.3628066700000, 598.6773223018580, 607.2012867208510,
		617.2724617258430, 625.9183206616810, 631.6425297119310, 636.9967130502720, 637.6511497076610, 638.4440887489340,
		639.9847319924860, 636.9671443863560, 631.0788306206600, 624.5905612461490, 620.3711259462000, 612.6642532373130,
		606.9803858250660, 603.5504136542150, 598.1684962141560, 594.6804040431330, 592.6589779791590, 590.9007193859220,
		586.0736699387440, 578.7519999955740, 571.1621459896420, 567.8780533963640, 566.0240475868220, 565.0479875886150,
		563.2558487447470, 564.9182040537270, 564.9419232335930, 564.9033172490690, 562.9714639654380, 561.3444190672690,
		562.0493702926070, 564.6897845007230, 568.2692111791840, 568.9249261644460, 568.3754084295420, 567.7114068935060,
		569.6721409681630, 574.4162120998910, 580.1287360645910, 587.4332425342710, 597.8730958467790, 607.9008554252090,
		617.8427593724600, 628.7577708984370, 636.7827864177300, 644.6366409719390, 652.1369657690020, 660.6526607728910,
		667.7922906004490, 677.7930709491740, 687.0432881814140, 693.6107068815170, 699.9132090201550, 702.8707414283780,
		699.5586247104060, 697.9134933988580, 698.4276967655690, 700.9738971925200, 700.8421111924630, 700.2525004338070,
		697.1157865027220, 693.2759670760080, 690.1145591352280, 685.5164675913590, 682.0540392673490, 683.0929241384810,
		682.8287321788800, 686.1920670589830, 690.2580846191070, 693.9387492751950, 695.1588249035000, 695.3438938025540,
		696.4641731616050, 699.3098411306030, 701.9106649481450, 702.7730246635760, 703.5041100516600, 704.9405620736560,
		708.8231880495950, 712.0786288263990, 717.3717979207160, 722.4268577073440,
	}
}

func TestTripleExponentialMovingAverageUpdate(t *testing.T) { //nolint: funlen, cyclop
	t.Parallel()

	check := func(index int, exp, act float64) {
		t.Helper()

		if math.Abs(exp-act) > 1e-3 {
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

	t.Run("length = 14, firstIsAverage = false (Metastock)", func(t *testing.T) {
		t.Parallel()

		const (
			i39value  = 84.721  // Index=39 value.
			i40value  = 84.089  // Index=40 value.
			i251value = 108.418 // Index=251 (last) value.
		)

		tema := testTripleExponentialMovingAverageCreateLength(l, false)

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

	t.Run("length = 26, firstIsAverage = false (Metastock)", func(t *testing.T) {
		t.Parallel()

		const (
			l          = 26
			lprimed    = 3*l - 3
			firstCheck = 216
		)

		tema := testTripleExponentialMovingAverageCreateLength(l, false)

		in := testTripleExponentialMovingAverageTascInput()
		exp := testTripleExponentialMovingAverageTascExpected()
		inlen := len(in)

		for i := 0; i < lprimed; i++ {
			checkNaN(i, tema.Update(in[i]))
		}

		for i := lprimed; i < inlen; i++ {
			act := tema.Update(in[i])

			if i >= firstCheck {
				check(i, exp[i], act)
			}
		}

		checkNaN(0, tema.Update(math.NaN()))
	})
}

func TestTripleExponentialMovingAverageUpdateEntity(t *testing.T) { //nolint: funlen
	t.Parallel()

	const (
		l        = 2
		lprimed  = 3*l - 3
		alpha    = 2. / float64(l+1)
		inp      = 3.
		expFalse = 2.888888888888889
		expTrue  = 2.6666666666666665
	)

	time := testTripleExponentialMovingAverageTime()
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
		tema := testTripleExponentialMovingAverageCreateLength(l, false)

		for i := 0; i < lprimed; i++ {
			tema.Update(0.)
		}

		check(expFalse, tema.UpdateScalar(&s))
	})

	t.Run("update bar", func(t *testing.T) {
		t.Parallel()

		b := data.Bar{Time: time, Close: inp}
		tema := testTripleExponentialMovingAverageCreateLength(l, true)

		for i := 0; i < lprimed; i++ {
			tema.Update(0.)
		}

		check(expTrue, tema.UpdateBar(&b))
	})

	t.Run("update quote", func(t *testing.T) {
		t.Parallel()

		q := data.Quote{Time: time, Bid: inp}
		tema := testTripleExponentialMovingAverageCreateLength(l, false)

		for i := 0; i < lprimed; i++ {
			tema.Update(0.)
		}

		check(expFalse, tema.UpdateQuote(&q))
	})

	t.Run("update trade", func(t *testing.T) {
		t.Parallel()

		r := data.Trade{Time: time, Price: inp}
		tema := testTripleExponentialMovingAverageCreateLength(l, true)

		for i := 0; i < lprimed; i++ {
			tema.Update(0.)
		}

		check(expTrue, tema.UpdateTrade(&r))
	})
}

func TestTripleExponentialMovingAverageIsPrimed(t *testing.T) { //nolint:dupl
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

func TestTripleExponentialMovingAverageMetadata(t *testing.T) {
	t.Parallel()

	check := func(what string, exp, act any) {
		t.Helper()

		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", what, exp, act)
		}
	}

	checkInstance := func(act indicator.Metadata, name string) {
		check("Type", indicator.TripleExponentialMovingAverage, act.Type)
		check("len(Outputs)", 1, len(act.Outputs))
		check("Outputs[0].Kind", int(TripleExponentialMovingAverageValue), act.Outputs[0].Kind)
		check("Outputs[0].Type", output.Scalar, act.Outputs[0].Type)
		check("Outputs[0].Name", name, act.Outputs[0].Name)
		check("Outputs[0].Description", "Triple exponential moving average "+name, act.Outputs[0].Description)
	}

	t.Run("length = 10, firstIsAverage = true", func(t *testing.T) {
		t.Parallel()

		tema := testTripleExponentialMovingAverageCreateLength(10, true)
		act := tema.Metadata()
		checkInstance(act, "tema(10)")
	})

	t.Run("alpha = 2/11 = 0.18181818..., firstIsAverage = false", func(t *testing.T) {
		t.Parallel()

		// α = 2 / (ℓ + 1) = 2/11 = 0.18181818...
		const alpha = 2. / 11.

		tema := testTripleExponentialMovingAverageCreateAlpha(alpha, false)
		act := tema.Metadata()
		checkInstance(act, "tema(10, 0.18181818)")
	})
}

func TestNewTripleExponentialMovingAverage(t *testing.T) { //nolint: funlen
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

	checkInstance := func(
		tema *TripleExponentialMovingAverage, name string, length int, alpha float64, firstIsAverage bool,
	) {
		check("name", name, tema.name)
		check("description", "Triple exponential moving average "+name, tema.description)
		check("firstIsAverage", firstIsAverage, tema.firstIsAverage)
		check("primed", false, tema.primed)
		check("length", length, tema.length)
		check("length2", length+length-1, tema.length2)
		check("length3", length+length+length-2, tema.length3)
		check("smoothingFactor", alpha, tema.smoothingFactor)
		check("count", 0, tema.count)
		check("sum", 0., tema.sum)
		check("ema1", 0., tema.ema1)
		check("ema2", 0., tema.ema2)
		check("ema3", 0., tema.ema3)
		check("barFunc == nil", false, tema.barFunc == nil)
		check("quoteFunc == nil", false, tema.quoteFunc == nil)
		check("tradeFunc == nil", false, tema.tradeFunc == nil)
	}

	t.Run("length > 1, firstIsAverage = false", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageLengthParams{
			Length: length, FirstIsAverage: false, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageLength(&params)
		check("err == nil", true, err == nil)
		checkInstance(tema, "tema(10)", length, alpha, false)
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

	t.Run("epsilon < α ≤ 1", func(t *testing.T) {
		t.Parallel()

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: alpha, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(tema, "tema(10, 0.18181818)", length, alpha, true)
	})

	t.Run("0 < α < epsilon", func(t *testing.T) {
		t.Parallel()

		const (
			alpha  = 0.00000001
			length = 199999999 // 2./0.00000001 - 1.
		)

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: alpha, FirstIsAverage: false, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(tema, "tema(199999999, 0.00000001)", length, alpha, false)
	})

	t.Run("α = 0", func(t *testing.T) {
		t.Parallel()

		const (
			alpha  = 0.00000001
			length = 199999999 // 2./0.00000001 - 1.
		)

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: 0, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(tema, "tema(199999999, 0.00000001)", length, alpha, true)
	})

	t.Run("α = 1", func(t *testing.T) {
		t.Parallel()

		const (
			alpha  = 1
			length = 1 // 2./1 - 1.
		)

		params := TripleExponentialMovingAverageSmoothingFactorParams{
			SmoothingFactor: alpha, FirstIsAverage: true, BarComponent: bc, QuoteComponent: qc, TradeComponent: tc,
		}

		tema, err := NewTripleExponentialMovingAverageSmoothingFactor(&params)
		check("err == nil", true, err == nil)
		checkInstance(tema, "tema(1, 1.00000000)", length, alpha, true)
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
