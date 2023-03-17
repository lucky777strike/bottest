package domain

import (
	"context"
	"time"
)

type Weather struct {
	City      string    `db:"city" json:"city"`
	Temp      int       `db:"temp" json:"temp"`
	Condition string    `db:"condition" json:"condition"`
	LastUpd   time.Time `db:"last_upd" json:"last_upd"`
}

type WeatherUsecase interface {
	GetWeather(ctx context.Context, city string) (Weather, error)
	AviableCities() []string
}
type WeatherRepository interface {
	GetWeather(ctx context.Context, city string) (Weather, error)
	SetWeather(ctx context.Context, w Weather) error
}
