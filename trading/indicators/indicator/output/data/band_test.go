//nolint:testpackage
package data

import (
	"math"
	"testing"
	"time"
)

var testBandTime = time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{})

func TestNewBand(t *testing.T) {
	t.Parallel()

	const (
		p1           = 1.
		p2           = 2.
		upperLarger  = "lower < upper"
		upperSmaller = "lower > upper"
	)

	check := func(condition, name string, exp, act any) {
		if exp != act {
			t.Errorf("(%s): %s is incorrect: expected %v, actual %v", condition, name, exp, act)
		}
	}

	b := NewBand(testBandTime, p1, p2)
	check(upperLarger, "Time", testBandTime, b.Time)
	check(upperLarger, "Lower", p1, b.Lower)
	check(upperLarger, "Upper", p2, b.Upper)

	b = NewBand(testBandTime, p2, p1)
	check(upperSmaller, "Time", testBandTime, b.Time)
	check(upperSmaller, "Lower", p1, b.Lower)
	check(upperSmaller, "Upper", p2, b.Upper)
}

func TestNewEmptyBand(t *testing.T) {
	t.Parallel()

	check := func(name string, exp, act any) {
		if exp != act {
			t.Errorf("%s is incorrect: expected %v, actual %v", name, exp, act)
		}
	}

	checkNaN := func(name string, act float64) {
		if !math.IsNaN(act) {
			t.Errorf("%s is incorrect: expected NaN, actual %v", name, act)
		}
	}

	b := NewEmptyBand(testBandTime)
	check("Time", testBandTime, b.Time)
	checkNaN("Lower", b.Lower)
	checkNaN("Upper", b.Upper)
}

func TestBandIsEmpty(t *testing.T) {
	t.Parallel()

	check := func(condition string, exp, act any) {
		if exp != act {
			t.Errorf("(%s): IsEmpty is incorrect: expected %v, actual %v", condition, exp, act)
		}
	}

	b := createTestBand()
	check("Lower and Upper not NaN", false, b.IsEmpty())

	b.Lower = math.NaN()
	check("Lower is NaN", true, b.IsEmpty())

	b.Upper = math.NaN()
	check("Lower and Upper are NaN", true, b.IsEmpty())

	b.Lower = 1.
	check("Upper is NaN", true, b.IsEmpty())
}

func TestBandString(t *testing.T) {
	t.Parallel()

	b := createTestBand()
	expected := "{2021-04-01 00:00:00, 1.000000, 2.000000}"

	if actual := b.String(); actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func createTestBand() Band {
	return Band{Time: testBandTime, Lower: 1., Upper: 2.}
}
