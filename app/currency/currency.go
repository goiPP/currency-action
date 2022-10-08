package currency

import (
	"app/restclient"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var apiResponse Body

func ExchangeRate(to string, from string) int {
	callApi(to, from)
	return apiResponse.Response.getRate()
}

func ExchangeResponse(to string, from string) string {
	callApi(to, from)
	return apiResponse.getResponse()
}

func callApi(to string, from string) {
	url := fmt.Sprintf("https://alpha-vantage.p.rapidapi.com/query?to_currency=%s&function=CURRENCY_EXCHANGE_RATE&from_currency=%s", to, from)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("No response from request")
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", "alpha-vantage.p.rapidapi.com")

	res, err := restclient.Client.Do(req)
	fmt.Print(req.URL)
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&apiResponse)
}
