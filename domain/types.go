package domain

// TODO пробуем встраивание
type Repository interface {
	GetStatRepo() StatisticsRepository
}

type Usecase interface {
	GetStatUcase() StatisticsUsecase
}
