//nolint:testpackage
package groupings

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		g    Grouping
		text string
	}{
		{FillToFill, fillToFill},
		{FlatToFlat, flatToFlat},
		{FlatToReduced, flatToReduced},
		{Grouping(0), unknown},
		{Grouping(9999), unknown},
		{Grouping(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.g.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.g, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		g       Grouping
		boolean bool
	}{
		{FillToFill, true},
		{FlatToFlat, true},
		{FlatToReduced, true},
		{Grouping(0), false},
		{Grouping(9999), false},
		{Grouping(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.g.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.g, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		g         Grouping
		json      string
		succeeded bool
	}{
		{FillToFill, "\"fillToFill\"", true},
		{FlatToFlat, "\"flatToFlat\"", true},
		{FlatToReduced, "\"flatToReduced\"", true},
		{Grouping(9999), nilstr, false},
		{Grouping(-9999), nilstr, false},
		{Grouping(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.g.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.g, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.g)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.g, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero Grouping
	tests := []struct {
		g         Grouping
		json      string
		succeeded bool
	}{
		{FillToFill, "\"fillToFill\"", true},
		{FlatToFlat, "\"flatToFlat\"", true},
		{FlatToReduced, "\"flatToReduced\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.g
		bs := []byte(tt.json)

		var g Grouping

		err := g.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %v", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != g {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, g)
		}
	}
}
