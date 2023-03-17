package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lucky777strike/bottest/domain"
)

func NewStatisticsPostgresRepository(db *sqlx.DB) domain.StatisticsRepository {
	return &statisticsPostgres{db}
}

func (s *statisticsPostgres) GetUserStatistics(ctx context.Context, userID int64) (*domain.UserStatistics, error) {
	userStats := &domain.UserStatistics{}
	query := `SELECT user_id, first_request_time, total_requests, last_request_time FROM user_statistics WHERE user_id = $1`

	err := s.db.GetContext(ctx, userStats, query, userID)
	if err != nil {
		return nil, err
	}

	return userStats, nil
}

func (s *statisticsPostgres) UpdateUserStatistics(ctx context.Context, userID int64, stats *domain.UserStatistics) error {
	query := `UPDATE user_statistics SET first_request_time = $1, total_requests = $2, last_request_time = $3 WHERE user_id = $4`

	_, err := s.db.ExecContext(ctx, query, stats.FirstRequestTime, stats.TotalRequests, stats.LastRequestTime, userID)
	return err
}

func (s *statisticsPostgres) ResetUserStatistics(ctx context.Context, userID int64) error {
	query := `UPDATE user_statistics SET first_request_time = NULL, total_requests = 0, last_request_time = NULL WHERE user_id = $1`

	_, err := s.db.ExecContext(ctx, query, userID)
	return err
}

func (s *statisticsPostgres) IncrementUserStatistics(ctx context.Context, userID int64) error {
	query := `UPDATE user_statistics SET total_requests = total_requests + 1, last_request_time = $1 WHERE user_id = $2`

	_, err := s.db.ExecContext(ctx, query, time.Now(), userID)
	return err
}
