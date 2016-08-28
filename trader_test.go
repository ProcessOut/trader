package trader

import (
	"testing"

	"github.com/shopspring/decimal"
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

func TestSetBaseCurrency(t *testing.T) {
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
