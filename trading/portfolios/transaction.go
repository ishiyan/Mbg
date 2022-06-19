package portfolios

//nolint:gofumpt
import (
	"time"

	"mbg/trading/currencies"
	"mbg/trading/portfolios/accounts/actions"
)

// Transaction is an immutable account transaction.
type Transaction struct {
	action          actions.Action
	time            time.Time
	currency        currencies.Currency
	amount          float64
	conversionRate  float64
	amountConverted float64
	note            string
}

// Action indicates if this is a deposit or a withdrawal.
func (t *Transaction) Action() actions.Action {
	return t.action
}

// Time is the date and time of this transaction.
func (t *Transaction) Time() time.Time {
	return t.time
}

// Currency is the the currency of this transaction.
func (t *Transaction) Currency() currencies.Currency {
	return t.currency
}

// Amount is the unsigned amount in of this transaction in the given currency.
// The sign is determided by the Action().
func (t *Transaction) Amount() float64 {
	return t.amount
}

// ConversionRate is an exchange rate from the given currency to the home currency of this account.
func (t *Transaction) ConversionRate() float64 {
	return t.conversionRate
}

// AmountConverted is the unsigned amount in of this transaction converted to the home currency of this account.
// The sign is determided by the Action().
func (t *Transaction) AmountConverted() float64 {
	return t.amountConverted
}

// Note is a free-text note of this transaction.
func (t *Transaction) Note() string {
	return t.note
}
