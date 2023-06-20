package app

import (
	"net/http"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/service"

	"github.com/gorilla/mux"
)

type StatusHandler struct {
	service service.StatusService
}

func (h StatusHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["project_id"]

	resource, appError := h.service.GetStatus(projectId)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, resource)
	}
}
