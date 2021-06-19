//nolint:testpackage
package granularities

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := last
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkUnits(b *testing.B) {
	act := last
	for i := 0; i < b.N; i++ {
		_ = act.Units()
	}
}

func BenchmarkDuration(b *testing.B) {
	act := last
	for i := 0; i < b.N; i++ {
		_ = act.Duration()
	}
}

func BenchmarkValue(b *testing.B) {
	act := last
	for i := 0; i < b.N; i++ {
		_ = act.Value()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := Tno100000
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var g Granularity

	bs := []byte("\"tno100000\"")
	for i := 0; i < b.N; i++ {
		_ = g.UnmarshalJSON(bs)
	}
}
