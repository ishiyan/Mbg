//nolint:testpackage
package granularities

import (
	"testing"
	"time"
)

//nolint:dupl
func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran Granularity
		text string
	}{
		{Aperiodic, aperiodic},
		{Sec1, sec1},
		{Sec3, sec3},
		{Sec5, sec5},
		{Sec10, sec10},
		{Sec15, sec15},
		{Sec30, sec30},
		{Min1, min1},
		{Min3, min3},
		{Min5, min5},
		{Min10, min10},
		{Min15, min15},
		{Min30, min30},
		{Hour1, hour1},
		{Hour3, hour3},
		{Hour6, hour6},
		{Hour12, hour12},
		{Day1, day1},
		{Week1, week1},
		{Month1, month1},
		{Month3, month3},
		{Month6, month6},
		{Year1, year1},
		{Pt1, pt1},
		{Pt10, pt10},
		{Pt100, pt100},
		{Pt1000, pt1000},
		{Vol1, vol1},
		{Vol10, vol10},
		{Vol100, vol100},
		{Vol1000, vol1000},
		{Vol10000, vol10000},
		{Vol100000, vol100000},
		{Tno1, tno1},
		{Tno10, tno10},
		{Tno100, tno100},
		{Tno1000, tno1000},
		{Tno10000, tno10000},
		{Tno100000, tno100000},
		{last, unknown},
		{Granularity(0), unknown},
		{Granularity(9999), unknown},
		{Granularity(-9999), unknown},
	}

	for _, tt := range tests {
		exp := tt.text
		act := tt.gran.String()

		if exp != act {
			t.Errorf("'%v'.String(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

//nolint:dupl
func TestIsTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran    Granularity
		boolean bool
	}{
		{Aperiodic, true},
		{Sec1, true},
		{Sec3, true},
		{Sec5, true},
		{Sec10, true},
		{Sec15, true},
		{Sec30, true},
		{Min1, true},
		{Min3, true},
		{Min5, true},
		{Min10, true},
		{Min15, true},
		{Min30, true},
		{Hour1, true},
		{Hour3, true},
		{Hour6, true},
		{Hour12, true},
		{Day1, true},
		{Week1, true},
		{Month1, true},
		{Month3, true},
		{Month6, true},
		{Year1, true},
		{Pt1, false},
		{Pt10, false},
		{Pt100, false},
		{Pt1000, false},
		{Vol1, false},
		{Vol10, false},
		{Vol100, false},
		{Vol1000, false},
		{Vol10000, false},
		{Vol100000, false},
		{Tno1, false},
		{Tno10, false},
		{Tno100, false},
		{Tno1000, false},
		{Tno10000, false},
		{Tno100000, false},
		{last, false},
		{Granularity(0), false},
		{Granularity(9999), false},
		{Granularity(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.gran.IsTime()

		if exp != act {
			t.Errorf("'%v'.IsTime(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

//nolint:dupl
func TestIsPoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran    Granularity
		boolean bool
	}{
		{Aperiodic, false},
		{Sec1, false},
		{Sec3, false},
		{Sec5, false},
		{Sec10, false},
		{Sec15, false},
		{Sec30, false},
		{Min1, false},
		{Min3, false},
		{Min5, false},
		{Min10, false},
		{Min15, false},
		{Min30, false},
		{Hour1, false},
		{Hour3, false},
		{Hour6, false},
		{Hour12, false},
		{Day1, false},
		{Week1, false},
		{Month1, false},
		{Month3, false},
		{Month6, false},
		{Year1, false},
		{Pt1, true},
		{Pt10, true},
		{Pt100, true},
		{Pt1000, true},
		{Vol1, false},
		{Vol10, false},
		{Vol100, false},
		{Vol1000, false},
		{Vol10000, false},
		{Vol100000, false},
		{Tno1, false},
		{Tno10, false},
		{Tno100, false},
		{Tno1000, false},
		{Tno10000, false},
		{Tno100000, false},
		{last, false},
		{Granularity(0), false},
		{Granularity(9999), false},
		{Granularity(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.gran.IsPoints()

		if exp != act {
			t.Errorf("'%v'.IsPoints(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

//nolint:dupl
func TestIsVolume(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran    Granularity
		boolean bool
	}{
		{Aperiodic, false},
		{Sec1, false},
		{Sec3, false},
		{Sec5, false},
		{Sec10, false},
		{Sec15, false},
		{Sec30, false},
		{Min1, false},
		{Min3, false},
		{Min5, false},
		{Min10, false},
		{Min15, false},
		{Min30, false},
		{Hour1, false},
		{Hour3, false},
		{Hour6, false},
		{Hour12, false},
		{Day1, false},
		{Week1, false},
		{Month1, false},
		{Month3, false},
		{Month6, false},
		{Year1, false},
		{Pt1, false},
		{Pt10, false},
		{Pt100, false},
		{Pt1000, false},
		{Vol1, true},
		{Vol10, true},
		{Vol100, true},
		{Vol1000, true},
		{Vol10000, true},
		{Vol100000, true},
		{Tno1, false},
		{Tno10, false},
		{Tno100, false},
		{Tno1000, false},
		{Tno10000, false},
		{Tno100000, false},
		{last, false},
		{Granularity(0), false},
		{Granularity(9999), false},
		{Granularity(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.gran.IsVolume()

		if exp != act {
			t.Errorf("'%v'.IsVolume(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

//nolint:dupl
func TestIsTurnover(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran    Granularity
		boolean bool
	}{
		{Aperiodic, false},
		{Sec1, false},
		{Sec3, false},
		{Sec5, false},
		{Sec10, false},
		{Sec15, false},
		{Sec30, false},
		{Min1, false},
		{Min3, false},
		{Min5, false},
		{Min10, false},
		{Min15, false},
		{Min30, false},
		{Hour1, false},
		{Hour3, false},
		{Hour6, false},
		{Hour12, false},
		{Day1, false},
		{Week1, false},
		{Month1, false},
		{Month3, false},
		{Month6, false},
		{Year1, false},
		{Pt1, false},
		{Pt10, false},
		{Pt100, false},
		{Pt1000, false},
		{Vol1, false},
		{Vol10, false},
		{Vol100, false},
		{Vol1000, false},
		{Vol10000, false},
		{Vol100000, false},
		{Tno1, true},
		{Tno10, true},
		{Tno100, true},
		{Tno1000, true},
		{Tno10000, true},
		{Tno100000, true},
		{last, false},
		{Granularity(0), false},
		{Granularity(9999), false},
		{Granularity(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.gran.IsTurnover()

		if exp != act {
			t.Errorf("'%v'.IsTurnover(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

//nolint:dupl
func TestUnits(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran  Granularity
		units GranularityUnits
	}{
		{Aperiodic, Time},
		{Sec1, Time},
		{Sec3, Time},
		{Sec5, Time},
		{Sec10, Time},
		{Sec15, Time},
		{Sec30, Time},
		{Min1, Time},
		{Min3, Time},
		{Min5, Time},
		{Min10, Time},
		{Min15, Time},
		{Min30, Time},
		{Hour1, Time},
		{Hour3, Time},
		{Hour6, Time},
		{Hour12, Time},
		{Day1, Time},
		{Week1, Time},
		{Month1, Time},
		{Month3, Time},
		{Month6, Time},
		{Year1, Time},
		{Pt1, Points},
		{Pt10, Points},
		{Pt100, Points},
		{Pt1000, Points},
		{Vol1, Volume},
		{Vol10, Volume},
		{Vol100, Volume},
		{Vol1000, Volume},
		{Vol10000, Volume},
		{Vol100000, Volume},
		{Tno1, Turnover},
		{Tno10, Turnover},
		{Tno100, Turnover},
		{Tno1000, Turnover},
		{Tno10000, Turnover},
		{Tno100000, Turnover},
		{last, Unknown},
		{Granularity(0), Unknown},
		{Granularity(9999), Unknown},
		{Granularity(-9999), Unknown},
	}

	for _, tt := range tests {
		exp := tt.units
		act := tt.gran.Units()

		if exp != act {
			t.Errorf("'%v'.Units(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

func TestDuration(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran     Granularity
		duration time.Duration
	}{
		{Aperiodic, time.Nanosecond},
		{Sec1, time.Second},
		{Sec3, time.Second * 3},
		{Sec5, time.Second * 5},
		{Sec10, time.Second * 10},
		{Sec15, time.Second * 15},
		{Sec30, time.Second * 30},
		{Min1, time.Minute},
		{Min3, time.Minute * 3},
		{Min5, time.Minute * 5},
		{Min10, time.Minute * 10},
		{Min15, time.Minute * 15},
		{Min30, time.Minute * 30},
		{Hour1, time.Hour},
		{Hour3, time.Hour * 3},
		{Hour6, time.Hour * 6},
		{Hour12, time.Hour * 12},
		{Day1, time.Hour * 24},
		{Week1, time.Hour * 24 * 7},
		{Month1, time.Hour * 24 * 30},
		{Month3, time.Hour * 24 * 90},
		{Month6, time.Hour * 24 * 180},
		{Year1, time.Hour * 24 * 365},
		{Pt1, time.Nanosecond},
		{Pt10, time.Nanosecond},
		{Pt100, time.Nanosecond},
		{Pt1000, time.Nanosecond},
		{Vol1, time.Nanosecond},
		{Vol10, time.Nanosecond},
		{Vol100, time.Nanosecond},
		{Vol1000, time.Nanosecond},
		{Vol10000, time.Nanosecond},
		{Vol100000, time.Nanosecond},
		{Tno1, time.Nanosecond},
		{Tno10, time.Nanosecond},
		{Tno100, time.Nanosecond},
		{Tno1000, time.Nanosecond},
		{Tno10000, time.Nanosecond},
		{Tno100000, time.Nanosecond},
		{last, time.Nanosecond},
		{Granularity(0), time.Nanosecond},
		{Granularity(9999), time.Nanosecond},
		{Granularity(-9999), time.Nanosecond},
	}

	for _, tt := range tests {
		exp := tt.duration
		act := tt.gran.Duration()

		if exp != act {
			t.Errorf("'%v'.Duration(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

func TestValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran  Granularity
		value int64
	}{
		{Aperiodic, 0},
		{Sec1, 1},
		{Sec3, 3},
		{Sec5, 5},
		{Sec10, 10},
		{Sec15, 15},
		{Sec30, 30},
		{Min1, 60},
		{Min3, 60 * 3},
		{Min5, 60 * 5},
		{Min10, 60 * 10},
		{Min15, 60 * 15},
		{Min30, 60 * 30},
		{Hour1, 3600},
		{Hour3, 3600 * 3},
		{Hour6, 3600 * 6},
		{Hour12, 3600 * 12},
		{Day1, 3600 * 24},
		{Week1, 3600 * 24 * 7},
		{Month1, 3600 * 24 * 30},
		{Month3, 3600 * 24 * 90},
		{Month6, 3600 * 24 * 180},
		{Year1, 3600 * 24 * 365},
		{Pt1, 1},
		{Pt10, 10},
		{Pt100, 100},
		{Pt1000, 1000},
		{Vol1, 1},
		{Vol10, 10},
		{Vol100, 100},
		{Vol1000, 1000},
		{Vol10000, 10000},
		{Vol100000, 100000},
		{Tno1, 1},
		{Tno10, 10},
		{Tno100, 100},
		{Tno1000, 1000},
		{Tno10000, 10000},
		{Tno100000, 100000},
		{last, 0},
		{Granularity(0), 0},
		{Granularity(9999), 0},
		{Granularity(-9999), 0},
	}

	for _, tt := range tests {
		exp := tt.value
		act := tt.gran.Value()

		if exp != act {
			t.Errorf("'%v'.Value(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

//nolint:dupl
func TestIsKnown(t *testing.T) {
	t.Parallel()

	tests := []struct {
		gran    Granularity
		boolean bool
	}{
		{Aperiodic, true},
		{Sec1, true},
		{Sec3, true},
		{Sec5, true},
		{Sec10, true},
		{Sec15, true},
		{Sec30, true},
		{Min1, true},
		{Min3, true},
		{Min5, true},
		{Min10, true},
		{Min15, true},
		{Min30, true},
		{Hour1, true},
		{Hour3, true},
		{Hour6, true},
		{Hour12, true},
		{Day1, true},
		{Week1, true},
		{Month1, true},
		{Month3, true},
		{Month6, true},
		{Year1, true},
		{Pt1, true},
		{Pt10, true},
		{Pt100, true},
		{Pt1000, true},
		{Vol1, true},
		{Vol10, true},
		{Vol100, true},
		{Vol1000, true},
		{Vol10000, true},
		{Vol100000, true},
		{Tno1, true},
		{Tno10, true},
		{Tno100, true},
		{Tno1000, true},
		{Tno10000, true},
		{Tno100000, true},
		{last, false},
		{Granularity(0), false},
		{Granularity(9999), false},
		{Granularity(-9999), false},
	}

	for _, tt := range tests {
		exp := tt.boolean
		act := tt.gran.IsKnown()

		if exp != act {
			t.Errorf("'%v'.IsKnown(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

//nolint:funlen
func TestMarshalJSON(t *testing.T) {
	t.Parallel()

	var nilstr string
	tests := []struct {
		gran      Granularity
		json      string
		succeeded bool
	}{
		{Aperiodic, "\"aperiodic\"", true},
		{Sec1, "\"sec1\"", true},
		{Sec3, "\"sec3\"", true},
		{Sec5, "\"sec5\"", true},
		{Sec10, "\"sec10\"", true},
		{Sec15, "\"sec15\"", true},
		{Sec30, "\"sec30\"", true},
		{Min1, "\"min1\"", true},
		{Min3, "\"min3\"", true},
		{Min5, "\"min5\"", true},
		{Min10, "\"min10\"", true},
		{Min15, "\"min15\"", true},
		{Min30, "\"min30\"", true},
		{Hour1, "\"hour1\"", true},
		{Hour3, "\"hour3\"", true},
		{Hour6, "\"hour6\"", true},
		{Hour12, "\"hour12\"", true},
		{Day1, "\"day1\"", true},
		{Week1, "\"week1\"", true},
		{Month1, "\"month1\"", true},
		{Month3, "\"month3\"", true},
		{Month6, "\"month6\"", true},
		{Year1, "\"year1\"", true},
		{Pt1, "\"pt1\"", true},
		{Pt10, "\"pt10\"", true},
		{Pt100, "\"pt100\"", true},
		{Pt1000, "\"pt1000\"", true},
		{Vol1, "\"vol1\"", true},
		{Vol10, "\"vol10\"", true},
		{Vol100, "\"vol100\"", true},
		{Vol1000, "\"vol1000\"", true},
		{Vol10000, "\"vol10000\"", true},
		{Vol100000, "\"vol100000\"", true},
		{Tno1, "\"tno1\"", true},
		{Tno10, "\"tno10\"", true},
		{Tno100, "\"tno100\"", true},
		{Tno1000, "\"tno1000\"", true},
		{Tno10000, "\"tno10000\"", true},
		{Tno100000, "\"tno100000\"", true},
		{last, nilstr, false},
		{Granularity(9999), nilstr, false},
		{Granularity(-9999), nilstr, false},
		{Granularity(0), nilstr, false},
	}

	for _, tt := range tests {
		exp := tt.json
		bs, err := tt.gran.MarshalJSON()

		if err != nil && tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected success '%v', got error %v", tt.gran, exp, err)

			continue
		}

		if err == nil && !tt.succeeded {
			t.Errorf("'%v'.MarshalJSON(): expected error, got success", tt.gran)

			continue
		}

		act := string(bs)
		if exp != act {
			t.Errorf("'%v'.MarshalJSON(): expected '%v', actual '%v'", tt.gran, exp, act)
		}
	}
}

//nolint:funlen
func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()

	var zerogr Granularity
	tests := []struct {
		gran      Granularity
		json      string
		succeeded bool
	}{
		{Aperiodic, "\"aperiodic\"", true},
		{Sec1, "\"sec1\"", true},
		{Sec3, "\"sec3\"", true},
		{Sec5, "\"sec5\"", true},
		{Sec10, "\"sec10\"", true},
		{Sec15, "\"sec15\"", true},
		{Sec30, "\"sec30\"", true},
		{Min1, "\"min1\"", true},
		{Min3, "\"min3\"", true},
		{Min5, "\"min5\"", true},
		{Min10, "\"min10\"", true},
		{Min15, "\"min15\"", true},
		{Min30, "\"min30\"", true},
		{Hour1, "\"hour1\"", true},
		{Hour3, "\"hour3\"", true},
		{Hour6, "\"hour6\"", true},
		{Hour12, "\"hour12\"", true},
		{Day1, "\"day1\"", true},
		{Week1, "\"week1\"", true},
		{Month1, "\"month1\"", true},
		{Month3, "\"month3\"", true},
		{Month6, "\"month6\"", true},
		{Year1, "\"year1\"", true},
		{Pt1, "\"pt1\"", true},
		{Pt10, "\"pt10\"", true},
		{Pt100, "\"pt100\"", true},
		{Pt1000, "\"pt1000\"", true},
		{Vol1, "\"vol1\"", true},
		{Vol10, "\"vol10\"", true},
		{Vol100, "\"vol100\"", true},
		{Vol1000, "\"vol1000\"", true},
		{Vol10000, "\"vol10000\"", true},
		{Vol100000, "\"vol100000\"", true},
		{Tno1, "\"tno1\"", true},
		{Tno10, "\"tno10\"", true},
		{Tno100, "\"tno100\"", true},
		{Tno1000, "\"tno1000\"", true},
		{Tno10000, "\"tno10000\"", true},
		{Tno100000, "\"tno100000\"", true},
		{zerogr, "\"unknown\"", false},
		{zerogr, "\"foobar\"", false},
	}

	for _, tt := range tests {
		exp := tt.gran
		bs := []byte(tt.json)

		var g Granularity

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
