package dto

import (
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

type CreateAccountInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	APIKey    string    `json:"api_key,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccountOutput(input *domain.Account) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}
