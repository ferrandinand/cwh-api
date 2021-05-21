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

type EnvironmentRepository interface {
	FindAll(project int, status string) ([]Environment, *errs.AppError)
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
	statusAsText := "active"
	if p.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
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
