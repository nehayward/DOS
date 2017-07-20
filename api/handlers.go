package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func AttackList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(attacks); err != nil {
		panic(err)
	}
}

func AttackShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var attackID int
	var err error
	if attackID, err = strconv.Atoi(vars["attack"]); err != nil {
		panic(err)
	}
	fmt.Println(attackID)
	// todo := RepoFindTodo(todoId)
	attack := Attack{}
	if attack.ID > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(attack); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func AttackCreate(w http.ResponseWriter, r *http.Request) {
	var attack Attack
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &attack); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	attack.ID = time.Now().UnixNano()
	attacks = append(attacks, attack)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	message := Response{Message: "Started attack!", URL: attack.URL, ID: time.Now().UnixNano()}

	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}

func AttackStop(w http.ResponseWriter, r *http.Request) {
	var attack Attack
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &attack); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	fmt.Println(attack.ID)

	message := Response{Message: "Stopping attack!", ID: attack.ID}

	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}
