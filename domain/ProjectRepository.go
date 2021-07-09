package domain

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/ferrandinand/cwh-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
)

type ProjectRepositoryDb struct {
	client         *sqlx.DB
	rabbitmqClient *amqp.Connection
}

func (d ProjectRepositoryDb) FindAll(status string, pageId int) (ProjectList, *errs.AppError) {
	var err error
	var projects ProjectList

	if status == "" {
		findAllSql := "select project_id, name, type, created_by, p.group, p.attributes,activities,status FROM projects p WHERE project_id > ? ORDER BY project_id LIMIT ?"
		err = d.client.Select(&projects.Items, findAllSql, strconv.Itoa(pageId), pageSize+1)
	} else {
		findAllSql := "select project_id, name, type, created_by, p.group, p.attributes,activities,status FROM projects p WHERE p.project_id > ? AND status = ? ORDER BY project_id LIMIT ?"
		err = d.client.Select(&projects.Items, findAllSql, strconv.Itoa(pageId), status, pageSize+1)
	}
	if err != nil {
		logger.Error("Error while querying projects table " + err.Error())
		return projects, errs.NewUnexpectedError("Unexpected database error")
	}

	if len(projects.Items) == pageSize+1 {
		projects.NextPageID = projects.Items[len(projects.Items)-1].Id
		projects.Items = projects.Items[:pageSize]
	}

	return projects, nil
}

func (d ProjectRepositoryDb) Save(p Project) (*Project, *errs.AppError) {
	sqlInsert := "INSERT INTO projects (name, type, created_by, `group`, attributes,activities,status) values (?, ?, ?, ?, ?, ?, ?)"

	attributes_json, err := json.Marshal(p.Attributes)
	if err != nil {
		logger.Error("Error while creating new project: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	activities_json, err := json.Marshal(p.Activities)
	if err != nil {
		logger.Error("Error while creating new project: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	_, err = d.client.Exec(sqlInsert, p.Name, p.Type, p.CreatedBy, p.Group, attributes_json, activities_json, p.Status)
	if err != nil {
		logger.Error("Error while creating new project: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &p, nil
}

func (d ProjectRepositoryDb) ById(id string) (*Project, *errs.AppError) {
	projectSql := "select p.project_id, p.name, p.created_by, p.group, p.attributes,p.activities,p.status from projects p where project_id = ?"

	var p Project
	err := d.client.Get(&p, projectSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Project not found")
		} else {
			logger.Error("Error while scanning project " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &p, nil
}

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

func (d ProjectRepositoryDb) DeleteProject(projectId string) (*Project, *errs.AppError) {

	projectSql := "select project_id, name, created_by, `group`, attributes,activities, status from projects where project_id = ?"

	var p Project
	err := d.client.Get(&p, projectSql, projectId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Project not found")
		} else {
			logger.Error("Error while scanning project " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	sqlUpdate := "UPDATE projects SET status=0 WHERE project_id=?"
	_, err = d.client.Exec(sqlUpdate, projectId)
	if err != nil {
		logger.Error("Error while deleting project: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	return &p, nil
}

func (d ProjectRepositoryDb) PublishProject(project Project) *errs.AppError {

	channel, err := d.rabbitmqClient.Channel()
	if err != nil {
		logger.Error("Error while creating rabbit mqchannel: " + err.Error())
		return errs.NewUnexpectedError("Unexpected error.")
	}

	_, err = channel.QueueDeclare(
		project.Type, // queue name
		true,         // durable
		false,        // auto delete
		false,        // exclusive
		false,        // no wait
		nil,          // arguments
	)
	if err != nil {
		logger.Error("Error while declaring queue: " + err.Error())
		return errs.NewUnexpectedError("Unexpected error.")
	}

	projectJSON, err := json.Marshal(project)
	if err != nil {
		logger.Error("Error while Marshalling project to json: " + err.Error())
	}
	// Create a message to publish.
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(projectJSON),
	}

	// Attempt to publish a message to the queue.
	if err := channel.Publish(
		"",           // exchange
		project.Type, // queue name
		false,        // mandatory
		false,        // immediate
		message,      // message to publish
	); err != nil {
		logger.Error("Error while publishing in queue: " + err.Error())
		return errs.NewUnexpectedError("Unexpected error.")
	}
	logger.Info(project.Name)

	return nil

}

func NewProjectRepositoryDb(dbClient *sqlx.DB, rabbitmqClient *amqp.Connection) ProjectRepositoryDb {
	return ProjectRepositoryDb{dbClient, rabbitmqClient}
}
