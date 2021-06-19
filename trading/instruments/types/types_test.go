//nolint:testpackage
package types

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		it   InstrumentType
		text string
	}{
		{Undefined, undefined},
		{Stock, stock},
		{Index, index},
		{INAV, inav},
		{ETF, etf},
		{ETC, etc},
		{Forex, forex},
		{Crypto, crypto},
		{last, unknown},
		{InstrumentType(0), unknown},
		{InstrumentType(9999), unknown},
		{InstrumentType(-9999), unknown},
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
		it      InstrumentType
		boolean bool
	}{
		{Undefined, false},
		{Stock, true},
		{Index, true},
		{INAV, true},
		{ETF, true},
		{ETC, true},
		{Forex, true},
		{Crypto, true},
		{last, false},
		{InstrumentType(0), false},
		{InstrumentType(9999), false},
		{InstrumentType(-9999), false},
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
		it        InstrumentType
		json      string
		succeeded bool
	}{
		{Undefined, "\"undefined\"", true},
		{Stock, "\"stock\"", true},
		{Index, "\"index\"", true},
		{INAV, "\"inav\"", true},
		{ETF, "\"etf\"", true},
		{ETC, "\"etc\"", true},
		{Forex, "\"forex\"", true},
		{Crypto, "\"crypto\"", true},
		{last, nilstr, false},
		{InstrumentType(9999), nilstr, false},
		{InstrumentType(-9999), nilstr, false},
		{InstrumentType(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.it.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %w", tt.it, exp, err)

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

	var zero InstrumentType
	tests := []struct {
		it        InstrumentType
		json      string
		succeeded bool
	}{
		{Undefined, "\"undefined\"", true},
		{Stock, "\"stock\"", true},
		{Index, "\"index\"", true},
		{INAV, "\"inav\"", true},
		{ETF, "\"etf\"", true},
		{ETC, "\"etc\"", true},
		{Forex, "\"forex\"", true},
		{Crypto, "\"crypto\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.it
		bs := []byte(tt.json)

		var it InstrumentType

		err := it.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %w", tt.json, exp, err)

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
