package domain

import (
	"context"
	"time"
)

type Currency struct {
	Name    string    `json:"name" db:"name" `
	Value   float32   `json:"value" db:"value" `
	LastUpd time.Time `json:"last_updated" db:"last_updated" `
}

type CurrencyRepository interface {
	GetCurrency(ctx context.Context, name string) (Currency, error)
	SetCurrency(ctx context.Context, c Currency) error
}

type CurrencyUsecase interface {
	AvailableCurrencies() []string
	GetCurrency(ctx context.Context, name string) (Currency, error)
	SetCurrency(ctx context.Context, currency Currency) error
}
