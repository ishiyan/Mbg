//nolint:testpackage
package portfolios

//nolint:gci
import (
	"mbg/trading/data/entities"
	"testing"
	"time"
)

func TestScalarHistoryEmpty(t *testing.T) {
	t.Parallel()

	sh := scalarHistory{}

	curr := sh.Current()
	if curr != 0 {
		t.Errorf("Current(): expected %v, actual %v", 0., curr)
	}

	hist := sh.History()
	if len(hist) != 0 {
		t.Errorf("History(): expected length %v, actual %v", 0, len(hist))
	}
}

//nolint:funlen
func TestScalarHistoryAdd(t *testing.T) {
	t.Parallel()

	s0 := entities.Scalar{Time: time.Now(), Value: 1.0}
	s1 := entities.Scalar{Time: s0.Time.AddDate(0, 0, -1), Value: 1.1}
	s2 := entities.Scalar{Time: s0.Time.AddDate(0, 0, -2), Value: 1.2}

	t.Run("increasing dates", func(t *testing.T) {
		t.Parallel()
		sh := scalarHistory{}
		sh.add(s2.Time, s2.Value)
		sh.add(s1.Time, s1.Value)
		sh.add(s0.Time, s0.Value)

		curr := sh.Current()
		if curr != s0.Value {
			t.Errorf("current(): expected %v, actual %v", s0, curr)
		}

		hist := sh.History()
		if len(hist) != 3 {
			t.Errorf("History(): expected length %v, actual %v", 3, len(hist))
		}

		if hist[0] != s2 {
			t.Errorf("History()[0]: expected %v, actual %v", s2, hist[0])
		}

		if hist[1] != s1 {
			t.Errorf("History()[1]: expected %v, actual %v", s1, hist[1])
		}

		if hist[2] != s0 {
			t.Errorf("History()[2]: expected %v, actual %v", s0, hist[2])
		}
	})

	t.Run("decreasing dates", func(t *testing.T) {
		t.Parallel()
		sh := scalarHistory{}
		sh.add(s0.Time, s0.Value)
		sh.add(s1.Time, s1.Value)
		sh.add(s2.Time, s2.Value)

		exp := entities.Scalar{Time: s0.Time, Value: s2.Value}
		curr := sh.Current()
		if curr != exp.Value {
			t.Errorf("Current(): expected %v, actual %v", exp.Value, curr)
		}

		hist := sh.History()
		if len(hist) != 1 {
			t.Errorf("copyHistory(): expected length %v, actual %v", 1, len(hist))
		}

		if hist[0] != exp {
			t.Errorf("History()[0]: expected %v, actual %v", exp, hist[0])
		}
	})

	t.Run("equal dates", func(t *testing.T) {
		t.Parallel()
		sh := scalarHistory{}
		sh.add(s0.Time, s0.Value)
		sh.add(s0.Time, s1.Value)
		sh.add(s0.Time, s2.Value)

		exp := entities.Scalar{Time: s0.Time, Value: s2.Value}
		curr := sh.Current()
		if curr != exp.Value {
			t.Errorf("Current(): expected %v, actual %v", exp.Value, curr)
		}

		hist := sh.History()
		if len(hist) != 1 {
			t.Errorf("History(): expected length %v, actual %v", 1, len(hist))
		}

		if hist[0] != exp {
			t.Errorf("History()[0]: expected %v, actual %v", exp, hist[0])
		}
	})
}

//nolint:funlen
func TestScalarHistoryAccumulate(t *testing.T) {
	t.Parallel()

	s0 := entities.Scalar{Time: time.Now(), Value: 1.0}
	s1 := entities.Scalar{Time: s0.Time.AddDate(0, 0, -1), Value: 1.1}
	s2 := entities.Scalar{Time: s0.Time.AddDate(0, 0, -2), Value: 1.2}

	t.Run("increasing dates", func(t *testing.T) {
		t.Parallel()
		sh := scalarHistory{}
		sh.accumulate(s2.Time, s2.Value)
		sh.accumulate(s1.Time, s1.Value)
		sh.accumulate(s0.Time, s0.Value)

		exp := entities.Scalar{Time: s0.Time, Value: s2.Value + s1.Value + s0.Value}
		curr := sh.Current()
		if curr != exp.Value {
			t.Errorf("Current(): expected %v, actual %v", exp.Value, curr)
		}

		hist := sh.History()
		if len(hist) != 3 {
			t.Errorf("History(): expected length %v, actual %v", 3, len(hist))
		}

		if hist[0] != s2 {
			t.Errorf("History()[0]: expected %v, actual %v", s2, hist[0])
		}

		exp1 := entities.Scalar{Time: s1.Time, Value: s2.Value + s1.Value}
		if hist[1] != exp1 {
			t.Errorf("History()[1]: expected %v, actual %v", exp1, hist[1])
		}

		if hist[2] != exp {
			t.Errorf("History()[2]: expected %v, actual %v", exp, hist[2])
		}
	})

	t.Run("decreasing dates", func(t *testing.T) {
		t.Parallel()
		sh := scalarHistory{}
		sh.accumulate(s0.Time, s0.Value)
		sh.accumulate(s1.Time, s1.Value)
		sh.accumulate(s2.Time, s2.Value)

		exp := entities.Scalar{Time: s0.Time, Value: s2.Value + s1.Value + s0.Value}
		curr := sh.Current()
		if curr != exp.Value {
			t.Errorf("Current(): expected %v, actual %v", exp.Value, curr)
		}

		hist := sh.History()
		if len(hist) != 1 {
			t.Errorf("History(): expected length %v, actual %v", 1, len(hist))
		}

		if hist[0] != exp {
			t.Errorf("History()[0]: expected %v, actual %v", exp, hist[0])
		}
	})

	t.Run("equal dates", func(t *testing.T) {
		t.Parallel()
		sh := scalarHistory{}
		sh.accumulate(s0.Time, s0.Value)
		sh.accumulate(s0.Time, s1.Value)
		sh.accumulate(s0.Time, s2.Value)

		exp := entities.Scalar{Time: s0.Time, Value: s2.Value + s1.Value + s0.Value}
		curr := sh.Current()
		if curr != exp.Value {
			t.Errorf("Current(): expected %v, actual %v", exp.Value, curr)
		}

		hist := sh.History()
		if len(hist) != 1 {
			t.Errorf("History(): expected length %v, actual %v", 1, len(hist))
		}

		if hist[0] != exp {
			t.Errorf("History()[0]: expected %v, actual %v", exp, hist[0])
		}
	})
}
