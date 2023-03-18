package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lucky777strike/bottest/domain"
)

func NewRepository(db *sqlx.DB) *domain.Repository {
	return &domain.Repository{Stat: NewStatisticsPostgresRepository(db),
		Weather: NewWeatherPostgresRepository(db)}
}
