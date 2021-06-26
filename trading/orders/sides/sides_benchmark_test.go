//nolint:testpackage
package sides

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := SellShortExempt
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := SellShortExempt
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var os OrderSide

	bs := []byte("\"sellShortExempt\"")
	for i := 0; i < b.N; i++ {
		_ = os.UnmarshalJSON(bs)
	}
}
