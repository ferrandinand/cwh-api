package app

import (
	"encoding/json"
	"net/http"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/service"
	"github.com/gorilla/mux"
)

type ServiceOrderHandler struct {
	service service.ServiceOrder
}

func (s ServiceOrderHandler) NewServiceOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["project_id"]
	environmentId := vars["environment_id"]

	var request dto.NewServiceOrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CreatedBy = r.Header.Get("User")
		request.Project = projectId
		request.Environment = environmentId

		service, appError := s.service.NewServiceOrder(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, service)
		}
	}
}

func (h ServiceOrderHandler) GetServiceOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceOrderId := vars["service_order_id"]

	service, appError := h.service.GetServiceOrder(serviceOrderId)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, service)
	}
}

func (h ServiceOrderHandler) GetEnvironmentServiceOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["project_id"]
	environmentId := vars["environment_id"]

	// decode incoming request
	var request dto.ServiceOrderRequest
	//build the request object
	request.Project = projectId
	request.Environment = environmentId

	service, appError := h.service.GetEnvironmentServiceOrders(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, service)
	}
}
