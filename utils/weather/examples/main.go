package main

import (
	"context"
	"fmt"

	"github.com/lucky777strike/bottest/utils/weather"
)

func main() {
	w := weather.New()
	fmt.Println(w.AviableCities())
	fmt.Println(w.GetWeather(context.Background(), "Санкт-Петербург"))
}
