package trader

import "testing"

func TestVerify(t *testing.T) {
	if ok := Verify("lel"); ok {
		t.Error("Currency isn't valid")
	}
	if ok := Verify("asdf"); ok {
		t.Error("Currency isn't valid")
	}
	if ok := Verify("CNY"); !ok {
		t.Error("Currency is valid")
	}
	if ok := Verify("cny"); !ok {
		t.Error("Currency is valid")
	}
}

func TestFullName(t *testing.T) {
	if c := FullName("AFA"); c != "Afghan Afghani" {
		t.Error("Wrong name")
	}
	if c := FullName("afa"); c != "Afghan Afghani" {
		t.Error("Wrong name")
	}
	if c := FullName("lel"); len(c) != 0 {
		t.Error("Wrong name")
	}
}
