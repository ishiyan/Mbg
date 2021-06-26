//nolint:testpackage
package status

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := Canceled
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := Canceled
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var os OrderStatus

	bs := []byte("\"canceled\"")
	for i := 0; i < b.N; i++ {
		_ = os.UnmarshalJSON(bs)
	}
}
