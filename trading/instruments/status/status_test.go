//nolint:testpackage
package status

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		is   InstrumentStatus
		text string
	}{
		{Active, active},
		{ActiveClosingOrdersOnly, activeClosingOrdersOnly},
		{Inactive, inactive},
		{Suspended, suspended},
		{PendingExpiry, pendingExpiry},
		{Expired, expired},
		{PendingDeletion, pendingDeletion},
		{Delisted, delisted},
		{KnockedOut, knockedOut},
		{KnockOutRevoked, knockOutRevoked},
		{last, unknown},
		{InstrumentStatus(0), unknown},
		{InstrumentStatus(9999), unknown},
		{InstrumentStatus(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.is.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.is, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		is      InstrumentStatus
		boolean bool
	}{
		{Active, true},
		{ActiveClosingOrdersOnly, true},
		{Inactive, true},
		{Suspended, true},
		{PendingExpiry, true},
		{Expired, true},
		{PendingDeletion, true},
		{Delisted, true},
		{KnockedOut, true},
		{KnockOutRevoked, true},
		{last, false},
		{InstrumentStatus(0), false},
		{InstrumentStatus(9999), false},
		{InstrumentStatus(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.is.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.is, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		is        InstrumentStatus
		json      string
		succeeded bool
	}{
		{Active, "\"active\"", true},
		{ActiveClosingOrdersOnly, "\"activeClosingOrdersOnly\"", true},
		{Inactive, "\"inactive\"", true},
		{Suspended, "\"suspended\"", true},
		{PendingExpiry, "\"pendingExpiry\"", true},
		{Expired, "\"expired\"", true},
		{PendingDeletion, "\"pendingDeletion\"", true},
		{Delisted, "\"delisted\"", true},
		{KnockedOut, "\"knockedOut\"", true},
		{KnockOutRevoked, "\"knockOutRevoked\"", true},
		{last, nilstr, false},
		{InstrumentStatus(9999), nilstr, false},
		{InstrumentStatus(-9999), nilstr, false},
		{InstrumentStatus(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.is.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.is, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.is)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.is, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero InstrumentStatus
	tests := []struct {
		is        InstrumentStatus
		json      string
		succeeded bool
	}{
		{Active, "\"active\"", true},
		{ActiveClosingOrdersOnly, "\"activeClosingOrdersOnly\"", true},
		{Inactive, "\"inactive\"", true},
		{Suspended, "\"suspended\"", true},
		{PendingExpiry, "\"pendingExpiry\"", true},
		{Expired, "\"expired\"", true},
		{PendingDeletion, "\"pendingDeletion\"", true},
		{Delisted, "\"delisted\"", true},
		{KnockedOut, "\"knockedOut\"", true},
		{KnockOutRevoked, "\"knockOutRevoked\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.is
		bs := []byte(tt.json)

		var is InstrumentStatus

		err := is.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %v", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != is {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, is)
		}
	}
}
