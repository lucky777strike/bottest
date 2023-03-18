package domain

type Service struct {
	Stat     StatisticsUsecase
	Weather  WeatherUsecase
	Currency CurrencyUsecase
}

type Repository struct {
	Stat     StatisticsRepository
	Weather  WeatherRepository
	Currency CurrencyRepository
}
