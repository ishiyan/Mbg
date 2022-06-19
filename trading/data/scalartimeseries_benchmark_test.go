//nolint:testpackage
package data

import (
	"testing"
	"time"
)

func BenchmarkScalarTimeSeriesAt(b *testing.B) {
	b.StopTimer()

	sts := ScalarTimeSeries{}
	dur4 := time.Duration(time.Minute * 4)
	dur5 := time.Duration(time.Minute * 5)

	t := time.Time{}
	for i := 1; i <= b.N; i++ {
		t = t.Add(dur5)
		sts.Add(t, float64(i))
	}

	t = time.Time{}
	for i := 0; i < b.N; i++ {
		t = t.Add(dur4)

		b.StartTimer()

		_ = sts.At(t)

		b.StopTimer()
	}
}
