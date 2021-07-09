package service

import (
	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
)

type GroupService interface {
	GetAllGroup(status string, pageId int) (dto.GroupResponseList, *errs.AppError)
	GetGroup(string) (*dto.GroupResponse, *errs.AppError)
}

type DefaultGroupService struct {
	repo domain.GroupRepository
}

func (s DefaultGroupService) GetAllGroup(status string, pageId int) (dto.GroupResponseList, *errs.AppError) {
	var response dto.GroupResponseList

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = "1"
	}
	groups, err := s.repo.FindAll(status, pageId)
	if err != nil {
		return response, err
	}

	groupResponseItems := make([]dto.GroupResponse, 0)
	for _, c := range groups.Items {
		groupResponseItems = append(groupResponseItems, c.ToDto())
	}

	response.NextPageID = groups.NextPageID
	response.Items = groupResponseItems

	return response, err
}

func (s DefaultGroupService) GetGroup(id string) (*dto.GroupResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewGroupService(repository domain.GroupRepository) DefaultGroupService {
	return DefaultGroupService{repository}
}
