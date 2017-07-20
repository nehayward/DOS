package api

import (
	"log"
	"net/http"
)

var attacks Attacks

func Start() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
