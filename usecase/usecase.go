package usecase

import (
	"github.com/lucky777strike/bottest/domain"
)

func NewService(statRepo domain.StatisticsRepository, weatherRepo domain.WeatherRepository) *domain.Service {
	return &domain.Service{
		Stat:    newStatisticsUsecase(statRepo),
		Weather: newWeatherUsecase(weatherRepo)}
}
