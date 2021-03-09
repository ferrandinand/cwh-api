package domain

import (
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/errs"
)

type User struct {
	Id         string `db:"user_id"`
	Name       string
	CreatedOn  string `db:"created_on"`
	Role       string
	Email      string
	Attributes string
	Status     int
}

func (u User) statusAsText() string {
	statusAsText := "active"
	if u.Status == 0 {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (u User) ToDto() dto.UserResponse {
	return dto.UserResponse{
		Id:         u.Id,
		Name:       u.Name,
		CreatedOn:  u.CreatedOn,
		Role:       u.Role,
		Email:      u.Email,
		Attributes: u.Attributes,
		Status:     u.statusAsText(),
	}
}

type UserRepository interface {
	FindAll(status string) ([]User, *errs.AppError)
	ById(string) (*User, *errs.AppError)
}
