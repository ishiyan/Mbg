//nolint:testpackage
package kaufman

import (
	"testing"
)

func TestAdaptiveMovingAverageOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    AdaptiveMovingAverageOutput
		text string
	}{
		{AdaptiveMovingAverageValue, adaptiveMovingAverageValue},
		{adaptiveMovingAverageLast, adaptiveMovingAverageUnknown},
		{AdaptiveMovingAverageOutput(0), adaptiveMovingAverageUnknown},
		{AdaptiveMovingAverageOutput(9999), adaptiveMovingAverageUnknown},
		{AdaptiveMovingAverageOutput(-9999), adaptiveMovingAverageUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestAdaptiveMovingAverageOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       AdaptiveMovingAverageOutput
		boolean bool
	}{
		{AdaptiveMovingAverageValue, true},
		{adaptiveMovingAverageLast, false},
		{AdaptiveMovingAverageOutput(0), false},
		{AdaptiveMovingAverageOutput(9999), false},
		{AdaptiveMovingAverageOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestAdaptiveMovingAverageOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         AdaptiveMovingAverageOutput
		json      string
		succeeded bool
	}{
		{AdaptiveMovingAverageValue, dqs + adaptiveMovingAverageValue + dqs, true},
		{adaptiveMovingAverageLast, nilstr, false},
		{AdaptiveMovingAverageOutput(9999), nilstr, false},
		{AdaptiveMovingAverageOutput(-9999), nilstr, false},
		{AdaptiveMovingAverageOutput(0), nilstr, false},
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

func TestAdaptiveMovingAverageOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero AdaptiveMovingAverageOutput
	tests := []struct {
		o         AdaptiveMovingAverageOutput
		json      string
		succeeded bool
	}{
		{AdaptiveMovingAverageValue, dqs + adaptiveMovingAverageValue + dqs, true},
		{zero, dqs + adaptiveMovingAverageUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o AdaptiveMovingAverageOutput

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
