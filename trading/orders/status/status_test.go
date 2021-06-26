//nolint:testpackage
package status

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		os   OrderStatus
		text string
	}{
		{Accepted, accepted},
		{PendingNew, pendingNew},
		{New, new},
		{Rejected, rejected},
		{PartiallyFilled, partiallyFilled},
		{Filled, filled},
		{Expired, expired},
		{PendingReplace, pendingReplace},
		{PendingCancel, pendingCancel},
		{Canceled, canceled},
		{OrderStatus(0), unknown},
		{OrderStatus(9999), unknown},
		{OrderStatus(-9999), unknown},
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
		os      OrderStatus
		boolean bool
	}{
		{Accepted, true},
		{PendingNew, true},
		{New, true},
		{Rejected, true},
		{PartiallyFilled, true},
		{Filled, true},
		{Expired, true},
		{PendingReplace, true},
		{PendingCancel, true},
		{Canceled, true},
		{OrderStatus(0), false},
		{OrderStatus(9999), false},
		{OrderStatus(-9999), false},
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
		os        OrderStatus
		json      string
		succeeded bool
	}{
		{Accepted, "\"accepted\"", true},
		{PendingNew, "\"pendingNew\"", true},
		{New, "\"new\"", true},
		{Rejected, "\"rejected\"", true},
		{PartiallyFilled, "\"partiallyFilled\"", true},
		{Filled, "\"filled\"", true},
		{Expired, "\"expired\"", true},
		{PendingReplace, "\"pendingReplace\"", true},
		{PendingCancel, "\"pendingCancel\"", true},
		{Canceled, "\"canceled\"", true},
		{OrderStatus(9999), nilstr, false},
		{OrderStatus(-9999), nilstr, false},
		{OrderStatus(0), nilstr, false},
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

	var zero OrderStatus
	tests := []struct {
		os        OrderStatus
		json      string
		succeeded bool
	}{
		{Accepted, "\"accepted\"", true},
		{PendingNew, "\"pendingNew\"", true},
		{New, "\"new\"", true},
		{Rejected, "\"rejected\"", true},
		{PartiallyFilled, "\"partiallyFilled\"", true},
		{Filled, "\"filled\"", true},
		{Expired, "\"expired\"", true},
		{PendingReplace, "\"pendingReplace\"", true},
		{PendingCancel, "\"pendingCancel\"", true},
		{Canceled, "\"canceled\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.os
		bs := []byte(tt.json)

		var os OrderStatus

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
