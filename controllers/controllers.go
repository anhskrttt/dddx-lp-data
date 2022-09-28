package controllers

import (
	"dddx-lp-data/abigen/pair"
	"dddx-lp-data/abigen/token"
	"dddx-lp-data/initializers"
	"dddx-lp-data/models"
	"dddx-lp-data/utils"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "pong",
	})
}

func GetUserInfo(c *gin.Context) {
	// Declare a model to response
	var user models.UserInformation
	user.Protocol = "DDDX"         // Default
	user.Tag = "Liquidity Provide" // Default

	// Get data off body req
	walletAddr := common.HexToAddress(c.Param("usrAddr"))
	user.Address = walletAddr.Hex()

	/**********************************************************************/
	/* Reading contract */
	pairAddress := common.HexToAddress("0x322Ba943c19f9ec1EF8ceAD8260b30E789Ca1846") // WBNB/BUSD
	instance, err := pair.NewPair(pairAddress, initializers.Client)
	if err != nil {
		log.Fatal(err)
	}

	user.PairName, err = instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	user.PairSymbol, err = instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get wallet's LP token balance
	bal, err := instance.BalanceOf(&bind.CallOpts{}, walletAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Get total supply of LP tokens
	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get the wallet's LP ratio
	lpRatio := new(big.Float).Quo(new(big.Float).SetInt(bal), new(big.Float).SetInt(totalSupply))

	// Get pool's token0 balance
	// Ex: WBNB

	// Get token0's address
	token0Addr, err := instance.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get pool's balance
	token0Instance, err := token.NewToken(common.HexToAddress(token0Addr.Hex()), initializers.Client)
	if err != nil {
		log.Fatal(err)
	}

	token0Bal, err := token0Instance.BalanceOf(&bind.CallOpts{}, pairAddress)
	if err != nil {
		log.Fatal(err)
	}

	token0BalRatio := new(big.Float).Mul(new(big.Float).SetInt(token0Bal), lpRatio)
	result0 := new(big.Int)
	token0BalRatio.Int(result0)

	user.Token0Bal = int(result0.Int64())

	// Get pool's token1 balance
	// Ex: WBNB
	// Get token0's address
	token1Addr, err := instance.Token1(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get pool's balance
	token1Instance, err := token.NewToken(common.HexToAddress(token1Addr.Hex()), initializers.Client)
	if err != nil {
		log.Fatal(err)
	}

	token1Bal, err := token1Instance.BalanceOf(&bind.CallOpts{}, pairAddress)
	if err != nil {
		log.Fatal(err)
	}

	_ = token1Bal

	token1BalRatio := new(big.Float).Mul(new(big.Float).SetInt(token1Bal), lpRatio)
	result1 := new(big.Int)
	token1BalRatio.Int(result1)

	user.Token1Bal = int(result1.Int64())

	// user.Token1Bal = new(big.Float).Mul(new(big.Float).SetInt(token1Bal), lpRatio)

	/* End of reading contract*/
	/**********************************************************************/

	/**********************************************************************/
	/* Query data from coingecko */

	// WBNB to USD
	price0, price1 := utils.GetUsdPriceInPair("wbnb", "busd")
	user.Token0BalInUsd = utils.WeiToEth(user.Token0Bal) * price0

	// BUSD to USD

	user.Token1BalInUsd = utils.WeiToEth(user.Token1Bal) * price1

	/* End of querying data from coingecko*/
	/**********************************************************************/

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
		"user":    user,
	})
}
