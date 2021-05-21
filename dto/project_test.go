package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_project_type_is_not_basic_or_advanced(t *testing.T) {
	request := NewProjectRequest{
		Type: "invalid project type",
	}
	appError := request.Validate()

	if appError.Message != "Project type can only be basic or advanced" {
		t.Error(" Invalid message while testing project type")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error(" Invalid code while testing project type")
	}
}

func Test_should_return_error_when_there_are_missing_project_fields(t *testing.T) {
	request := NewProjectRequest{
		Type: "basic",
	}
	appError := request.Validate()

	if appError.Message != "Mandatory fields project name cannot be empty" {
		t.Error("Repository error message when cannot be empty")
	}
}

func Test_should_return_error_when_there_are_missing_project_name_field(t *testing.T) {
	request := NewProjectRequest{
		Name: "",
		Type: "basic",
	}
	appError := request.Validate()

	if appError.Message != "Mandatory fields project name cannot be empty" {
		t.Error("Repository error message when Name cannot be empty")
	}
}
