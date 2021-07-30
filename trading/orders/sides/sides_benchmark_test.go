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

func BenchmarkIsKnown(b *testing.B) {
	act := SellShortExempt
	for i := 0; i < b.N; i++ {
		_ = act.IsKnown()
	}
}

func BenchmarkIsBuy(b *testing.B) {
	act := SellShortExempt
	for i := 0; i < b.N; i++ {
		_ = act.IsBuy()
	}
}

func BenchmarkIsSell(b *testing.B) {
	act := SellShortExempt
	for i := 0; i < b.N; i++ {
		_ = act.IsSell()
	}
}

func BenchmarkIsShort(b *testing.B) {
	act := SellShortExempt
	for i := 0; i < b.N; i++ {
		_ = act.IsShort()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := SellShortExempt
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var s Side

	bs := []byte("\"sellShortExempt\"")
	for i := 0; i < b.N; i++ {
		_ = s.UnmarshalJSON(bs)
	}
}
