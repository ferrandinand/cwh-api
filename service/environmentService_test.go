package service

import (
	"testing"
	"time"

	realdomain "github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/mocks/domain"
	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/golang/mock/gomock"
)

type NewEnvironmentRequest struct {
	Name        string
	Environment int
}

func Test_environment_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewEnvironmentRequest{
		Name: "",
	}
	service := NewEnvironmentService(nil)
	// Act
	_, appError := service.NewEnvironment(request)
	// Assert
	if appError == nil {
		t.Error("failed while testing the new environment validation")
	}
}

var mockEnvRepo *domain.MockEnvironmentRepository
var serviceEnvironment EnvironmentService

func setupEnvironment(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockEnvRepo = domain.NewMockEnvironmentRepository(ctrl)
	serviceEnvironment = NewEnvironmentService(mockEnvRepo)
	return func() {
		serviceEnvironment = nil
		defer ctrl.Finish()
	}
}

func Test_environment_should_return_an_error_from_the_server_side_if_the_new_environment_cannot_be_created(t *testing.T) {
	// Arrange
	teardown := setupEnvironment(t)
	defer teardown()

	req := dto.NewEnvironmentRequest{
		Name:    "test",
		Project: 1,
	}
	environment := realdomain.Environment{
		Name:      "test",
		Project:   1,
		CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
	}
	mockEnvRepo.EXPECT().Save(environment).Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	// Act
	_, appError := serviceEnvironment.NewEnvironment(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for new environment")
	}

}

func Test_should_return_new_environment_response_when_a_new_environment_is_saved_successfully(t *testing.T) {
	// Arrange
	teardown := setupEnvironment(t)
	defer teardown()

	req := dto.NewEnvironmentRequest{
		Name:    "test",
		Project: 1,
	}

	environment := realdomain.Environment{
		Name:      "test",
		Project:   1,
		CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
	}
	environmentWithName := environment
	environmentWithName.Name = "test"
	mockEnvRepo.EXPECT().Save(environment).Return(&environmentWithName, nil)
	// Act
	newEnvironment, appError := serviceEnvironment.NewEnvironment(req)

	// Assert
	if appError != nil {
		t.Error("Test failed while creating new environment")
	}
	if newEnvironment.Name != environmentWithName.Name {
		t.Error("Failed while mathching new environment name")
	}
}
