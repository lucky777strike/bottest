package domain

type Repository interface {
	GetStatRepo() StatisticsRepository
}

type Usecase interface {
	GetStatUcase() StatisticsUsecase
}
