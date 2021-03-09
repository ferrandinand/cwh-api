package dto

import (
	"github.com/ferrandinand/cwh-api/errs"
)

type NewProjectRequest struct {
	Name      string
	CreatedBy string `db:"created_by"`
	Group     string
	RepoURL   string `db:"repo_url"`
}

func (r NewProjectRequest) Validate() *errs.AppError {
	if r.Group == "" {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	return nil
}
