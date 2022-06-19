package portfolios

//nolint:gci
import (
	"math"
	"sync"
	"time"

	"mbg/trading/data"
)

// Drawdown contains a time series of drawdown amount, percentage and their maximal values.
type Drawdown struct {
	mu sync.RWMutex

	watermark     *data.Scalar
	amount        *data.Scalar
	percentage    *data.Scalar
	amountMax     *data.Scalar
	percentageMax *data.Scalar

	watermarkHistory     []*data.Scalar
	amountHistory        []*data.Scalar
	percentageHistory    []*data.Scalar
	amountMaxHistory     []*data.Scalar
	percentageMaxHistory []*data.Scalar
}

// Watermark returns the current high watermark amount
// in positive values or zero if not initialized.
func (d *Drawdown) Watermark() float64 {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if v := d.watermark; v != nil {
		return v.Value
	}

	return 0
}

// WatermarkHistory returns the high watermark amount time series
// in positive values or an empty slice if not initialized.
func (d *Drawdown) WatermarkHistory() []data.Scalar {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.watermarkHistory == nil || len(d.watermarkHistory) == 0 {
		return []data.Scalar{}
	}

	v := make([]data.Scalar, len(d.watermarkHistory))
	for i, s := range d.watermarkHistory {
		v[i] = *s
	}

	return v
}

// Amount returns the current drawdown amount
// in negative values or zero if not initialized.
func (d *Drawdown) Amount() float64 {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if v := d.amount; v != nil {
		return v.Value
	}

	return 0
}

// AmountHistory returns the drawdown amount time series
// in negative values or an empty slice if not initialized.
func (d *Drawdown) AmountHistory() []data.Scalar {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.amountHistory == nil || len(d.amountHistory) == 0 {
		return []data.Scalar{}
	}

	v := make([]data.Scalar, len(d.amountHistory))
	for i, s := range d.amountHistory {
		v[i] = *s
	}

	return v
}

// Percentage returns the current drawdown percentage
// in range [-100, 0] or zero if not initialized.
func (d *Drawdown) Percentage() float64 {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if v := d.percentage; v != nil {
		return v.Value * 100 //nolint:gomnd
	}

	return 0
}

// PercentageHistory returns the drawdown percentage time series
// in range [-100, 0] or an empty slice if not initialized.
func (d *Drawdown) PercentageHistory() []data.Scalar {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.percentageHistory == nil || len(d.percentageHistory) == 0 {
		return []data.Scalar{}
	}

	v := make([]data.Scalar, len(d.percentageHistory))
	for i, s := range d.percentageHistory {
		v[i] = *s
		v[i].Value *= 100
	}

	return v
}

// MaxAmount returns the current maximal drawdown amount
// (the minimal negative historical value until now)
// in negative values or zero if not initialized.
func (d *Drawdown) MaxAmount() float64 {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if v := d.amountMax; v != nil {
		return v.Value
	}

	return 0
}

// MaxAmountHistory returns the maximal drawdown amount time series
// (the minimal negative historical value before every sample)
// in negative values or an empty slice if not initialized.
func (d *Drawdown) MaxAmountHistory() []data.Scalar {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.amountMaxHistory == nil || len(d.amountMaxHistory) == 0 {
		return []data.Scalar{}
	}

	v := make([]data.Scalar, len(d.amountMaxHistory))
	for i, s := range d.amountMaxHistory {
		v[i] = *s
	}

	return v
}

// MaxPercentage returns the current maximal drawdown percentage
// (the minimal negative historical value until now)
// in range [-100, 0] or zero if not initialized.
func (d *Drawdown) MaxPercentage() float64 {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if v := d.percentageMax; v != nil {
		return v.Value * 100 //nolint:gomnd
	}

	return 0
}

// MaxPercentageHistory returns the maximal drawdown percentage time series
// (the minimal negative historical value before every sample)
// in range [-100, 0] or an empty slice if not initialized.
func (d *Drawdown) MaxPercentageHistory() []data.Scalar {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.percentageMaxHistory == nil || len(d.percentageMaxHistory) == 0 {
		return []data.Scalar{}
	}

	v := make([]data.Scalar, len(d.percentageMaxHistory))
	for i, s := range d.percentageMaxHistory {
		v[i] = *s
		v[i].Value *= 100
	}

	return v
}

// add adds a new sample to the time series
// if the new sample time is later the last time of the time series.
// Otherwise (if the new sample time is less or equal to the last time),
// the last time series value will be updated.
func (d *Drawdown) add(time time.Time, value float64) {
	d.mu.Lock()
	defer d.mu.Unlock()

	switch {
	case d.watermark == nil:
		s := data.Scalar{Time: time, Value: value}
		d.watermarkHistory = append(d.watermarkHistory, &s)
		d.watermark = &s

		return
	case !d.watermark.Time.Before(time):
		return
	case d.watermark.Value < value:
		s := data.Scalar{Time: time, Value: value}
		d.watermarkHistory = append(d.watermarkHistory, &s)
		d.watermark = &s
	}

	var a, f float64

	if d.watermark.Value != 0 {
		a = math.Min(value-d.watermark.Value, 0)
		f = math.Min(a/d.watermark.Value, 0)
	}

	switch {
	case d.amount == nil:
		d.amount = &data.Scalar{Time: time, Value: a}
		d.amountHistory = append(d.amountHistory, d.amount)

		d.percentage = &data.Scalar{Time: time, Value: f}
		d.percentageHistory = append(d.percentageHistory, d.percentage)

		d.amountMax = &data.Scalar{Time: time, Value: a}
		d.amountMaxHistory = append(d.amountMaxHistory, d.amountMax)

		d.percentageMax = &data.Scalar{Time: time, Value: f}
		d.percentageMaxHistory = append(d.percentageMaxHistory, d.percentageMax)
	case d.amount.Time.Before(time):
		d.amount = &data.Scalar{Time: time, Value: a}
		d.amountHistory = append(d.amountHistory, d.amount)

		d.percentage = &data.Scalar{Time: time, Value: f}
		d.percentageHistory = append(d.percentageHistory, d.percentage)

		d.amountMax = &data.Scalar{Time: time, Value: math.Min(a, d.amountMax.Value)}
		d.amountMaxHistory = append(d.amountMaxHistory, d.amountMax)

		d.percentageMax = &data.Scalar{Time: time, Value: math.Min(f, d.percentageMax.Value)}
		d.percentageMaxHistory = append(d.percentageMaxHistory, d.percentageMax)
	}
}
