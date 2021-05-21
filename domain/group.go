package domain

import (
	"time"

	"github.com/ferrandinand/cwh-api/dto"
)

type Group struct {
	GroupId    int    `db:"group_id"`
	Name       string `db:"name"`
	CreatedOn  string `db:"created_on"`
	Attributes string
}

func (g Group) ToDto() dto.GroupRequest {
	return dto.GroupRequest{
		Name:      g.Name,
		CreatedOn: g.CreatedOn,
	}
}

func NewGroup(name string) Group {
	return Group{
		Name:      name,
		CreatedOn: time.Now().Format("2006-01-02 15:04:05"),
	}
}
