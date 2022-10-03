package utils

import (
	"dddx-lp-data/abigen/pair"
	"dddx-lp-data/abigen/token"
	"dddx-lp-data/initializers"
	"dddx-lp-data/models"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Which is the better practice?  Return values only or return values with error?
// func GetPoolInstance(address string) (*pair.Pair, error) {
// 	return pair.NewPair(common.HexToAddress(address), initializers.Client)
// }

// func GetTokenInstance(address string) (*token.Token, error) {
// 	return token.NewToken(common.HexToAddress(address), initializers.Client)
// }

func GetPoolInstance(address string) *pair.Pair {
	pool, err := pair.NewPair(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return pool
}

func GetTokenInstance(address string) *token.Token {
	token, err := token.NewToken(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return token
}

func GetTokenPairInstances(poolInstance *pair.Pair) (*token.Token, *token.Token) {
	token0Address, token1Address, err := poolInstance.Tokens(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return GetTokenInstance(token0Address.String()), GetTokenInstance(token1Address.String())
}

func GetPoolFromAddress(poolAddress string) (ps models.PoolSimple) {
	var err error

	var poolToken models.Token
	instance := GetPoolInstance(poolAddress)

	poolToken.Name, err = instance.Name(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	poolToken.Symbol, err = instance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	poolToken.Decimals, err = instance.Decimals(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	ps.LPToken = poolToken
	ps.ChainName = "BNB Smart Chain"
	ps.PoolAddress = poolAddress

	return
}

func GetTokenPairBalOfUser(userAddress string, poolAddress string) (models.TokenBalance, models.TokenBalance) {
	// Initialize instances
	poolInstance := GetPoolInstance(poolAddress)
	token0Instance, token1Instance := GetTokenPairInstances(poolInstance)

	// Get symbols
	symbol0, err := token0Instance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	symbol1, err := token1Instance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	// Get the wallet's LP ratio
	lpRatio := GetLpRatio(poolInstance, userAddress)

	// Balance of pool
	token0BalOfPool, err := token0Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
	if err != nil {
		panic(err)
	}

	token1BalOfPool, err := token1Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
	if err != nil {
		panic(err)
	}

	token0BalOfUser := GetRatioOfBal(lpRatio, token0BalOfPool) // = ratio * total amount
	token1BalOfUser := GetRatioOfBal(lpRatio, token1BalOfPool) // = ratio * total amount

	// Balance in USD
	/**********************************************************************/
	/* Query data from coingecko */

	// WBNB to USD
	fmt.Println(symbol0, symbol1)
	price0, price1 := GetUsdPriceInPair(symbol0, symbol1)
	Token0BalInUsd := WeiToEth(token0BalOfUser) * price0

	// BUSD to USD
	Token1BalInUsd := WeiToEth(token1BalOfUser) * price1

	/* End of querying data from coingecko*/
	/**********************************************************************/

	token0Bal := models.TokenBalance{
		TokenSymbol:  symbol0,
		Balance:      token0BalOfUser,
		BalanceInUSD: Token0BalInUsd,
	}

	token1Bal := models.TokenBalance{
		TokenSymbol:  symbol1,
		Balance:      token1BalOfUser,
		BalanceInUSD: Token1BalInUsd,
	}

	return token0Bal, token1Bal

}

// Get ratio amount of total
// return: ratio * total
func GetRatioOfBal(ratio *big.Float, total *big.Int) *big.Int {
	ratioBal := new(big.Float).Mul(new(big.Float).SetInt(total), ratio)
	result := new(big.Int)
	ratioBal.Int(result)
	return result
}

func GetLpRatio(poolInstance *pair.Pair, userAddress string) *big.Float {
	// Get wallet's LP token balance
	bal, err := poolInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	if err != nil {
		panic(err)
	}

	// Get total supply of LP tokens
	totalSupply, err := poolInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	// This equals to: "return *big.Float(bal/totalSupply)"
	return DivBigIntToBigFloat(bal, totalSupply)
}

func GetTokenPairBalOfPool(poolAddress string) (*big.Int, *big.Int) {
	poolInstance := GetPoolInstance(poolAddress)

	// Get addresses of token0, token1
	token0Addr, token1Addr, err := poolInstance.Tokens(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	// Get pool's balance
	token0Instance, err := token.NewToken(common.HexToAddress(token0Addr.Hex()), initializers.Client)
	if err != nil {
		panic(err)
	}

	token0Bal, err := token0Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
	if err != nil {
		panic(err)
	}

	// Get pool's balance
	token1Instance, err := token.NewToken(common.HexToAddress(token1Addr.Hex()), initializers.Client)
	if err != nil {
		panic(err)
	}

	token1Bal, err := token1Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
	if err != nil {
		panic(err)
	}

	return token0Bal, token1Bal
}
