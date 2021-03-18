package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-api/errs"
)

type JSONField map[string]interface{}

type Project struct {
	Id         string `db:"project_id"`
	Name       string
	CreatedBy  string `db:"created_by"`
	CreatedOn  string `db:"created_on"`
	Group      string
	RepoURL    string `db:"repo_url"`
	Attributes JSONField
	Activities JSONField `db:"activities"`
	Status     string
}

func (p Project) statusAsText() string {
	statusAsText := "active"
	if p.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (p Project) ToDto() dto.ProjectResponse {
	return dto.ProjectResponse{
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

func NewProject(name string, user string, group string, repoURL string) Project {
	var jsonEmpty map[string]interface{}

	return Project{
		Name:       name,
		CreatedBy:  user,
		Group:      group,
		RepoURL:    repoURL,
		Attributes: jsonEmpty,
		Activities: jsonEmpty,
		Status:     "1",
	}
}

func (jf *JSONField) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &jf)
		return nil
	case string:
		json.Unmarshal([]byte(v), &jf)
		return nil
	default:
		return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
func (jf *JSONField) Value() (driver.Value, error) {
	return json.Marshal(jf)
}
