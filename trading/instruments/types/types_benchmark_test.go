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
	act := Crypto
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var it InstrumentType

	bs := []byte("\"crypto\"")
	for i := 0; i < b.N; i++ {
		_ = it.UnmarshalJSON(bs)
	}
}
