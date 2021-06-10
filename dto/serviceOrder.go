package dto

import (
	"github.com/ferrandinand/cwh-lib/errs"
)

type NewServiceOrderRequest struct {
	Service     string
	Environment string
	Project     string
	CreatedBy   string
}

type ServiceOrderRequest struct {
	Service     string
	Environment string
	Project     string
	CreatedBy   string
}

type ServiceOrderResponse struct {
	Id          string
	Service     string
	Environment string
	Project     string
	CreatedBy   string
	CreatedOn   string
	Attributes  map[string]interface{}
}

func (r NewServiceOrderRequest) Validate() *errs.AppError {
	return nil
}
