//nolint:testpackage
package holidays

import (
	"testing"
)

func TestCalendarString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		c    Calendar
		text string
	}{
		{NoHolidays, noHolidays},
		{WeekendsOnly, weekendsOnly},
		{TARGET, target},
		{Euronext, euronext},
		{UnitedStates, unitedStates},
		{Switzerland, switzerland},
		{Sweden, sweden},
		{Denmark, denmark},
		{Norway, norway},
		{Iceland, iceland},
		{last, unknown},
		{Calendar(0), noHolidays},
		{Calendar(9999), unknown},
		{Calendar(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.c.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.c, exp, act)
		}
	}
}

func TestCalendarIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		c       Calendar
		boolean bool
	}{
		{NoHolidays, true},
		{WeekendsOnly, true},
		{TARGET, true},
		{Euronext, true},
		{UnitedStates, true},
		{Switzerland, true},
		{Sweden, true},
		{Denmark, true},
		{Norway, true},
		{Iceland, true},
		{last, false},
		{Calendar(0), true},
		{Calendar(9999), false},
		{Calendar(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.c.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.c, exp, act)
		}
	}
}

func TestCalendarMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		c         Calendar
		json      string
		succeeded bool
	}{
		{NoHolidays, "\"noHolidays\"", true},
		{WeekendsOnly, "\"weekendsOnly\"", true},
		{TARGET, "\"target\"", true},
		{Euronext, "\"euronext\"", true},
		{UnitedStates, "\"unitedStates\"", true},
		{Switzerland, "\"switzerland\"", true},
		{Sweden, "\"sweden\"", true},
		{Denmark, "\"denmark\"", true},
		{Norway, "\"norway\"", true},
		{Iceland, "\"iceland\"", true},
		{last, nilstr, false},
		{Calendar(9999), nilstr, false},
		{Calendar(-9999), nilstr, false},
		{Calendar(0), "\"noHolidays\"", true},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.c.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %w", tt.c, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.c)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.c, exp, act)
		}
	}
}

func TestCalendarUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zero Calendar
	tests := []struct {
		c         Calendar
		json      string
		succeeded bool
	}{
		{NoHolidays, "\"noHolidays\"", true},
		{WeekendsOnly, "\"weekendsOnly\"", true},
		{TARGET, "\"target\"", true},
		{Euronext, "\"euronext\"", true},
		{UnitedStates, "\"unitedStates\"", true},
		{Switzerland, "\"switzerland\"", true},
		{Sweden, "\"sweden\"", true},
		{Denmark, "\"denmark\"", true},
		{Norway, "\"norway\"", true},
		{Iceland, "\"iceland\"", true},
		{zero, "\"unknown\"", false},
		{zero, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.c
		bs := []byte(tt.json)

		var c Calendar

		err := c.UnmarshalJSON(bs)
		if err != nil && tt.succeeded {
			t.Errorf("UnmarshalJSON('%v'): expected success '%v', got error %w", tt.json, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("MarshalJSON('%v'): expected error, got success", tt.json)

			continue
		}

		if exp != c {
			t.Errorf("MarshalJSON('%v'): expected '%v', actual '%v'", tt.json, exp, c)
		}
	}
}
