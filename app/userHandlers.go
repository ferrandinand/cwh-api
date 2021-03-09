package app

import (
	"encoding/json"
	"net/http"

	"github.com/ferrandinand/cwh-api/service"
	"github.com/gorilla/mux"
)

type UserHandlers struct {
	service service.UserService
}

func (ch *UserHandlers) getAllUsers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	users, err := ch.service.GetAllUser(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, users)
	}
}

func (ch *UserHandlers) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["user_id"]

	user, err := ch.service.GetUser(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, user)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}