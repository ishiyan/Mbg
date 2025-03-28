//nolint:testpackage
package portfolios

//nolint:gofumpt
import (
	"testing"
	"time"

	"mbg/trading/data"
)

func TestDrawdownEmpty(t *testing.T) {
	t.Parallel()

	const (
		fmtVal = "%v(): expected 0, actual %v"
		fmtLen = "%vHistory(): expected length 0, actual %v"
	)

	d := Drawdown{}

	if d.Watermark() != 0 {
		t.Errorf(fmtVal, "Watermark", d.Watermark())
	}

	if len(d.WatermarkHistory()) != 0 {
		t.Errorf(fmtLen, "Watermark", len(d.WatermarkHistory()))
	}

	if d.Amount() != 0 {
		t.Errorf(fmtVal, "Amount", d.Amount())
	}

	if len(d.AmountHistory()) != 0 {
		t.Errorf(fmtLen, "Amount", len(d.AmountHistory()))
	}

	if d.Percentage() != 0 {
		t.Errorf(fmtVal, "Percentage", d.Percentage())
	}

	if len(d.PercentageHistory()) != 0 {
		t.Errorf(fmtLen, "Percentage", len(d.PercentageHistory()))
	}

	if d.MaxAmount() != 0 {
		t.Errorf(fmtVal, "MaxAmount", d.MaxAmount())
	}

	if len(d.MaxAmountHistory()) != 0 {
		t.Errorf(fmtLen, "MaxAmount", len(d.MaxAmountHistory()))
	}

	if d.MaxPercentage() != 0 {
		t.Errorf(fmtVal, "MaxPercentage", d.MaxPercentage())
	}

	if len(d.MaxPercentageHistory()) != 0 {
		t.Errorf(fmtLen, "MaxPercentage", len(d.MaxPercentageHistory()))
	}
}

//nolint:funlen,gocognit
func TestDrawdownAdd(t *testing.T) {
	t.Parallel()

	const (
		fmtVal        = "%v(): expected %v, actual %v"
		fmtLen        = "%vHistory(): expected length %v, actual %v"
		fmtIdx        = "%vHistory()[%v]: expected %v, actual %v"
		watermark     = "Watermark"
		amount        = "Amount"
		percentage    = "Percentage"
		maxAmount     = "MaxAmount"
		maxPercentage = "MaxPercentage"
	)

	t0 := time.Now()
	d := Drawdown{}
	wh := []data.Scalar{}
	ah := []data.Scalar{}
	ph := []data.Scalar{}
	amh := []data.Scalar{}
	pmh := []data.Scalar{}

	tests := []struct {
		s             data.Scalar
		watermark     float64
		amount        float64
		percentage    float64
		amountMax     float64
		percentageMax float64
		add           bool
		addWatermark  bool
	}{
		{data.Scalar{Time: t0.AddDate(0, 0, 1), Value: 100}, 100, 0, 0, 0, 0, false, true},
		{data.Scalar{Time: t0.AddDate(0, 0, 2), Value: 110}, 110, 0, 0, 0, 0, true, true},
		{data.Scalar{Time: t0.AddDate(0, 0, 3), Value: 105}, 110, -5, -50. / 11, -5, -50. / 11, true, false},
		{data.Scalar{Time: t0.AddDate(0, 0, 4), Value: 90}, 110, -20, -200. / 11, -20, -200. / 11, true, false},
		{data.Scalar{Time: t0.AddDate(0, 0, 5), Value: 115}, 115, 0, 0, -20, -200. / 11, true, true},
		{data.Scalar{Time: t0.AddDate(0, 0, 4), Value: 999}, 115, 0, 0, -20, -200. / 11, false, false},
		{data.Scalar{Time: t0.AddDate(0, 0, 5), Value: 999}, 115, 0, 0, -20, -200. / 11, false, false},
	}

	for _, tt := range tests {
		d.add(tt.s.Time, tt.s.Value)

		if tt.addWatermark {
			wh = append(wh, data.Scalar{Time: tt.s.Time, Value: tt.watermark})
		}

		if tt.add {
			ah = append(ah, data.Scalar{Time: tt.s.Time, Value: tt.amount})
			ph = append(ph, data.Scalar{Time: tt.s.Time, Value: tt.percentage})
			amh = append(amh, data.Scalar{Time: tt.s.Time, Value: tt.amountMax})
			pmh = append(pmh, data.Scalar{Time: tt.s.Time, Value: tt.percentageMax})
		}

		if d.Watermark() != tt.watermark {
			t.Errorf(fmtVal, watermark, tt.watermark, d.Watermark())
		}

		if d.Amount() != tt.amount {
			t.Errorf(fmtVal, amount, tt.amount, d.Amount())
		}

		if d.Percentage() != tt.percentage {
			t.Errorf(fmtVal, percentage, tt.percentage, d.Percentage())
		}

		if d.MaxAmount() != tt.amountMax {
			t.Errorf(fmtVal, maxAmount, tt.amountMax, d.MaxAmount())
		}

		if d.MaxPercentage() != tt.percentageMax {
			t.Errorf(fmtVal, maxPercentage, tt.percentageMax, d.MaxPercentage())
		}

		h := d.WatermarkHistory()
		if len(h) != len(wh) {
			t.Errorf(fmtLen, watermark, len(wh), len(h))
		} else {
			for i, v := range wh {
				if h[i] != v {
					t.Errorf(fmtIdx, watermark, i, h[i], v)
				}
			}
		}

		h = d.AmountHistory()
		if len(h) != len(ah) {
			t.Errorf(fmtLen, amount, len(ah), len(h))
		} else {
			for i, v := range ah {
				if h[i] != v {
					t.Errorf(fmtIdx, amount, i, h[i], v)
				}
			}
		}

		h = d.PercentageHistory()
		if len(h) != len(ph) {
			t.Errorf(fmtLen, percentage, len(ph), len(h))
		} else {
			for i, v := range ph {
				if h[i] != v {
					t.Errorf(fmtIdx, percentage, i, h[i], v)
				}
			}
		}

		h = d.MaxAmountHistory()
		if len(h) != len(amh) {
			t.Errorf(fmtLen, maxAmount, len(amh), len(h))
		} else {
			for i, v := range amh {
				if h[i] != v {
					t.Errorf(fmtIdx, maxAmount, i, h[i], v)
				}
			}
		}

		h = d.MaxPercentageHistory()
		if len(h) != len(pmh) {
			t.Errorf(fmtLen, maxPercentage, len(pmh), len(h))
		} else {
			for i, v := range pmh {
				if h[i] != v {
					t.Errorf(fmtIdx, maxPercentage, i, h[i], v)
				}
			}
		}
	}
}
