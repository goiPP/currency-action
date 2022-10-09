package currency

import (
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestExchangeRate(t *testing.T) {
	defer gock.Off()

	t.Setenv("API_KEY", "secret")
	mockResponse := map[string]map[string]string{
		"Realtime Currency Exchange Rate": {
			"1. From_Currency Code": "EUR",
			"2. From_Currency Name": "Euro",
			"3. To_Currency Codee":  "THB",
			"4. To_Currency Name":   "Thai Baht",
			"5. Exchange Rate":      "36.57900000",
			"6. Last Refreshed":     "2022-10-07 11:24:07",
			"7. Time Zone":          "UTC",
			"8. Bid Price":          "36.57900000",
			"9. Ask Price":          "38",
		},
	}

	gock.New("https://alpha-vantage.p.rapidapi.com").
		Get("/query").
		MatchParam("to_currency", "THB").
		MatchParam("function", "CURRENCY_EXCHANGE_RATE").
		MatchParam("from_currency", "EUR").
		MatchHeader("X-RapidAPI-Key", "secret").
		MatchHeader("X-RapidAPI-Host", "alpha-vantage.p.rapidapi.com").
		Reply(200).
		JSON(mockResponse)

	value := ExchangeRate("THB", "EUR")
	if value != 38 {
		t.Errorf("Expected '38', got %d", value)
	}
}
