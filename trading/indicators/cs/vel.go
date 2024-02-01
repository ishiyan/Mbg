package main

import (
	"fmt"
	"math"
)

type TSeries []float64

// JVEL

func JVELaux1(SrcA TSeries, Depth int) TSeries {
	Result := make(TSeries, len(SrcA))
	jrc04 := 0.0
	jrc05 := 0.0
	jrc06 := 0.0
	jrc07 := 0.0
	jrc08 := 0.0
	jrc09 := 0.0
	jrc01 := SrcA
	jrc02 := Depth
	jrc04 = float64(jrc02 + 1)
	jrc05 = jrc04 * (jrc04 + 1) / 2
	jrc06 = jrc05 * (2*jrc04 + 1) / 3
	jrc07 = jrc05*jrc05*jrc05 - jrc06*jrc06
	for Bar := jrc02; Bar < len(SrcA); Bar++ {
		jrc08 = 0.0
		jrc09 = 0.0
		for jrc10 := 0; jrc10 <= jrc02; jrc10++ {
			jrc08 += jrc01[Bar-jrc10] * (jrc04 - float64(jrc10))
			jrc09 += jrc01[Bar-jrc10] * (jrc04 - float64(jrc10)) * (jrc04 - float64(jrc10))
		}
		Result[Bar] = (jrc09*jrc05 - jrc08*jrc06) / jrc07
	}
	return Result
}

func JVELaux3(SrcA TSeries) TSeries {
	Result := make(TSeries, len(SrcA))
	JR02 := 0.0
	JR04 := 0.0
	JR05 := 0.0
	JR08 := 0.0
	JR09 := 0.0
	JR10 := 0.0
	JR12 := 0.0
	JR13 := 0.0
	JR14 := 0.0
	JR16 := 0.0
	JR17 := 0.0
	JR19 := 0.0
	JR20 := 0.0
	JR22 := 0.0
	JR23 := 0.0
	JR28 := 0.0
	JR01 := 0
	JR03 := 0
	JR06 := 0
	JR07 := 0
	JR11 := 0
	//JR15 := 0
	JR18 := 0
	JR24 := 0
	JR25 := 0
	JR26 := 0
	JR27 := 0
	//JR29 := 0
	JR21 := 0.0
	JR21a := 0.0
	JR21b := 0.0
	JR40 := make([]float64, 100)
	JR41 := make([]float64, 100)
	JR01 = 30
	JR02 = 0.0001
	for Bar := JR01; Bar < len(SrcA); Bar++ {
		JR27 = Bar
		if Bar == JR01 {
			JR28 = 0.0
			for JR29 := 1; JR29 <= JR01-1; JR29++ {
				if SrcA[Bar-JR29] == SrcA[Bar-JR29-1] {
					JR28 += 1
				}
			}
			if JR28 < float64(JR01-1) {
				JR26 = JR27 - JR01
			} else {
				JR26 = JR27
			}
			JR18 = 0
			JR25 = 0
			JR21 = SrcA[Bar-1]
			JR03 = 3
			JR04 = 0.86 - 0.55/math.Sqrt(float64(JR03))
			JR05 = 1 - math.Exp(-math.Log(4)/float64(JR03)/2)
			JR06 = JR01 + 1 //float64(JR01+1)
			JR07 = 3
			JR08 = (SrcA[Bar] - SrcA[Bar-JR07]) / float64(JR07)
			JR11 = int(math.Trunc(math.Min(1+float64(JR27-JR26), float64(JR06))))
			for JR15 := JR11 - 1; JR15 >= 1; JR15-- {
				if JR25 <= 0 {
					JR25 = 100
				}
				JR25 -= 1
				JR41[JR25] = SrcA[Bar-JR15]
			}
		}
		if JR25 <= 0 {
			JR25 = 100
		}
		JR25 -= 1
		JR41[JR25] = SrcA[Bar]
		if JR11 <= JR01 {
			if Bar == JR01 {
				JR21 = SrcA[Bar]
			} else {
				JR21 = math.Sqrt(JR05)*SrcA[Bar] + (1-math.Sqrt(JR05))*JR21a
			}
			if Bar > JR01+1 {
				JR08 = (JR21 - JR21b) / 2
			} else {
				JR08 = 0
			}
			JR11 += 1
		} else {
			if JR11 <= JR06 {
				JR12 = float64(JR11*(JR11+1)*(JR11-1)) / 12
				JR13 = float64(JR11+1) / 2
				JR14 = float64(JR11-1) / 2
				JR09 = 0.0
				JR10 = 0
				for JR15 := JR11 - 1; JR15 >= 0; JR15-- {
					JR24 = (JR25 + JR15) % 100
					JR09 += JR41[JR24]
					JR10 += JR41[JR24] * (JR14 - float64(JR15))
				}
				JR16 = JR10 / JR12
				JR17 = (JR09 / float64(JR11)) - (JR16 * JR13)
				JR19 = 0.0
				for JR15 := JR11 - 1; JR15 >= 0; JR15-- {
					JR17 += JR16
					JR24 = (JR25 + JR15) % 100
					JR19 += math.Abs(JR41[JR24] - JR17)
				}
				JR20 = (JR19 / float64(JR11)) * math.Pow(float64(JR06)/float64(JR11), 0.25)
				JR11 += 1
			} else {
				if (Bar % 1000) == 0 {
					JR09 = 0.0
					JR10 = 0
					for JR15 := JR06 - 1; JR15 >= 0; JR15-- {
						JR24 = (JR25 + JR15) % 100
						JR09 += JR41[JR24]
						JR10 += JR41[JR24] * (JR14 - float64(JR15))
					}
				} else {
					JR24 = (JR25 + JR06) % 100
					JR10 = JR10 - JR09 + JR41[JR24]*JR13 + SrcA[Bar]*JR14
					JR09 = JR09 - JR41[JR24] + SrcA[Bar]
				}
				if JR18 <= 0 {
					JR18 = JR06
				}
				JR18 -= 1
				JR19 -= JR40[JR18]
				JR16 = JR10 / JR12
				JR17 = (JR09 / float64(JR06)) + (JR16 * JR14)
				JR40[JR18] = math.Abs(SrcA[Bar] - JR17)
				JR19 = math.Max(JR02, (JR19 + JR40[JR18]))
				JR20 += ((JR19 / float64(JR06)) - JR20) * JR05
			}
			JR20 = math.Max(JR02, JR20)
			JR22 = SrcA[Bar] - (JR21 + JR08*JR04)
			JR23 = 1 - math.Exp(-math.Abs(JR22)/JR20/float64(JR03))
			JR08 = JR23*JR22 + JR08*JR04
			JR21 += JR08
		}
		JR21b = JR21a
		JR21a = JR21
		Result[Bar] = JR21
	}
	return Result
}

