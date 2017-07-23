package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nehayward/dos/core"
)

func AttackList(w http.ResponseWriter, r *http.Request) {
	if len(attacks) > 0 {
		sendJSON(w, attacks, http.StatusOK)
		return
	}
	sendJSON(w, jsonError{Code: http.StatusNotFound, Message: "No active attacks"}, http.StatusOK)
}

func AttackShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	attackID, err := strconv.Atoi(vars["attackid"])
	if err != nil {
		panic(err)
	}

	attack, exists, _ := attacks.FindAttack(attackID)
	if exists {
		sendJSON(w, attack, http.StatusOK)
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Code: http.StatusNotFound, Message: "Not Found"}); err != nil {
		panic(err)
	}

}

func AttackCreate(w http.ResponseWriter, r *http.Request) {
	var attack core.Attack
	body := readBody(r.Body)

	e := readJSON(body, &attack)
	if e != nil {
		response := jsonError{Message: "Bad JSON"}
		sendJSON(w, response, http.StatusUnprocessableEntity)
		return
	}

	if attack.CheckURL() {
		attack.Info = &core.Info{}
		attack.ID = attacks.GetNewID()
		attack.NewAttack()

		attacks.AddAttack(attack)
		message := Response{Message: "Started attack!", URL: attack.URL, ID: attack.ID}
		sendJSON(w, message, http.StatusCreated)
	}
}

func AttackStop(w http.ResponseWriter, r *http.Request) {
	var attack core.Attack
	body := readBody(r.Body)

	if err := json.Unmarshal(body, &attack); err != nil {
		response := jsonError{Message: "Bad JSON"}
		sendJSON(w, response, http.StatusUnprocessableEntity)
		return
	}

	a, exists, _ := attacks.FindAttack(attack.ID)
	if exists {
		a.Worker.StopWorker()
		fmt.Println(attacks)
		attacks.RemoveAttack(attack.ID)
		fmt.Println(attacks)
		response := Response{Message: "Stopping attack!", ID: attack.ID}
		sendJSON(w, response, http.StatusAccepted)
		return
	}
	response := Response{Message: "ID Not Found", ID: attack.ID}
	sendJSON(w, response, http.StatusAccepted)
}
