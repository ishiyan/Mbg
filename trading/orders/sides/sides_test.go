//nolint:testpackage
package sides

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		os   OrderSide
		text string
	}{
		{Buy, buy},
		{Sell, sell},
		{BuyMinus, buyMinus},
		{SellPlus, sellPlus},
		{SellShort, sellShort},
		{SellShortExempt, sellShortExempt},
		{OrderSide(0), unknown},
		{OrderSide(9999), unknown},
		{OrderSide(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.os.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.os, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		os      OrderSide
		boolean bool
	}{
		{Buy, true},
		{Sell, true},
		{BuyMinus, true},
		{SellPlus, true},
		{SellShort, true},
		{SellShortExempt, true},
		{OrderSide(0), false},
		{OrderSide(9999), false},
		{OrderSide(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.os.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.os, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		os        OrderSide
		json      string
		succeeded bool
	}{
		{Buy, "\"buy\"", true},
		{Sell, "\"sell\"", true},
		{BuyMinus, "\"buyMinus\"", true},
		{SellPlus, "\"sellPlus\"", true},
		{SellShort, "\"sellShort\"", true},
		{SellShortExempt, "\"sellShortExempt\"", true},
		{OrderSide(9999), nilstr, false},
		{OrderSide(-9999), nilstr, false},
		{OrderSide(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.os.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %w", tt.os, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.os)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.os, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero OrderSide
	tests := []struct {
		os        OrderSide
		json      string
		succeeded bool
	}{
		{Buy, "\"buy\"", true},
		{Sell, "\"sell\"", true},
		{BuyMinus, "\"buyMinus\"", true},
		{SellPlus, "\"sellPlus\"", true},
		{SellShort, "\"sellShort\"", true},
		{SellShortExempt, "\"sellShortExempt\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.os
		bs := []byte(tt.json)

		var os OrderSide

		err := os.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %w", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != os {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, os)
		}
	}
}
