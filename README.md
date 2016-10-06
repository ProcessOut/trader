# trader

[![GoDoc](https://godoc.org/github.com/ProcessOut/trader?status.svg)](https://godoc.org/github.com/ProcessOut/trader)

Trader takes charge of the amounts handling and currency conversions. It makes
operations on amounts in different currencies easy to perform, and uses
arbitrary-precision fixed-point decimal numbers.

## Installation

```bash
go get gopkg.in/processout/trader
```

## Usage

- All ISO 4217 currencies are supported, and they can be found in `currency-list.go`
- Error handling isn't included here for brevity's sake
- It makes sense that your base currency should always have a comparative value of 1
- Take a look at example_test.go for more examples

`import "github.com/processout/trader"`

```go
// We first want to define the currencies we support
usd, _ := trader.NewCurrency("USD", decimal.NewFromFloat(1)) // Will be base currency
eur, _ := trader.NewCurrency("eur", decimal.NewFromFloat(0.8))

// We add them to the list
currencies := trader.Currencies{usd, eur}

// We may now create our trader, and set its base currency to the US dollar
t, _ := trader.New(currencies, "usd")

// Now that we have our trader, we can create amounts
amount, _ := t.NewAmount(decimal.NewFromFloat(42.42), "USD")
// or amount, _ := t.NewAmountFromString("42.42", "USD")
// or amount, _ := t.NewAmountFromFloat(42.42, "USD")

// With this amount, we can now do currency conversions
amountEUR, _ := amount.ToCurrency("EUR") // = "33.936"

// We could also add two amounts with different currencies
USDPlusEUR, _ := amount.Add(amountEUR)

// The result is:
// USD(42.42) + EUR(33.936) == USD(42.42) + USD(42.42) = USD(84.84)
fmt.Println(USDPlusEUR.String(2)) // Prints 84.84
```

## Notes

Trader uses the package [github.com/processout/decimal](github.com/processout/decimal)
to perform its arbitrary-precision calculations.
