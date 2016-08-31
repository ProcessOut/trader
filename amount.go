package trader

import (
	"fmt"

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

// RateTo returns the rate that would be applied to convert the amount of a
// to the target currency. An error is returned if the provided code is not
// in the Amount Trader currency list
func (a Amount) RateTo(code string) (*decimal.Decimal, error) {
	c, err := a.Trader.Currencies.Find(code)
	if err != nil {
		return nil, err
	}

	r := c.Value.Div(*a.Currency.Value)
	return &r, nil
}

// ToCurrency converts the Amount to the given Currency. If the given Currency
// is the same as the currency one of the Amount, the Amount is returned
// directly
func (a Amount) ToCurrency(code string) (*Amount, error) {
	if a.Currency.Is(code) {
		return &a, nil
	}

	rate, err := a.RateTo(code)
	if err != nil {
		return nil, err
	}

	v := a.Value.Mul(*rate)
	return a.Trader.NewAmount(&v, code)
}

// Add returns a new Amount corresponding to the sum of a and b. The
// currency of the returned amount is the same as the Currency of a.
// The returned Amount will use the Trader of a for any future operation.
// If the trader of a and b is not the same, an error is returned
func (a Amount) Add(b *Amount) (*Amount, error) {
	if !a.Trader.Is(b.Trader) {
		return nil, fmt.Errorf("The trader of a and b are not the same.")
	}

	n, _ := b.ToCurrency(a.Currency.Code)
	r := a.Value.Add(*n.Value)
	return a.Trader.NewAmount(&r, a.Currency.Code)
}

// Sub returns a new Amount corresponding to the substraction of b from a. The
// currency of the returned amount is the same as the Currency of a.
// The returned Amount will use the Trader of a for any future operation.
// If the trader of a and b is not the same, an error is returned
func (a Amount) Sub(b *Amount) (*Amount, error) {
	if !a.Trader.Is(b.Trader) {
		return nil, fmt.Errorf("The trader of a and b are not the same.")
	}

	n, _ := b.ToCurrency(a.Currency.Code)
	r := a.Value.Sub(*n.Value)
	return a.Trader.NewAmount(&r, a.Currency.Code)
}

// String returns the amount value with the given number of decimals
func (a Amount) String(decimals int32) string {
	return a.Value.StringFixed(decimals)
}
