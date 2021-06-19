//nolint:testpackage
package symbology

import (
	"testing"
)

func BenchmarkValidateCheckDigitISIN(b *testing.B) {
	isin := ISIN("NGAFRINSURE4")
	for i := 0; i < b.N; i++ {
		_ = isin.ValidateCheckDigit()
	}
}

func BenchmarkCalculateCheckDigitISIN(b *testing.B) {
	isin := ISIN("NGAFRINSURE4")
	for i := 0; i < b.N; i++ {
		_, _ = isin.CalculateCheckDigit()
	}
}

func BenchmarkValidateCountryISIN(b *testing.B) {
	isin := ISIN("SZ")
	for i := 0; i < b.N; i++ {
		_ = isin.ValidateCountry()
	}
}
