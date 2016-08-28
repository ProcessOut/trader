# trader

[![GoDoc](https://godoc.org/github.com/ProcessOut/trader?status.svg)](https://godoc.org/github.com/ProcessOut/trader)

Trader takes charge of the amounts handling and currency conversions. It makes
operations on amounts in different currencies easy to perform, and uses
Arbitrary-precision fixed-point decimal numbers.

## Usage

```go
// We first want to define the currencies we support
usd := decimal.NewFromString("1") // USD will be our base currency
eur := decimal.NewFromString("0.8")
currencies := Currencies{
    NewCurrency("USD", &usd),
    NewCurrency("EUR", &eur),
}

// We may now create our trader, and set its base currency to the US dollar
trader, _ := New(currencies, "usd")

// Now that we have our trader, we can create amounts
price := decimal.NewFromString("42.42")
amount := trader.NewAmount(&price, "USD")
// or amount := trader.NewAmountFromString("42.42", "USD")
// or amount := trader.NewAmountFromFloat(42.42, "USD")

// With this amount, we can now do currency conversions
amountEUR := amount.ToCurrency("EUR")

// We could also add two amounts with different currencies
USDPlusEUR := amount.Add(amountEUR)

// The result is:
// USD(42.00) + EUR(USD(42.00)) == USD(42.00) + USD(42.00) = USD(82.00)
fmt.Println(USDPlusEUR.String())
```

## Notes

Trader uses the package github.com/shopspring/decimal to perform its calculations.
