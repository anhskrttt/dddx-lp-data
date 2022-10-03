package utils

import "math/big"

func DivBigIntToBigFloat(x *big.Int, y *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(x), new(big.Float).SetInt(y))
}
