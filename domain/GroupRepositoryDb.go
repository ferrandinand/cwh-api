package domain

import (
	"database/sql"
	"strconv"

	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/ferrandinand/cwh-lib/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type GroupRepositoryDb struct {
	client *sqlx.DB
}

//Belongs to group repository
func (d GroupRepositoryDb) SaveGroup(g Group) (*Group, *errs.AppError) {

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

func (d GroupRepositoryDb) FindAll(status string, pageId int) (GroupList, *errs.AppError) {
	var err error
	var groups GroupList

	if status == "" {
		findAllSql := "select group_id, name, created_on, g.attributes FROM `groups` g WHERE group_id > ? ORDER BY group_id LIMIT ?"
		err = d.client.Select(&groups.Items, findAllSql, strconv.Itoa(pageId), pageSize+1)
	} else {
		findAllSql := "select group_id, name, created_on, g.attributes, status FROM `groups` g WHERE g.group_id > ? AND status = ? ORDER BY group_id LIMIT ?"
		err = d.client.Select(&groups.Items, findAllSql, strconv.Itoa(pageId), status, pageSize+1)
	}
	if err != nil {
		logger.Error("Error while querying groups table " + err.Error())
		return groups, errs.NewUnexpectedError("Unexpected database error")
	}

	if len(groups.Items) == pageSize+1 {
		groups.NextPageID = groups.Items[len(groups.Items)-1].GroupId
		groups.Items = groups.Items[:pageSize]
	}

	return groups, nil
}

func (d GroupRepositoryDb) ById(id string) (*Group, *errs.AppError) {
	groupSql := "select p.group_id, p.name, p.created_on, p.attributes from groups p where group_id = ?"

	var p Group
	err := d.client.Get(&p, groupSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Group not found")
		} else {
			logger.Error("Error while scanning group " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &p, nil
}
