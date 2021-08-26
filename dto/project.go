package dto

import (
	"github.com/ferrandinand/cwh-lib/errs"
)

const BASIC = "basic"
const ADVANCED = "advanced"

type NewProjectRequest struct {
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	CreatedBy  string                 `json:"created_by"`
	Group      int                    `json:"group"`
	Attributes map[string]interface{} `json:"attributes"`
}

type ProjectResponse struct {
	Id         int                    `json:"id"`
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	CreatedBy  string                 `json:"created_by"`
	CreatedOn  string                 `json:"created_on"`
	Group      int                    `json:"group"`
	Attributes map[string]interface{} `json:"attributes"`
	Activities map[string]interface{} `json:"activities"`
	Status     string                 `json:"status"`
}

type ProjectResponseList struct {
	Items      []ProjectResponse `json:"items"`
	NextPageID int               `json:"next_page_id,omitempty" example:"10"`
	PrevPageID int               `json:"prev_page_id,omitempty" example:"10"`
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
