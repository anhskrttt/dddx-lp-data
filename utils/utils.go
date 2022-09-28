package utils

import (
	"math"
	"math/big"
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
