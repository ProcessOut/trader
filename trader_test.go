package trader

import (
	"testing"

	"github.com/processout/decimal"
)

func TestNew(t *testing.T) {
	usd := decimal.NewFromFloat(1)
	eur := decimal.NewFromFloat(0.8)
	currencies := Currencies{}
	trader, err := New(currencies, "usd")
	if err == nil {
		t.Error("An error should have been returned")
	}
	if trader != nil {
		t.Error("The trader should have been nil")
	}

	currencies = Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}
	trader, err = New(currencies, "bad")
	if err == nil {
		t.Error("An error should have been returned")
	}
	if trader != nil {
		t.Error("The trader should have been nil")
	}

	trader, err = New(currencies, "usd")
	if err != nil {
		t.Error("No error should have happened")
	}
	if trader == nil {
		t.Error("The trader should have been set")
	}
}

func TestTrader_SetBaseCurrency(t *testing.T) {
	usd := decimal.NewFromFloat(1)
	eur := decimal.NewFromFloat(0.8)
	currencies := Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}
	trader, err := New(currencies, "usd")
	if trader.BaseCurrency.Code != "USD" {
		t.Error("The default base currency was not correctly set")
	}

	err = trader.SetBaseCurrency("bad")
	if err == nil {
		t.Error("There should have been an error")
	}
	if trader.BaseCurrency.Code != "USD" {
		t.Error("The base currency was updated but shouldn't have been")
	}

	err = trader.SetBaseCurrency("eur")
	if err != nil {
		t.Error("There shouldn't have been an error")
	}
	if trader.BaseCurrency.Code != "EUR" {
		t.Error("The base currency was not correctly set")
	}
}

func TestTrader_Is(t *testing.T) {
	usd := decimal.NewFromFloat(1)
	eur := decimal.NewFromFloat(0.8)
	bad := decimal.NewFromFloat(0.5)
	currencies := Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}
	currencies2 := Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
	}
	trader, _ := New(currencies, "usd")
	trader2, _ := New(currencies2, "usd")

	if !trader.Is(trader2) {
		t.Error("The traders should have been equal")
	}

	trader2, _ = New(currencies2, "eur")
	if trader.Is(trader2) {
		t.Error("The traders should have not been equal")
	}

	currencies2 = Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &bad),
	}

	trader2, _ = New(currencies2, "usd")
	if trader.Is(trader2) {
		t.Error("The traders should have not been equal")
	}

	currencies2 = Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("BAD", &eur),
	}

	trader2, _ = New(currencies2, "usd")
	if trader.Is(trader2) {
		t.Error("The traders should have not been equal")
	}

	currencies2 = Currencies{
		NewCurrency("USD", &usd),
		NewCurrency("EUR", &eur),
		NewCurrency("BAD", &bad),
	}

	trader2, _ = New(currencies2, "usd")
	if trader.Is(trader2) {
		t.Error("The traders should have not been equal")
	}
}
