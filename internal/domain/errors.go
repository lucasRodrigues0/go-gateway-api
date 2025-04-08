package domain

import "errors"

var (
	ErrAccountNotFound    = errors.New("account not found")
	ErrDuplicatedApiKey   = errors.New("duplicated api key")
	ErrInvoiceNotFound    = errors.New("invoice not found")
	ErrUnauthorizedAccess = errors.New("unauthorized access")
	ErrInvalidAmount      = errors.New("invalid amount")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidStatus      = errors.New("invalid status")
)
