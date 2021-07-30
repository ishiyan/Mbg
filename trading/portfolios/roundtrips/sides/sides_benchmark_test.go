//nolint:testpackage
package sides

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := Short
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := Short
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var s Side

	bs := []byte("\"short\"")
	for i := 0; i < b.N; i++ {
		_ = s.UnmarshalJSON(bs)
	}
}
