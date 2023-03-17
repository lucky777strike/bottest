package usecase

import (
	"github.com/lucky777strike/bottest/domain"
)

type CompositeUsecase struct {
	StatUcase domain.StatisticsUsecase
}

func NewUsecase(Repo domain.Repository) domain.Usecase {
	return &CompositeUsecase{
		StatUcase: newStatisticsUsecase(Repo.GetStatRepo())}
}

func (r *CompositeUsecase) GetStatUcase() domain.StatisticsUsecase {
	return r.StatUcase
}
