package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/errs"
)

type AttributesUser map[string]interface{}

type User struct {
	Id         string `db:"user_id"`
	Name       string
	Password   string
	CreatedOn  string `db:"created_on"`
	Role       string
	Email      string
	Attributes AttributesUser
	Status     string
}

func (u User) statusAsText() string {
	statusAsText := "active"
	if u.Status == "0" {
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

func (ua *AttributesUser) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &ua)
		return nil
	case string:
		json.Unmarshal([]byte(v), &ua)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (ua *AttributesUser) Value() (driver.Value, error) {
	return json.Marshal(ua)
}
