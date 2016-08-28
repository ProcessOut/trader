// Package trader takes charge of the amounts handling and currency conversions.
package trader

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
