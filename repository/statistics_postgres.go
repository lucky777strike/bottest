package repository

import (
	"context"
	"time"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lucky777strike/bottest/domain"
)

type statisticsPostgres struct {
	db *sqlx.DB
}

func NewStatisticsPostgresRepository(db *sqlx.DB) domain.StatisticsRepository {
	return &statisticsPostgres{db}
}
func (s *statisticsPostgres) SetUserStatistics(ctx context.Context, userStats domain.UserStatistics) error {
	query := `
		INSERT INTO user_statistics (user_id, first_request_time, total_requests, last_request_time)
		VALUES ($1, $2, $3, $4)
	`
	_, err := s.db.ExecContext(ctx, query, userStats.UserID, userStats.FirstRequestTime, userStats.TotalRequests, userStats.LastRequestTime)
	return err
}

func (s *statisticsPostgres) GetUserStatistics(ctx context.Context, userID int64) (*domain.UserStatistics, error) {
	userStats := &domain.UserStatistics{}
	query := `SELECT user_id, first_request_time, total_requests, last_request_time FROM user_statistics WHERE user_id = $1`

	err := s.db.GetContext(ctx, userStats, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return userStats, nil
}

func (s *statisticsPostgres) UpdateUserStatistics(ctx context.Context, stats domain.UserStatistics) error {
	query := `UPDATE user_statistics SET first_request_time = $1, total_requests = $2, last_request_time = $3 WHERE user_id = $4`

	_, err := s.db.ExecContext(ctx, query, stats.FirstRequestTime, stats.TotalRequests, stats.LastRequestTime, stats.UserID)
	return err
}

func (s *statisticsPostgres) IncrementUserStatistics(ctx context.Context, userID int64) error {
	query := `UPDATE user_statistics SET total_requests = total_requests + 1, last_request_time = $1 WHERE user_id = $2`

	res, err := s.db.ExecContext(ctx, query, time.Now(), userID)
	a, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if a == 0 {
		return domain.ErrUserNotFound
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrUserNotFound
		}
		return err
	}
	return err
}
