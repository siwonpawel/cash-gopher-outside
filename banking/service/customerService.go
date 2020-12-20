package service

import (
	"github.com/siwonpawel/cash-gopher-outside/banking/domain"
	"github.com/siwonpawel/cash-gopher-outside/banking/dto"
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]*dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]*dto.CustomerResponse, *errs.AppError) {

	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	customersDto := make([]*dto.CustomerResponse, 0, len(customers))

	for _, c := range customers {
		customersDto = append(customersDto, c.ToDto())
	}

	return customersDto, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	return c.ToDto(), nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
