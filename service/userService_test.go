package service

import (
	"testing"
	"time"

	realdomain "github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/mocks/domain"
	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/golang/mock/gomock"
)

type NewUserRequest struct {
	Name       string
	Password   string
	Role       string
	Email      string
	Attributes map[string]interface{}
	Status     string
}

func Test_user_should_return_a_validation_error_response_when_the_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewUserRequest{
		Name:       "",
		Password:   "", //must be not null
		Role:       "Admin",
		Email:      "e0@test.com",
		Attributes: nil,
		Status:     "active",
	}
	userService := NewUserService(nil)
	// Act
	_, appError := userService.NewUser(request)
	// Assert
	if appError == nil {
		t.Error("failed while testing the new user validation")
	}
}

var mockUserRepo *domain.MockUserRepository
var userService UserService

func setupUser(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockUserRepo = domain.NewMockUserRepository(ctrl)
	userService = NewUserService(mockUserRepo)
	return func() {
		userService = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_user_cannot_be_created(t *testing.T) {
	// Arrange
	teardown := setupUser(t)
	defer teardown()

	req := dto.NewUserRequest{
		Name:       "test",
		Password:   "test",
		Role:       "Admin",
		Email:      "dsdd@dsds.com",
		Attributes: nil,
		Status:     "active",
	}

	user := realdomain.User{
		Name:      "test",
		Password:  "test",
		Role:      "Admin",
		Email:     "dsdd@dsds.com",
		Status:    "1",
		CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
	}

	mockUserRepo.EXPECT().NewUser(user).Return(nil, errs.NewUnexpectedError("Unexpected database error"))
	// Act
	_, appError := userService.NewUser(req)

	// Assert
	if appError == nil {
		t.Error("Test failed while validating error for new user")
	}

}

func Test_should_return_new_user_response_when_a_new_user_is_saved_successfully(t *testing.T) {
	// Arrange
	teardown := setupUser(t)
	defer teardown()

	req := dto.NewUserRequest{
		Name:       "test",
		Password:   "test",
		Role:       "admin",
		Email:      "email@email.com",
		Attributes: nil,
		Status:     "active",
	}

	user := realdomain.User{
		Name:       "test",
		CreatedOn:  time.Now().Format("2006-01-02 15:04:05"),
		Password:   "test",
		Role:       "admin",
		Email:      "email@email.com",
		Attributes: nil,
		Status:     "1",
	}

	userWithName := user
	userWithName.Name = "test"
	mockUserRepo.EXPECT().NewUser(user).Return(&userWithName, nil)
	// Act
	newUser, appError := userService.NewUser(req)

	// Assert
	if appError != nil {
		t.Error("Test failed while creating new project")
	}
	if newUser.Name != userWithName.Name {
		t.Error("Failed while mathching new user name")
	}
}
