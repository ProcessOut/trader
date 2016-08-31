package trader

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func Example() {
	// We first want to define the currencies we support
	usd, _ := decimal.NewFromString("1") // USD will be our base currency
	eur, _ := decimal.NewFromString("0.8")
	currencies := Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}

	// We may now create our trader, and set its base currency to the US dollar
	t, _ := New(currencies, "usd")

	// Now that we have our trader, we can create amounts
	amount, _ := t.NewAmountFromString("42.42", "USD")

	// With this amount, we can now do currency conversions
	amountEUR, _ := amount.ToCurrency("EUR")

	// We could also add two amounts with different currencies
	USDPlusEUR, _ := amount.Add(amountEUR)

	// The result is:
	// USD(42.42) + EUR(USD(42.42)) == USD(42.42) + USD(42.42) = USD(82.42)
	fmt.Println(USDPlusEUR.String(2))
	// Output: 84.84
}

func ExampleNew() {
	// We first want to define the currencies we support
	usd, _ := decimal.NewFromString("1") // USD will be our base currency
	eur, _ := decimal.NewFromString("0.8")
	currencies := Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}

	// We may now create our trader, and set its base currency to the US dollar
	New(currencies, "USD")
}

func ExampleTrader_NewAmount() {
	// We first want to define the currencies we support
	usd, _ := decimal.NewFromString("1") // USD will be our base currency
	eur, _ := decimal.NewFromString("0.8")

	// Create a mock Trader
	t, _ := New(Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}, "USD")

	d, _ := decimal.NewFromString("42.42")
	t.NewAmount(&d, "usd")
}

func ExampleTrader_NewAmountFromString() {
	// We first want to define the currencies we support
	usd, _ := decimal.NewFromString("1") // USD will be our base currency
	eur, _ := decimal.NewFromString("0.8")

	// Create a mock Trader
	t, _ := New(Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}, "USD")

	t.NewAmountFromString("42.42", "usd")
}

func ExampleTrader_NewAmountFromFloat() {
	// We first want to define the currencies we support
	usd, _ := decimal.NewFromString("1") // USD will be our base currency
	eur, _ := decimal.NewFromString("0.8")

	// Create a mock Trader
	t, _ := New(Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}, "USD")

	t.NewAmountFromFloat(42.42, "usd")
}

func ExampleAmount_String() {
	// We first want to define the currencies we support
	usd, _ := decimal.NewFromString("1") // USD will be our base currency
	eur, _ := decimal.NewFromString("0.8")

	// Create a mock Trader
	t, _ := New(Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}, "USD")

	a, _ := t.NewAmountFromFloat(42.42, "usd")

	fmt.Println(a.String(3))
	// Output: 42.420
}
