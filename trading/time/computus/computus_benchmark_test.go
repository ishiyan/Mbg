//nolint:testpackage
package computus

import (
	"testing"
)

func BenchmarkEasterSunday(b *testing.B) {
	j := easterSundayPerYearFirstYear
	for i := 0; i < b.N; i++ {
		if j > easterSundayPerYearLastYear {
			j = easterSundayPerYearLastYear
		}

		_, _ = EasterSunday(j)
		j++
	}
}

func BenchmarkEasterSundayKnuth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = EasterSunday(i + easterSundayPerYearLastYear + 1)
	}
}

func BenchmarkOrthodoxEasterSunday(b *testing.B) {
	j := easterSundayPerYearFirstYear
	for i := 0; i < b.N; i++ {
		if j > easterSundayPerYearLastYear {
			j = easterSundayPerYearLastYear
		}

		_, _ = OrthodoxEasterSunday(j)
		j++
	}
}

func BenchmarkEasterSundayYearDay(b *testing.B) {
	j := easterSundayPerYearFirstYear
	for i := 0; i < b.N; i++ {
		if j > easterSundayPerYearLastYear {
			j = easterSundayPerYearLastYear
		}

		_, _ = EasterSundayYearDay(j)
		j++
	}
}

func BenchmarkEasterSundayYearDayKnuth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = EasterSundayYearDay(i + easterSundayPerYearLastYear + 1)
	}
}

func BenchmarkOrthodoxEasterSundayYearDay(b *testing.B) {
	j := easterSundayPerYearFirstYear
	for i := 0; i < b.N; i++ {
		if j > easterSundayPerYearLastYear {
			j = easterSundayPerYearLastYear
		}

		_, _ = OrthodoxEasterSundayYearDay(j)
		j++
	}
}
