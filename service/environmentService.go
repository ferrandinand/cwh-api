package service

import (
	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
)

//go:generate mockgen -destination=../mocks/service/mockEnvironmentService.go -package=service github.com/ferrandinand/cwh-api/service EnvironmentService
//go:generate mockgen -destination=../mocks/domain/mockEnvironmentRepository.go -package=domain github.com/ferrandinand/cwh-api/domain EnvironmentRepository
type EnvironmentService interface {
	NewEnvironment(request dto.NewEnvironmentRequest) (*dto.EnvironmentResponse, *errs.AppError)
	GetEnvironment(string) (*dto.EnvironmentResponse, *errs.AppError)
	GetAllEnvironment(int, string) ([]dto.EnvironmentResponse, *errs.AppError)
}

type DefaultEnvironmentService struct {
	repo domain.EnvironmentRepository
}

func (s DefaultEnvironmentService) NewEnvironment(req dto.NewEnvironmentRequest) (*dto.EnvironmentResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	//New environment
	environment := domain.NewEnvironment(req.Name, req.Project)
	newEnvironment, err := s.repo.Save(environment)
	if err != nil {
		return nil, err
	}
	response := newEnvironment.ToDto()
	return &response, nil
}

func (s DefaultEnvironmentService) GetAllEnvironment(project int, status string) ([]dto.EnvironmentResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = "1"
	}
	environments, err := s.repo.FindAll(project, status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.EnvironmentResponse, 0)
	for _, c := range environments {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultEnvironmentService) GetEnvironment(id string) (*dto.EnvironmentResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewEnvironmentService(repository domain.EnvironmentRepository) DefaultEnvironmentService {
	return DefaultEnvironmentService{repository}
}
