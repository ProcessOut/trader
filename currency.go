package trader

import (
	"errors"
	"fmt"
	"strings"

	"github.com/processout/decimal"
)

// CurrencyCode represents a currency code in the norm ISO 4217
type CurrencyCode string

// format sets the CurrencyCode to the right format
func (c CurrencyCode) format() CurrencyCode {
	return CurrencyCode(strings.ToUpper(string(c)))
}

// String to implement Stringer interface
func (c CurrencyCode) String() string {
	return string(c.format())
}

// Currency represents a currency and its value relative to the dollar
type Currency struct {
	// Code is the ISO 4217 code of the currency
	Code CurrencyCode `json:"code"`
	// Value is the value of the currency, relative to the base currency
	Value decimal.Decimal `json:"value"`
}

// emptyCurrency represents an empty currency
var emptyCurrency = Currency{}

// NewCurrency creates a new Currency structure, but returns an error if the
// currency code is not part of ISO 4217
func NewCurrency(code CurrencyCode, v decimal.Decimal) (Currency, error) {
	// Verify code:
	if ok := code.Verify(); !ok {
		return emptyCurrency, errors.New("Currency `" + code.String() + "' does not exist")
	}
	return Currency{
		Code:  code.format(),
		Value: v,
	}, nil
}

// Currencies represents a slice of Currencies
type Currencies []Currency

// Find finds a Currency within the Currencies slice from the given
// currency code, or returns an error if the currency code was not found
func (c Currencies) Find(code CurrencyCode) (Currency, error) {
	for _, v := range c {
		if v.Is(code) {
			return v, nil
		}
	}

	return emptyCurrency, fmt.Errorf("The currency code %s could not be found.", code)
}

// Is returns true if the given code is the code of the Currency, false
// otherwise
func (c Currency) Is(code CurrencyCode) bool {
	return c.Code == code.format()
}

// DecimalPlaces returns the number of decimal places a currency has
// e.g. for USD there are 2 ($12.25), for JPY there are 0 (5412)
func (c Currency) DecimalPlaces() int {
	// Here we just test for the currencies that don't have 2 decimal places
	return c.Code.Information().Places
}
