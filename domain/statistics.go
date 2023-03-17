package domain

import (
	"context"
	"time"
)

type StatisticsUsecase interface {
	GetUserStatistics(ctx context.Context, userID int64) (*UserStatistics, error)
	UpdateUserStatistics(ctx context.Context, stats UserStatistics) error
	ResetUserStatistics(ctx context.Context, userID int64) error
	IncrementUserStatistics(ctx context.Context, userID int64) error
}
type StatisticsRepository interface {
	SetUserStatistics(ctx context.Context, userStats UserStatistics) error
	GetUserStatistics(ctx context.Context, userID int64) (*UserStatistics, error)
	UpdateUserStatistics(ctx context.Context, stats UserStatistics) error
	IncrementUserStatistics(ctx context.Context, userID int64) error
}

type UserStatistics struct {
	UserID           int64     `db:"user_id"`
	FirstRequestTime time.Time `db:"first_request_time"`
	TotalRequests    int       `db:"total_requests"`
	LastRequestTime  time.Time `db:"last_request_time"`
}
