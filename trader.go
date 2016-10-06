// Package trader takes charge of the amounts handling and currency conversions.
package trader

import "errors"

// Trader is the structure containing the conversions values used to
// handle the amount conversions
type Trader struct {
	Currencies   Currencies `json:"currencies"`
	BaseCurrency *Currency  `json:"base_currency"`
}

// New creates a new Trader structure, and sets the base currency to the
// given currency code. This currency code should be provided in the
// currencies slice, otherwise an error is returned
func New(currencies Currencies, code string) (*Trader, error) {
	// Verify currencies
	for _, c := range currencies {
		if c != nil {
			if ok := Verify(c.Code); !ok {
				return nil, errors.New("Currency `" + c.Code + "' isn't part of ISO 4217")
			}
		}
	}
	if ok := Verify(code); !ok {
		return nil, errors.New("Currency `" + code + "' isn't part of ISO 4217")
	}

	// Init trader
	c, err := currencies.Find(code)
	if err != nil {
		return nil, err
	}

	return &Trader{
		Currencies:   currencies,
		BaseCurrency: c,
	}, nil
}

// SetBaseCurrency sets the base currency for the Trader, from the given
// Currency code. If the currency code was not found in the currencies of the
// Trader, an error is returned
func (t *Trader) SetBaseCurrency(code string) error {
	c, err := t.Currencies.Find(code)
	if err != nil {
		return err
	}

	t.BaseCurrency = c
	return nil
}

// Is compares two trader. If the base currency of t and trader are not the same,
// returns false. If trader does not contain a currency from t, returns false.
// If one of the currencies of trader does not have the same value as the one
// of t, returns false. If the number of currencies supported by t and trader
// is not the same, returns false. Returns true otherwise
func (t Trader) Is(trader *Trader) bool {
	if !t.BaseCurrency.Is(trader.BaseCurrency.Code) ||
		t.BaseCurrency.Value.Cmp(*trader.BaseCurrency.Value) != 0 {

		return false
	}

	if len(t.Currencies) != len(trader.Currencies) {
		return false
	}

	for _, c := range t.Currencies {
		tc, err := trader.Currencies.Find(c.Code)
		if err != nil {
			return false
		}

		if c.Value.Cmp(*tc.Value) != 0 {
			return false
		}
	}

	return true
}
