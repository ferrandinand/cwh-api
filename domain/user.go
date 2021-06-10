package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"database/sql/driver"

	"golang.org/x/crypto/bcrypt"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
)

type AttributesUser map[string]interface{}

type User struct {
	Id         string `db:"user_id"`
	Name       string
	LastName   string `db:"last_name"`
	Password   string
	CreatedOn  string `db:"created_on"`
	Role       string
	Email      string
	Attributes AttributesUser
	Status     string
}

func setUserPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword)
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
		LastName:   u.LastName,
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
	ByUsername(string) (*User, *errs.AppError)
	NewUser(user User) (*User, *errs.AppError)
	UpdateUser(userId string, user User) (*User, *errs.AppError)
	DeleteUser(userId string) (*User, *errs.AppError)
}

func NewUser(name string, lastName string, password string, role string, email string) User {
	var jsonEmpty map[string]interface{}

	return User{
		Name:       name,
		LastName:   lastName,
		CreatedOn:  time.Now().Format("2006-01-02 15:04:05"),
		Password:   password,
		Role:       role,
		Email:      email,
		Attributes: jsonEmpty,
		Status:     "1",
	}
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
