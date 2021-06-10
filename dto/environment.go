package dto

import (
	"github.com/ferrandinand/cwh-lib/errs"
)

type NewEnvironmentRequest struct {
	Name    string
	Project int
}

type EnvironmentResponse struct {
	EnvironmentId int
	Name          string
	Project       int
	CreatedOn     string
	Status        string
	Attributes    map[string]interface{}
}

type EnvironmentResponseList struct {
	Items      []EnvironmentResponse `json:"items"`
	NextPageID int                   `json:"next_page_id,omitempty" example:"10"`
}

func (r NewEnvironmentRequest) NameIsNull() bool {
	if r.Name == "" {
		return true
	}
	return false
}

func (r NewEnvironmentRequest) Validate() *errs.AppError {

	if r.NameIsNull() {
		return errs.NewValidationError("Environment name cannot be null")
	}
	return nil
}
