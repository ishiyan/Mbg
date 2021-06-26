//nolint:testpackage
package types

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := LimitOnClose
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := LimitOnClose
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var ot OrderType

	bs := []byte("\"limitOnClose\"")
	for i := 0; i < b.N; i++ {
		_ = ot.UnmarshalJSON(bs)
	}
}
