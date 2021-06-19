// Package granularities enumerates time granularities for aggregation
// of sequential temporal data points into a single statistical value.
//
// Temporal data points can be aggregated by time, by number of data points, by volume and by turnover.
//
// See the following references for the definition of the time granularity.
//
// Jerome Euzenat, Angelo Montanari. Time granularity.
// Michael Fisher, Dov Gabbay, Lluis Vila. Handbook of temporal reasoning in artiﬁcial intelligence,
// Elsevier, pp.59-118, 2005, Foundations of artiﬁcial intelligence, 0-444-51493-7 https://hal.inria.fr/hal-00922282
//
// Bettini C., Dyreson C.E., Evans W.S., Snodgrass R.T., Wang X.S.
// A Glossary of Time Granularity Concepts.
// In: Etzion O., Jajodia S., Sripada S. (eds)
// Temporal Databases: Research and Practice. Lecture Notes in Computer Science, vol 1399.
// Springer, Berlin, Heidelberg, 1998, ISBN Print ISBN 978-3-540-64519-1 https://doi.org/10.1007/BFb0053711
package granularities

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

// GranularityUnits enumerates the units of aggregation.
type GranularityUnits int

const (
	// Unknown aggregation units.
	Unknown GranularityUnits = iota
	// Time represents an aggregation by time.
	Time
	// Points represent an aggregation by a number of data points.
	Points
	// Volume represents an aggregation by volume.
	Volume
	// Turnover represents an aggregation by turnover.
	Turnover
)

// Granularity enumerates an aggregation of temporal data points into a single statistical value.
type Granularity int

const (
	// Aperiodic is an irregular time granularity.
	Aperiodic Granularity = iota + 1
	// Sec1 is a time granularity of 1 second.
	Sec1
	// Sec3 is a time granularity of 3 seconds.
	Sec3
	// Sec5 is a time granularity of 5 seconds.
	Sec5
	// Sec10 is a time granularity of 10 seconds.
	Sec10
	// Sec15 is a time granularity of 15 seconds.
	Sec15
	// Sec30 is a time granularity of 30 seconds.
	Sec30
	// Min1 is a time granularity of 1 minute.
	Min1
	// Min3 is a time granularity of 3 minutes.
	Min3
	// Min5 is a time granularity of 5 minutes.
	Min5
	// Min10 is a time granularity of 10 minutes.
	Min10
	// Min15 is a time granularity of 15 minutes.
	Min15
	// Min30 is a time granularity of 30 minutes.
	Min30
	// Hour1 is a time granularity of 1 hour.
	Hour1
	// Hour3 is a time granularity of 3 hours.
	Hour3
	// Hour6 is a time granularity of 6 hours.
	Hour6
	// Hour12 is a time granularity of 12 hours.
	Hour12
	// Day1 is a time granularity of 1 day.
	Day1
	// Week1 is a time granularity of 1 week.
	Week1
	// Month1 is a time granularity of 1 month.
	Month1
	// Month3 is a time granularity of 3 months.
	Month3
	// Month6 is a time granularity of 6 months.
	Month6
	// Year1 is a time granularity of 1 year.
	Year1
	// Pt1 is a time granularity of 1 data point.
	Pt1
	// Pt10 is a time granularity of 10 data points.
	Pt10
	// Pt100 is a time granularity of 100 data points.
	Pt100
	// Pt1000 is a time granularity of 1000 data points.
	Pt1000
	// Vol1 is a time granularity of 1 volume unit.
	Vol1
	// Vol10 is a time granularity of 10 volume units.
	Vol10
	// Vol100 is a time granularity of 100 volume units.
	Vol100
	// Vol1000 is a time granularity of 1000 volume units.
	Vol1000
	// Vol10000 is a time granularity of 10000 volume units.
	Vol10000
	// Vol100000 is a time granularity of 100000 volume units.
	Vol100000
	// Tno1 is a time granularity of 1 turnover unit.
	Tno1
	// Tno10 is a time granularity of 10 turnover units.
	Tno10
	// Tno100 is a time granularity of 100 turnover units.
	Tno100
	// Tno1000 is a time granularity of 1000 turnover units.
	Tno1000
	// Tno10000 is a time granularity of 10000 turnover units.
	Tno10000
	// Tno100000 is a time granularity of 100000 turnover units.
	Tno100000
	last
)

