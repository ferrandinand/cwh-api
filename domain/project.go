package domain

import (
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/errs"
)

type Project struct {
	Id         string `db:"project_id"`
	Name       string
	CreatedBy  string `db:"created_by"`
	CreatedOn  string `db:"created_on"`
	Group      int
	RepoURL    string `db:"repo_url"`
	Attributes string
	Activities string
	Status     int
}

func (p Project) statusAsText() string {
	statusAsText := "active"
	if p.Status == 0 {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (p Project) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:         p.Id,
		Name:       p.Name,
		CreatedBy:  p.CreatedBy,
		CreatedOn:  p.CreatedOn,
		Group:      p.Group,
		RepoURL:    p.RepoURL,
		Attributes: p.Attributes,
		Activities: p.Activities,
		Status:     p.statusAsText(),
	}
}

type ProjectRepository interface {
	FindAll(status string) ([]Project, *errs.AppError)
	ById(string) (*Project, *errs.AppError)
	Save(project Project) (*Project, *errs.AppError)
	SaveEnvironment(environment Environment) (*Environment, *errs.AppError)
	FindEnvironmentBy(projectId string) ([]Environment, *errs.AppError)
}

func NewProject(name string, user string, group int, repoURL string) Project {
	return Project{
		Name:       name,
		CreatedBy:  user,
		Group:      group,
		RepoURL:    repoURL,
		Attributes: "",
		Activities: "",
		Status:     1,
	}
}
