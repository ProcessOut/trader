package trader

import (
	"github.com/shopspring/decimal"
)

// Amount represents an amount in a given currency
type Amount struct {
	Trader   *Trader          `json:"trader"`
	Value    *decimal.Decimal `json:"value"`
	Currency *Currency        `json:"currency"`
}

// NewAmount creates a new amount structure from a decimal and a currency
func (t *Trader) NewAmount(d *decimal.Decimal, code string) (*Amount, error) {
	c, err := t.Currencies.Find(code)
	if err != nil {
		return nil, err
	}

	return &Amount{
		Trader:   t,
		Value:    d,
		Currency: c,
	}, nil
}

// NewAmountFromFloat creates a new amount structure from a float and
// a currency
func (t *Trader) NewAmountFromFloat(f float64, c string) (*Amount, error) {
	d := decimal.NewFromFloat(f)
	return t.NewAmount(&d, c)
}

// NewAmountFromString creates a new amount structure from a string
// and a currency. Returns an error if the string could not be parsed
func (t *Trader) NewAmountFromString(s, c string) (*Amount, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return nil, err
	}
	return t.NewAmount(&d, c)
}

// BaseCurrencyValue returns the value converted to the base currency
// of the Trader. If the Currency of the Amount is already the base currency,
// the value is returned directly
func (a Amount) BaseCurrencyValue() *decimal.Decimal {
	if a.Currency.Is(a.Trader.BaseCurrency.Code) {
		return a.Value
	}

	v := a.Value.Div(*a.Currency.Value)
	return &v
}

// BaseCurrencyAmount returns a new Amount representing the Amount converted
// to the base currency of the Trader. If the Currency of the Amount is already
// the base currency, the Amount is returned directly
func (a *Amount) BaseCurrencyAmount() *Amount {
	if a.Currency.Is(a.Trader.BaseCurrency.Code) {
		return a
	}

	v := a.Value.Div(*a.Currency.Value)
	r, _ := a.Trader.NewAmount(&v, a.Trader.BaseCurrency.Code)
	return r
}

// ToCurrency converts the Amount to the given Currency. If the given Currency
// is the same as the currency one of the Amount, the Amount is returned
// directly
func (a *Amount) ToCurrency(code string) (*Amount, error) {
	if a.Currency.Is(code) {
		return a, nil
	}

	b := a.BaseCurrencyAmount()
	if b.Currency.Is(code) {
		return b, nil
	}

	c, err := a.Trader.Currencies.Find(code)
	if err != nil {
		return nil, err
	}

	v := b.Value.Mul(*c.Value)
	return a.Trader.NewAmount(&v, code)
}
