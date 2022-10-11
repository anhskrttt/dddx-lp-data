package utils

import (
	"dddx-lp-data/abigen/token"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetTokenPairBalOfPool(poolAddress string, token0Instance *token.Token, token1Instance *token.Token) (*big.Int, *big.Int) {
	// Balance of pool
	token0BalOfPool, err := token0Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
	if err != nil {
		panic(err)
	}

	token1BalOfPool, err := token1Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
	if err != nil {
		panic(err)
	}

	return token0BalOfPool, token1BalOfPool
}

func GetTokenPairSymbolOfPool(poolAddress string, token0Instance *token.Token, token1Instance *token.Token) (string, string) {
	symbol0, err := token0Instance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	symbol1, err := token1Instance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return symbol0, symbol1
}
