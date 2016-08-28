package trader

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewCurrency(t *testing.T) {
	d := &decimal.Decimal{}
	c := NewCurrency("test", d)
	if c == nil {
		t.Error("The currency was nil")
	}
	if c.Code != "TEST" {
		t.Error("The code was not correctly set")
	}
	if *c.Value != *d {
		t.Error("The currency value was not correctly set")
	}
}

func TestFind(t *testing.T) {
	d := decimal.NewFromFloat(1)
	c := NewCurrency("test", &d)
	cs := Currencies{c}

	f, err := cs.Find("bad")
	if err == nil {
		t.Error("An error should have been returned")
	}
	if f != nil {
		t.Error("No currency should have been found")
	}

	f, err = cs.Find("test")
	if err != nil {
		t.Error("No error should have happened")
	}
	if f == nil {
		t.Error("The currency should have bene found")
	}

	f, err = cs.Find("TeSt")
	if err != nil {
		t.Error("No error should have happened")
	}
	if f == nil {
		t.Error("The currency should have bene found")
	}
}
