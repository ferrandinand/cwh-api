package dto

import "github.com/ferrandinand/cwh-lib/errs"

type UserResponse struct {
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	LastName   string                 `json:"last_name"`
	CreatedOn  string                 `json:"created_on"`
	Role       string                 `json:"role"`
	Email      string                 `json:"email"`
	Attributes map[string]interface{} `json:"attributes"`
	Status     string                 `json:"status"`
}

type UserRequest struct {
	Name       string
	LastName   string
	Role       string
	Email      string
	Attributes map[string]interface{}
	Status     string
}

type NewUserRequest struct {
	Name       string
	LastName   string
	Password   string
	Role       string
	Email      string
	Attributes map[string]interface{}
	Status     string
}

func (r NewUserRequest) Validate() *errs.AppError {
	if r.Name == "" || r.Password == "" || r.Email == "" {
		return errs.NewValidationError("Mandatory fields cannot be empty")
	}
	return nil
}

func (r UserRequest) Validate() *errs.AppError {
	if r.Name == "" || r.Email == "" {
		return errs.NewValidationError("Mandatory fields cannot be empty")
	}
	return nil
}