func JVELSeries(SrcA TSeries, ADepth int) TSeries {
	return JVELaux3(JVELaux1(SrcA, ADepth))
}

func main() {
	SrcA := testInput()

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=2)")
	Result := JVELSeries(SrcA, 2) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=3)")
	Result = JVELSeries(SrcA, 3) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=4)")
	Result = JVELSeries(SrcA, 4) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=5)")
	Result = JVELSeries(SrcA, 5) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=6)")
	Result = JVELSeries(SrcA, 6) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=7)")
	Result = JVELSeries(SrcA, 7) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=8)")
	Result = JVELSeries(SrcA, 8) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=9)")
	Result = JVELSeries(SrcA, 9) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=10)")
	Result = JVELSeries(SrcA, 10) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=11)")
	Result = JVELSeries(SrcA, 11) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=12)")
	Result = JVELSeries(SrcA, 12) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=13)")
	Result = JVELSeries(SrcA, 13) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=14)")
	Result = JVELSeries(SrcA, 14) // SrcA, ADepth
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JVELSeries(test_input, ADepth=15)")
	Result = JVELSeries(SrcA, 15) // SrcA, ADepth
	testOutput(Result)
}

func testInput() TSeries {
	return TSeries{
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

func testOutput(arr TSeries) {
	i0 := 0
	fmt.Print("[]float64 {")
	for i := 0; i < 42; i++ {
		for j := 0; j < 6; j++ {
			fmt.Printf(" %.15f,", arr[i0+j])
		}

		i0 += 6
		fmt.Println("")
	}

	fmt.Println("}")
}
