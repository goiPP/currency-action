package currency

import (
    "fmt"
    "strconv"
)

type Detail struct {
	FromCurrencyCode string `json:"1. From_Currency Code"`
	FromCurrencyName string `json:"2. From_Currency Name"`
	ToCurrencyCode   string `json:"3. To_Currency Code"`
	ToCurrencyName   string `json:"4. To_Currency Name"`
	ExchangeRate     string `json:"5. Exchange Rate"`
	LastRefreshed    string `json:"6. Last Refreshed"`
	TimeZone         string `json:"7. Time Zone"`
	BidPrice         string `json:"8. Bid Price"`
	AskPrice         string `json:"9. Ask Price"`
}
type Body struct {
	Response Detail `json:"Realtime Currency Exchange Rate"`
}

func (b Body) getResponse() string {
	return fmt.Sprintf("1 %s = %s %s", b.Response.FromCurrencyCode, b.Response.AskPrice, b.Response.ToCurrencyCode)
	return fmt.Sprintf("1 %s = %s %s", b.Response.FromCurrencyCode, b.Response.AskPrice, b.Response.ToCurrencyCode)
}

func (d Detail) getRate() int {
    intVar, err := strconv.Atoi(d.AskPrice)
    if err != nil {
        fmt.Println("Cannot convert string to int")
    }
	return intVar
}

