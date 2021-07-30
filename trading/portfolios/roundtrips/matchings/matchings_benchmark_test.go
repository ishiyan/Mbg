//nolint:testpackage
package matchings

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := LastInFirstOut
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := LastInFirstOut
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var m Matching

	bs := []byte("\"lastInFirstOut\"")
	for i := 0; i < b.N; i++ {
		_ = m.UnmarshalJSON(bs)
	}
}
