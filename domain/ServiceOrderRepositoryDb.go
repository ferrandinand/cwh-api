package domain

import (
	"database/sql"
	"encoding/json"

	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/ferrandinand/cwh-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ServiceOrderRepositoryDb struct {
	client *sqlx.DB
}

func (d ServiceOrderRepositoryDb) FindAll(project string, environment string) ([]ServiceOrder, *errs.AppError) {
	var err error
	serviceOrders := make([]ServiceOrder, 0)

	findAllSql := "select service_order_id, service, environment, project, created_by, status, attributes from service_orders WHERE project = ? AND environment = ? AND status = 1"
	err = d.client.Select(&serviceOrders, findAllSql, project, environment)
	if err != nil {
		logger.Error("Error while querying service_orders table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return serviceOrders, nil
}

func (d ServiceOrderRepositoryDb) Save(p ServiceOrder) (*ServiceOrder, *errs.AppError) {
	sqlInsert := "INSERT INTO service_orders (service, environment, project, created_by, status, attributes) values (?, ?, ?, ?, ?, ?)"

	attributes_json, err := json.Marshal(p.Attributes)
	if err != nil {
		logger.Error("Error while creating new serviceOrder: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	_, err = d.client.Exec(sqlInsert, p.Service, p.Environment, p.Project, p.CreatedBy, p.Status, attributes_json)
	if err != nil {
		logger.Error(string(p.Service + p.Environment + p.Project))
		logger.Error("Error while creating new serviceOrder: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &p, nil
}

func (d ServiceOrderRepositoryDb) ById(id string) (*ServiceOrder, *errs.AppError) {
	serviceOrderSql := "select service_order_id, service, environment, project, created_by, status, attributes from service_orders where service_order_id = ? AND status=1"

	var p ServiceOrder
	err := d.client.Get(&p, serviceOrderSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Service Order not found")
		} else {
			logger.Error("Error while scanning user " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &p, nil
}

func NewServiceOrderRepositoryDb(dbClient *sqlx.DB) ServiceOrderRepositoryDb {
	return ServiceOrderRepositoryDb{dbClient}
}
