package usecase

import (
	"github.com/lucky777strike/bottest/domain"
)

type Service struct {
	domain.StatisticsUsecase
}

func NewService(statRepo domain.StatisticsRepository) domain.Usecase {
	return &Service{
		StatisticsUsecase: newStatisticsUsecase(statRepo)}
}
