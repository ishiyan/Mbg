//nolint:testpackage
package status

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
	act := KnockOutRevoked
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var is InstrumentStatus

	bs := []byte("\"knockOutRevoked\"")
	for i := 0; i < b.N; i++ {
		_ = is.UnmarshalJSON(bs)
	}
}
