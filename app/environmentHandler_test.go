package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/mocks/service"

	"github.com/ferrandinand/cwh-lib/errs"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var routerEnv *mux.Router
var chEnv EnvironmentHandler
var mockEnvService *service.MockEnvironmentService

func setupEnv(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockEnvService = service.NewMockEnvironmentService(ctrl)
	chEnv = EnvironmentHandler{mockEnvService}
	routerEnv = mux.NewRouter()
	routerEnv.HandleFunc("/project/{project_id}/environments", chEnv.GetAllEnvironment)
	return func() {
		routerEnv = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_projects_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setupEnv(t)
	defer teardown()

	jsonMock := map[string]interface{}{
		"test": "1",
	}

	var body io.Reader
	body = strings.NewReader("{\"Name\":\"test\"},{\"Status\":\"active\"}")

	dummyEnvironments := []dto.EnvironmentResponse{
		{1, "master", 2, "2021-01-01 00:00:00", "active", jsonMock},
		{3, "test90909", 2, "2021-03-30 12:54:47", "inactive", jsonMock},
		{4, "test90909", 2, "2021-03-30 12:56:06", "active", jsonMock},
	}

	mockEnvService.EXPECT().GetAllEnvironment(1, "", 1).Return(dummyEnvironments, nil)
	request, _ := http.NewRequest(http.MethodGet, "/project/1/environments", body)

	// Act
	recorder := httptest.NewRecorder()
	routerEnv.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("%v", recorder.Code)
		t.Error("Failed while testing the status code")
	}

}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setupEnv(t)
	defer teardown()

	var body io.Reader
	body = strings.NewReader("{\"status\":\"active\"}")

	mockEnvService.EXPECT().GetAllEnvironment(1, "", 1).Return(nil, errs.NewUnexpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodGet, "/project/1/environments", body)

	// Act
	recorder := httptest.NewRecorder()
	routerEnv.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
