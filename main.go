package main

import (
	"fmt"
	"app/currency"
	githubactions "github.com/sethvargo/go-githubactions"
)

func main() {

	action := githubactions.New()
	to_currency := action.GetInput("to_currency")
	if to_currency == "" {
	  githubactions.Fatalf("missing 'to_currency'")
	}

	from_currency := action.GetInput("from_currency")
	if to_currency == "" {
	  githubactions.Fatalf("missing 'from_currency'")
	}

    output := action.GetInput("from_currency")
    if output == "response" {
        fmt.Printf(`::set-output name=response::%s`, currency.ExchangeResponse(to_currency, from_currency))
    } else {
        fmt.Printf(`::set-output name=rate::%d`, currency.ExchangeRate(to_currency, from_currency))
    }

}
