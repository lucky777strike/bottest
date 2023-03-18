package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lucky777strike/bottest/domain"
)

type currencyPostgres struct {
	db *sqlx.DB
}

func NewCurrencyPostgresRepository(db *sqlx.DB) domain.CurrencyRepository {
	return &currencyPostgres{db}
}

func (c *currencyPostgres) GetCurrency(ctx context.Context, name string) (domain.Currency, error) {
	var currency domain.Currency
	query := `SELECT name, value, last_updated FROM currency WHERE name = $1`

	err := c.db.GetContext(ctx, &currency, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Currency{}, domain.ErrNoCurrencyInBase
		}
		return domain.Currency{}, err
	}

	return currency, nil
}

func (c *currencyPostgres) SetCurrency(ctx context.Context, currency domain.Currency) error {
	query := `
        INSERT INTO currency (name, value, last_updated)
        VALUES ($1, $2, $3)
        ON CONFLICT (name) DO UPDATE
        SET value = $2, last_updated = $3
    `

	_, err := c.db.ExecContext(ctx, query, currency.Name, currency.Value, currency.LastUpd)
	return err
}
