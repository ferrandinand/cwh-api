package domain

import (
	"time"

	"github.com/ferrandinand/cwh-api/dto"
	"github.com/ferrandinand/cwh-lib/errs"
)

type ServiceOrder struct {
	Id          string `db:"service_order_id"`
	Service     string
	Environment string
	Project     string
	CreatedBy   string `db:"created_by"`
	CreatedOn   string `db:"created_on"`
	Attributes  JSONField
	Status      string
}

func (s ServiceOrder) statusAsText() string {
	status, _ := commonStatusAsText(s.Status)
	return status
}

func (s ServiceOrder) ToDto() dto.ServiceOrderResponse {
	return dto.ServiceOrderResponse{
		Id:          s.Id,
		Service:     s.Service,
		Environment: s.Environment,
		Project:     s.Project,
		CreatedBy:   s.CreatedBy,
		CreatedOn:   s.CreatedOn,
		Attributes:  s.Attributes,
	}
}

type ServiceOrderRepository interface {
	FindAll(project string, environment string) ([]ServiceOrder, *errs.AppError)
	ById(serviceOrderId string) (*ServiceOrder, *errs.AppError)
	Save(serviceOrder ServiceOrder) (*ServiceOrder, *errs.AppError)
}

func NewServiceOrder(service string, environment string, project string, user string) ServiceOrder {
	var jsonEmpty map[string]interface{}

	return ServiceOrder{
		Service:     service,
		Environment: environment,
		Project:     project,
		CreatedBy:   user,
		CreatedOn:   time.Now().Format("2006-01-02 15:04:05"),
		Attributes:  jsonEmpty,
		Status:      "2", //Created
	}
}
