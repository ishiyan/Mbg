//nolint:testpackage
package types

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		it   IndicatorType
		text string
	}{
		{SimpleMovingAverage, simpleMovingAverage},
		{ExponentialMovingAverage, exponentialMovingAverage},
		{BollingerBands, bollingerBands},
		{Variance, variance},
		{StandardDeviation, standardDeviation},
		{GoertzelSpectrum, goertzelSpectrum},
		{last, unknown},
		{IndicatorType(0), unknown},
		{IndicatorType(9999), unknown},
		{IndicatorType(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.it.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.it, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		it      IndicatorType
		boolean bool
	}{
		{SimpleMovingAverage, true},
		{ExponentialMovingAverage, true},
		{BollingerBands, true},
		{Variance, true},
		{StandardDeviation, true},
		{GoertzelSpectrum, true},
		{last, false},
		{IndicatorType(0), false},
		{IndicatorType(9999), false},
		{IndicatorType(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.it.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.it, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		it        IndicatorType
		json      string
		succeeded bool
	}{
		{SimpleMovingAverage, "\"simpleMovingAverage\"", true},
		{ExponentialMovingAverage, "\"exponentialMovingAverage\"", true},
		{BollingerBands, "\"bollingerBands\"", true},
		{Variance, "\"variance\"", true},
		{StandardDeviation, "\"standardDeviation\"", true},
		{GoertzelSpectrum, "\"goertzelSpectrum\"", true},
		{last, nilstr, false},
		{IndicatorType(9999), nilstr, false},
		{IndicatorType(-9999), nilstr, false},
		{IndicatorType(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.it.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.it, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.it)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.it, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero IndicatorType
	tests := []struct {
		it        IndicatorType
		json      string
		succeeded bool
	}{
		{SimpleMovingAverage, "\"simpleMovingAverage\"", true},
		{ExponentialMovingAverage, "\"exponentialMovingAverage\"", true},
		{BollingerBands, "\"bollingerBands\"", true},
		{Variance, "\"variance\"", true},
		{StandardDeviation, "\"standardDeviation\"", true},
		{GoertzelSpectrum, "\"goertzelSpectrum\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.it
		bs := []byte(tt.json)

		var it IndicatorType

		err := it.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %v", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != it {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, it)
		}
	}
}
