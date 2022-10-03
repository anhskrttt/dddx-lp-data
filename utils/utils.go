package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"strings"
)

var TokenSymbolToId map[string]string = map[string]string{
	"wbnb": "wbnb",
	"busd": "busd",
	"eth":  "ethereum",
	"btc":  "bitcoin",
}

func WeiToEth(balance *big.Int) float64 {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	bnbValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	// WARNING: Should check accuracy?
	result, _ := bnbValue.Float64()

	return result
}

func GetUsdPrice(token string) float64 {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + token + "&vs_currencies=usd")
	if err != nil {
		log.Fatalln(err)
	}
	// Client must close response body when finished with it
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]interface{}
	errParse := json.Unmarshal([]byte(body), &data)
	if errParse != nil {
		panic(err)
	}

	price := data[token].(map[string]interface{})["usd"]
	return price.(float64)
}

func GetUsdPriceInPair(token0 string, token1 string) (float64, float64) {
	// Lower case all parameter strings
	token0_id := TokenSymbolToId[strings.ToLower(token0)]
	token1_id := TokenSymbolToId[strings.ToLower(token1)]

	fmt.Println(token0_id, token1_id)

	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + token0_id + "," + token1_id + "&vs_currencies=usd")
	if err != nil {
		panic(err)
	}
	// Client must close response body when finished with it
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]interface{}
	errParse := json.Unmarshal([]byte(body), &data)
	if errParse != nil {
		panic(err)
	}

	price0 := data[token0_id].(map[string]interface{})["usd"]
	price1 := data[token1_id].(map[string]interface{})["usd"]
	return price0.(float64), price1.(float64)
}
