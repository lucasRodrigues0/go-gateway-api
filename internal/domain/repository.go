package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	Update(account *Account) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindById(id string) (*Invoice, error)
	FindByAccountID(AccountID string) (*[]Invoice, error)
	UpdateStatus(invoice *Invoice) error
}
