package domain

import (
	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/ferrandinand/cwh-lib/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type EnvironmentRepositoryDb struct {
	client *sqlx.DB
}

func (d EnvironmentRepositoryDb) FindAll(projecId int, status string, pageId int) (EnvironmentList, *errs.AppError) {
	var environments EnvironmentList

	sqlGetEnvironment := "SELECT environment_id, name, project, created_on, attributes from environments where project = ? AND environment_id > ? AND status = ? ORDER BY environment_id LIMIT ?"

	err := d.client.Select(&environments.Items, sqlGetEnvironment, projecId, pageId, status, pageSize+1)
	if err != nil {
		logger.Error("Error while fetching environments information: " + err.Error())
		return environments, errs.NewUnexpectedError("Unexpected database error")
	}

	if len(environments.Items) == pageSize+1 {
		environments.NextPageID = environments.Items[len(environments.Items)-1].EnvironmentId
		environments.Items = environments.Items[:pageSize]
	}

	return environments, nil
}

// Find Environments by environment id
func (d EnvironmentRepositoryDb) ById(environment_id string) (*Environment, *errs.AppError) {
	sqlGetEnvironment := "SELECT environment_id, name, project, created_on, attributes from environments where environment_id = ?"

	var e Environment
	err := d.client.Get(&e, sqlGetEnvironment, environment_id)

	if err != nil {
		logger.Error("Error while fetching environment information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &e, nil
}

func (d EnvironmentRepositoryDb) Save(e Environment) (*Environment, *errs.AppError) {
	// inserting new env
	_, err := d.client.Exec(`INSERT INTO environments (name, project) values (?, ?)`, e.Name, e.Project)
	if err != nil {
		logger.Error("Error while creating new environment: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &e, nil
}

func NewEnvironmentRepositoryDb(dbClient *sqlx.DB) EnvironmentRepositoryDb {
	return EnvironmentRepositoryDb{dbClient}
}
