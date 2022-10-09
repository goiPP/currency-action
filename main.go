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

    output := action.GetInput("output")
    if output == "response" {
	    fmt.Print("now")
	    action.Infof(currency.ExchangeResponse(to_currency, from_currency))
    } 

}
