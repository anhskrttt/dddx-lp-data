package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
)

func WeiToEth(balance int) float64 {
	fbalance := new(big.Float)
	fbalance.SetString(strconv.Itoa(balance))
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
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=" + token0 + "," + token1 + "&vs_currencies=usd")
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

	price0 := data[token0].(map[string]interface{})["usd"]
	price1 := data[token1].(map[string]interface{})["usd"]
	return price0.(float64), price1.(float64)
}
