package domain

import (
	"context"
	"time"
)

type StatisticsUsecase interface {
	GetUserStatistics(ctx context.Context, userID int64) (*UserStatistics, error)
	UpdateUserStatistics(ctx context.Context, userID int64, stats *UserStatistics) error
	ResetUserStatistics(ctx context.Context, userID int64) error
	IncrementUserStatistics(ctx context.Context, userID int64) error
}
type StatisticsRepository interface {
	GetUserStatistics(ctx context.Context, userID int64) (*UserStatistics, error)
	UpdateUserStatistics(ctx context.Context, userID int64, stats *UserStatistics) error
	ResetUserStatistics(ctx context.Context, userID int64) error
	IncrementUserStatistics(ctx context.Context, userID int64) error
}

type UserStatistics struct {
	UserID           int64
	FirstRequestTime time.Time
	TotalRequests    int
	LastRequestTime  time.Time
}
