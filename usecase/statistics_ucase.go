package usecase

import (
	"context"
	"errors"
	"time"

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
		if errors.Is(err, domain.ErrUserNotFound) {
			newUserStats := &domain.UserStatistics{
				UserID:           userID,
				FirstRequestTime: time.Now(),
				LastRequestTime:  time.Now(),
			}
			err = s.statsRepo.SetUserStatistics(ctx, *newUserStats)
			if err != nil {
				return nil, err
			}
			return newUserStats, nil
		}
		return nil, err
	}
	return userStats, nil
}

func (s *StatisticsService) UpdateUserStatistics(ctx context.Context, stats domain.UserStatistics) error {
	return s.statsRepo.UpdateUserStatistics(ctx, stats)
}

func (s *StatisticsService) ResetUserStatistics(ctx context.Context, userID int64) error {
	newUserStats := domain.UserStatistics{
		UserID:           userID,
		FirstRequestTime: time.Now(),
		TotalRequests:    0,
		LastRequestTime:  time.Now(),
	}
	return s.statsRepo.UpdateUserStatistics(ctx, newUserStats)
}

func (s *StatisticsService) IncrementUserStatistics(ctx context.Context, userID int64) error {
	err := s.statsRepo.IncrementUserStatistics(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newUserStats := &domain.UserStatistics{
				UserID:           userID,
				FirstRequestTime: time.Now(),
				LastRequestTime:  time.Now(),
				TotalRequests:    1,
			}
			err = s.statsRepo.SetUserStatistics(ctx, *newUserStats)
			if err != nil {
				return err
			}
			return nil
		}

	}
	return err
}
