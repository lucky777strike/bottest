package currency

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/anaskhan96/soup"
)

type CurrencyService struct {
	client *http.Client
	paths  map[string]string
}

type CurrencyRes struct {
	Name  string
	Value float32
}

func New() *CurrencyService {
	cur := make(map[string]string)
	cur["usd"] = "/usd/rub"
	cur["eur"] = "/usd/eur"
	client := &http.Client{}
	return &CurrencyService{client: client, paths: cur}
}

func (c *CurrencyService) ParseCurrencyValue(ctx context.Context, currency string) (CurrencyRes, error) {
	res := CurrencyRes{}
	if p, ok := c.paths[currency]; ok {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()
		req, err := http.NewRequestWithContext(ctxWithTimeout, http.MethodGet, "https://www.currency.me.uk/convert"+p, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
		if err != nil {
			return res, err
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return res, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return res, err
		}
		doc := soup.HTMLParse(string(body))
		parsed := doc.FindAll("input", "id", "answer")
		if len(parsed) != 1 {
			return res, ErrParsingErr
		}
		if v, ok := parsed[0].Attrs()["value"]; ok {
			value, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return res, ErrParsingErr
			}
			res.Name = currency
			res.Value = float32(value)
			return res, nil

		}
		//value, err := strconv.ParseFloat(parsed[0].Text(), 32)
	}
	return res, ErrCurNotFound

}

func (c *CurrencyService) AvailableCurrencies() []string {
	keys := make([]string, 0, len(c.paths))
	for k := range c.paths {
		keys = append(keys, k)
	}
	return keys
}
