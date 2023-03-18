package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/lucky777strike/bottest/domain"
	"github.com/lucky777strike/bottest/utils/currency"
)

type CurrencyService struct {
	tool *currency.CurrencyService
	repo domain.CurrencyRepository
}

func NewCurrencyUsecase(repo domain.CurrencyRepository) domain.CurrencyUsecase {
	return &CurrencyService{tool: currency.New(), repo: repo}
}

func (c *CurrencyService) GetCurrency(ctx context.Context, name string) (domain.Currency, error) {
	cur, err := c.repo.GetCurrency(ctx, name)
	if err != nil {
		if errors.Is(err, domain.ErrNoCurrencyInBase) {
			cur, err = c.UpdateCurrencyFromAPI(ctx, name)
			if err != nil {
				return domain.Currency{}, err
			}
		} else {
			return domain.Currency{}, err
		}
	}
	if time.Since(cur.LastUpd) > 1*time.Hour { //Если с последнего апдейта в базе прошло больше 1х часов парсим заново
		cur, err = c.UpdateCurrencyFromAPI(ctx, name)
		if err != nil {
			return domain.Currency{}, domain.ErrNoCurrencyInBase
		}
	}

	return cur, nil
}

func (c *CurrencyService) SetCurrency(ctx context.Context, currency domain.Currency) error {

	return c.repo.SetCurrency(ctx, currency)
}

func (c *CurrencyService) UpdateCurrencyFromAPI(ctx context.Context, name string) (domain.Currency, error) {
	curValue, err := c.tool.ParseCurrencyValue(ctx, name)
	if err != nil {
		return domain.Currency{}, err
	}

	cur := domain.Currency{
		Name:    name,
		Value:   curValue.Value,
		LastUpd: time.Now(),
	}

	err = c.repo.SetCurrency(ctx, cur)
	if err != nil {
		return domain.Currency{}, err
	}

	return cur, nil
}

func (c *CurrencyService) AvailableCurrencies() []string {
	return c.tool.AvailableCurrencies()
}
