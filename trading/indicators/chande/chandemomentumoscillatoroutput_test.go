//nolint:testpackage,dupl
package chande

import (
	"testing"
)

func TestChandeMomentumOscillatorOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    ChandeMomentumOscillatorOutput
		text string
	}{
		{ChandeMomentumOscillatorValue, chandeMomentumOscillatorValue},
		{chandeMomentumOscillatorLast, chandeMomentumOscillatorUnknown},
		{ChandeMomentumOscillatorOutput(0), chandeMomentumOscillatorUnknown},
		{ChandeMomentumOscillatorOutput(9999), chandeMomentumOscillatorUnknown},
		{ChandeMomentumOscillatorOutput(-9999), chandeMomentumOscillatorUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestChandeMomentumOscillatorOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       ChandeMomentumOscillatorOutput
		boolean bool
	}{
		{ChandeMomentumOscillatorValue, true},
		{chandeMomentumOscillatorLast, false},
		{ChandeMomentumOscillatorOutput(0), false},
		{ChandeMomentumOscillatorOutput(9999), false},
		{ChandeMomentumOscillatorOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestChandeMomentumOscillatorOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         ChandeMomentumOscillatorOutput
		json      string
		succeeded bool
	}{
		{ChandeMomentumOscillatorValue, dqs + chandeMomentumOscillatorValue + dqs, true},
		{chandeMomentumOscillatorLast, nilstr, false},
		{ChandeMomentumOscillatorOutput(9999), nilstr, false},
		{ChandeMomentumOscillatorOutput(-9999), nilstr, false},
		{ChandeMomentumOscillatorOutput(0), nilstr, false},
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

func TestChandeMomentumOscillatorOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero ChandeMomentumOscillatorOutput
	tests := []struct {
		o         ChandeMomentumOscillatorOutput
		json      string
		succeeded bool
	}{
		{ChandeMomentumOscillatorValue, dqs + chandeMomentumOscillatorValue + dqs, true},
		{zero, dqs + chandeMomentumOscillatorUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o ChandeMomentumOscillatorOutput

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
