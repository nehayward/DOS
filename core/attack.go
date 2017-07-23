package core

import (
	"fmt"
	"net/http"
	"net/url"
)

// NewAttack setups new attack
func (a *Attack) NewAttack() {
	if a.Method == "" {
		a.Method = http.MethodGet
	}
	if a.Requests == 0 {
		a.Requests = 1000
	}

	a.Worker.NewWorker()
	a.Worker.StartWorker(a.Info, a.Requests, a.URL)
}

// CheckURL determines string is url
func (a *Attack) CheckURL() bool {
	fmt.Println(a.URL)
	_, err := url.ParseRequestURI(a.URL)
	if err != nil {
		return false
	}
	return true
}

// Exists checks if attack is empty
func (a Attack) Exists() bool {
	if a != (Attack{}) {
		return true
	}
	return false
}
