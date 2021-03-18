package domain

import (
	"database/sql"

	"github.com/ferrandinand/cwh-api/errs"
	"github.com/ferrandinand/cwh-api/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProjectRepositoryDb struct {
	client *sqlx.DB
}

func (d ProjectRepositoryDb) FindAll(status string) ([]Project, *errs.AppError) {
	var err error
	projects := make([]Project, 0)

	if status == "" {
		findAllSql := "select project_id, name, created_by, p.group, repo_url,p.attributes,activities,status from projects p"
		err = d.client.Select(&projects, findAllSql)
	} else {
		findAllSql := "select project_id, name, created_by, p.group, repo_url,p.attributes,activities,status from projects p where status = ?"
		err = d.client.Select(&projects, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying projects table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return projects, nil
}

func (d ProjectRepositoryDb) Save(p Project) (*Project, *errs.AppError) {
	sqlInsert := "INSERT INTO projects (name, created_by, group, repo_url,attributes,activities,status) values (?, ?, ?, ?, ?, ?, ?)"

	_, err := d.client.Exec(sqlInsert, p.Name, p.CreatedBy, p.Group, p.RepoURL, p.Attributes, p.Activities, p.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &p, nil
}

func (d ProjectRepositoryDb) ById(id string) (*Project, *errs.AppError) {
	projectSql := "select p.project_id, p.name, p.created_by, p.group, p.repo_url,p.attributes,p.activities,p.status from projects p where project_id = ?"

	var p Project
	err := d.client.Get(&p, projectSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Project not found")
		} else {
			logger.Error("Error while scanning user " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &p, nil
}

/**
 * environment = make an entry in the environment table + update the balance in the projects table
 */
func (d ProjectRepositoryDb) SaveEnvironment(e Environment) (*Environment, *errs.AppError) {
	// inserting new env
	_, err := d.client.Exec(`INSERT INTO environments (name, project, atrributes) values (?, ?, ?)`, e.Name, e.Project, e.Attributes)
	if err != nil {
		logger.Error("Error while creating new environment: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &e, nil
}

// Find Environments by project id
func (d ProjectRepositoryDb) FindEnvironmentBy(project_id string) ([]Environment, *errs.AppError) {
	sqlGetProject := "SELECT environment_id, name, project, created_on, attributes from environments where project = ?"

	environments := make([]Environment, 0)
	err := d.client.Select(&environments, sqlGetProject, project_id)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return environments, nil
}

func NewProjectRepositoryDb(dbClient *sqlx.DB) ProjectRepositoryDb {
	return ProjectRepositoryDb{dbClient}
}
