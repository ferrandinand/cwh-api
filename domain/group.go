package domain

import (
	"time"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
	"github.com/jmoiron/sqlx"
)

type Group struct {
	GroupId    int    `db:"group_id"`
	Name       string `db:"name"`
	CreatedOn  string `db:"created_on"`
	Attributes string
	Status     string
}

func (g Group) statusAsText() string {
	statusAsText := "active"
	if g.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

type GroupList struct {
	Items      []Group `json:"items"`
	NextPageID int     `json:"next_page_id,omitempty" example:"10"`
}

type GroupRepository interface {
	FindAll(status string, pageId int) (GroupList, *errs.AppError)
	ById(string) (*Group, *errs.AppError)
	SaveGroup(group Group) (*Group, *errs.AppError)
}

func (g Group) ToDto() dto.GroupResponse {
	return dto.GroupResponse{
		Id:        g.GroupId,
		Name:      g.Name,
		CreatedOn: g.CreatedOn,
		Status:    g.statusAsText(),
	}
}

func NewGroup(name string) Group {
	return Group{
		Name:      name,
		CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func NewGroupRepositoryDb(dbClient *sqlx.DB) GroupRepositoryDb {
	return GroupRepositoryDb{dbClient}
}
