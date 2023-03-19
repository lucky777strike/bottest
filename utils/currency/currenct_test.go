package currency_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/lucky777strike/bottest/utils/currency"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return nil, errors.New("unimplemented")
}

func TestParseCurrencyValue(t *testing.T) {
	testCases := []struct {
		name         string
		currency     string
		mockResponse string
		mockError    error
		expected     currency.CurrencyRes
		expectedErr  error
	}{
		{
			name:         "USD success",
			currency:     "usd",
			mockResponse: `<html> <head></head> <body> <input id="answer" value="75.23" /> </body> </html> `,
			expected:     currency.CurrencyRes{Name: "usd", Value: 75.23},
		},
		{
			name:         "EUR success",
			currency:     "eur",
			mockResponse: `<html> <head></head> <body> <input id="answer" value="83.45" /> </body> </html>`,
			expected:     currency.CurrencyRes{Name: "eur", Value: 83.45},
		},
		{
			name:        "invalid currency",
			currency:    "invalid",
			expectedErr: currency.ErrCurNotFound,
		},
		{
			name:         "parsing error",
			currency:     "usd",
			mockResponse: `<html><head></head><body></body></html>`,
			expectedErr:  currency.ErrParsingErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := &mockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					if tc.mockError != nil {
						return nil, tc.mockError
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(strings.NewReader(tc.mockResponse)),
					}, nil
				},
			}

			cs := currency.NewWithClient(mockClient)

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()

			res, err := cs.ParseCurrencyValue(ctx, tc.currency)

			if tc.expectedErr != nil {
				require.Error(t, err)
				assert.EqualError(t, err, tc.expectedErr.Error(), "expected error: %v, got: %v", tc.expectedErr, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, res)
			}
		})
	}
}
func TestAvailableCurrencies(t *testing.T) {
	cs := currency.New()

	expectedCurrencies := []string{"usd", "eur"}

	availableCurrencies := cs.AvailableCurrencies()

	assert.Equal(t, len(expectedCurrencies), len(availableCurrencies), "unexpected number of available currencies")

	for _, expectedCurrency := range expectedCurrencies {
		assert.Contains(t, availableCurrencies, expectedCurrency, "expected currency not found in available currencies")
	}
}
