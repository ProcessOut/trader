package trader

import (
	"fmt"
	"testing"

	"github.com/processout/decimal"
)

func TestExample(*testing.T) { Example() }
func Example() {
	// We first want to define the currencies we support
	usd, _ := NewCurrency("USD", decimal.NewFromFloat(1)) // Will be base currency
	eur, _ := NewCurrency("eur", decimal.NewFromFloat(0.8))

	// We add them to the list
	currencies := Currencies{usd, eur}

	// We may now create our trader, and set its base currency to the US dollar
	t, _ := New(currencies, "usd")

	// Now that we have our trader, we can create amounts
	amount, _ := t.NewAmount(decimal.NewFromFloat(42.42), "USD")
	// or amount, _ := NewAmountFromString("42.42", "USD")
	// or amount, _ := NewAmountFromFloat(42.42, "USD")

	// With this amount, we can now do currency conversions
	amountEUR, _ := amount.ToCurrency("EUR") // = "33.936"

	// We could also add two amounts with different currencies
	USDPlusEUR, _ := amount.Add(amountEUR)

	// The result is:
	// USD(42.42) + EUR(33.936) == USD(42.42) + USD(42.42) = USD(84.84)
	fmt.Println(USDPlusEUR.String(2)) // Prints 84.84
}

func TestExampleNew(*testing.T) { ExampleNew() }
func ExampleNew() {
	// We first want to define the currencies we support
	usd, _ := NewCurrency("USD", decimal.NewFromFloat(1))
	eur, _ := NewCurrency("EUR", decimal.NewFromFloat(0.8))
	currencies := Currencies{usd, eur}

	// We may now create our trader, and set its base currency to the US dollar
	New(currencies, "USD")
}

func TestExampleTrader_NewAmount(*testing.T) { ExampleTrader_NewAmount() }
func ExampleTrader_NewAmount() {
	// We first want to define the currencies we support
	usdRate, _ := decimal.NewFromString("1") // USD will be our base currency
	eurRate, _ := decimal.NewFromString("0.8")
	usd, _ := NewCurrency("USD", usdRate)
	eur, _ := NewCurrency("EUR", eurRate)

	// Create a mock Trader
	t, _ := New(Currencies{usd, eur}, "USD")

	d, _ := decimal.NewFromString("42.42")
	t.NewAmount(d, "usd")
}

func TestExampleTrader_NewAmountFromString(*testing.T) { ExampleTrader_NewAmountFromString() }
func ExampleTrader_NewAmountFromString() {
	// We first want to define the currencies we support
	usdRate, _ := decimal.NewFromString("1") // USD will be our base currency
	eurRate, _ := decimal.NewFromString("0.8")
	usd, _ := NewCurrency("USD", usdRate)
	eur, _ := NewCurrency("EUR", eurRate)

	// Create a mock Trader
	t, _ := New(Currencies{usd, eur}, "USD")

	t.NewAmountFromString("42.42", "usd")
}

func TestExampleTrader_NewAmountFromFloat(*testing.T) { ExampleTrader_NewAmountFromFloat() }
func ExampleTrader_NewAmountFromFloat() {
	// We first want to define the currencies we support
	usdRate, _ := decimal.NewFromString("1") // USD will be our base currency
	eurRate, _ := decimal.NewFromString("0.8")
	usd, _ := NewCurrency("USD", usdRate)
	eur, _ := NewCurrency("EUR", eurRate)

	// Create a mock Trader
	t, _ := New(Currencies{usd, eur}, "USD")

	t.NewAmountFromFloat(42.42, "usd")
}

func TestExampleAmount_String(*testing.T) { ExampleAmount_String() }
func ExampleAmount_String() {
	// We first want to define the currencies we support
	usdRate, _ := decimal.NewFromString("1") // USD will be our base currency
	eurRate, _ := decimal.NewFromString("0.8")
	usd, _ := NewCurrency("USD", usdRate)
	eur, _ := NewCurrency("EUR", eurRate)

	// Create a mock Trader
	t, _ := New(Currencies{usd, eur}, "USD")

	a, _ := t.NewAmountFromFloat(42.42, "usd")

	fmt.Println(a.String(3))
	// Output: 42.420
}
