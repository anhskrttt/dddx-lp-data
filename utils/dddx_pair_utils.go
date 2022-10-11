package utils

import (
	"dddx-lp-data/abigen/pair"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetLPBalFromPool(userAddress string, poolAddress string) *big.Int {
	poolInstance := GetPoolInstance(poolAddress)

	lpBal, err := poolInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	if err != nil {
		panic(err)
	}

	return lpBal
}

func GetLPBalFromPoolInstance(userAddress string, poolInstance *pair.Pair) *big.Int {
	lpBal, err := poolInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	if err != nil {
		panic(err)
	}

	return lpBal
}
