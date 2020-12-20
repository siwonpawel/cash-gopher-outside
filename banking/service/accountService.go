package service

import (
	"github.com/siwonpawel/cash-gopher-outside/banking/domain"
	"github.com/siwonpawel/cash-gopher-outside/banking/dto"
	"github.com/siwonpawel/cash-gopher-outside/banking/errs"
)

type DefaultAccountService struct {
	repo domain.AccountRepository
}

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account, err := domain.FromNewAccountRequest(req)
	if err != nil {
		return nil, err
	}

	savedAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	return savedAccount.ToNewAccountResponseDto(), nil
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
