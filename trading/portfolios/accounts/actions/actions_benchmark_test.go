//nolint:testpackage
package actions

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := Debit
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := Debit
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var a Action

	bs := []byte("\"debit\"")
	for i := 0; i < b.N; i++ {
		_ = a.UnmarshalJSON(bs)
	}
}
