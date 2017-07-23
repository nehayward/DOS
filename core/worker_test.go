package core

import "testing"

func TestNewWorker(t *testing.T) {
	w := Worker{}
	w.NewWorker()

	if w.ID == 0 {
		t.Error("Expected ID, got", w.ID)
	}
}

func TestStartWorker(t *testing.T) {
	a := Attack{}
	w := Worker{}

	w.NewWorker()
	w.StartWorker(a.Info, 100, "http://bin.org/get")
}
