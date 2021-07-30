//nolint:testpackage
package groupings

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := FlatToReduced
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := FlatToReduced
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var g Grouping

	bs := []byte("\"flatToReduced\"")
	for i := 0; i < b.N; i++ {
		_ = g.UnmarshalJSON(bs)
	}
}
