//nolint:testpackage
package ehlers

import (
	"testing"
)

func TestMesaAdaptiveMovingAverageOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    MesaAdaptiveMovingAverageOutput
		text string
	}{
		{MesaAdaptiveMovingAverageValue, mesaAdaptiveMovingAverageValue},
		{mesaAdaptiveMovingAverageLast, mesaAdaptiveMovingAverageUnknown},
		{MesaAdaptiveMovingAverageOutput(0), mesaAdaptiveMovingAverageUnknown},
		{MesaAdaptiveMovingAverageOutput(9999), mesaAdaptiveMovingAverageUnknown},
		{MesaAdaptiveMovingAverageOutput(-9999), mesaAdaptiveMovingAverageUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestMesaAdaptiveMovingAverageOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       MesaAdaptiveMovingAverageOutput
		boolean bool
	}{
		{MesaAdaptiveMovingAverageValue, true},
		{mesaAdaptiveMovingAverageLast, false},
		{MesaAdaptiveMovingAverageOutput(0), false},
		{MesaAdaptiveMovingAverageOutput(9999), false},
		{MesaAdaptiveMovingAverageOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestMesaAdaptiveMovingAverageOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         MesaAdaptiveMovingAverageOutput
		json      string
		succeeded bool
	}{
		{MesaAdaptiveMovingAverageValue, dqs + mesaAdaptiveMovingAverageValue + dqs, true},
		{mesaAdaptiveMovingAverageLast, nilstr, false},
		{MesaAdaptiveMovingAverageOutput(9999), nilstr, false},
		{MesaAdaptiveMovingAverageOutput(-9999), nilstr, false},
		{MesaAdaptiveMovingAverageOutput(0), nilstr, false},
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

func TestMesaAdaptiveMovingAverageOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero MesaAdaptiveMovingAverageOutput
	tests := []struct {
		o         MesaAdaptiveMovingAverageOutput
		json      string
		succeeded bool
	}{
		{MesaAdaptiveMovingAverageValue, dqs + mesaAdaptiveMovingAverageValue + dqs, true},
		{zero, dqs + mesaAdaptiveMovingAverageUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o MesaAdaptiveMovingAverageOutput

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
