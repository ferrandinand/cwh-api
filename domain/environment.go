package domain

import (
	"time"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
)

type Environment struct {
	EnvironmentId int    `db:"environment_id"`
	Name          string `db:"name"`
	Project       int
	CreatedOn     string `db:"created_on"`
	Status        string
	Attributes    JSONField
}

type EnvironmentList struct {
	Items      []Environment `json:"items"`
	NextPageID int           `json:"next_page_id,omitempty" example:"10"`
}

type EnvironmentRepository interface {
	FindAll(project int, status string, pageId int) (EnvironmentList, *errs.AppError)
	ById(string) (*Environment, *errs.AppError)
	Save(environment Environment) (*Environment, *errs.AppError)
}

func NewEnvironment(name string, project int) Environment {
	var jsonEmpty map[string]interface{}

	return Environment{
		Name:       name,
		Project:    project,
		CreatedOn:  time.Now().Format("2006-01-02 15:04:05"),
		Attributes: jsonEmpty,
	}
}

func (p Environment) statusAsText() string {
	status, _ := commonStatusAsText(p.Status)
	return status
}

func (t Environment) ToDto() dto.EnvironmentResponse {
	return dto.EnvironmentResponse{
		EnvironmentId: t.EnvironmentId,
		Name:          t.Name,
		Project:       t.Project,
		CreatedOn:     t.CreatedOn,
		Attributes:    t.Attributes,
		Status:        t.statusAsText(),
	}
}
