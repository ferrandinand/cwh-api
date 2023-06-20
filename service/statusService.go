package service

import (
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/domain"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/dto"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/errs"
)

type StatusService interface {
	GetStatus(projectId string) (dto.StatusResponseList, *errs.AppError)
}

type DefaultStatusService struct {
	adapter domain.StatusAdapter
}

func (s DefaultStatusService) GetStatus(projectId string) (dto.StatusResponseList, *errs.AppError) {
	var response dto.StatusResponseList

	status, err := s.adapter.GetStatus(projectId)
	if err != nil {
		return response, err
	}

	statusItems := make([]dto.StatusResponse, 0)
	for _, c := range status.Items {
		statusItems = append(statusItems, c.ToDto())
	}
	response.Items = statusItems

	return response, nil
}

func NewStatusService(apiAdapter domain.GatusAPIAdapter) DefaultStatusService {
	return DefaultStatusService{apiAdapter}
}
