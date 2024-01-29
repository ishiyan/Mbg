//nolint:testpackage,dupl
package indicators

import (
	"testing"
)

func TestMomentumOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    MomentumOutput
		text string
	}{
		{MomentumValue, momentumValue},
		{momentumLast, momentumUnknown},
		{MomentumOutput(0), momentumUnknown},
		{MomentumOutput(9999), momentumUnknown},
		{MomentumOutput(-9999), momentumUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestMomentumOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       MomentumOutput
		boolean bool
	}{
		{MomentumValue, true},
		{momentumLast, false},
		{MomentumOutput(0), false},
		{MomentumOutput(9999), false},
		{MomentumOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestMomentumOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         MomentumOutput
		json      string
		succeeded bool
	}{
		{MomentumValue, dqs + momentumValue + dqs, true},
		{momentumLast, nilstr, false},
		{MomentumOutput(9999), nilstr, false},
		{MomentumOutput(-9999), nilstr, false},
		{MomentumOutput(0), nilstr, false},
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

func TestMomentumOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero MomentumOutput
	tests := []struct {
		o         MomentumOutput
		json      string
		succeeded bool
	}{
		{MomentumValue, dqs + momentumValue + dqs, true},
		{zero, dqs + momentumUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o MomentumOutput

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
