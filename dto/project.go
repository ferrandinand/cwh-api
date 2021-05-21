package dto

import (
	"github.com/ferrandinand/cwh-lib/errs"
)

const BASIC = "basic"
const ADVANCED = "advanced"

type NewProjectRequest struct {
	Name      string
	Type      string
	CreatedBy string
	Group     int
}

func (r NewProjectRequest) IsProjectTypeBasic() bool {
	return r.Type == BASIC
}

func (r NewProjectRequest) IsProjectTypeAdvanced() bool {
	return r.Type == ADVANCED
}

func (r NewProjectRequest) Validate() *errs.AppError {
	if !r.IsProjectTypeBasic() && !r.IsProjectTypeAdvanced() {
		return errs.NewValidationError("Project type can only be basic or advanced")
	}

	if r.Name == "" {
		return errs.NewValidationError("Mandatory fields project name cannot be empty")
	}
	return nil
}

type ProjectResponse struct {
	Id         string
	Name       string
	Type       string
	CreatedBy  string
	CreatedOn  string
	Group      int
	Attributes map[string]interface{}
	Activities map[string]interface{}
	Status     string
}
