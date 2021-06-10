package app

import (
	"encoding/json"
	"net/http"

	"github.com/ferrandinand/cwh-api/dto"
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

func (ch *UserHandlers) getCurrentUser(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User")

	user, err := ch.service.GetUser(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, user)
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

func (ch *UserHandlers) NewUser(w http.ResponseWriter, r *http.Request) {
	var request dto.NewUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		user, appError := ch.service.NewUser(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, user)
		}
	}
}

func (ch *UserHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var request dto.UserRequest
	vars := mux.Vars(r)
	userId := vars["user_id"]

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		user, appError := ch.service.UpdateUser(userId, request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, user)
		}
	}
}

func (ch *UserHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]

	user, appError := ch.service.DeleteUser(userId)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, user)
	}
}
