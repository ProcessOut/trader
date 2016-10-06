package trader

import (
	"testing"

	"github.com/processout/decimal"
)

func TestNewCurrency(t *testing.T) {
	d := decimal.Decimal{}
	c, err := NewCurrency("test", d)
	if err == nil {
		t.Error("Should have been error")
	} else if c != emptyCurrency {
		t.Error("The currency was not nil")
	}

	c, err = NewCurrency("usd", d)
	if err != nil {
		t.Error("Shouldn't have been an error")
	}
	if c.Code != "USD" {
		t.Error("The code was not correctly set")
	}
	if c.Value != d {
		t.Error("The currency value was not correctly set")
	}
}

func TestCurrencies_Find(t *testing.T) {
	d := decimal.NewFromFloat(1)
	c, _ := NewCurrency("usd", d)
	cs := Currencies{c}

	f, err := cs.Find("bad")
	if err == nil {
		t.Error("An error should have been returned")
	}
	if f != emptyCurrency {
		t.Error("No currency should have been found")
	}

	f, err = cs.Find("usd")
	if err != nil {
		t.Error("No error should have happened")
	}
	if f == emptyCurrency {
		t.Error("The currency should have been found")
	}

	f, err = cs.Find("uSd")
	if err != nil {
		t.Error("No error should have happened")
	}
	if f == emptyCurrency {
		t.Error("The currency should have been found")
	}
}

func TestCurrency_Is(t *testing.T) {
	d := decimal.NewFromFloat(1)
	c, _ := NewCurrency("usd", d)

	if c.Is("bad") {
		t.Error("Should have returned false")
	}
	if !c.Is("usd") {
		t.Error("Should have returned true")
	}
	if !c.Is("UsD") {
		t.Error("Should have returned true")
	}
}

func TestCurrency_DecimalPlaces(t *testing.T) {
	d := decimal.NewFromFloat(1)

	c, _ := NewCurrency("usd", d)
	if c.DecimalPlaces() != 2 {
		t.Error("The decimal places should have been 2")
	}

	c, _ = NewCurrency("bif", d)
	if c.DecimalPlaces() != 0 {
		t.Error("The decimal places should have been 0")
	}

	c, _ = NewCurrency("bhd", d)
	if c.DecimalPlaces() != 3 {
		t.Error("The decimal places should have been 3")
	}

	c, _ = NewCurrency("clf", d)
	if c.DecimalPlaces() != 4 {
		t.Error("The decimal places should have been 4")
	}
}
