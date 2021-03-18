package domain

import (
	"database/sql"

	"github.com/ferrandinand/cwh-api/errs"
	"github.com/ferrandinand/cwh-api/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	client *sqlx.DB
}

func (d UserRepositoryDb) FindAll(status string) ([]User, *errs.AppError) {
	var err error
	users := make([]User, 0)

	if status == "" {
		findAllSql := "select user_id, name,created_on, role, email, attributes, status from users"
		err = d.client.Select(&users, findAllSql)
	} else {
		findAllSql := "select user_id, name,created_on, role, email, attributes, status from users where status = ?"
		err = d.client.Select(&users, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying users table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return users, nil
}

func (d UserRepositoryDb) ById(id string) (*User, *errs.AppError) {
	userSql := "select user_id, name,created_on, role, email, attributes,status from users where user_id = ?"

	var c User
	err := d.client.Get(&c, userSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User not found")
		} else {
			logger.Error("Error while scanning user " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewUserRepositoryDb(dbClient *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{dbClient}
}
