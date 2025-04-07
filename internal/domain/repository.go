package domain

type AccountRepository interface {
	// CreateAccount(name, email string) (*Account, error)
	// GetAccountByAPIKey(apiKey string) (*Account, error)
	// UpdateAccount(account *Account) error
	Save(account *Account)
	FindByAPIKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	Update(account *Account) error
}
