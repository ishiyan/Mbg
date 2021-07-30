//nolint:testpackage
package currencies

import "testing"

func BenchmarkConvert(b *testing.B) {
	c := createBrnchmarkInstance()

	for i := 0; i < b.N; i++ {
		_, _ = c.Convert(1, GBP, CHF)
	}
}

func BenchmarkKnownBaseCurrencies(b *testing.B) {
	c := createBrnchmarkInstance()

	for i := 0; i < b.N; i++ {
		_ = c.KnownBaseCurrencies(CHF)
	}
}

func BenchmarkKnownTermCurrencies(b *testing.B) {
	c := createBrnchmarkInstance()

	for i := 0; i < b.N; i++ {
		_ = c.KnownTermCurrencies(GBP)
	}
}

func BenchmarkUpdate(b *testing.B) {
	uc := NewUpdatableConverter()

	for i := 0; i < b.N; i++ {
		uc.Update(GBP, CHF, 1)
	}
}

func createBrnchmarkInstance() Converter {
	uc := NewUpdatableConverter()
	uc.Update(GBP, EUR, gbpEur)
	uc.Update(GBP, CHF, gbpChf)
	uc.Update(CHF, EUR, chfEur)
	uc.Update(EUR, CHF, eurChf)

	return uc
}
