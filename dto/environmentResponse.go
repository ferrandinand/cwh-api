package dto

type EnvironmentResponse struct {
	EnvironmentId string `db:"environment_id"`
	Name          string `db:"name"`
	Project       string
	CreatedOn     string `db:"created_on"`
	Attributes    string
}
