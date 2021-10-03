//nolint:testpackage
package portfolios

//nolint:gci
import (
	"mbg/trading/data/entities"
	"testing"
	"time"
)

func TestPnLEmpty(t *testing.T) {
	t.Parallel()

	const (
		fmtVal = "%v(): expected 0, actual %v"
		fmtLen = "%vHistory(): expected length 0, actual %v"
	)

	p := newPnL()

	if p.Amount() != 0 {
		t.Errorf(fmtVal, "Amount", p.Amount())
	}

	if len(p.AmountHistory()) != 0 {
		t.Errorf(fmtLen, "Amount", len(p.AmountHistory()))
	}

	if p.UnrealizedAmount() != 0 {
		t.Errorf(fmtVal, "UnrealizedAmount", p.UnrealizedAmount())
	}

	if len(p.UnrealizedAmountHistory()) != 0 {
		t.Errorf(fmtLen, "UnrealizedAmount", len(p.UnrealizedAmountHistory()))
	}

	if p.Percentage() != 0 {
		t.Errorf(fmtVal, "Percentage", p.Percentage())
	}

	if len(p.PercentageHistory()) != 0 {
		t.Errorf(fmtLen, "Percentage", len(p.PercentageHistory()))
	}
}

//nolint:funlen,gocognit
func TestPnLadd(t *testing.T) {
	t.Parallel()

	const (
		fmtVal           = "%v(): expected %v, actual %v"
		fmtLen           = "%vHistory(): expected length %v, actual %v"
		fmtIdx           = "%vHistory()[%v]: expected %v, actual %v"
		amount           = "Amount"
		amountUnrealized = "UnrealizedAmount"
		percentage       = "Percentage"
	)

	t0 := time.Now()
	p := newPnL()
	ah := []entities.Scalar{}
	uah := []entities.Scalar{}
	ph := []entities.Scalar{}

	tests := []struct {
		t                   time.Time
		entryAmount         float64
		amount              float64
		unrealizedAmount    float64
		cashFlow            float64
		pnlAmount           float64
		pnlUnrealizedAmount float64
		pnlPercentage       float64
		add                 bool
	}{
		{t0.AddDate(0, 0, 1), -100, -100, -200, -20, -120, -200, 120, true},
		{t0.AddDate(0, 0, 2), -100, 50, 100, -10, 40, 100, -40, true},
	}

	for _, tt := range tests {
		p.add(tt.t, tt.entryAmount, tt.amount, tt.unrealizedAmount, tt.cashFlow)

		if tt.add {
			ah = append(ah, entities.Scalar{Time: tt.t, Value: tt.amount + tt.cashFlow})
			uah = append(uah, entities.Scalar{Time: tt.t, Value: tt.unrealizedAmount})

			if tt.entryAmount != 0 {
				ph = append(ph, entities.Scalar{Time: tt.t, Value: (tt.amount + tt.cashFlow) / tt.entryAmount * hundred})
			} else {
				ph = append(ph, entities.Scalar{Time: tt.t, Value: 0})
			}
		}

		if p.Amount() != tt.pnlAmount {
			t.Errorf(fmtVal, amount, tt.pnlAmount, p.Amount())
		}

		if p.UnrealizedAmount() != tt.pnlUnrealizedAmount {
			t.Errorf(fmtVal, amountUnrealized, tt.pnlUnrealizedAmount, p.UnrealizedAmount())
		}

		if p.Percentage() != tt.pnlPercentage {
			t.Errorf(fmtVal, percentage, tt.pnlPercentage, p.Percentage())
		}

		h := p.AmountHistory()
		if len(h) != len(ah) {
			t.Errorf(fmtLen, amount, len(ah), len(h))
		} else {
			for i, v := range ah {
				if h[i] != v {
					t.Errorf(fmtIdx, amount, i, h[i], v)
				}
			}
		}

		h = p.UnrealizedAmountHistory()
		if len(h) != len(uah) {
			t.Errorf(fmtLen, amountUnrealized, len(uah), len(h))
		} else {
			for i, v := range uah {
				if h[i] != v {
					t.Errorf(fmtIdx, amountUnrealized, i, h[i], v)
				}
			}
		}

		h = p.PercentageHistory()
		if len(h) != len(ph) {
			t.Errorf(fmtLen, percentage, len(ph), len(h))
		} else {
			for i, v := range ph {
				if h[i] != v {
					t.Errorf(fmtIdx, percentage, i, h[i], v)
				}
			}
		}
	}
}
