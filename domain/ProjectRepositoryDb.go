package domain

import (
	"strconv"

	"github.com/ferrandinand/cwh-api/errs"
	"github.com/ferrandinand/cwh-api/logger"
	"github.com/jmoiron/sqlx"
)

type ProjectRepositoryDb struct {
	client *sqlx.DB
}


Id        int `db:"project_id"`
Name      string
CreatedBy string `db:"created_by"`
CreatedOn string `db:"created_on"`
Group        int
RepoURL      string `db:"repo_url"`
Attributes   string
Activities   string
Status int

func (d ProjectRepositoryDb) Save(p Project) (*Project, *errs.AppError) {
	sqlInsert := "INSERT INTO projects (name, created_by, group, repo_url,attributes,activities,status) values (?, ?, ?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, p.Name, p.CreatedBy, p.Group, p.RepoURL, p.Attributes, p.Activities,p.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.ProjectId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d ProjectRepositoryDb) ById(id string) (*Project, *errs.AppError) {
	userSql := "select project_id, name, created_by, group, repo_url,attributes,activities,status from projects where project_id = ?"

	var p Project
	err := d.client.Get(&p, userSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Project not found")
		} else {
			logger.Error("Error while scanning user " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

/**
 * environment = make an entry in the environment table + update the balance in the projects table
 */
func (d ProjectRepositoryDb) SaveEnvironment(e Environment) (*Environment, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO environments (name, project, atrributes)
											values (?, ?, ?, ?)`, e.Name, e.Project, e.Attributes)

	// updating project Activities
	//if t.IsWithdrawal() {
	//	_, err = tx.Exec(`UPDATE projects SET amount = amount - ? where account_id = ?`, t.Amount, t.ProjectId)
	//} else {
	//	_, err = tx.Exec(`UPDATE projects SET amount = amount + ? where account_id = ?`, t.Amount, t.ProjectId)
	//}

	// in case of error Rollback, and changes from both the tables will be reverted
	//if err != nil {
	//	tx.Rollback()
	//	logger.Error("Error while saving transaction: " + err.Error())
	//	return nil, errs.NewUnexpectedError("Unexpected database error")
	//}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	//transactionId, err := result.LastInsertId()
	//if err != nil {
	//	logger.Error("Error while getting the last transaction id: " + err.Error())
	//	return nil, errs.NewUnexpectedError("Unexpected database error")
	//}
//
	//// Getting the latest account information from the projects table
	//account, appErr := d.FindBy(t.ProjectId)
	//if appErr != nil {
	//	return nil, appErr
	//}
	//t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	//t.Amount = account.Amount
	return &t, nil
}

// Find Environments by project id
func (d ProjectRepositoryDb) FindEnvironmentBy(project_id string) ([]Environment, *errs.AppError) {
	sqlGetProject := "SELECT environment_id, name, project, created_on, attributes from environments where project = ?"
	
	environments := make([]Environment, 0)
	err = d.client.Select(&environments, sqlGetProject, project_id)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &environments, nil
}


func NewProjectRepositoryDb(dbClient *sqlx.DB) ProjectRepositoryDb {
	return ProjectRepositoryDb{dbClient}
}
