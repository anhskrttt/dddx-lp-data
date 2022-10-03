package utils

import "math/big"

func DivBigIntToBigFloat(x *big.Int, y *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt(y))
}

// Get ratio amount of total
// return: ratio * total
func GetBalOfRatio(ratio *big.Float, total *big.Int) *big.Int {
	ratioBal := new(big.Float).Mul(new(big.Float).SetInt(total), ratio)
	result := new(big.Int)
	ratioBal.Int(result)
	return result
}
