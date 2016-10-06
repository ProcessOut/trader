package trader

import (
	"fmt"
	"math"

	"github.com/processout/decimal"
)

// Amount represents an amount in a given currency
type Amount struct {
	Trader   Trader          `json:"-"`
	Value    decimal.Decimal `json:"value"`
	Currency Currency        `json:"currency"`
}

// emptyAmount represents an empty amount
var emptyAmount = Amount{}

// NewAmount creates a new amount structure from a decimal and a currency
func (t Trader) NewAmount(d decimal.Decimal, code CurrencyCode) (Amount, error) {
	c, err := t.Currencies.Find(code)
	if err != nil {
		return emptyAmount, err
	}

	return Amount{
		Trader:   t,
		Value:    d,
		Currency: c,
	}, nil
}

// NewAmountFromFloat creates a new amount structure from a float and
// a currency
func (t *Trader) NewAmountFromFloat(f float64, c CurrencyCode) (Amount, error) {
	return t.NewAmount(decimal.NewFromFloat(f), c)
}

// NewAmountFromString creates a new amount structure from a string
// and a currency. Returns an error if the string could not be parsed
func (t *Trader) NewAmountFromString(s string, c CurrencyCode) (Amount, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return emptyAmount, err
	}
	return t.NewAmount(d, c)
}

// RateTo returns the rate that would be applied to convert the amount of a
// to the target currency. An error is returned if the provided code is not
// in the Amount Trader currency list
func (a Amount) RateTo(code CurrencyCode) (decimal.Decimal, error) {
	c, err := a.Trader.Currencies.Find(code)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return c.Value.Div(a.Currency.Value), nil
}

// ToCurrency converts the Amount to the given Currency. If the given Currency
// is the same as the currency one of the Amount, the Amount is returned
// directly
func (a Amount) ToCurrency(code CurrencyCode) (Amount, error) {
	if a.Currency.Is(code) {
		return a, nil
	}

	rate, err := a.RateTo(code)
	if err != nil {
		return emptyAmount, err
	}

	return a.Trader.NewAmount(a.Value.Mul(rate), code)
}

// Add returns a new Amount corresponding to the sum of a and b. The
// currency of the returned amount is the same as the Currency of a.
// The returned Amount will use the Trader of a for any future operation.
// If the trader of a and b is not the same, an error is returned
func (a Amount) Add(b Amount) (Amount, error) {
	if !a.Trader.Is(b.Trader) {
		return emptyAmount, fmt.Errorf("The trader of a and b are not the same.")
	}

	n, _ := b.ToCurrency(a.Currency.Code)
	return a.Trader.NewAmount(a.Value.Add(n.Value), a.Currency.Code)
}

// Sub returns a new Amount corresponding to the substraction of b from a. The
// currency of the returned amount is the same as the Currency of a.
// The returned Amount will use the Trader of a for any future operation.
// If the trader of a and b is not the same, an error is returned
func (a Amount) Sub(b Amount) (Amount, error) {
	if !a.Trader.Is(b.Trader) {
		return emptyAmount, fmt.Errorf("The trader of a and b are not the same.")
	}

	n, _ := b.ToCurrency(a.Currency.Code)
	return a.Trader.NewAmount(a.Value.Sub(n.Value), a.Currency.Code)
}

// Cmp compares a and b precisely in this order.
// Returns:
//	- = 0 if a is equal to b
//  - < 0 if a is smaller than b
//  - > 0 if a is greater than b
// To compare a and b, b is first converted to the currency of a
func (a Amount) Cmp(b Amount) (int, error) {
	if !a.Trader.Is(b.Trader) {
		return 0, fmt.Errorf("The trader of a and b are not the same.")
	}

	n, _ := b.ToCurrency(a.Currency.Code)
	return a.Value.Cmp(n.Value), nil
}

// IsEmpty returns true if the amount is empty
func (a Amount) IsEmpty() bool {
	return a.Currency == emptyAmount.Currency &&
		a.Value.Equals(emptyAmount.Value) &&
		a.Trader.Is(emptyAmount.Trader)
}

// Int64 translates an amount into an in64 by adjusting its amount to the
// lowest possible decimal of its currency (ex: USD: 10.23 -> 1023)
func (a Amount) Int64() int64 {
	return a.Value.Mul(
		decimal.NewFromFloat(math.Pow10(a.Currency.DecimalPlaces())),
	).IntPart()
}

// String returns the amount value with the given number of decimals
func (a Amount) String(decimals int32) string {
	return a.Value.StringFixed(decimals)
}
