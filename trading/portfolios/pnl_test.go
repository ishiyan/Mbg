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

	if p.NetAmount() != 0 {
		t.Errorf(fmtVal, "NetAmount", p.NetAmount())
	}

	if len(p.NetAmountHistory()) != 0 {
		t.Errorf(fmtLen, "NetAmount", len(p.NetAmountHistory()))
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

	if p.NetPercentage() != 0 {
		t.Errorf(fmtVal, "NetPercentage", p.NetPercentage())
	}

	if len(p.NetPercentageHistory()) != 0 {
		t.Errorf(fmtLen, "NetPercentage", len(p.NetPercentageHistory()))
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
		amountNet        = "NetAmount"
		amountUnrealized = "UnrealizedAmount"
		percentage       = "Percentage"
		percentageNet    = "NetPercentage"
	)

	t0 := time.Now()
	p := newPnL()
	ah := []entities.Scalar{}
	nah := []entities.Scalar{}
	uah := []entities.Scalar{}
	ph := []entities.Scalar{}
	nph := []entities.Scalar{}

	tests := []struct {
		t                   time.Time
		entryAmount         float64
		amount              float64
		unrealizedAmount    float64
		cashFlow            float64
		netCashFlow         float64
		pnlAmount           float64
		pnlNetAmount        float64
		pnlUnrealizedAmount float64
		pnlPercentage       float64
		pnlNetPercentage    float64
		add                 bool
	}{
		{t0.AddDate(0, 0, 1), -100, -100, -200, -20, -10, -120, -110, -200, 120, 110, true},
		{t0.AddDate(0, 0, 2), -100, 50, 100, -10, -5, 40, 45, 100, -40, -45, true},
	}

	for _, tt := range tests {
		p.add(tt.t, tt.entryAmount, tt.amount, tt.unrealizedAmount, tt.cashFlow, tt.netCashFlow)

		if tt.add {
			ah = append(ah, entities.Scalar{Time: tt.t, Value: tt.amount + tt.cashFlow})
			nah = append(nah, entities.Scalar{Time: tt.t, Value: tt.amount + tt.netCashFlow})
			uah = append(uah, entities.Scalar{Time: tt.t, Value: tt.unrealizedAmount})

			if tt.entryAmount != 0 {
				ph = append(ph, entities.Scalar{Time: tt.t, Value: (tt.amount + tt.cashFlow) / tt.entryAmount * hundred})
				nph = append(nph, entities.Scalar{Time: tt.t, Value: (tt.amount + tt.netCashFlow) / tt.entryAmount * hundred})
			} else {
				ph = append(ph, entities.Scalar{Time: tt.t, Value: 0})
				nph = append(nph, entities.Scalar{Time: tt.t, Value: 0})
			}
		}

		if p.Amount() != tt.pnlAmount {
			t.Errorf(fmtVal, amount, tt.pnlAmount, p.Amount())
		}

		if p.NetAmount() != tt.pnlNetAmount {
			t.Errorf(fmtVal, amountNet, tt.pnlNetAmount, p.NetAmount())
		}

		if p.UnrealizedAmount() != tt.pnlUnrealizedAmount {
			t.Errorf(fmtVal, amountUnrealized, tt.pnlUnrealizedAmount, p.UnrealizedAmount())
		}

		if p.Percentage() != tt.pnlPercentage {
			t.Errorf(fmtVal, percentage, tt.pnlPercentage, p.Percentage())
		}

		if p.NetPercentage() != tt.pnlNetPercentage {
			t.Errorf(fmtVal, percentageNet, tt.pnlNetPercentage, p.NetPercentage())
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

		h = p.NetAmountHistory()
		if len(h) != len(nah) {
			t.Errorf(fmtLen, amountNet, len(nah), len(h))
		} else {
			for i, v := range nah {
				if h[i] != v {
					t.Errorf(fmtIdx, amountNet, i, h[i], v)
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

		h = p.NetPercentageHistory()
		if len(h) != len(nph) {
			t.Errorf(fmtLen, percentageNet, len(nph), len(h))
		} else {
			for i, v := range nph {
				if h[i] != v {
					t.Errorf(fmtIdx, percentageNet, i, h[i], v)
				}
			}
		}
	}
}
