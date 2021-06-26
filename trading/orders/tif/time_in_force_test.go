//nolint:testpackage
package tif

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		otf  OrderTimeInForce
		text string
	}{
		{Day, day},
		{ImmediateOrCancel, immediateOrCancel},
		{FillOrKill, fillOrKill},
		{GoodTillCanceled, goodTillCanceled},
		{GoodTillDate, goodTillDate},
		{AtOpen, atOpen},
		{AtClose, atClose},
		{OrderTimeInForce(0), unknown},
		{OrderTimeInForce(9999), unknown},
		{OrderTimeInForce(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.otf.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.otf, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		otf     OrderTimeInForce
		boolean bool
	}{
		{Day, true},
		{ImmediateOrCancel, true},
		{FillOrKill, true},
		{GoodTillCanceled, true},
		{GoodTillDate, true},
		{AtOpen, true},
		{AtClose, true},
		{OrderTimeInForce(0), false},
		{OrderTimeInForce(9999), false},
		{OrderTimeInForce(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.otf.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.otf, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		otf       OrderTimeInForce
		json      string
		succeeded bool
	}{
		{Day, "\"day\"", true},
		{ImmediateOrCancel, "\"immediateOrCancel\"", true},
		{FillOrKill, "\"fillOrKill\"", true},
		{GoodTillCanceled, "\"goodTillCanceled\"", true},
		{GoodTillDate, "\"goodTillDate\"", true},
		{AtOpen, "\"atOpen\"", true},
		{AtClose, "\"atClose\"", true},
		{OrderTimeInForce(9999), nilstr, false},
		{OrderTimeInForce(-9999), nilstr, false},
		{OrderTimeInForce(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.otf.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %w", tt.otf, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.otf)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.otf, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero OrderTimeInForce
	tests := []struct {
		otf       OrderTimeInForce
		json      string
		succeeded bool
	}{
		{Day, "\"day\"", true},
		{ImmediateOrCancel, "\"immediateOrCancel\"", true},
		{FillOrKill, "\"fillOrKill\"", true},
		{GoodTillCanceled, "\"goodTillCanceled\"", true},
		{GoodTillDate, "\"goodTillDate\"", true},
		{AtOpen, "\"atOpen\"", true},
		{AtClose, "\"atClose\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.otf
		bs := []byte(tt.json)

		var otf OrderTimeInForce

		err := otf.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %w", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != otf {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, otf)
		}
	}
}
