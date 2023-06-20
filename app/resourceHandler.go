package app

import (
	"fmt"
	"net/http"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/logger"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/service"

	"github.com/gorilla/mux"
)

type CostHandler struct {
	service service.CostService
}

func (h CostHandler) GetResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["project_id"]

	resource, appError := h.service.GetCost(projectId)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, resource)
	}
}

func (h CostHandler) GetAllResource(w http.ResponseWriter, r *http.Request) {

	//Get pageID from the context created in the pagination

	pageID := r.Context().Value("page_id")
	logger.Info(fmt.Sprintf("In the handler"))

	costs, appError := h.service.GetAllCost(pageID.(int))
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, costs)
	}

}
