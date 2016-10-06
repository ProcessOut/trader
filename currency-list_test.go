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

func TestFullName(t *testing.T) {
	if c := CurrencyCode("AFA").FullName(); c != "Afghan Afghani" {
		t.Error("Wrong name")
	}
	if c := CurrencyCode("afa").FullName(); c != "Afghan Afghani" {
		t.Error("Wrong name")
	}
	if c := CurrencyCode("lel").FullName(); len(c) != 0 {
		t.Error("Wrong name")
	}
}
