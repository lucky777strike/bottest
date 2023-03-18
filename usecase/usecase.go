package usecase

import (
	"github.com/lucky777strike/bottest/domain"
)

func NewService(repo *domain.Repository) *domain.Service {
	return &domain.Service{
		Stat:     newStatisticsUsecase(repo.Stat),
		Weather:  newWeatherUsecase(repo.Weather),
		Currency: NewCurrencyUsecase(repo.Currency)}
}
