package domain

import "github.com/ferrandinand/cwh-api/dto"

const WITHDRAWAL = "withdrawal"

type Environment struct {
	EnvironmentId string `db:"environment_id"`
	Name          string `db:"name"`
	Project       string
	CreatedOn     string `db:"created_on"`
	Attributes    string
}

func (t Environment) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (t Environment) ToDto() dto.EnvironmentResponse {
	return dto.TransactionResponse{
		EnvironmentId: t.EnvironmentId,
		Name:          t.Name,
		Project:       t.Project,
		CreatedOn:     t.CreatedOn,
		Attributes:    t.Attributes,
	}
}
