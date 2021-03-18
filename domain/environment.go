package domain

import "github.com/ferrandinand/cwh-api/dto"

type Environment struct {
	EnvironmentId string `db:"environment_id"`
	Name          string `db:"name"`
	Project       string
	CreatedOn     string `db:"created_on"`
	Attributes    string
}

func (t Environment) ToDto() dto.EnvironmentResponse {
	return dto.EnvironmentResponse{
		EnvironmentId: t.EnvironmentId,
		Name:          t.Name,
		Project:       t.Project,
		CreatedOn:     t.CreatedOn,
		Attributes:    t.Attributes,
	}
}
