package trader

import (
	"testing"

	"github.com/shopspring/decimal"
)

func getTrader() *Trader {
	usd := decimal.NewFromFloat(1)
	eur := decimal.NewFromFloat(0.8)
	currencies := Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}
	trader, _ := New(currencies)
	return trader
}

func TestNewAmount(t *testing.T) {
	trader := getTrader()
	d := decimal.NewFromFloat(4.2)

	amount, err := trader.NewAmount(&d, "bad")
	if err == nil {
		t.Error("There should have been an error")
	}

	amount, err = trader.NewAmount(&d, "usd")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	if amount == nil {
		t.Error("The returned amount shouldn't have been nil")
	}

	amount, err = trader.NewAmount(&d, "eur")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	if amount == nil {
		t.Error("The returned amount shouldn't have been nil")
	}
}

func TestNewAmountFromFloat(t *testing.T) {
	trader := getTrader()

	amount, err := trader.NewAmountFromFloat(4.2, "bad")
	if err == nil {
		t.Error("There should have been an error")
	}

	amount, err = trader.NewAmountFromFloat(4.2, "usd")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	if amount == nil {
		t.Error("The returned amount shouldn't have been nil")
	}

	amount, err = trader.NewAmountFromFloat(4.2, "eur")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	if amount == nil {
		t.Error("The returned amount shouldn't have been nil")
	}
}

func TestNewAmountFromString(t *testing.T) {
	trader := getTrader()

	amount, err := trader.NewAmountFromString("bad", "usd")
	if err == nil {
		t.Error("There should have been an error")
	}

	amount, err = trader.NewAmountFromString("4.2", "bad")
	if err == nil {
		t.Error("There should have been an error")
	}

	amount, err = trader.NewAmountFromString("4.2", "usd")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	if amount == nil {
		t.Error("The returned amount shouldn't have been nil")
	}

	amount, err = trader.NewAmountFromString("4.2", "eur")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	if amount == nil {
		t.Error("The returned amount shouldn't have been nil")
	}
}

func TestBaseCurrencyValue(t *testing.T) {
	trader := getTrader()
	amount, _ := trader.NewAmountFromString("1", "usd")

	v := amount.BaseCurrencyValue()
	if v != amount.Value {
		t.Error("The base currency value should have stayed the same")
	}

	amount, _ = trader.NewAmountFromString("1", "eur")
	f, _ := amount.BaseCurrencyValue().Float64()
	if f != 1.25 {
		t.Error("The base currency value was wrongly converted")
	}
}

func TestBaseCurrencyAmount(t *testing.T) {
	trader := getTrader()
	amount, _ := trader.NewAmountFromString("1", "usd")

	a := amount.BaseCurrencyAmount()
	if a != amount {
		t.Error("The base currency amount should have stayed the same")
	}

	amount, _ = trader.NewAmountFromString("1", "eur")
	a = amount.BaseCurrencyAmount()
	if a == amount {
		t.Error("The base currency amount should have changed")
	}
	f, _ := a.Value.Float64()
	if f != 1.25 {
		t.Error("The base currency amount was wrongly converted")
	}
}

func TestToCurrency(t *testing.T) {
	trader := getTrader()
	amount, _ := trader.NewAmountFromString("1", "usd")

	a, err := amount.ToCurrency("bad")
	if err == nil {
		t.Error("There should have been an error")
	}

	a, err = amount.ToCurrency("eur")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	f, _ := a.Value.Float64()
	if f != 0.8 {
		t.Error("The amount was wrongly converted")
	}

	a2, err := a.ToCurrency("usd")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	f, _ = a2.Value.Float64()
	if f != 1 {
		t.Error("The amount was wrongly converted bad")
	}
}
