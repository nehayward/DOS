package api

import (
	"log"
	"net/http"

	"github.com/nehayward/dos/core"
)

var attacks core.Attacks

func Start() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
