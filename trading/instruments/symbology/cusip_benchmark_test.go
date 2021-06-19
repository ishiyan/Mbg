//nolint:testpackage
package symbology

import (
	"testing"
)

func BenchmarkValidateCUSIP(b *testing.B) {
	cusip := CUSIP("DUS0421C5")
	for i := 0; i < b.N; i++ {
		_ = cusip.Validate()
	}
}

func BenchmarkCalculateCheckDigitCUSIP(b *testing.B) {
	cusip := CUSIP("DUS0421C5")
	for i := 0; i < b.N; i++ {
		_, _ = cusip.CalculateCheckDigit()
	}
}
