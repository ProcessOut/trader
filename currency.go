package trader

import (
	"fmt"
	"strings"

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
func NewCurrency(code string, v *decimal.Decimal) *Currency {
	return &Currency{
		Code:  strings.ToUpper(code),
		Value: v,
	}
}

// Currencies represents a slice of Currencies
type Currencies []*Currency

// Find finds a Currency within the Currencies slice from the given
// currency code, or returns an error if the currency code was not found
func (c Currencies) Find(code string) (*Currency, error) {
	code = strings.ToUpper(code)
	for _, v := range c {
		if v.Code == code {
			return v, nil
		}
	}

	return nil, fmt.Errorf("The currency code %s could not be found.", code)
}
