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
	GetAllEnvironment(projectId int, status string, pageId int) (dto.EnvironmentResponseList, *errs.AppError)
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

func (s DefaultEnvironmentService) GetAllEnvironment(project int, status string, pageId int) (dto.EnvironmentResponseList, *errs.AppError) {
	var response dto.EnvironmentResponseList

	status, statusError := statusToNumber(status)
	if statusError != nil {
		return response, errs.NewValidationError(statusError.Error())
	}

	environments, err := s.repo.FindAll(project, status, pageId)
	if err != nil {
		return response, err
	}

	environmentResponseItems := make([]dto.EnvironmentResponse, 0)
	for _, c := range environments.Items {
		environmentResponseItems = append(environmentResponseItems, c.ToDto())
	}

	response.NextPageID = environments.NextPageID
	response.Items = environmentResponseItems

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
