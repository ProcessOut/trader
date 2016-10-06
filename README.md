# trader

[![GoDoc](https://godoc.org/github.com/ProcessOut/trader?status.svg)](https://godoc.org/github.com/ProcessOut/trader)

Trader takes charge of the amounts handling and currency conversions. It makes
operations on amounts in different currencies easy to perform, and uses
arbitrary-precision fixed-point decimal numbers.

## Installation

```bash
go get gopkg.in/processout/trader.v1
```

## Usage

Notes:

- All ISO 4217 currencies are supported, and they can be found in `currency-list.go`
- Error handling isn't included here for brevity's sake

```go
// We first want to define the currencies we support
usdV := decimal.NewFromFloat(1) // USD will be our base currency
eurV := decimal.NewFromFloat(0.8)
usd, _ := trader.NewCurrency("USD", &usdV) // returns error if
eur, _ := trader.NewCurrency("eur", &eurV) // currency isn't ISO 4217

// We add them to the list
currencies := trader.Currencies{usd, eur}

// We may now create our trader, and set its base currency to the US dollar
t, _ := trader.New(currencies, "usd")

// Now that we have our trader, we can create amounts
price, _ := decimal.NewFromFloat(42.42)
amount, _ := t.NewAmount(&price, "USD")
// or amount, _ := trader.NewAmountFromString("42.42", "USD")
// or amount, _ := trader.NewAmountFromFloat(42.42, "USD")

// With this amount, we can now do currency conversions
amountEUR, _ := amount.ToCurrency("EUR")

// We could also add two amounts with different currencies
USDPlusEUR, _ := amount.Add(amountEUR)

// The result is:
// USD(42.00) + EUR(USD(42.00)) == USD(42.00) + USD(42.00) = USD(82.00)
fmt.Println(USDPlusEUR.String())
```

## Notes

Trader uses the package [github.com/shopspring/decimal](github.com/shopspring/decimal)
to perform its calculations.
