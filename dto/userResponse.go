package dto

type UserResponse struct {
	Id         string `db:"user_id"`
	Name       string
	CreatedOn  string `db:"created_on"`
	Role       string
	Email      string
	Attributes map[string]interface{}
	Status     string
}
