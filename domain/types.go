package domain

type Service struct {
	Stat    StatisticsUsecase
	Weather WeatherUsecase
}
