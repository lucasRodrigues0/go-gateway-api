package service

import (
	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"github.com/devfullcycle/imersao22/go-gateway/internal/dto"
	"github.com/devfullcycle/imersao22/go-gateway/internal/repository"
)

type AccountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (s *AccountService) CreateAccount(input *dto.CreateAccountInput) (*dto.AccountOutput, error) {

	account := dto.ToAccount(input)

	existingAccount, err := s.accountRepository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicatedApiKey
	}

	err = s.accountRepository.Save(account)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.accountRepository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)

	err = s.accountRepository.UpdateBalance(account)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.accountRepository.FindByAPIKey(apiKey)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}

func (s *AccountService) FindById(id string) (*dto.AccountOutput, error) {
	account, err := s.accountRepository.FindByAPIKey(id)

	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}
