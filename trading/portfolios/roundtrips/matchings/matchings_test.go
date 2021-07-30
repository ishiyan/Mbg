//nolint:testpackage
package matchings

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		m    Matching
		text string
	}{
		{FirstInFirstOut, firstInFirstOut},
		{LastInFirstOut, lastInFirstOut},
		{Matching(0), unknown},
		{Matching(9999), unknown},
		{Matching(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.m.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.m, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		m       Matching
		boolean bool
	}{
		{FirstInFirstOut, true},
		{LastInFirstOut, true},
		{Matching(0), false},
		{Matching(9999), false},
		{Matching(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.m.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.m, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		m         Matching
		json      string
		succeeded bool
	}{
		{FirstInFirstOut, "\"firstInFirstOut\"", true},
		{LastInFirstOut, "\"lastInFirstOut\"", true},
		{Matching(9999), nilstr, false},
		{Matching(-9999), nilstr, false},
		{Matching(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.m.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %w", tt.m, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.m)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.m, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero Matching
	tests := []struct {
		m         Matching
		json      string
		succeeded bool
	}{
		{FirstInFirstOut, "\"firstInFirstOut\"", true},
		{LastInFirstOut, "\"lastInFirstOut\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.m
		bs := []byte(tt.json)

		var m Matching

		err := m.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %w", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != m {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, m)
		}
	}
}
