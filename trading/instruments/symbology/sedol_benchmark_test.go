//nolint:testpackage
package symbology

import (
	"testing"
)

func BenchmarkValidateSEDOL(b *testing.B) {
	sedol := SEDOL("B0WNLY7")
	for i := 0; i < b.N; i++ {
		_ = sedol.Validate()
	}
}

func BenchmarkCalculateCheckDigitSEDOL(b *testing.B) {
	sedol := SEDOL("B0WNLY7")
	for i := 0; i < b.N; i++ {
		_, _ = sedol.CalculateCheckDigit()
	}
}
