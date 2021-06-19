// Package holidays enumerates holiday schedules for a specific exchange or a country.
package holidays

import (
	"bytes"
	"errors"
	"fmt"
)

// Calendar enumerates holiday schedules for a specific exchange or a country.
type Calendar int

const (
	// NoHolidays is a physical days (no holidays) schedule.
	NoHolidays Calendar = iota

	// WeekendsOnly is a weekends-only holiday schedule.
	WeekendsOnly

	// TARGET is the 'Trans-european Automated Real-time Gross settlement Express Transfer' holiday schedule.
	TARGET

	// Euronext is the Euronext exchange holiday schedule.
	Euronext

	// UnitedStates is a generic US exchange holiday schedule.
	UnitedStates

	// Switzerland is the generic Swiss exchange holiday schedule.
	Switzerland

	// Sweden is the generic Swedish exchange holiday schedule.
	Sweden

	// Denmark is the generic Danish exchange holiday schedule.
	Denmark

	// Norway is the generic Norwegian exchange holiday schedule.
	Norway

	// Iceland is the generic Icelandic exchange holiday schedule.
	Iceland
	last
)

const (
	unknown      = "unknown"
	noHolidays   = "noHolidays"
	weekendsOnly = "weekendsOnly"
	target       = "target"
	euronext     = "euronext"
	unitedStates = "unitedStates"
	switzerland  = "switzerland"
	sweden       = "sweden"
	denmark      = "denmark"
	norway       = "norway"
	iceland      = "iceland"
)

var errUnknownCalendar = errors.New("unknown holiday calendar")

//nolint:exhaustive,cyclop
// String implements the Stringer interface.
func (c Calendar) String() string {
	switch c {
	case NoHolidays:
		return noHolidays
	case WeekendsOnly:
		return weekendsOnly
	case TARGET:
		return target
	case Euronext:
		return euronext
	case UnitedStates:
		return unitedStates
	case Switzerland:
		return switzerland
	case Sweden:
		return sweden
	case Denmark:
		return denmark
	case Norway:
		return norway
	case Iceland:
		return iceland
	default:
		return unknown
	}
}

// IsKnown determines if this holiday calendar is known.
func (c Calendar) IsKnown() bool {
	return c >= NoHolidays && c < last
}

// MarshalJSON implements the Marshaler interface.
func (c Calendar) MarshalJSON() ([]byte, error) {
	s := c.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownCalendar)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

//nolint:cyclop
// UnmarshalJSON implements the Unmarshaler interface.
func (c *Calendar) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case noHolidays:
		*c = NoHolidays
	case weekendsOnly:
		*c = WeekendsOnly
	case target:
		*c = TARGET
	case euronext:
		*c = Euronext
	case unitedStates:
		*c = UnitedStates
	case switzerland:
		*c = Switzerland
	case sweden:
		*c = Sweden
	case denmark:
		*c = Denmark
	case norway:
		*c = Norway
	case iceland:
		*c = Iceland
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownCalendar)
	}

	return nil
}
