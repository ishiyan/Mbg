//nolint:testpackage
package portfolios

//nolint:gofumpt
import (
	"testing"
	"time"

	"mbg/trading/currencies"
	"mbg/trading/portfolios/accounts/actions"
)

func TestTransaction(t *testing.T) {
	t.Parallel()

	const (
		fmt             = "%v(): expected %v, actual %v"
		action          = actions.Debit
		currency        = currencies.CHF
		amount          = 12.34
		amountConverted = 56.78
		conversionRate  = 0.9
		note            = "foo"
	)

	time := time.Now()
	tr := Transaction{
		time:            time,
		action:          action,
		currency:        currency,
		amount:          amount,
		amountConverted: amountConverted,
		conversionRate:  conversionRate,
		note:            note,
	}

	if tr.Time() != time {
		t.Errorf(fmt, "Time", time, tr.Time())
	}

	if tr.Action() != action {
		t.Errorf(fmt, "Action", action, tr.Action())
	}

	if tr.Currency() != currency {
		t.Errorf(fmt, "Currency", currency, tr.Currency())
	}

	if tr.Amount() != amount {
		t.Errorf(fmt, "Amount", amount, tr.Amount())
	}

	if tr.AmountConverted() != amountConverted {
		t.Errorf(fmt, "AmountConverted", amountConverted, tr.AmountConverted())
	}

	if tr.ConversionRate() != conversionRate {
		t.Errorf(fmt, "ConversionRate", conversionRate, tr.ConversionRate())
	}

	if tr.Note() != note {
		t.Errorf(fmt, "Note", note, tr.Note())
	}
}
