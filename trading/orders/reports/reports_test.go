//nolint:testpackage
package reports

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ort  OrderReportType
		text string
	}{
		{PendingNew, pendingNew},
		{New, new},
		{Rejected, rejected},
		{PartiallyFilled, partiallyFilled},
		{Filled, filled},
		{Expired, expired},
		{PendingReplace, pendingReplace},
		{Replaced, replaced},
		{ReplaceRejected, replaceRejected},
		{PendingCancel, pendingCancel},
		{Canceled, canceled},
		{CancelRejected, cancelRejected},
		{OrderStatus, orderStatus},
		{OrderReportType(0), unknown},
		{OrderReportType(9999), unknown},
		{OrderReportType(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.ort.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.ort, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ort     OrderReportType
		boolean bool
	}{
		{PendingNew, true},
		{New, true},
		{Rejected, true},
		{PartiallyFilled, true},
		{Filled, true},
		{Expired, true},
		{PendingReplace, true},
		{Replaced, true},
		{ReplaceRejected, true},
		{PendingCancel, true},
		{Canceled, true},
		{CancelRejected, true},
		{OrderStatus, true},
		{OrderReportType(0), false},
		{OrderReportType(9999), false},
		{OrderReportType(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.ort.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.ort, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		ort       OrderReportType
		json      string
		succeeded bool
	}{
		{PendingNew, "\"pendingNew\"", true},
		{New, "\"new\"", true},
		{Rejected, "\"rejected\"", true},
		{PartiallyFilled, "\"partiallyFilled\"", true},
		{Filled, "\"filled\"", true},
		{Expired, "\"expired\"", true},
		{PendingReplace, "\"pendingReplace\"", true},
		{Replaced, "\"replaced\"", true},
		{ReplaceRejected, "\"replaceRejected\"", true},
		{PendingCancel, "\"pendingCancel\"", true},
		{Canceled, "\"canceled\"", true},
		{CancelRejected, "\"cancelRejected\"", true},
		{OrderStatus, "\"orderStatus\"", true},
		{OrderReportType(9999), nilstr, false},
		{OrderReportType(-9999), nilstr, false},
		{OrderReportType(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.ort.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %w", tt.ort, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.ort)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.ort, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero OrderReportType
	tests := []struct {
		ort       OrderReportType
		json      string
		succeeded bool
	}{
		{PendingNew, "\"pendingNew\"", true},
		{New, "\"new\"", true},
		{Rejected, "\"rejected\"", true},
		{PartiallyFilled, "\"partiallyFilled\"", true},
		{Filled, "\"filled\"", true},
		{Expired, "\"expired\"", true},
		{PendingReplace, "\"pendingReplace\"", true},
		{Replaced, "\"replaced\"", true},
		{ReplaceRejected, "\"replaceRejected\"", true},
		{PendingCancel, "\"pendingCancel\"", true},
		{Canceled, "\"canceled\"", true},
		{CancelRejected, "\"cancelRejected\"", true},
		{OrderStatus, "\"orderStatus\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.ort
		bs := []byte(tt.json)

		var ort OrderReportType

		err := ort.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %w", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != ort {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, ort)
		}
	}
}
