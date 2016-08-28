package trader

import (
	"github.com/shopspring/decimal"
)

// Currency represents a currency and its value relative to the dollar
type Currency struct {
	// Code is the ISO 4217 code of the currency
	Code string `json:"code"`
	// Value is the value of the currency, relative to the base currency
	Value *decimal.Decimal `json:"value"`
}

// NewCurrency creates a new Currency structure
func NewCurrency(c string, v *decimal.Decimal) *Currency {
	return &Currency{
		Code:  c,
		Value: v,
	}
}
