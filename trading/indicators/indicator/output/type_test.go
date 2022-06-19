//nolint:testpackage
package output

import (
	"testing"
)

func TestTypeString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		t    Type
		text string
	}{
		{Scalar, scalar},
		{Band, band},
		{Heatmap, heatmap},
		{last, unknown},
		{Type(0), unknown},
		{Type(9999), unknown},
		{Type(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.t.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.t, exp, act)
		}
	}
}

func TestTypeIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		t       Type
		boolean bool
	}{
		{Scalar, true},
		{Band, true},
		{Heatmap, true},
		{last, false},
		{Type(0), false},
		{Type(9999), false},
		{Type(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.t.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.t, exp, act)
		}
	}
}

func TestTypeMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		t         Type
		json      string
		succeeded bool
	}{
		{Scalar, "\"scalar\"", true},
		{Band, "\"band\"", true},
		{Heatmap, "\"heatmap\"", true},
		{last, nilstr, false},
		{Type(9999), nilstr, false},
		{Type(-9999), nilstr, false},
		{Type(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.t.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.t, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.t)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.t, exp, act)
		}
	}
}

func TestTypeUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero Type
	tests := []struct {
		t         Type
		json      string
		succeeded bool
	}{
		{Scalar, "\"scalar\"", true},
		{Band, "\"band\"", true},
		{Heatmap, "\"heatmap\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.t
		bs := []byte(tt.json)

		var act Type

		err := act.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %v", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != act {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, act)
		}
	}
}
