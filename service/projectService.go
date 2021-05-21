package service

import (
	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
)

//go:generate mockgen -destination=../mocks/service/mockProjectService.go -package=service github.com/ferrandinand/cwh-api/service ProjectService
//go:generate mockgen -destination=../mocks/domain/mockProjectRepository.go -package=domain github.com/ferrandinand/cwh-api/domain ProjectRepository
type ProjectService interface {
	NewProject(request dto.NewProjectRequest) (*dto.ProjectResponse, *errs.AppError)
	GetAllProject(string) ([]dto.ProjectResponse, *errs.AppError)
	GetProject(string) (*dto.ProjectResponse, *errs.AppError)
}

type DefaultProjectService struct {
	repo domain.ProjectRepository
}

func (s DefaultProjectService) NewProject(req dto.NewProjectRequest) (*dto.ProjectResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	//New project
	project := domain.NewProject(req.Name, req.Type, req.CreatedBy, req.Group)
	newProject, err := s.repo.Save(project)
	if err != nil {
		return nil, err
	}
	response := newProject.ToDto()
	return &response, nil
}

func (s DefaultProjectService) GetAllProject(status string) ([]dto.ProjectResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = "1"
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

func NewProjectService(repository domain.ProjectRepository) DefaultProjectService {
	return DefaultProjectService{repository}
}
