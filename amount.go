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
	c, err := t.FindCurrency(code)
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
