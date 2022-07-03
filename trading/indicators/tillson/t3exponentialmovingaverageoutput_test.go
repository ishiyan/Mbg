//nolint:testpackage
package tillson

import (
	"testing"
)

func TestT3ExponentialMovingAverageOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    T3ExponentialMovingAverageOutput
		text string
	}{
		{T3ExponentialMovingAverageValue, t3ExponentialMovingAverageValue},
		{t3ExponentialMovingAverageLast, t3ExponentialMovingAverageUnknown},
		{T3ExponentialMovingAverageOutput(0), t3ExponentialMovingAverageUnknown},
		{T3ExponentialMovingAverageOutput(9999), t3ExponentialMovingAverageUnknown},
		{T3ExponentialMovingAverageOutput(-9999), t3ExponentialMovingAverageUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestT3ExponentialMovingAverageOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       T3ExponentialMovingAverageOutput
		boolean bool
	}{
		{T3ExponentialMovingAverageValue, true},
		{t3ExponentialMovingAverageLast, false},
		{T3ExponentialMovingAverageOutput(0), false},
		{T3ExponentialMovingAverageOutput(9999), false},
		{T3ExponentialMovingAverageOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestT3ExponentialMovingAverageOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         T3ExponentialMovingAverageOutput
		json      string
		succeeded bool
	}{
		{T3ExponentialMovingAverageValue, dqs + t3ExponentialMovingAverageValue + dqs, true},
		{t3ExponentialMovingAverageLast, nilstr, false},
		{T3ExponentialMovingAverageOutput(9999), nilstr, false},
		{T3ExponentialMovingAverageOutput(-9999), nilstr, false},
		{T3ExponentialMovingAverageOutput(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.o.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.o, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.o)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestT3ExponentialMovingAverageOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero T3ExponentialMovingAverageOutput
	tests := []struct {
		o         T3ExponentialMovingAverageOutput
		json      string
		succeeded bool
	}{
		{T3ExponentialMovingAverageValue, dqs + t3ExponentialMovingAverageValue + dqs, true},
		{zero, dqs + t3ExponentialMovingAverageUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o T3ExponentialMovingAverageOutput

		err := o.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %v", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != o {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, o)
		}
	}
}
