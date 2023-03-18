package domain

import (
	"context"
	"time"
)

type Currency struct {
	Name    string    `json:"name"`
	Value   float32   `json:"value"`
	LastUpd time.Time `json:"last_updated"`
}

type CurrencyRepository interface {
	GetCurrency(ctx context.Context, name string) (Currency, error)
	SetCurrency(ctx context.Context, c Currency) error
	UpdateCurrency(ctx context.Context, c Currency) error
}

type CurrencyUsecase interface {
	GetCurrency(ctx context.Context, name string) (*Currency, error)
	SetCurrency(ctx context.Context, name, value string) error
	UpdateCurrency(ctx context.Context, name, value string) error
}
