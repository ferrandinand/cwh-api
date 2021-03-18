package app

import (
	"encoding/json"
	"net/http"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/service"
	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	service service.ProjectService
}

//name CreatedBy Group RepoURL
func (h ProjectHandler) NewProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectName := vars["project_name"]
	var request dto.NewProjectRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.Name = projectName
		project, appError := h.service.NewProject(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, project)
		}
	}
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

func (h ProjectHandler) GetAllProject(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	users, err := h.service.GetAllProject(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, users)
	}
}

// New environment
// /customers/2000/accounts/90720
//func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
//	// get the account_id and customer_id from the URL
//	vars := mux.Vars(r)
//	accountId := vars["account_id"]
//	customerId := vars["customer_id"]
//
//	// decode incoming request
//	var request dto.TransactionRequest
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		writeResponse(w, http.StatusBadRequest, err.Error())
//	} else {
//
//		//build the request object
//		request.AccountId = accountId
//		request.CustomerId = customerId
//
//		// make transaction
//		account, appError := h.service.MakeTransaction(request)
//
//		if appError != nil {
//			writeResponse(w, appError.Code, appError.AsMessage())
//		} else {
//			writeResponse(w, http.StatusOK, account)
//		}
//	}
//
//}