const (
	unknown   = "unknown"
	aperiodic = "aperiodic"
	sec1      = "sec1"
	sec3      = "sec3"
	sec5      = "sec5"
	sec10     = "sec10"
	sec15     = "sec15"
	sec30     = "sec30"
	min1      = "min1"
	min3      = "min3"
	min5      = "min5"
	min10     = "min10"
	min15     = "min15"
	min30     = "min30"
	hour1     = "hour1"
	hour3     = "hour3"
	hour6     = "hour6"
	hour12    = "hour12"
	day1      = "day1"
	week1     = "week1"
	month1    = "month1"
	month3    = "month3"
	month6    = "month6"
	year1     = "year1"
	pt1       = "pt1"
	pt10      = "pt10"
	pt100     = "pt100"
	pt1000    = "pt1000"
	vol1      = "vol1"
	vol10     = "vol10"
	vol100    = "vol100"
	vol1000   = "vol1000"
	vol10000  = "vol10000"
	vol100000 = "vol100000"
	tno1      = "tno1"
	tno10     = "tno10"
	tno100    = "tno100"
	tno1000   = "tno1000"
	tno10000  = "tno10000"
	tno100000 = "tno100000"
)

var errUnknownGranularity = errors.New("unknown granularity")

//gocyclo:ignore
//nolint:funlen,cyclop,exhaustive
// String implements the Stringer interface.
func (g Granularity) String() string {
	switch g {
	case Aperiodic:
		return aperiodic
	case Sec1:
		return sec1
	case Sec3:
		return sec3
	case Sec5:
		return sec5
	case Sec10:
		return sec10
	case Sec15:
		return sec15
	case Sec30:
		return sec30
	case Min1:
		return min1
	case Min3:
		return min3
	case Min5:
		return min5
	case Min10:
		return min10
	case Min15:
		return min15
	case Min30:
		return min30
	case Hour1:
		return hour1
	case Hour3:
		return hour3
	case Hour6:
		return hour6
	case Hour12:
		return hour12
	case Day1:
		return day1
	case Week1:
		return week1
	case Month1:
		return month1
	case Month3:
		return month3
	case Month6:
		return month6
	case Year1:
		return year1
	case Pt1:
		return pt1
	case Pt10:
		return pt10
	case Pt100:
		return pt100
	case Pt1000:
		return pt1000
	case Vol1:
		return vol1
	case Vol10:
		return vol10
	case Vol100:
		return vol100
	case Vol1000:
		return vol1000
	case Vol10000:
		return vol10000
	case Vol100000:
		return vol100000
	case Tno1:
		return tno1
	case Tno10:
		return tno10
	case Tno100:
		return tno100
	case Tno1000:
		return tno1000
	case Tno10000:
		return tno10000
	case Tno100000:
		return tno100000
	default:
		return unknown
	}
}

// IsTime determines if this granularity aggregates by time.
func (g Granularity) IsTime() bool {
	return g >= Aperiodic && g < Pt1
}

// IsPoints determines if this granularity aggregates by a number of data points.
func (g Granularity) IsPoints() bool {
	return g >= Pt1 && g < Vol1
}

// IsVolume determines if this granularity aggregates by volume.
func (g Granularity) IsVolume() bool {
	return g >= Vol1 && g < Tno1
}

// IsTurnover determines if this granularity aggregates by turnover.
func (g Granularity) IsTurnover() bool {
	return g >= Tno1 && g < last
}

// Units returns a type of the agregation units of this granularity.
func (g Granularity) Units() GranularityUnits {
	if g.IsTime() {
		return Time
	}

	if g.IsPoints() {
		return Points
	}

	if g.IsVolume() {
		return Volume
	}

	if g.IsTurnover() {
		return Turnover
	}

	return Unknown
}

//nolint:gomnd,funlen,cyclop,exhaustive
// Duration determines the time duration of this granularity in nanoseconds.
// Returns 1 ns (the lowest duration) if granularity is aperiodic or aggregates data points, volume or turnover.
func (g Granularity) Duration() time.Duration {
	switch g {
	case Sec1:
		return time.Second
	case Sec3:
		return time.Second * 3
	case Sec5:
		return time.Second * 5
	case Sec10:
		return time.Second * 10
	case Sec15:
		return time.Second * 15
	case Sec30:
		return time.Second * 30
	case Min1:
		return time.Minute
	case Min3:
		return time.Minute * 3
	case Min5:
		return time.Minute * 5
	case Min10:
		return time.Minute * 10
	case Min15:
		return time.Minute * 15
	case Min30:
		return time.Minute * 30
	case Hour1:
		return time.Hour
	case Hour3:
		return time.Hour * 3
	case Hour6:
		return time.Hour * 6
	case Hour12:
		return time.Hour * 12
	case Day1:
		return time.Hour * 24
	case Week1:
		return time.Hour * 24 * 7
	case Month1:
		return time.Hour * 24 * 30
	case Month3:
		return time.Hour * 24 * 90
	case Month6:
		return time.Hour * 24 * 180
	case Year1:
		return time.Hour * 24 * 365
	default:
		return time.Nanosecond
	}
}

