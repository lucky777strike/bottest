package usecase

import (
	"github.com/lucky777strike/bottest/domain"
)

type Service struct {
	domain.StatisticsUsecase
	domain.WeatherUsecase
}

func NewService(statRepo domain.StatisticsRepository, weatherRepo domain.WeatherRepository) domain.Usecase {
	return &Service{
		StatisticsUsecase: newStatisticsUsecase(statRepo),
		WeatherUsecase:    newWeatherUsecase(weatherRepo)}
}
