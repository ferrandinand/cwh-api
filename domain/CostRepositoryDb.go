package domain

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/errs"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/logger"
)

type CostRepositoryDb struct {
	client *sqlx.DB
}

func (c CostRepositoryDb) FindAll(pageId int) (CostList, *errs.AppError) {
	var err error
	var costs CostList

	findAllSql := `project_name,installation,project,cluster,reported_on,tenant,budget,
cpu,ebs_cost,ebs_gb,efs_cost,efs_gb,elb_cost,elb_gb,mem_gb,pods,pods_cost,rds_cost,s3_cost,s3_gb,secrets_cost,
transfer_cost,others_cost,total FROM costs WHERE reported_on = ? LIMIT ?`

	err = c.client.Select(&costs.Items, findAllSql, strconv.Itoa(pageId), pageSize+1)

	if err != nil {
		logger.Error("Error while querying costs table " + err.Error())
		return costs, errs.NewUnexpectedError("Unexpected database error")
	}

	if len(costs.Items) == pageSize+1 {
		costs.NextPageID = costs.Items[len(costs.Items)-1].Id
		costs.Items = costs.Items[:pageSize]
	}

	return costs, nil
}

func (c CostRepositoryDb) ById(id string) (CostList, *errs.AppError) {
	var cl CostList

	vulnerabilitySql := `SELECT project_name,installation,project,cluster,reported_on,tenant,budget,
cpu,ebs_cost,ebs_gb,efs_cost,efs_gb,elb_cost,elb_gb,mem_gb,pods,pods_cost,rds_cost,s3_cost,s3_gb,secrets_cost,transfer_cost,others_cost,total
FROM costs c
WHERE reported_on = (
  SELECT MAX(reported_on)
  FROM costs
  WHERE project = ?
)
AND c.project = ?
ORDER BY c.reported_on ASC`

	err := c.client.Select(&cl.Items, vulnerabilitySql, id, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return cl, errs.NewNotFoundError("Resource Usage not found")
		} else {
			logger.Error("Error while scanning project " + err.Error())
			return cl, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return cl, nil
}
