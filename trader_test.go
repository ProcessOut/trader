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
	c1, _ := NewCurrency("USD", &usd)
	c2, _ := NewCurrency("EUR", &eur)

	currencies = Currencies{c1, c2}
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
	c1, _ := NewCurrency("USD", &usd)
	c2, _ := NewCurrency("EUR", &eur)
	currencies := Currencies{c1, c2}
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
	c1, _ := NewCurrency("USD", &usd)
	c2, _ := NewCurrency("EUR", &eur)
	currencies := Currencies{c1, c2}
	currencies2 := Currencies{c1, c2}
	trader, _ := New(currencies, "usd")
	trader2, _ := New(currencies2, "usd")

	if !trader.Is(trader2) {
		t.Error("The traders should have been equal")
	}

	trader2, _ = New(currencies2, "eur")
	if trader.Is(trader2) {
		t.Error("The traders should have not been equal")
	}

	currencies2 = Currencies{c1, c2}

	trader2, _ = New(currencies2, "usd")
	if trader.Is(trader2) {
		t.Error("The traders should have not been equal")
	}

	c3, _ := NewCurrency("GHC", &bad)
	currencies2 = Currencies{c1, c3}

	trader2, _ = New(currencies2, "usd")
	if trader.Is(trader2) {
		t.Error("The traders should have not been equal")
	}

	currencies2 = Currencies{c1, c2, c3}

	trader2, _ = New(currencies2, "usd")
	if trader.Is(trader2) {
		t.Error("The traders should have not been equal")
	}
}
