//nolint:testpackage,dupl
package indicators

import (
	"testing"
)

func TestRateOfChangePercentOutputString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o    RateOfChangePercentOutput
		text string
	}{
		{RateOfChangePercentValue, rateOfChangePercentValue},
		{rateOfChangePercentLast, rateOfChangePercentUnknown},
		{RateOfChangePercentOutput(0), rateOfChangePercentUnknown},
		{RateOfChangePercentOutput(9999), rateOfChangePercentUnknown},
		{RateOfChangePercentOutput(-9999), rateOfChangePercentUnknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.o.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestRateOfChangePercentOutputIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		o       RateOfChangePercentOutput
		boolean bool
	}{
		{RateOfChangePercentValue, true},
		{rateOfChangePercentLast, false},
		{RateOfChangePercentOutput(0), false},
		{RateOfChangePercentOutput(9999), false},
		{RateOfChangePercentOutput(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.o.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.o, exp, act)
		}
	}
}

func TestRateOfChangePercentOutputMarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var nilstr string
	tests := []struct {
		o         RateOfChangePercentOutput
		json      string
		succeeded bool
	}{
		{RateOfChangePercentValue, dqs + rateOfChangePercentValue + dqs, true},
		{rateOfChangePercentLast, nilstr, false},
		{RateOfChangePercentOutput(9999), nilstr, false},
		{RateOfChangePercentOutput(-9999), nilstr, false},
		{RateOfChangePercentOutput(0), nilstr, false},
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

func TestRateOfChangePercentOutputUnmarshalJSON(t *testing.T) {
	t.Parallel()

	const dqs = "\""

	var zero RateOfChangePercentOutput
	tests := []struct {
		o         RateOfChangePercentOutput
		json      string
		succeeded bool
	}{
		{RateOfChangePercentValue, dqs + rateOfChangePercentValue + dqs, true},
		{zero, dqs + rateOfChangePercentUnknown + dqs, false},
		{zero, dqs + "foobar" + dqs, false},
	}

	for _, tt := range tests {
		exp := tt.o
		bs := []byte(tt.json)

		var o RateOfChangePercentOutput

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
