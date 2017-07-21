package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func sendJSON(w http.ResponseWriter, response interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func readBody(r io.ReadCloser) []byte {
	body, err := ioutil.ReadAll(io.LimitReader(r, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Close(); err != nil {
		panic(err)
	}
	return body
}

func readJSON(body []byte, v interface{}) error {
	if err := json.Unmarshal(body, &v); err != nil {
		return err
	}
	return nil
}
