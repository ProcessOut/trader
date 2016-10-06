package trader

import "testing"

func TestVerify(t *testing.T) {
	if ok := CurrencyCode("lel").Verify(); ok {
		t.Error("Currency isn't valid")
	}
	if ok := CurrencyCode("asdf").Verify(); ok {
		t.Error("Currency isn't valid")
	}
	if ok := CurrencyCode("CNY").Verify(); !ok {
		t.Error("Currency is valid")
	}
	if ok := CurrencyCode("cny").Verify(); !ok {
		t.Error("Currency is valid")
	}
}

func TestInformation(t *testing.T) {
	if c := CurrencyCode("AED").Information(); c == nil ||
		c.FullName != "United Arab Emirates dirham" ||
		c.Number != 784 ||
		c.Places != 2 ||
		c.Countries[0] != "United Arab Emirates" {
		t.Error("Wrong information")
	}
	if c := CurrencyCode("aEd").Information(); c == nil || c.FullName != "United Arab Emirates dirham" {
		t.Error("Wrong information")
	}
	if c := CurrencyCode("lel").Information(); c != nil {
		t.Error("False positive")
	}
}
