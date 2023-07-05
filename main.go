package main

import (

	"fmt"
	"github.com/reactioon/rgo-rapi-lib-go/rapi"

)

func main() {

	// setup keys
	apiKey := []byte("{reactioon-api-key}")
	apiSecret := []byte("{reactioon-api-secret}")

	// load library
	r := rapi.Load(apiKey, apiSecret)

	// execute context with GET
	arrDataGet := make(map[string]string)
	requestReturnGet, requestErrGet := r.Request("GET", "api/v2/bots/spot/all", arrDataGet)

	fmt.Println("return", requestReturnGet)
	fmt.Println("error", requestErrGet)

	// execute context with POST
	arrDataPost := make(map[string]string)
	arrDataPost["exchange"] = "binance"
	arrDataPost["symbol"] = "BTCUSDT"
	arrDataPost["currency"] = "USDT"
	requestReturnPost, requestErrPost := r.Request("POST", "api/v2/watchlist/market/info", arrDataPost)

	fmt.Println("return", requestReturnPost)
	fmt.Println("error", requestErrPost)

}