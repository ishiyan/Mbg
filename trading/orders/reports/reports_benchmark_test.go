//nolint:testpackage
package reports

import (
	"testing"
)

func BenchmarkString(b *testing.B) {
	act := CancelRejected
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := CancelRejected
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	var ort OrderReportType

	bs := []byte("\"cancelRejected\"")
	for i := 0; i < b.N; i++ {
		_ = ort.UnmarshalJSON(bs)
	}
}
