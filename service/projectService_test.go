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

type NewProjectRequest struct {
	Name      string
	Type      string
	CreatedBy string
	Group     int
	RepoURL   string
}

func Test_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewProjectRequest{
		Name:  "test",
		Type:  "eo", //must be basic or advanced
		Group: 2,
	}
	service := NewProjectService(nil)
	// Act
	_, appError := service.NewProject(request)
	// Assert
	if appError == nil {
		t.Error("failed while testing the new project validation")
	}
}

var mockProjectRepo *domain.MockProjectRepository
var service ProjectService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockProjectRepo = domain.NewMockProjectRepository(ctrl)
	service = NewProjectService(mockProjectRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_project_cannot_be_created(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewProjectRequest{
		Name:  "test",
		Type:  "advanced",
		Group: 2,
	}
	var jsonEmpty map[string]interface{}

	project := realdomain.Project{
		Name:       "test",
		Type:       "advanced",
		Group:      2,
		CreatedOn:  time.Now().Format("2006-01-02 15:04:05"),
		Attributes: jsonEmpty,
		Activities: jsonEmpty,
		Status:     "1",
	}

	mockProjectRepo.EXPECT().Save(project).Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	// Act
	_, appError := service.NewProject(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for new project")
	}

}

func Test_should_return_new_project_response_when_a_new_project_is_saved_successfully(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.NewProjectRequest{
		Name:  "test",
		Type:  "advanced",
		Group: 2,
	}

	project := realdomain.Project{
		Name:      req.Name,
		Type:      req.Type,
		Group:     req.Group,
		CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
		Status:    "1",
	}

	projectWithName := project
	projectWithName.Name = "test"
	mockProjectRepo.EXPECT().Save(project).Return(&projectWithName, nil)
	// Act
	newProject, appError := service.NewProject(req)

	// Assert
	if appError != nil {
		t.Error("Test failed while creating new project")
	}
	if newProject.Name != projectWithName.Name {
		t.Error("Failed while mathching new project name")
	}
}
