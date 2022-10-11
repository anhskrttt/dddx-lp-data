package utils

import (
	"dddx-lp-data/abigen/pair"
	"dddx-lp-data/abigen/ve"
	"dddx-lp-data/models"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GetTokenPairBalOfUser return token balance info of 02 tokens in pool.
// models.TokenBalance includes fields: (1) token symbol; (2) Balance (in wei); and (3) Balance in USD.
func GetTokenPairBalOfUser(userAddress string, gaugeAddress string, isFarming bool) (models.TokenBalance, models.TokenBalance) {
	// Declare instances: gauge, pool
	gaugeInstance := GetGaugeInstance(gaugeAddress)

	// Get pair contract address from gauge contract
	poolAddress, err := gaugeInstance.Stake(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	poolInstance := GetPoolInstance(poolAddress.Hex())

	token0Instance, token1Instance := GetTokenPairInstances(poolInstance)

	// Get pool's balance of token0, token1
	token0BalOfPool, token1BalOfPool := GetTokenPairBalOfPool(poolAddress.Hex(), token0Instance, token1Instance)

	var bal *big.Int
	if isFarming {
		bal, err = gaugeInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
		if err != nil {
			panic(err)
		}
	} else {
		bal, err = poolInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
		if err != nil {
			panic(err)
		}
	}

	totalSupply, err := poolInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get the wallet's LP ratio
	lpRatio := DivBigIntToBigFloat(bal, totalSupply)

	token0BalOfUser := GetFractionOfNum(lpRatio, token0BalOfPool) // = ratio * total amount
	token1BalOfUser := GetFractionOfNum(lpRatio, token1BalOfPool) // = ratio * total amount

	// token0BalOfUser1 := GetBal(userBal, token0BalOfPool)

	// Balance in USD
	/**********************************************************************/
	/* Query data from coingecko */

	// Get symbols
	symbol0, symbol1 := GetTokenPairSymbolOfPool(poolAddress.Hex(), token0Instance, token1Instance)

	// WBNB to USD
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

func GetTokenPairBalOfUserFrom(bal *big.Int, poolAddress string, poolInstance *pair.Pair) (models.TokenBalance, models.TokenBalance) {

	token0Instance, token1Instance := GetTokenPairInstances(poolInstance)

	// Get pool's balance of token0, token1
	token0BalOfPool, token1BalOfPool := GetTokenPairBalOfPool(poolAddress, token0Instance, token1Instance)

	totalSupply, err := poolInstance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get the wallet's LP ratio
	lpRatio := DivBigIntToBigFloat(bal, totalSupply)

	token0BalOfUser := GetFractionOfNum(lpRatio, token0BalOfPool) // = ratio * total amount
	token1BalOfUser := GetFractionOfNum(lpRatio, token1BalOfPool) // = ratio * total amount

	// token0BalOfUser1 := GetBal(userBal, token0BalOfPool)

	// Balance in USD
	/**********************************************************************/
	/* Query data from coingecko */

	// Get symbols
	symbol0, symbol1 := GetTokenPairSymbolOfPool(poolAddress, token0Instance, token1Instance)

	// WBNB to USD
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

// GetStakedBalances return an array of staked token information.
// Staked token information includes: Balance (in wei) and time to unlock fund.
func GetStakedBalances(userAddress string, veAddress string) []models.UserInformationLocked {
	var result []models.UserInformationLocked

	veInstance := GetVeInstance(veAddress)
	// Get user's total veNFT
	userBal, err := veInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	if err != nil {
		panic(err)
	}

	num := int(userBal.Int64())

	for i := 0; i < num; i++ {
		set_i := big.NewInt(int64(i))
		var test models.UserInformationLocked
		tokenId, err := veInstance.TokenOfOwnerByIndex(&bind.CallOpts{}, common.HexToAddress(userAddress), set_i)
		if err != nil {
			panic(err)
		}

		test, err = veInstance.Locked(&bind.CallOpts{}, tokenId)
		if err != nil {
			panic(err)
		}

		result = append(result, test)
	}

	return result
}

func GetStakedBalancesOf(userAddress string, veInstance *ve.Ve) []models.UserInformationLocked {
	var result []models.UserInformationLocked

	// Get user's total veNFT
	userBal, err := veInstance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(userAddress))
	if err != nil {
		panic(err)
	}

	num := int(userBal.Int64())

	for i := 0; i < num; i++ {
		set_i := big.NewInt(int64(i))
		var test models.UserInformationLocked
		tokenId, err := veInstance.TokenOfOwnerByIndex(&bind.CallOpts{}, common.HexToAddress(userAddress), set_i)
		if err != nil {
			panic(err)
		}

		test, err = veInstance.Locked(&bind.CallOpts{}, tokenId)
		if err != nil {
			panic(err)
		}

		result = append(result, test)
	}

	return result
}

func GetTokenSymbolFromVeAddress(veAddress string) string {
	veInstance := GetVeInstance(veAddress)
	tokenAddress, err := veInstance.Token(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	tokenInstance := GetTokenInstance(tokenAddress.Hex())
	symbol, err := tokenInstance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return symbol
}

func GetTokenSymbolFromVeInstance(veInstance *ve.Ve) string {
	tokenAddress, err := veInstance.Token(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	tokenInstance := GetTokenInstance(tokenAddress.Hex())
	symbol, err := tokenInstance.Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return symbol
}

/**********************************************************************/
/* Generate instances of models. For testing purpose only. */

// GetPoolFromAddress generates a new poolSimple model based on pool contract address.
// Default data is: DDDX protocol, BSC chain.
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

func GetPoolFromInstance(address string, instance *pair.Pair) (ps models.PoolSimple) {
	var err error

	var poolToken models.Token

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
	ps.PoolAddress = address

	return
}

// GetPoolFromGaugeAddress generates a new poolSimple model based on gauge contract address.
// Default data is: DDDX protocol, BSC chain.
func GetPoolFromGaugeAddress(gaugeAddress string) models.PoolSimple {
	gaugeInstance := GetGaugeInstance(gaugeAddress)

	// Get pair contract address from gauge contract
	poolAddress, err := gaugeInstance.Stake(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return GetPoolFromAddress(poolAddress.Hex())
}

/* End of generating instances of models */
/**********************************************************************/
