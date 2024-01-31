//nolint:testpackage,dupl
package indicators

import (
	"testing"
)

func TestRateOfChangeOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    RateOfChangeOutput
		text string
	}{
		{RateOfChangeValue, rateOfChangeValue},
		{rateOfChangeLast, rateOfChangeUnknown},
		{RateOfChangeOutput(0), rateOfChangeUnknown},
		{RateOfChangeOutput(9999), rateOfChangeUnknown},
		{RateOfChangeOutput(-9999), rateOfChangeUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestRateOfChangeOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       RateOfChangeOutput
		boolean bool
	}{
		{RateOfChangeValue, true},
		{rateOfChangeLast, false},
		{RateOfChangeOutput(0), false},
		{RateOfChangeOutput(9999), false},
		{RateOfChangeOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestRateOfChangeOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         RateOfChangeOutput
		json      string
		succeeded bool
	}{
		{RateOfChangeValue, dqs + rateOfChangeValue + dqs, true},
		{rateOfChangeLast, nilstr, false},
		{RateOfChangeOutput(9999), nilstr, false},
		{RateOfChangeOutput(-9999), nilstr, false},
		{RateOfChangeOutput(0), nilstr, false},
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

func TestRateOfChangeOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero RateOfChangeOutput
	tests := []struct {
		o         RateOfChangeOutput
		json      string
		succeeded bool
	}{
		{RateOfChangeValue, dqs + rateOfChangeValue + dqs, true},
		{zero, dqs + rateOfChangeUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o RateOfChangeOutput

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
