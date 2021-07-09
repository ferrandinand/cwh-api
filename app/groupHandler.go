package app

import (
	"net/http"

	"github.com/ferrandinand/cwh-api/service"

	"github.com/gorilla/mux"
)

type GroupHandler struct {
	service service.GroupService
}

func (h GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId := vars["group_id"]

	group, appError := h.service.GetGroup(groupId)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, group)
	}
}

func (h GroupHandler) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	//Get pageID from the context created in the pagination
	pageID := r.Context().Value("page_id")

	groups, appError := h.service.GetAllGroup(status, pageID.(int))
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, groups)
	}

}
