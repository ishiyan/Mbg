//nolint:testpackage
package sides

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s    Side
		text string
	}{
		{Long, long},
		{Short, short},
		{Side(0), unknown},
		{Side(9999), unknown},
		{Side(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.s.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.s, exp, act)
		}
	}
}

func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s       Side
		boolean bool
	}{
		{Long, true},
		{Short, true},
		{Side(0), false},
		{Side(9999), false},
		{Side(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.s.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.s, exp, act)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		s         Side
		json      string
		succeeded bool
	}{
		{Long, "\"long\"", true},
		{Short, "\"short\"", true},
		{Side(9999), nilstr, false},
		{Side(-9999), nilstr, false},
		{Side(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.s.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.s, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.s)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.s, exp, act)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero Side
	tests := []struct {
		s         Side
		json      string
		succeeded bool
	}{
		{Long, "\"long\"", true},
		{Short, "\"short\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.s
		bs := []byte(tt.json)

		var s Side

		err := s.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %v", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != s {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, s)
		}
	}
}