//gocyclo:ignore
//nolint:gomnd,funlen,cyclop
// Value determines the value of this granularity in units.
// Returns the number of seconds for time granularities, 0 for aperiodic time granularity.
func (g Granularity) Value() int64 {
	switch g {
	case Sec1:
		return 1
	case Sec3:
		return 3
	case Sec5:
		return 5
	case Sec10:
		return 10
	case Sec15:
		return 15
	case Sec30:
		return 30
	case Min1:
		return 60
	case Min3:
		return 60 * 3
	case Min5:
		return 60 * 5
	case Min10:
		return 60 * 10
	case Min15:
		return 60 * 15
	case Min30:
		return 60 * 30
	case Hour1:
		return 3600
	case Hour3:
		return 3600 * 3
	case Hour6:
		return 3600 * 6
	case Hour12:
		return 3600 * 12
	case Day1:
		return 3600 * 24
	case Week1:
		return 3600 * 24 * 7
	case Month1:
		return 3600 * 24 * 30
	case Month3:
		return 3600 * 24 * 90
	case Month6:
		return 3600 * 24 * 180
	case Year1:
		return 3600 * 24 * 365
	case Pt1:
		return 1
	case Pt10:
		return 10
	case Pt100:
		return 100
	case Pt1000:
		return 1000
	case Vol1:
		return 1
	case Vol10:
		return 10
	case Vol100:
		return 100
	case Vol1000:
		return 1000
	case Vol10000:
		return 10000
	case Vol100000:
		return 100000
	case Tno1:
		return 1
	case Tno10:
		return 10
	case Tno100:
		return 100
	case Tno1000:
		return 1000
	case Tno10000:
		return 10000
	case Tno100000:
		return 100000
	case Aperiodic, last:
		return 0
	default:
		return 0
	}
}

// IsKnown determines if this granularity is predefined.
func (g Granularity) IsKnown() bool {
	return g >= Aperiodic && g < last
}

// MarshalJSON implements the Marshaler interface.
func (g Granularity) MarshalJSON() ([]byte, error) {
	s := g.String()
	if s == unknown {
		return nil, fmt.Errorf("cannot marshal '%s': %w", s, errUnknownGranularity)
	}

	const extra = 2 // Two bytes for quotes.

	b := make([]byte, 0, len(s)+extra)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')

	return b, nil
}

//gocyclo:ignore
//nolint:funlen,cyclop
// UnmarshalJSON implements the Unmarshaler interface.
func (g *Granularity) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)

	switch s {
	case aperiodic:
		*g = Aperiodic
	case sec1:
		*g = Sec1
	case sec3:
		*g = Sec3
	case sec5:
		*g = Sec5
	case sec10:
		*g = Sec10
	case sec15:
		*g = Sec15
	case sec30:
		*g = Sec30
	case min1:
		*g = Min1
	case min3:
		*g = Min3
	case min5:
		*g = Min5
	case min10:
		*g = Min10
	case min15:
		*g = Min15
	case min30:
		*g = Min30
	case hour1:
		*g = Hour1
	case hour3:
		*g = Hour3
	case hour6:
		*g = Hour6
	case hour12:
		*g = Hour12
	case day1:
		*g = Day1
	case week1:
		*g = Week1
	case month1:
		*g = Month1
	case month3:
		*g = Month3
	case month6:
		*g = Month6
	case year1:
		*g = Year1
	case pt1:
		*g = Pt1
	case pt10:
		*g = Pt10
	case pt100:
		*g = Pt100
	case pt1000:
		*g = Pt1000
	case vol1:
		*g = Vol1
	case vol10:
		*g = Vol10
	case vol100:
		*g = Vol100
	case vol1000:
		*g = Vol1000
	case vol10000:
		*g = Vol10000
	case vol100000:
		*g = Vol100000
	case tno1:
		*g = Tno1
	case tno10:
		*g = Tno10
	case tno100:
		*g = Tno100
	case tno1000:
		*g = Tno1000
	case tno10000:
		*g = Tno10000
	case tno100000:
		*g = Tno100000
	default:
		return fmt.Errorf("cannot unmarshal '%s': %w", s, errUnknownGranularity)
	}

	return nil
}
