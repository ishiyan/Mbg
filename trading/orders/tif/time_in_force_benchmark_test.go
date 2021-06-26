//nolint:testpackage
package tif

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := AtClose
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := AtClose
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var otf OrderTimeInForce

	bs := []byte("\"atClose\"")
	for i := 0; i < b.N; i++ {
		_ = otf.UnmarshalJSON(bs)
	}
}
