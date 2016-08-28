# trader

[![GoDoc](https://godoc.org/github.com/ProcessOut/trader?status.svg)](https://godoc.org/github.com/ProcessOut/trader)

Trader takes charge of the amounts handling and currency conversions. It makes
operations on amounts in different currencies easy to perform, and uses
arbitrary-precision fixed-point decimal numbers.

## Installation

```
go get github.com/processout/trader
```

## Usage

```go
// We first want to define the currencies we support
usd, _ := decimal.NewFromString("1") // USD will be our base currency
eur, _ := decimal.NewFromString("0.8")
currencies := trader.Currencies{
    trader.NewCurrency("USD", &usd),
    trader.NewCurrency("EUR", &eur),
}

// We may now create our trader, and set its base currency to the US dollar
t, _ := trader.New(currencies, "usd")

// Now that we have our trader, we can create amounts
price, _ := decimal.NewFromString("42.42")
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
