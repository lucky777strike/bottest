package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/lucky777strike/bottest/domain"
	"github.com/lucky777strike/bottest/utils/weather"
)

type WeatherService struct {
	tool *weather.WeatherService
	repo domain.WeatherRepository
}

func newWeatherUsecase(repo domain.WeatherRepository) domain.WeatherUsecase {
	return &WeatherService{tool: weather.New(), repo: repo}
}

func (w *WeatherService) GetWeather(ctx context.Context, city string) (domain.Weather, error) {
	var res domain.Weather
	res, err := w.repo.GetWeather(ctx, city)
	if err != nil {
		log.Println(err)
		if errors.Is(err, domain.ErrNoWeatherInBase) { //Если город нет в базе
			res, err = w.ParseWeather(ctx, city)
			if err != nil {
				if errors.Is(err, weather.ErrCityNotFound) {
					return domain.Weather{}, domain.ErrCityNotFound
				}
				return domain.Weather{}, err
			}

		}
	}
	if time.Since(res.LastUpd) > 2*time.Hour { //Если с последнего апдейта в базе прошло больше 2х часов парсим заново
		res, err = w.ParseWeather(ctx, city)
		if err != nil {
			return domain.Weather{}, domain.ErrCityNotFound
		}
	}
	return res, nil

}

func (w *WeatherService) ParseWeather(ctx context.Context, city string) (domain.Weather, error) {

	wres, err := w.tool.GetWeather(ctx, city)
	if err != nil {
		if errors.Is(err, weather.ErrCityNotFound) {
			return domain.Weather{}, domain.ErrCityNotFound
		}
		return domain.Weather{}, err
	}
	res := domain.Weather{
		City:      city,
		LastUpd:   time.Now(),
		Temp:      wres.Temp,
		Condition: wres.Condition,
	}
	w.repo.SetWeather(ctx, res)

	return res, nil

}

func (w *WeatherService) AviableCities() []string {
	return w.tool.AviableCities()
}
