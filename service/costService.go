package service

import (
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/domain"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/dto"
	"bitbucket.org/iccgit/icc-cwh-backstage/cwh-api/lib/errs"
)

type CostService interface {
	GetCost(projectId string) (dto.CostResponseList, *errs.AppError)
	GetAllCost(ProjectpageId int) (dto.CostResponseList, *errs.AppError)
}

type DefaultCostService struct {
	repo domain.CostRepository
}

func (s DefaultCostService) GetCost(id string) (dto.CostResponseList, *errs.AppError) {
	var response dto.CostResponseList

	costs, err := s.repo.ById(id)
	if err != nil {
		return response, err
	}
	costsResponseItems := make([]dto.CostResponse, 0)
	for _, c := range costs.Items {
		costsResponseItems = append(costsResponseItems, c.ToDto())
	}

	response.Items = costsResponseItems
	return response, nil
}

func (s DefaultCostService) GetAllCost(pageId int) (dto.CostResponseList, *errs.AppError) {
	var response dto.CostResponseList

	costs, err := s.repo.FindAll(pageId)
	if err != nil {
		return response, err
	}

	costsResponseItems := make([]dto.CostResponse, 0)
	for _, c := range costs.Items {
		costsResponseItems = append(costsResponseItems, c.ToDto())
	}

	response.NextPageID = costs.NextPageID
	response.Items = costsResponseItems

	return response, err
}

func NewCostService(repository domain.CostRepository) DefaultCostService {
	return DefaultCostService{repository}
}
