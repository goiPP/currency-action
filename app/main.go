package main

import (
	"fmt"
	"app/currency"
)


func main() {
	var to_currency string
	var from_currency string

	fmt.Print("Enter your to_currency: ")
	fmt.Scan(&to_currency)
	fmt.Print("Enter your from_currency: ")
	fmt.Scan(&from_currency)

    fmt.Print(currency.ExchangeResponse(to_currency, from_currency))
}
