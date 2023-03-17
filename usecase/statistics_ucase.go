package usecase

import (
	"context"

	"github.com/lucky777strike/bottest/domain"
)

type StatisticsService struct {
	statsRepo domain.StatisticsRepository
}

func newStatisticsUsecase(statsRepo domain.StatisticsRepository) domain.StatisticsUsecase {
	return &StatisticsService{statsRepo}
}

func (s *StatisticsService) GetUserStatistics(ctx context.Context, userID int64) (*domain.UserStatistics, error) {
	userStats, err := s.statsRepo.GetUserStatistics(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userStats, nil
}

func (s *StatisticsService) UpdateUserStatistics(ctx context.Context, userID int64, stats *domain.UserStatistics) error {
	return s.statsRepo.UpdateUserStatistics(ctx, userID, stats)
}

func (s *StatisticsService) ResetUserStatistics(ctx context.Context, userID int64) error {
	return s.statsRepo.ResetUserStatistics(ctx, userID)
}

func (s *StatisticsService) IncrementUserStatistics(ctx context.Context, userID int64) error {
	return s.statsRepo.IncrementUserStatistics(ctx, userID)
}
