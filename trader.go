package trader

import (
	"fmt"
)

// Trader is the structure containing the conversions values used to
// handle the amount conversions
type Trader struct {
	Currencies   []*Currency `json:"currencies"`
	BaseCurrency *Currency   `json:"base_currency"`
}

// New creates a new Trader structure. The default base Currency of the Trader
// is set to the first Currency of the given currencies slice
func New(currencies []*Currency) *Trader {
	return &Trader{
		Currencies:   currencies,
		BaseCurrency: currencies[0],
	}
}

// FindCurrency finds a Currency in the currencies of the Trader, from the
// given Currency code. If the currency code was not found in the currencies
// of the Trader, an error is returned
func (t Trader) FindCurrency(code string) (*Currency, error) {
	for _, v := range t.Currencies {
		if v.Code == code {
			return v, nil
		}
	}

	return nil, fmt.Errorf("The currency code %s could not be found.", code)
}

// SetBaseCurrency sets the base currency for the Trader, from the given
// Currency code. If the currency code was not found in the currencies of the
// Trader, an error is returned
func (t Trader) SetBaseCurrency(code string) error {
	c, err := t.FindCurrency(code)
	if err != nil {
		return err
	}

	t.BaseCurrency = c
	return nil
}
