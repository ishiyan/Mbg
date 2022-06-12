//nolint:testpackage
package actions

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a    Action
		text string
	}{
		{Credit, credit},
		{Debit, debit},
		{Action(0), unknown},
		{Action(9999), unknown},
		{Action(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.a.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.a, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a       Action
		boolean bool
	}{
		{Credit, true},
		{Debit, true},
		{Action(0), false},
		{Action(9999), false},
		{Action(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.a.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.a, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		a         Action
		json      string
		succeeded bool
	}{
		{Credit, "\"credit\"", true},
		{Debit, "\"debit\"", true},
		{Action(9999), nilstr, false},
		{Action(-9999), nilstr, false},
		{Action(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.a.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.a, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.a)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.a, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero Action
	tests := []struct {
		a         Action
		json      string
		succeeded bool
	}{
		{Credit, "\"credit\"", true},
		{Debit, "\"debit\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.a
		bs := []byte(tt.json)

		var a Action

		err := a.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %v", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != a {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, a)
		}
	}
}
