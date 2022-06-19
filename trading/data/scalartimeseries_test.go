//nolint:testpackage
package data

//nolint:gci
import (
	"testing"
	"time"
)

//nolint:funlen,dupl
func TestScalarTimeSeries(t *testing.T) {
	t.Parallel()

	verifyCurr := func(exp, act float64) {
		t.Helper()

		if act != exp {
			t.Errorf("Current(): expected %v, actual %v", exp, act)
		}
	}

	verifyHistLen := func(exp int, act []Scalar) {
		t.Helper()

		if len(act) != exp {
			t.Errorf("History(): expected length %v, actual %v", exp, act)
		}
	}

	verifyHistElem := func(idx int, tim time.Time, val float64, hist []Scalar) {
		t.Helper()

		exp := Scalar{Time: tim, Value: val}
		if hist[idx] != exp {
			t.Errorf("History()[%v]: expected %v, actual %v", idx, exp, hist[idx])
		}
	}

	verifyAt := func(tim string, exp, act float64) {
		t.Helper()

		if act != exp {
			t.Errorf("At(%v): expected %v, actual %v", tim, exp, act)
		}
	}

	s0 := Scalar{Time: time.Now(), Value: 1}
	s1 := Scalar{Time: s0.Time.AddDate(0, 0, 2), Value: 2}
	s2 := Scalar{Time: s0.Time.AddDate(0, 0, 1), Value: 3}
	s3 := Scalar{Time: s0.Time.AddDate(0, 0, 2), Value: 4}
	s4 := Scalar{Time: s0.Time.AddDate(0, 0, -1), Value: 5}
	s5 := Scalar{Time: s0.Time.AddDate(0, 0, 3), Value: 6}
	s6 := Scalar{Time: s0.Time.AddDate(0, 0, 2), Value: 7}

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}

		verifyCurr(0, sts.Current())
		verifyHistLen(0, sts.History())
		verifyAt("s0", 0, sts.At(s0.Time))
	})

	t.Run("add increasing dates", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}
		sts.Add(s0.Time, s0.Value) // 1(s0)
		sts.Add(s2.Time, s2.Value) // 1(s0), 3(s2)
		sts.Add(s1.Time, s1.Value) // 1(s0), 3(s2), 2(s1)

		// Expected: 1(s0), 3(s2), 2(s1).
		verifyCurr(2, sts.Current())

		h := sts.History()
		verifyHistLen(3, h)
		verifyHistElem(0, s0.Time, s0.Value, h)
		verifyHistElem(1, s2.Time, s2.Value, h)
		verifyHistElem(2, s1.Time, s1.Value, h)

		verifyAt("s0-1h", 0, sts.At(s0.Time.Add(-time.Hour)))
		verifyAt("s0", s0.Value, sts.At(s0.Time))
		verifyAt("s0+1h", s0.Value, sts.At(s0.Time.Add(time.Hour)))
		verifyAt("s2", s2.Value, sts.At(s2.Time))
		verifyAt("s2+1h", s2.Value, sts.At(s2.Time.Add(time.Hour)))
		verifyAt("s1", s1.Value, sts.At(s1.Time))
		verifyAt("s1+1h", s1.Value, sts.At(s1.Time.Add(time.Hour)))
	})

	t.Run("accumulate increasing dates", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}
		sts.Accumulate(s0.Time, s0.Value) // 1(s0)
		sts.Accumulate(s2.Time, s2.Value) // 1(s0), 4(s2+s0)
		sts.Accumulate(s1.Time, s1.Value) // 1(s0), 4(s2+s0), 6(s1+s2+s0)

		// Expected: 1(s0), 4(s2+s0), 6(s1+s2+s0).
		verifyCurr(s1.Value+s2.Value+s0.Value, sts.Current())

		h := sts.History()
		verifyHistLen(3, h)
		verifyHistElem(0, s0.Time, s0.Value, h)
		verifyHistElem(1, s2.Time, s2.Value+s0.Value, h)
		verifyHistElem(2, s1.Time, s1.Value+s2.Value+s0.Value, h)

		verifyAt("s0-1h", 0, sts.At(s0.Time.Add(-time.Hour)))
		verifyAt("s0", s0.Value, sts.At(s0.Time))
		verifyAt("s0+1h", s0.Value, sts.At(s0.Time.Add(time.Hour)))
		verifyAt("s2", s2.Value+s0.Value, sts.At(s2.Time))
		verifyAt("s2+1h", s2.Value+s0.Value, sts.At(s2.Time.Add(time.Hour)))
		verifyAt("s1", s1.Value+s2.Value+s0.Value, sts.At(s1.Time))
		verifyAt("s1+1h", s1.Value+s2.Value+s0.Value, sts.At(s1.Time.Add(time.Hour)))
	})

	t.Run("add equal dates", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}
		sts.Add(s0.Time, s0.Value) // 1(s0)
		sts.Add(s0.Time, s2.Value) // 3(s0)
		sts.Add(s0.Time, s1.Value) // 2(s0)

		// Expected: 2(s0).
		verifyCurr(s1.Value, sts.Current())

		h := sts.History()
		verifyHistLen(1, h)
		verifyHistElem(0, s0.Time, s1.Value, h)

		verifyAt("s0-1h", 0, sts.At(s0.Time.Add(-time.Hour)))
		verifyAt("s0", s1.Value, sts.At(s0.Time))
		verifyAt("s0+1h", s1.Value, sts.At(s0.Time.Add(time.Hour)))
	})

	t.Run("accumulate equal dates", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}
		sts.Accumulate(s0.Time, s0.Value) // 1(s0)
		sts.Accumulate(s0.Time, s2.Value) // 4(s0)
		sts.Accumulate(s0.Time, s1.Value) // 6(s0)

		// Expected: 6(s0).
		verifyCurr(s1.Value+s2.Value+s0.Value, sts.Current())

		h := sts.History()
		verifyHistLen(1, h)
		verifyHistElem(0, s0.Time, s1.Value+s2.Value+s0.Value, h)

		verifyAt("s0-1h", 0, sts.At(s0.Time.Add(-time.Hour)))
		verifyAt("s0", s1.Value+s2.Value+s0.Value, sts.At(s0.Time))
		verifyAt("s0+1h", s1.Value+s2.Value+s0.Value, sts.At(s0.Time.Add(time.Hour)))
	})

	t.Run("add decreasing dates", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}
		sts.Add(s1.Time, s1.Value) // 2(s1)
		sts.Add(s2.Time, s2.Value) // 3(s2), 2(s1)
		sts.Add(s0.Time, s0.Value) // 1(s0), 3(s2), 2(s1)

		// Expected: 1(s0), 3(s2), 2(s1).
		verifyCurr(s1.Value, sts.Current())

		h := sts.History()
		verifyHistLen(3, h)
		verifyHistElem(0, s0.Time, s0.Value, h)
		verifyHistElem(1, s2.Time, s2.Value, h)
		verifyHistElem(2, s1.Time, s1.Value, h)

		verifyAt("s0-1h", 0, sts.At(s0.Time.Add(-time.Hour)))
		verifyAt("s0", s0.Value, sts.At(s0.Time))
		verifyAt("s0+1h", s0.Value, sts.At(s0.Time.Add(time.Hour)))
		verifyAt("s2", s2.Value, sts.At(s2.Time))
		verifyAt("s2+1h", s2.Value, sts.At(s2.Time.Add(time.Hour)))
		verifyAt("s1", s1.Value, sts.At(s1.Time))
		verifyAt("s1+1h", s1.Value, sts.At(s1.Time.Add(time.Hour)))
	})

	t.Run("accumulate decreasing dates", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}
		sts.Accumulate(s1.Time, s1.Value) // 2(s1)
		sts.Accumulate(s2.Time, s2.Value) // 3(s2), 5(s1+s2)
		sts.Accumulate(s0.Time, s0.Value) // 1(s0), 4(s2+s0), 6(s1+s2+s0)

		// Expected: 1(s0), 4(s2+s0), 6(s1+s2+s0).
		verifyCurr(s1.Value+s2.Value+s0.Value, sts.Current())

		h := sts.History()
		verifyHistLen(3, h)
		verifyHistElem(0, s0.Time, s0.Value, h)
		verifyHistElem(1, s2.Time, s2.Value+s0.Value, h)
		verifyHistElem(2, s1.Time, s1.Value+s2.Value+s0.Value, h)

		verifyAt("s0-1h", 0, sts.At(s0.Time.Add(-time.Hour)))
		verifyAt("s0", s0.Value, sts.At(s0.Time))
		verifyAt("s0+1h", s0.Value, sts.At(s0.Time.Add(time.Hour)))
		verifyAt("s2", s2.Value+s0.Value, sts.At(s2.Time))
		verifyAt("s2+1h", s2.Value+s0.Value, sts.At(s2.Time.Add(time.Hour)))
		verifyAt("s1", s1.Value+s2.Value+s0.Value, sts.At(s1.Time))
		verifyAt("s1+1h", s1.Value+s2.Value+s0.Value, sts.At(s1.Time.Add(time.Hour)))
	})

	t.Run("add random dates", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}
		sts.Add(s0.Time, s0.Value) // 1(s0)
		sts.Add(s1.Time, s1.Value) // 1(s0), 2(s1)
		sts.Add(s2.Time, s2.Value) // 1(s0), 3(s2), 2(s1)
		sts.Add(s3.Time, s3.Value) // 1(s0), 3(s2), 4(s3,s1)
		sts.Add(s4.Time, s4.Value) // 5(s4), 1(s0), 3(s2), 4(s3,s1)
		sts.Add(s5.Time, s5.Value) // 5(s4), 1(s0), 3(s2), 4(s3,s1), 6(s5)
		sts.Add(s6.Time, s6.Value) // 5(s4), 1(s0), 3(s2), 7(s6,s3,s1), 6(s5)

		// Expected: 5(s4), 1(s0), 3(s2), 7(s6,s3,s1), 6(s5).
		verifyCurr(6, sts.Current())

		h := sts.History()
		verifyHistLen(5, h)
		verifyHistElem(0, s4.Time, s4.Value, h)
		verifyHistElem(1, s0.Time, s0.Value, h)
		verifyHistElem(2, s2.Time, s2.Value, h)
		verifyHistElem(3, s6.Time, s6.Value, h)
		verifyHistElem(4, s5.Time, s5.Value, h)

		verifyAt("s4-1h", 0, sts.At(s4.Time.Add(-time.Hour)))
		verifyAt("s4", s4.Value, sts.At(s4.Time))
		verifyAt("s4+1h", s4.Value, sts.At(s4.Time.Add(time.Hour)))
		verifyAt("s0+1h", s0.Value, sts.At(s0.Time.Add(time.Hour)))
		verifyAt("s2", s2.Value, sts.At(s2.Time))
		verifyAt("s6", s6.Value, sts.At(s6.Time))
		verifyAt("s5", s5.Value, sts.At(s5.Time))
		verifyAt("s5+1h", s5.Value, sts.At(s5.Time.Add(time.Hour)))
	})

	t.Run("accumulate random dates", func(t *testing.T) {
		t.Parallel()

		sts := ScalarTimeSeries{}
		sts.Accumulate(s0.Time, s0.Value) // 1(s0)
		sts.Accumulate(s1.Time, s1.Value) // 1(s0), 3(s1)
		sts.Accumulate(s2.Time, s2.Value) // 1(s0), 4(s2), 6(s1)
		sts.Accumulate(s3.Time, s3.Value) // 1(s0), 4(s2), 10(s3,s1)
		sts.Accumulate(s4.Time, s4.Value) // 5(s4), 6(s0), 9(s2), 15(s3,s1)
		sts.Accumulate(s5.Time, s5.Value) // 5(s4), 6(s0), 9(s2), 15(s3,s1), 21(s5)
		sts.Accumulate(s6.Time, s6.Value) // 5(s4), 6(s0), 9(s2), 22(s6,s3,s1), 28(s5)

		// Expected: 5(s4), 6(s0), 9(s2), 22(s6,s3,s1), 28(s5).
		verifyCurr(28, sts.Current())

		h := sts.History()
		verifyHistLen(5, h)
		verifyHistElem(0, s4.Time, 5, h)
		verifyHistElem(1, s0.Time, 6, h)
		verifyHistElem(2, s2.Time, 9, h)
		verifyHistElem(3, s6.Time, 22, h)
		verifyHistElem(4, s5.Time, 28, h)

		verifyAt("s4-1h", 0, sts.At(s4.Time.Add(-time.Hour)))
		verifyAt("s4", 5, sts.At(s4.Time))
		verifyAt("s4+1h", 5, sts.At(s4.Time.Add(time.Hour)))
		verifyAt("s0+1h", 6, sts.At(s0.Time.Add(time.Hour)))
		verifyAt("s2", 9, sts.At(s2.Time))
		verifyAt("s6", 22, sts.At(s6.Time))
		verifyAt("s5", 28, sts.At(s5.Time))
		verifyAt("s5+1h", 28, sts.At(s5.Time.Add(time.Hour)))
	})
}
