package dto

type ProjectResponse struct {
	Id         string `db:"project_id"`
	Name       string
	CreatedBy  string `db:"created_by"`
	CreatedOn  string `db:"created_on"`
	Group      int
	RepoURL    string `db:"repo_url"`
	Attributes string
	Activities string
	Status     int
}
