package domain

import (
	"database/sql"

	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/ferrandinand/cwh-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type GroupRepositoryDb struct {
	client *sqlx.DB
}

//Belongs to project repository
func (d ProjectRepositoryDb) SaveGroup(g Group) (*Group, *errs.AppError) {

	sqlInsertGroup := "INSERT INTO `groups` (name) values (?)"

	logger.Info("Generating new group" + string(g.Name))
	_, err := d.client.Exec(sqlInsertGroup, g.Name)
	if err != nil {
		logger.Error("Error while creating new group: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	// Get latest group generated
	sqlGetGroup := "SELECT max(group_id) as group_id FROM `groups` g WHERE g.name=? LIMIT 1"
	err = d.client.Get(&g, sqlGetGroup, g.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Group not found")
		} else {
			logger.Error("Error while scanning group " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &g, nil
}
