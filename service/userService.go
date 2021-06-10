package service

import (
	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
)

//go:generate mockgen -destination=../mocks/service/mockUserService.go -package=service github.com/ferrandinand/cwh-api/service UserService
//go:generate mockgen -destination=../mocks/domain/mockUserRepository.go -package=domain github.com/ferrandinand/cwh-api/domain UserRepository
type UserService interface {
	GetAllUser(string) ([]dto.UserResponse, *errs.AppError)
	GetUser(string) (*dto.UserResponse, *errs.AppError)
	GetUserByUsername(string) (*dto.UserResponse, *errs.AppError)
	NewUser(request dto.NewUserRequest) (*dto.UserResponse, *errs.AppError)
	UpdateUser(userId string, request dto.UserRequest) (*dto.UserResponse, *errs.AppError)
	DeleteUser(userId string) (*dto.UserResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (s DefaultUserService) GetAllUser(status string) ([]dto.UserResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	users, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.UserResponse, 0)
	for _, c := range users {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultUserService) GetUser(id string) (*dto.UserResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultUserService) GetUserByUsername(username string) (*dto.UserResponse, *errs.AppError) {

	c, err := s.repo.ByUsername(username)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultUserService) NewUser(req dto.NewUserRequest) (*dto.UserResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	//New user
	user := domain.NewUser(req.Name, req.LastName, req.Password, req.Role, req.Email)
	newUser, err := s.repo.NewUser(user)
	if err != nil {
		return nil, err
	}
	response := newUser.ToDto()
	return &response, nil
}

func (s DefaultUserService) UpdateUser(userId string, req dto.UserRequest) (*dto.UserResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	//New user
	user := domain.NewUser(req.Name, req.LastName, "", req.Role, req.Email)
	updateUser, err := s.repo.UpdateUser(userId, user)
	if err != nil {
		return nil, err
	}
	response := updateUser.ToDto()
	return &response, nil
}

func (s DefaultUserService) DeleteUser(userId string) (*dto.UserResponse, *errs.AppError) {
	//Delete user
	deleteUser, err := s.repo.DeleteUser(userId)
	if err != nil {
		return nil, err
	}
	response := deleteUser.ToDto()
	return &response, nil
}

func NewUserService(repository domain.UserRepository) DefaultUserService {
	return DefaultUserService{repository}
}
