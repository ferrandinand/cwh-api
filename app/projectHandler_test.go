package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/mocks/service"

	"github.com/ferrandinand/cwh-lib/errs"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var router *mux.Router
var ch ProjectHandler
var mockService *service.MockProjectService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockProjectService(ctrl)
	ch = ProjectHandler{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/projects", ch.GetAllProject)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_project_should_return_projects_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	jsonMock := map[string]interface{}{
		"test": "1",
	}

	dummyProjects := []dto.ProjectResponse{
		{"2", "test-2", "basic", "stan", "01/01/2021", 1, jsonMock, jsonMock, "1"},
		{"3", "test-2", "basic", "stan", "01/01/2021", 1, jsonMock, jsonMock, "1"},
		{"4", "test-2", "basic", "stan", "01/01/2021", 1, jsonMock, jsonMock, "1"},
		{"5", "test-2", "basic", "stan", "01/01/2021", 1, jsonMock, jsonMock, "1"},
	}

	mockService.EXPECT().GetAllProject("").Return(dummyProjects, nil)
	request, _ := http.NewRequest(http.MethodGet, "/projects", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_project_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	mockService.EXPECT().GetAllProject("").Return(nil, errs.NewUnexpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodGet, "/projects", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
