//nolint:testpackage
package entities

import (
	"testing"
	"time"
)

func TestOhlcvMedian(t *testing.T) {
	t.Parallel()

	o := Ohlcv{Time: time.Time{}, Open: 0, High: 3.0, Low: 2.0, Close: 0, Volume: 0}
	expected := (o.Low + o.High) / 2

	if actual := o.Median(); actual != expected {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}

func TestOhlcvTypical(t *testing.T) {
	t.Parallel()

	o := Ohlcv{Time: time.Time{}, Open: 0, High: 4.0, Low: 2.0, Close: 3.0, Volume: 0}
	expected := (o.Low + o.High + o.Close) / 3

	if actual := o.Typical(); actual != expected {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}

func TestOhlcvWeighted(t *testing.T) {
	t.Parallel()

	o := Ohlcv{Time: time.Time{}, Open: 0, High: 4.0, Low: 2.0, Close: 3.0, Volume: 0}
	expected := (o.Low + o.High + o.Close + o.Close) / 4

	if actual := o.Weighted(); actual != expected {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}

func TestOhlcvAverage(t *testing.T) {
	t.Parallel()

	o := Ohlcv{Time: time.Time{}, Open: 3.0, High: 5.0, Low: 2.0, Close: 4.0, Volume: 0}
	expected := (o.Low + o.High + o.Open + o.Close) / 4

	if actual := o.Average(); actual != expected {
		t.Errorf("expected %f, actual %f", expected, actual)
	}
}

func TestOhlcvIsRising(t *testing.T) {
	t.Parallel()

	o := Ohlcv{Time: time.Time{}, Open: 2.0, High: 0, Low: 0, Close: 3.0, Volume: 0}
	expected := true
	actual := o.IsRising()

	if actual != expected {
		t.Errorf("rising: expected %v, actual %v", expected, actual)
	}

	o = Ohlcv{Time: time.Time{}, Open: 3.0, High: 0, Low: 0, Close: 2.0, Volume: 0}
	expected = false
	actual = o.IsRising()

	if actual != expected {
		t.Errorf("falling: expected %v, actual %v", expected, actual)
	}

	o = Ohlcv{Time: time.Time{}, Open: 0, High: 0, Low: 0, Close: 0, Volume: 0}
	actual = o.IsRising()

	if actual != expected {
		t.Errorf("flat: expected %v, actual %v", expected, actual)
	}
}

func TestOhlcvIsFalling(t *testing.T) {
	t.Parallel()

	o := Ohlcv{Time: time.Time{}, Open: 2.0, High: 0, Low: 0, Close: 3.0, Volume: 0}
	expected := false
	actual := o.IsFalling()

	if actual != expected {
		t.Errorf("rising: expected %v, actual %v", expected, actual)
	}

	o = Ohlcv{Time: time.Time{}, Open: 3.0, High: 0, Low: 0, Close: 2.0, Volume: 0}
	expected = true
	actual = o.IsFalling()

	if actual != expected {
		t.Errorf("falling: expected %v, actual %v", expected, actual)
	}

	o = Ohlcv{Time: time.Time{}, Open: 0, High: 0, Low: 0, Close: 0, Volume: 0}
	expected = false
	actual = o.IsFalling()

	if actual != expected {
		t.Errorf("flat: expected %v, actual %v", expected, actual)
	}
}

func TestOhlcvString(t *testing.T) {
	t.Parallel()

	o := Ohlcv{
		Time: time.Date(2021, time.April, 1, 0, 0, 0, 0, &time.Location{}),
		Open: 2.0, High: 3.0, Low: 4.0, Close: 5.0, Volume: 6.0,
	}
	expected := "Ohlcv(2021-04-01 00:00:00 +0000 UTC, 2.000000, 3.000000, 4.000000, 5.000000, 6.000000)"

	if actual := o.String(); actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
