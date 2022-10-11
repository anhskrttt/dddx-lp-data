package utils

import (
	"math/big"
)

// DivBigIntToBigFloat calculates and returns z = x/y.
// With typeOf(z) = *big.Float
// With typeOf(x) = *big.Int
// With typeOf(y) = *big.Int
func DivBigIntToBigFloat(x *big.Int, y *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt(y))
}

// GetFractionOfNum returns a fraction of total amount.
// E.g., to find 3/4 of 12, GetFractionOfNum(0.75, 12) will return 9.
func GetFractionOfNum(ratio *big.Float, total *big.Int) *big.Int {
	ratioBal := new(big.Float).Mul(new(big.Float).SetInt(total), ratio)
	result := new(big.Int)
	ratioBal.Int(result)
	return result
}

func IntToBigInt(i int) *big.Int {
	return big.NewInt(int64(i))
}
