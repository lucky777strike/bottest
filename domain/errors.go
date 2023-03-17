package domain

import "errors"

// stat
var ErrUserNotFound = errors.New("user not found")

// weather
var ErrCityNotFound = errors.New("city not found")
var ErrNoWeatherInBase = errors.New("database not contain this city")
