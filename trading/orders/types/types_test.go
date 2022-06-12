//nolint:testpackage
package types

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ot   OrderType
		text string
	}{
		{Market, market},
		{MarketIfTouched, marketIfTouched},
		{Limit, limit},
		{Stop, stop},
		{StopLimit, stopLimit},
		{TrailingStop, trailingStop},
		{MarketOnClose, marketOnClose},
		{MarketToLimit, marketToLimit},
		{LimitIfTouched, limitIfTouched},
		{LimitOnClose, limitOnClose},
		{OrderType(0), unknown},
		{OrderType(9999), unknown},
		{OrderType(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.ot.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.ot, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ot      OrderType
		boolean bool
	}{
		{Market, true},
		{MarketIfTouched, true},
		{Limit, true},
		{Stop, true},
		{StopLimit, true},
		{TrailingStop, true},
		{MarketOnClose, true},
		{MarketToLimit, true},
		{LimitIfTouched, true},
		{LimitOnClose, true},
		{OrderType(0), false},
		{OrderType(9999), false},
		{OrderType(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.ot.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.ot, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		ot        OrderType
		json      string
		succeeded bool
	}{
		{Market, "\"market\"", true},
		{MarketIfTouched, "\"marketIfTouched\"", true},
		{Limit, "\"limit\"", true},
		{Stop, "\"stop\"", true},
		{StopLimit, "\"stopLimit\"", true},
		{TrailingStop, "\"trailingStop\"", true},
		{MarketOnClose, "\"marketOnClose\"", true},
		{MarketToLimit, "\"marketToLimit\"", true},
		{LimitIfTouched, "\"limitIfTouched\"", true},
		{LimitOnClose, "\"limitOnClose\"", true},
		{OrderType(9999), nilstr, false},
		{OrderType(-9999), nilstr, false},
		{OrderType(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.ot.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.ot, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.ot)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.ot, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero OrderType
	tests := []struct {
		ot        OrderType
		json      string
		succeeded bool
	}{
		{Market, "\"market\"", true},
		{MarketIfTouched, "\"marketIfTouched\"", true},
		{Limit, "\"limit\"", true},
		{Stop, "\"stop\"", true},
		{StopLimit, "\"stopLimit\"", true},
		{TrailingStop, "\"trailingStop\"", true},
		{MarketOnClose, "\"marketOnClose\"", true},
		{MarketToLimit, "\"marketToLimit\"", true},
		{LimitIfTouched, "\"limitIfTouched\"", true},
		{LimitOnClose, "\"limitOnClose\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.ot
		bs := []byte(tt.json)

		var ot OrderType

		err := ot.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %v", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != ot {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, ot)
		}
	}
}
