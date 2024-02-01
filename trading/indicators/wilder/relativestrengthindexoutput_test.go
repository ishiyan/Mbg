//nolint:testpackage,dupl
package wilder

import (
	"testing"
)

func TestRelativeStrengthIndexOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    RelativeStrengthIndexOutput
		text string
	}{
		{RelativeStrengthIndexValue, relativeStrengthIndexValue},
		{relativeStrengthIndexLast, relativeStrengthIndexUnknown},
		{RelativeStrengthIndexOutput(0), relativeStrengthIndexUnknown},
		{RelativeStrengthIndexOutput(9999), relativeStrengthIndexUnknown},
		{RelativeStrengthIndexOutput(-9999), relativeStrengthIndexUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestRelativeStrengthIndexOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       RelativeStrengthIndexOutput
		boolean bool
	}{
		{RelativeStrengthIndexValue, true},
		{relativeStrengthIndexLast, false},
		{RelativeStrengthIndexOutput(0), false},
		{RelativeStrengthIndexOutput(9999), false},
		{RelativeStrengthIndexOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestRelativeStrengthIndexOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         RelativeStrengthIndexOutput
		json      string
		succeeded bool
	}{
		{RelativeStrengthIndexValue, dqs + relativeStrengthIndexValue + dqs, true},
		{relativeStrengthIndexLast, nilstr, false},
		{RelativeStrengthIndexOutput(9999), nilstr, false},
		{RelativeStrengthIndexOutput(-9999), nilstr, false},
		{RelativeStrengthIndexOutput(0), nilstr, false},
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

func TestRelativeStrengthIndexOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero RelativeStrengthIndexOutput
	tests := []struct {
		o         RelativeStrengthIndexOutput
		json      string
		succeeded bool
	}{
		{RelativeStrengthIndexValue, dqs + relativeStrengthIndexValue + dqs, true},
		{zero, dqs + relativeStrengthIndexUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o RelativeStrengthIndexOutput

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
