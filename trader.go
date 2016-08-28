// Package trader takes charge of the amounts handling and currency conversions.
//
// Working with Trader is as follows:
//
// 		// We first want to define the currencies we support
// 		usd, _ := decimal.NewFromString("1") // USD will be our base currency
// 		eur, _ := decimal.NewFromString("0.8")
// 		currencies := trader.Currencies{
//     		trader.NewCurrency("USD", &usd),
//     		trader.NewCurrency("EUR", &eur),
// 		}
//
//		// We may now create our trader, and set its base currency to the US dollar
// 		t, _ := trader.New(currencies, "usd")
//
// 		// Now that we have our trader, we can create amounts
// 		amount, _ := t.NewAmountFromString("42.42", "USD")
//
// 		// With this amount, we can now do currency conversions
// 		amountEUR, _ := amount.ToCurrency("EUR")
//
// 		// We could also add two amounts with different currencies
// 		USDPlusEUR, _ := amount.Add(amountEUR)
//
// 		// The result is:
// 		// USD(42.00) + EUR(USD(42.00)) == USD(42.00) + USD(42.00) = USD(82.00)
// 		fmt.Println(USDPlusEUR.String())
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
