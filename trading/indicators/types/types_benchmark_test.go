//nolint:testpackage
package types

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := last
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := GoertzelSpectrum
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var it IndicatorType

	bs := []byte("\"goertzelSpectrum\"")
	for i := 0; i < b.N; i++ {
		_ = it.UnmarshalJSON(bs)
	}
}
