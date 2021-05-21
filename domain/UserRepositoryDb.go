package domain

import (
	"database/sql"
	"encoding/json"

	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/ferrandinand/cwh-lib/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	client *sqlx.DB
}

func (d UserRepositoryDb) FindAll(status string) ([]User, *errs.AppError) {
	var err error
	users := make([]User, 0)

	findAllSql := "select user_id, name,created_on, role, email, attributes, status from users where status = 1"
	err = d.client.Select(&users, findAllSql)
	if err != nil {
		logger.Error("Error while querying users table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return users, nil
}

func (d UserRepositoryDb) ById(userId string) (*User, *errs.AppError) {
	userSql := "SELECT user_id, name,created_on, role, email, attributes,status from users WHERE user_id = ? AND status=1"

	var c User
	err := d.client.Get(&c, userSql, userId)
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

func (d UserRepositoryDb) NewUser(u User) (*User, *errs.AppError) {

	sqlInsert := "INSERT INTO users (name, password, created_on,role,email, attributes, status) values (?, ?, ?, ?, ?, ?, 1)"

	cyptedPass := setUserPassword(u.Password)
	attributes_json, err := json.Marshal(u.Attributes)
	if err != nil {
		logger.Error("Error while creating new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	_, err = d.client.Exec(sqlInsert, u.Name, cyptedPass, u.CreatedOn, u.Role, u.Email, attributes_json, u.Status)
	if err != nil {
		logger.Error("Error while creating new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	return &u, nil
}

func (d UserRepositoryDb) UpdateUser(userId string, u User) (*User, *errs.AppError) {

	sqlUpdate := "UPDATE users SET name=?, role=?, email=?, attributes=?, status=? WHERE user_id=? AND status=1"

	attributes_json, err := json.Marshal(u.Attributes)
	if err != nil {
		logger.Error("Error while updating user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	_, err = d.client.Exec(sqlUpdate, u.Name, u.Role, u.Email, attributes_json, u.Status, userId)
	if err != nil {
		logger.Error("Error while updating user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	u.Id = userId

	return &u, nil
}

func (d UserRepositoryDb) DeleteUser(userId string) (*User, *errs.AppError) {

	userSql := "select user_id, name,created_on, role, email, attributes,status from users where user_id = ?"

	var u User
	err := d.client.Get(&u, userSql, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("User not found")
		} else {
			logger.Error("Error while scanning user " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	sqlUpdate := "UPDATE users SET status=0 WHERE user_id=?"
	_, err = d.client.Exec(sqlUpdate, userId)
	if err != nil {
		logger.Error("Error while deleting user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &u, nil
}

func NewUserRepositoryDb(dbClient *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{dbClient}
}
