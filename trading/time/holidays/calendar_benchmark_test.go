//nolint:testpackage
package holidays

import (
	"testing"
)

func BenchmarkCalendarString(b *testing.B) {
	act := last
	for i := 0; i < b.N; i++ {
		_ = act.String()
	}
}

func BenchmarkCalendarMarshalJSON(b *testing.B) {
	act := Iceland
	for i := 0; i < b.N; i++ {
		_, _ = act.MarshalJSON()
	}
}

func BenchmarkCalendarUnmarshalJSON(b *testing.B) {
	var c Calendar

	bs := []byte("\"iceland\"")
	for i := 0; i < b.N; i++ {
		_ = c.UnmarshalJSON(bs)
	}
}
