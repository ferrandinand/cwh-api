package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/service"
	"github.com/ferrandinand/cwh-lib/logger"

	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	service service.ProjectService
}

func (h ProjectHandler) NewProject(w http.ResponseWriter, r *http.Request) {
	var request dto.NewProjectRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CreatedBy = r.Header.Get("User")
		project, appError := h.service.NewProject(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, project)
		}
	}
	logger.Error("Request " + fmt.Sprintf("%s", request))

}

func (h ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["project_id"]

	project, appError := h.service.GetProject(projectId)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, project)
	}
}

func (h ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["project_id"]

	project, appError := h.service.DeleteProject(projectId)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, project)
	}
}

func (h ProjectHandler) GetAllProject(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	//Get pageID from the context created in the pagination
	pageID := r.Context().Value("page_id")

	projects, appError := h.service.GetAllProject(status, pageID.(int))
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, projects)
	}

}
