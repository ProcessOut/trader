package trader

import "errors"

// Trader is the structure containing the conversions values used to
// handle the amount conversions
type Trader struct {
	Currencies   Currencies `json:"currencies"`
	BaseCurrency *Currency  `json:"base_currency"`
}

// New creates a new Trader structure. The default base Currency of the Trader
// is set to the first Currency of the given currencies slice. The currencies
// slice must contain at least 1 currency
func New(currencies Currencies) (*Trader, error) {
	if len(currencies) == 0 {
		return nil, errors.New("The currencies must contain at least one currency.")
	}

	return &Trader{
		Currencies:   currencies,
		BaseCurrency: currencies[0],
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
