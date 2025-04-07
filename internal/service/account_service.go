package service

import "github.com/devfullcycle/imersao22/go-gateway/internal/repository"

type AccountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}
