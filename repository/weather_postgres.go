package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/lucky777strike/bottest/domain"
)

type weatherPostgres struct {
	db *sqlx.DB
}

func NewWeatherPostgresRepository(db *sqlx.DB) domain.WeatherRepository {
	return &weatherPostgres{db}
}

func (w *weatherPostgres) GetWeather(ctx context.Context, city string) (domain.Weather, error) {
	var weather domain.Weather
	query := `SELECT city, temp, condition, last_upd FROM weather WHERE city = $1`

	err := w.db.GetContext(ctx, &weather, query, city)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Weather{}, domain.ErrNoWeatherInBase
		}
		return domain.Weather{}, err
	}

	return weather, nil
}

func (w *weatherPostgres) SetWeather(ctx context.Context, weather domain.Weather) error {
	query := `
		INSERT INTO weather (city, temp, condition, last_upd)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (city) DO UPDATE
		SET temp = $2, condition = $3, last_upd = $4
	`

	_, err := w.db.ExecContext(ctx, query, weather.City, weather.Temp, weather.Condition, weather.LastUpd)
	return err
}
