package core

import "testing"

func TestCheckURL(t *testing.T) {
	a := Attack{URL: "http://google.com"}
	v := a.CheckURL()
	if !v {
		t.Error("Expected true, got", false)
	}

	a.URL = "notAURL"
	v = a.CheckURL()
	if v {
		t.Error("Expected false, got", true)
	}
}

func TestExists(t *testing.T) {
	var a Attack
	var v bool

	v = a.Exists()
	if v {
		t.Error("Expected false, got", true)
	}

	a.ID = 1
	v = a.Exists()
	if !v {
		t.Error("Expected true, got", false)
	}
}

// TODO: Reformat
func TestNewAttack(t *testing.T) {
	a := Attack{URL: "http://google.com"}
	a.NewAttack()
	if a.Worker.ID == 0 {
		t.Error("Expected new attack, but not started")
	}
}
