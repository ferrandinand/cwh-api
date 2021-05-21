package dto

import "github.com/ferrandinand/cwh-lib/errs"

type UserResponse struct {
	Id         string
	Name       string
	CreatedOn  string
	Role       string
	Email      string
	Attributes map[string]interface{}
	Status     string
}

type UserRequest struct {
	Name       string
	Role       string
	Email      string
	Attributes map[string]interface{}
	Status     string
}

type NewUserRequest struct {
	Name       string
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
