package domain

type Service struct {
	Stat    StatisticsUsecase
	Weather WeatherUsecase
}

type Repository struct {
	Stat    StatisticsRepository
	Weather WeatherRepository
}
