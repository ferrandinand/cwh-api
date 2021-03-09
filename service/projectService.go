package service

import (
	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/errs"
)

//go:generate mockgen -destination=../mocks/service/mockProjectService.go -package=service github.com/ferrandinand/cwh-api/service ProjectService
type ProjectService interface {
	GetAllProject(string) ([]dto.ProjectResponse, *errs.AppError)
	GetProject(string) (*dto.ProjectResponse, *errs.AppError)
	GetAllEnvironment(string) ([]dto.ProjectResponse, *errs.AppError)
}

type DefaultProjectService struct {
	repo domain.ProjectRepository
}

func (s DefaultProjectService) NewProject(req dto.NewProjectRequest) (*dto.NewProjectResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	project := domain.NewProject(name, user, group, repoURL)
	if newProject, err := s.repo.Save(account); err != nil {
		return nil, err
	} else {
		return newProject.ToNewProjectResponseDto(), nil
	}
}

func (s DefaultProjectService) GetAllProject(status string) ([]dto.ProjectResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	projects, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.ProjectResponse, 0)
	for _, c := range projects {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultProjectService) GetProject(id string) (*dto.ProjectResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultProjectService) GetAllEnvironments(project_id string) ([]dto.EnvironmentResponse, *errs.AppError) {
	environments, err := s.repo.FindEnvironmentBy(project_id)
	if err != nil {
		return nil, err
	}

	response := make([]dto.EnvironmentResponse, 0)
	for _, c := range environments {
		response = append(response, c.ToDto())
	}

	return response, err
}

func NewProjectService(repository domain.ProjectRepository) DefaultProjectService {
	return DefaultProjectService{repository}
}
