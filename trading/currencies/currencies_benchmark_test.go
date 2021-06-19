//nolint:testpackage
package currencies

import "testing"

func BenchmarkSymbol(b *testing.B) {
	instance := Currency("FOO")
	for i := 0; i < b.N; i++ {
		_ = instance.Symbol()
	}
}

func BenchmarkDecimals(b *testing.B) {
	instance := Currency("FOO")
	for i := 0; i < b.N; i++ {
		_ = instance.Decimals()
	}
}

func BenchmarkRoundString(b *testing.B) {
	instance := Currency("FOO")
	for i := 0; i < b.N; i++ {
		_ = instance.RoundString(123.456)
	}
}
