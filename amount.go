package trader

import (
	"github.com/shopspring/decimal"
)

// Amount represents an amount in a given currency
type Amount struct {
	Trader   *Trader          `json:"-"`
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
func (a Amount) BaseCurrencyAmount() *Amount {
	if a.Currency.Is(a.Trader.BaseCurrency.Code) {
		return &a
	}

	v := a.Value.Div(*a.Currency.Value)
	r, _ := a.Trader.NewAmount(&v, a.Trader.BaseCurrency.Code)
	return r
}

// ToCurrency converts the Amount to the given Currency. If the given Currency
// is the same as the currency one of the Amount, the Amount is returned
// directly
func (a Amount) ToCurrency(code string) (*Amount, error) {
	if a.Currency.Is(code) {
		return &a, nil
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

// Add returns a new Amount corresponding to the sum of a and b. The
// Currency of the returned amount is the same as the Currency of a (structure
// on which Add(b) is called). The Trader of b is used to convert b to the
// currency of a in order to do the addition. An error is returned if the
// currency of a is not found in the Currencies slice of the Trader of b
func (a Amount) Add(b *Amount) (*Amount, error) {
	n, err := b.ToCurrency(a.Currency.Code)
	if err != nil {
		return nil, err
	}

	c := a.Value.Add(*n.Value)
	n.Value = &c
	return n, nil
}

// Sub returns a new Amount corresponding to the substraction of b from a. The
// Currency of the returned amount is the same as the Currency of a (structure
// on which Sub(b) is called). The Trader of b is used to convert b to the
// currency of a in order to do the substraction. An error is returned if the
// currency of a is not found in the Currencies slice of the Trader of b
func (a Amount) Sub(b *Amount) (*Amount, error) {
	n, err := b.ToCurrency(a.Currency.Code)
	if err != nil {
		return nil, err
	}

	c := a.Value.Sub(*b.Value)
	n.Value = &c
	return n, nil
}

// Mul returns a new Amount corresponding to the multiplication of a and b. The
// Currency of the returned amount is the same as the Currency of a (structure
// on which Mul(b) is called). The Trader of b is used to convert b to the
// currency of a in order to do the multiplication. An error is returned if the
// currency of a is not found in the Currencies slice of the Trader of b
func (a Amount) Mul(b *Amount) (*Amount, error) {
	n, err := b.ToCurrency(a.Currency.Code)
	if err != nil {
		return nil, err
	}

	c := a.Value.Mul(*n.Value)
	n.Value = &c
	return n, nil
}

// Div returns a new Amount corresponding to the division of a by b. The
// Currency of the returned amount is the same as the Currency of a (structure
// on which Div(b) is called). The Trader of b is used to convert b to the
// currency of a in order to do the division. An error is returned if the
// currency of a is not found in the Currencies slice of the Trader of b.
// Warning: The division isn't precise, the division precision is 16 decimals
func (a Amount) Div(b *Amount) (*Amount, error) {
	n, err := b.ToCurrency(a.Currency.Code)
	if err != nil {
		return nil, err
	}

	c := a.Value.Div(*n.Value)
	n.Value = &c
	return n, nil
}
