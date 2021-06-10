package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/service"

	"github.com/gorilla/mux"
)

type EnvironmentHandler struct {
	service service.EnvironmentService
}

func (h EnvironmentHandler) NewEnvironment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var request dto.NewEnvironmentRequest

	projectId := vars["project_id"]

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		//build the request object
		request.Project, _ = strconv.Atoi(projectId)

		// create envs
		environment, appError := h.service.NewEnvironment(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, environment)
		}
	}
}

func (h EnvironmentHandler) GetEnviroment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	environmentId := vars["environment_id"]

	project, appError := h.service.GetEnvironment(environmentId)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, project)
	}
}

func (h EnvironmentHandler) GetAllEnvironment(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	vars := mux.Vars(r)
	projectId := vars["project_id"]

	//Get pageID from the context
	pageID := r.Context().Value("page_id")

	t, _ := strconv.Atoi(projectId)

	environments, err := h.service.GetAllEnvironment(t, status, pageID.(int))
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, environments)
	}
}
