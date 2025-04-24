//nolint:testpackage
package ehlers

import (
	"testing"
)

func TestFractalAdaptiveMovingAverageOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    FractalAdaptiveMovingAverageOutput
		text string
	}{
		{FractalAdaptiveMovingAverageValue, fractalAdaptiveMovingAverageValue},
		{FractalAdaptiveMovingAverageValueFdim, fractalAdaptiveMovingAverageValueFdim},
		{fractalAdaptiveMovingAverageLast, fractalAdaptiveMovingAverageUnknown},
		{FractalAdaptiveMovingAverageOutput(0), fractalAdaptiveMovingAverageUnknown},
		{FractalAdaptiveMovingAverageOutput(9999), fractalAdaptiveMovingAverageUnknown},
		{FractalAdaptiveMovingAverageOutput(-9999), fractalAdaptiveMovingAverageUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestFractalAdaptiveMovingAverageOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       FractalAdaptiveMovingAverageOutput
		boolean bool
	}{
		{FractalAdaptiveMovingAverageValue, true},
		{FractalAdaptiveMovingAverageValueFdim, true},
		{fractalAdaptiveMovingAverageLast, false},
		{FractalAdaptiveMovingAverageOutput(0), false},
		{FractalAdaptiveMovingAverageOutput(9999), false},
		{FractalAdaptiveMovingAverageOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestFractalAdaptiveMovingAverageOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         FractalAdaptiveMovingAverageOutput
		json      string
		succeeded bool
	}{
		{FractalAdaptiveMovingAverageValue, dqs + fractalAdaptiveMovingAverageValue + dqs, true},
		{FractalAdaptiveMovingAverageValueFdim, dqs + fractalAdaptiveMovingAverageValueFdim + dqs, true},
		{fractalAdaptiveMovingAverageLast, nilstr, false},
		{FractalAdaptiveMovingAverageOutput(9999), nilstr, false},
		{FractalAdaptiveMovingAverageOutput(-9999), nilstr, false},
		{FractalAdaptiveMovingAverageOutput(0), nilstr, false},
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

func TestFractalAdaptiveMovingAverageOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero FractalAdaptiveMovingAverageOutput
	tests := []struct {
		o         FractalAdaptiveMovingAverageOutput
		json      string
		succeeded bool
	}{
		{FractalAdaptiveMovingAverageValue, dqs + fractalAdaptiveMovingAverageValue + dqs, true},
		{FractalAdaptiveMovingAverageValueFdim, dqs + fractalAdaptiveMovingAverageValueFdim + dqs, true},
		{zero, dqs + fractalAdaptiveMovingAverageUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o FractalAdaptiveMovingAverageOutput

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
