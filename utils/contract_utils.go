package utils

import (
	"dddx-lp-data/abigen/gauge"
	"dddx-lp-data/abigen/pair"
	"dddx-lp-data/abigen/token"
	"dddx-lp-data/abigen/ve"
	"dddx-lp-data/initializers"
	"dddx-lp-data/models"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

/**********************************************************************/
/* Generate instances */

// GetTokenInstance generates a new erc20 contract instance.
func GetTokenInstance(address string) *token.Token {
	token, err := token.NewToken(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return token
}

// GetPoolInstance generates a new pool contract instance.
func GetPoolInstance(address string) *pair.Pair {
	pool, err := pair.NewPair(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return pool
}

// GetGaugeInstance generates a new gauge contract instance.
func GetGaugeInstance(address string) *gauge.Gauge {
	gauge, err := gauge.NewGauge(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return gauge
}

// GetVeInstance generates a new vote-escrowed-nft contract instance.
func GetVeInstance(address string) *ve.Ve {
	veNFT, err := ve.NewVe(common.HexToAddress(address), initializers.Client)
	if err != nil {
		panic(err)
	}

	return veNFT
}

// GetTokenPairInstances generates two new erc20 contract instances from pool Instance.
func GetTokenPairInstances(poolInstance *pair.Pair) (*token.Token, *token.Token) {
	token0Address, token1Address, err := poolInstance.Tokens(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	return GetTokenInstance(token0Address.String()), GetTokenInstance(token1Address.String())
}

/* End of Generate instances */
/**********************************************************************/

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

	token0BalOfUser := GetBalOfRatio(lpRatio, token0BalOfPool) // = ratio * total amount
	token1BalOfUser := GetBalOfRatio(lpRatio, token1BalOfPool) // = ratio * total amount

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

func GetPoolFromGauge(gaugeAddress string) models.PoolSimple {
	gaugeInstance := GetGaugeInstance(gaugeAddress)

	// Get pair contract address from gauge contract
	poolAddress, err := gaugeInstance.Stake(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	return GetPoolFromAddress(poolAddress.Hex())
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

// func GetTokenPairBalOfPool(poolAddress string) (*big.Int, *big.Int) {
// 	poolInstance := GetPoolInstance(poolAddress)

// 	// Get addresses of token0, token1
// 	token0Addr, token1Addr, err := poolInstance.Tokens(&bind.CallOpts{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Get pool's balance
// 	token0Instance, err := token.NewToken(common.HexToAddress(token0Addr.Hex()), initializers.Client)
// 	if err != nil {
// 		panic(err)
// 	}

// 	token0Bal, err := token0Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Get pool's balance
// 	token1Instance, err := token.NewToken(common.HexToAddress(token1Addr.Hex()), initializers.Client)
// 	if err != nil {
// 		panic(err)
// 	}

// 	token1Bal, err := token1Instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(poolAddress))
// 	if err != nil {
// 		panic(err)
// 	}

// 	return token0Bal, token1Bal
// }

/**********************************************************************/
/* Quick access to instance's funcs. */

func GetPoolAddressFromGauge(gaugeAddress string) string {
	gaugeInstance := GetGaugeInstance(gaugeAddress)

	// Get pair contract address from gauge contract
	poolAddress, err := gaugeInstance.Stake(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	return poolAddress.Hex()
}

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

/* End of quick access to instance's funcs. */
/**********************************************************************/

/**********************************************************************/
/* Generate instances of models. For testing purpose only. */

// GetPoolFromAddress generates a new pool model based on pool contract address.
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

/* End of generating instances of models */
/**********************************************************************/
