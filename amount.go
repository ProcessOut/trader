package trader

import (
	"github.com/shopspring/decimal"
)

// Amount represents an amount in a given currency
type Amount struct {
	Value    *decimal.Decimal `json:"value"`
	Currency string           `json:"currency"`
}

// NewAmount creates a new amount structure from a decimal and a currency
func NewAmount(d *decimal.Decimal, c string) *Amount {
	return &Amount{
		Value:    d,
		Currency: c,
	}
}

// NewAmountFromFloat creates a new amount structure from a float and
// a currency
func NewAmountFromFloat(f float64, c string) *Amount {
	d := decimal.NewFromFloat(f)
	return NewAmount(&d, c)
}

// NewAmountFromString creates a new amount structure from a string
// and a currency. Returns an error if the string could not be parsed
func NewAmountFromString(s, c string) (*Amount, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return nil, err
	}
	return NewAmount(&d, c), nil
}
