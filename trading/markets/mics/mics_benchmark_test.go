//nolint:testpackage
package mics

import "testing"

func BenchmarkOperationalMIC(b *testing.B) {
	instance := MIC("FOO")
	for i := 0; i < b.N; i++ {
		_ = instance.OperationalMIC()
	}
}

func BenchmarkTimeZoneSeconds(b *testing.B) {
	instance := MIC("FOO")
	for i := 0; i < b.N; i++ {
		_ = instance.TimeZoneSeconds()
	}
}

func BenchmarkIsPredefined(b *testing.B) {
	instance := MIC("FOO")
	for i := 0; i < b.N; i++ {
		_ = instance.IsPredefined()
	}
}
