//nolint:testpackage
package indicator

import (
	"testing"
)

func BenchmarkTypeString(b *testing.B) {
	act := last
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkTypeMarshalJSON(b *testing.B) {
	act := GoertzelSpectrum
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkTypeUnmarshalJSON(b *testing.B) {
	var t Type

	bs := []byte("\"goertzelSpectrum\"")
	for i := 0; i < b.N; i++ {
		_ = t.UnmarshalJSON(bs)
	}
}
