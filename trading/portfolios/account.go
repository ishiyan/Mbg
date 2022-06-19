package portfolios

//nolint:gofumpt
import (
	"sync"
	"time"

	"mbg/trading/currencies"
	"mbg/trading/data"
	"mbg/trading/portfolios/accounts/actions"
)

// Account is a single-entry account holding transactions in home currency.
// Transactions in a foreign currency are converted to the home currency.
type Account struct {
	mu           sync.RWMutex
	holder       string
	currency     currencies.Currency
	converter    currencies.Converter
	balance      data.ScalarTimeSeries
	transactions []*Transaction

	// TO-DO: currency conversion commission (fixed % from converted to home + min/max absolute) ???
	// TO-DO: should be an interface with different implementations
}

// newAccount creates a new account in a given home currency.
// This is the only correct way to create an account instance.
func newAccount(holder string, currency currencies.Currency, converter currencies.Converter) *Account {
	return &Account{
		holder:    holder,
		currency:  currency,
		converter: converter,
		balance:   data.ScalarTimeSeries{},
	}
}

// Holder is the holder of this account.
func (a *Account) Holder() string {
	return a.holder
}

// Currency is the home currency of this account.
func (a *Account) Currency() currencies.Currency {
	return a.currency
}

// Balance is the total ot all transactions expressed in the home currency.
func (a *Account) Balance() float64 {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.balance.Current()
}

// BalanceHistory is a time series of the total ot all transactions expressed in the home currency.
func (a *Account) BalanceHistory() []data.Scalar {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return a.balance.History()
}

// TransactionHistory returns a collection of the account transactions.
func (a *Account) TransactionHistory() []*Transaction {
	a.mu.RLock()
	defer a.mu.RUnlock()

	ts := make([]*Transaction, 0, len(a.transactions))
	for _, t := range a.transactions {
		ts = append(ts, t)
	}

	return ts
}

// add deposits (withdraws) an amount of money in the specified currency
// into (from) this account.
//
// It does not check if the balance becomes negative.
//
// The negative amount means debit withdrawal,
// the positive amount means credit deposit.
//
// The amount will be converted into the home currency
// if the indicated currency differs from the home one.
func (a *Account) add(time time.Time, amount float64, currency currencies.Currency, note string) {
	var action actions.Action

	switch {
	case amount < 0:
		action = actions.Debit
		amount = -amount
	case amount > 0:
		action = actions.Credit
	default:
		return
	}

	var conv, rate float64

	switch {
	case currency == a.currency:
		rate = 1
		conv = amount
	default:
		conv, rate = a.converter.Convert(amount, currency, a.currency)
	}

	t := &Transaction{
		action:          action,
		time:            time,
		currency:        currency,
		amount:          amount,
		conversionRate:  rate,
		amountConverted: conv,
		note:            note,
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	switch {
	case action == actions.Debit:
		a.balance.Accumulate(time, -amount)
	default:
		a.balance.Accumulate(time, amount)
	}

	a.transactions = append(a.transactions, t)
}

//nolint:funlen
// addExecution deposits (withdraws) an amount of money associated
// with an order execution into (from) this account.
//
// It does not check if the balance becomes negative.
//
// The order and commission amounts will be converted into the home currency
// if the indicated currencies differ from the home one.
func (a *Account) addExecution(exec *Execution) {
	var action actions.Action

	amount := exec.cashFlow + exec.debt

	switch {
	case amount < 0:
		action = actions.Debit
		amount = -amount
	case amount > 0:
		action = actions.Credit
	default:
		return
	}

	var conv, rate float64

	switch {
	case exec.currency == a.currency:
		rate = 1
		conv = amount
	default:
		conv, rate = a.converter.Convert(amount, exec.currency, a.currency)
	}

	t := &Transaction{
		action:          action,
		time:            exec.reportTime,
		currency:        exec.currency,
		amount:          amount,
		conversionRate:  rate,
		amountConverted: conv,
		note:            "order execution",
	}

	var tc *Transaction

	if exec.commission != 0 {
		switch {
		case exec.commissionCurrency == a.currency:
			rate = 1
			conv = exec.commission
		default:
			conv, rate = a.converter.Convert(exec.commission, exec.commissionCurrency, a.currency)
		}

		tc = &Transaction{
			action:          actions.Debit,
			time:            exec.reportTime,
			currency:        exec.commissionCurrency,
			amount:          exec.commission,
			conversionRate:  rate,
			amountConverted: conv,
			note:            "order execution commission",
		}
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	switch {
	case action == actions.Debit:
		a.balance.Accumulate(exec.reportTime, -amount)
	default:
		a.balance.Accumulate(exec.reportTime, amount)
	}

	a.transactions = append(a.transactions, t)

	if tc != nil {
		a.balance.Accumulate(exec.reportTime, -tc.amountConverted)
		a.transactions = append(a.transactions, tc)
	}
}
