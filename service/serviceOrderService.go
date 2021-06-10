package service

import (
	"github.com/ferrandinand/cwh-api/domain"
	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
)

//go:generate mockgen -destination=../mocks/service/mockServiceOrder.go -package=service github.com/ferrandinand/cwh-api/service ServiceOrder
//go:generate mockgen -destination=../mocks/domain/mockServiceOrderRepository.go -package=domain github.com/ferrandinand/cwh-api/domain ServiceOrderRepository
type ServiceOrder interface {
	NewServiceOrder(request dto.NewServiceOrderRequest) (*dto.ServiceOrderResponse, *errs.AppError)
	GetEnvironmentServiceOrders(request dto.ServiceOrderRequest) ([]dto.ServiceOrderResponse, *errs.AppError)
	GetServiceOrder(string) (*dto.ServiceOrderResponse, *errs.AppError)
}

type DefaultServiceOrder struct {
	repo domain.ServiceOrderRepository
}

func (s DefaultServiceOrder) NewServiceOrder(req dto.NewServiceOrderRequest) (*dto.ServiceOrderResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	serviceOrder := domain.NewServiceOrder(req.Service, req.Environment, req.Project, req.CreatedBy)
	newServiceOrder, err := s.repo.Save(serviceOrder)
	if err != nil {
		return nil, err
	}
	response := newServiceOrder.ToDto()
	return &response, nil
}

func (s DefaultServiceOrder) GetEnvironmentServiceOrders(req dto.ServiceOrderRequest) ([]dto.ServiceOrderResponse, *errs.AppError) {
	services, err := s.repo.FindAll(req.Project, req.Environment)
	if err != nil {
		return nil, err
	}
	response := make([]dto.ServiceOrderResponse, 0)
	for _, c := range services {
		response = append(response, c.ToDto())
	}
	return response, err
}

func (s DefaultServiceOrder) GetServiceOrder(id string) (*dto.ServiceOrderResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewServiceOrderService(repository domain.ServiceOrderRepository) DefaultServiceOrder {
	return DefaultServiceOrder{repository}
}
