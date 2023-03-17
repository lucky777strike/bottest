package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lucky777strike/bottest/domain"
)

type statisticsPostgres struct {
	db *sqlx.DB
}

type CompositeRepository struct {
	StatRepo domain.StatisticsRepository
}

func (r *CompositeRepository) GetStatRepo() domain.StatisticsRepository {
	return r.StatRepo
}

func NewRepository(db *sqlx.DB) domain.Repository {
	return &CompositeRepository{
		StatRepo: NewStatisticsPostgresRepository(db)}
}
