package main

import (
	"fmt"
	"math"
)

type TSeries []float64

// JRSX

func JRSXSeries(SrcA TSeries, LenA int) TSeries {
	var Result TSeries
	Result = make(TSeries, len(SrcA))
	f0 := 0
	f88 := 0
	f90 := 0
	f8 := 0.0
	f10 := 0.0
	f18 := 0.0
	f20 := 0.0
	f28 := 0.0
	f30 := 0.0
	f38 := 0.0
	f40 := 0.0
	f48 := 0.0
	f50 := 0.0
	f58 := 0.0
	f60 := 0.0
	f68 := 0.0
	f70 := 0.0
	f78 := 0.0
	f80 := 0.0
	v4 := 0.0
	v8 := 0.0
	vC := 0.0
	v10 := 0.0
	v14 := 0.0
	v18 := 0.0
	v1C := 0.0
	v20 := 0.0

	for Bar := 0; Bar < len(SrcA); Bar++ {
		if f90 == 0 {
			f90 = 1
			f0 = 0
			if LenA-1 >= 5 {
				f88 = LenA - 1
			} else {
				f88 = 5
			}
			f8 = 100 * SrcA[Bar]
			f18 = 3.0 / float64(LenA+2)
			f20 = 1 - f18
		} else {
			if f88 <= f90 {
				f90 = f88 + 1
			} else {
				f90 = f90 + 1
			}
			f10 = f8
			f8 = 100 * SrcA[Bar]
			v8 = f8 - f10
			f28 = f20*f28 + f18*v8
			f30 = f18*f28 + f20*f30
			vC = f28*1.5 - f30*0.5
			f38 = f20*f38 + f18*vC
			f40 = f18*f38 + f20*f40
			v10 = f38*1.5 - f40*0.5
			f48 = f20*f48 + f18*v10
			f50 = f18*f48 + f20*f50
			v14 = f48*1.5 - f50*0.5
			f58 = f20*f58 + f18*math.Abs(v8)
			f60 = f18*f58 + f20*f60
			v18 = f58*1.5 - f60*0.5
			f68 = f20*f68 + f18*v18
			f70 = f18*f68 + f20*f70
			v1C = f68*1.5 - f70*0.5
			f78 = f20*f78 + f18*v1C
			f80 = f18*f78 + f20*f80
			v20 = f78*1.5 - f80*0.5
			if f88 >= f90 && f8 != f10 {
				f0 = 1
			}
			if f88 == f90 && f0 == 0 {
				f90 = 0
			}
		}
		if f88 < f90 && v20 > 1.0e-10 {
			v4 = (v14/v20 + 1) * 50
			if v4 > 100 {
				v4 = 100
			}
			if v4 < 0 {
				v4 = 0
			}
		} else {
			v4 = 50
		}
		Value := v4
		Result[Bar] = Value
	}
	return Result
}

func main() {
	SrcA := testInput()

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=2)")
	Result := JRSXSeries(SrcA, 2) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=3)")
	Result = JRSXSeries(SrcA, 3) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=4)")
	Result = JRSXSeries(SrcA, 4) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=5)")
	Result = JRSXSeries(SrcA, 5) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=6)")
	Result = JRSXSeries(SrcA, 6) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=7)")
	Result = JRSXSeries(SrcA, 7) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=8)")
	Result = JRSXSeries(SrcA, 8) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=9)")
	Result = JRSXSeries(SrcA, 9) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=10)")
	Result = JRSXSeries(SrcA, 10) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=11)")
	Result = JRSXSeries(SrcA, 11) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=12)")
	Result = JRSXSeries(SrcA, 12) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=13)")
	Result = JRSXSeries(SrcA, 13) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=14)")
	Result = JRSXSeries(SrcA, 14) // SrcA, ALength
	testOutput(Result)

	fmt.Println("//////////////////////////////////////////////////")
	fmt.Println("// JRSXSeries(test_input, ALength=15)")
	Result = JRSXSeries(SrcA, 15) // SrcA, ALength
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
