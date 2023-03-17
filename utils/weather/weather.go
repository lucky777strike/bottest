package weather

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

type WeatherService struct {
	client *http.Client
	paths  map[string]string
}

type WeatherRes struct {
	Temp      int
	Condition string
}

func New() *WeatherService {
	cities := make(map[string]string)
	// cities["Санкт-Петербург"] = "/pogoda/russia/saint_petersburg/"
	// cities["Адлер"] = "/pogoda/russia/adler/"
	// cities["Анапа"] = "/pogoda/russia/anapa/"
	// cities["Архангельск"] = "/pogoda/russia/arkhangelsk/"
	// cities["Астрахань"] = "/pogoda/russia/astrakhan/"
	cities["Санкт-Петербург"] = "/weather/overview/sankt-peterburg/"
	cities["Москва"] = "/weather/overview/moskva"
	cities["yakutsk"] = "/weather/overview/yakutsk"
	client := &http.Client{}
	return &WeatherService{client: client, paths: cities}
}

// func (w *WeatherService) GetWeather(ctx context.Context, city string) (string, string, error) { //TODO ERROR HANDLING
// 	if p, ok := w.paths[city]; ok {
// 		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://world-weather.ru"+p, nil)
// 		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
// 		if err != nil {
// 			return "", "", err
// 		}
// 		resp, err := w.client.Do(req)
// 		if err != nil {
// 			return "", "", err
// 		}
// 		defer resp.Body.Close()

// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return "", "", err
// 		}
// 		doc := soup.HTMLParse(string(body))
// 		number := doc.Find("div", "id", "weather-now-number").Text()
// 		weather := "неизвестно"
// 		if w, ok := doc.Find("span", "id", "weather-now-icon").Attrs()["title"]; ok {
// 			weather = w
// 		}

// 		return number, weather, nil
// 	}
// 	return "", "", ErrCityNotFound
// }

func (w *WeatherService) GetWeather(ctx context.Context, city string) (WeatherRes, error) { //TODO ERROR HANDLING
	res := WeatherRes{}
	if p, ok := w.paths[city]; ok {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()
		req, err := http.NewRequestWithContext(ctxWithTimeout, http.MethodGet, "https://www.meteoservice.ru"+p, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
		if err != nil {
			return res, err
		}
		resp, err := w.client.Do(req)
		if err != nil {
			return res, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}
		//fmt.Println(string(body))
		doc := soup.HTMLParse(string(body))
		number := "неизвестно"
		weather := "неизвестно"
		digitsl := doc.FindAllStrict("div", "class", "temperature margin-bottom-0")
		if len(digitsl) > 0 {
			a := digitsl[0].FindAll("span", "class", "value")
			if len(a) > 0 {
				number = a[0].Text()
			}

		}
		weathersl := doc.FindAll("div", "class", "icon")
		for _, elem := range weathersl {
			//sl := strings.Split(elem.Attrs()["class"], " ")
			if strings.Contains(elem.Attrs()["class"], "has-tip tip-top") {
				weather = elem.Attrs()["title"]
				break
			}

		}

		//weather := doc.Find("div", "class", "col-16 text-500").Text()
		number = strings.Trim(number, "°")
		temp, err := strconv.Atoi(number)
		if err != nil {
			return res, err
		}

		return WeatherRes{Temp: temp, Condition: weather}, nil
	}
	return res, ErrCityNotFound
}

func (w *WeatherService) AviableCities() []string {
	keys := make([]string, 0, len(w.paths))
	for k := range w.paths {
		keys = append(keys, k)
	}
	return keys
}
