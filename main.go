package main

import (
	"dddx-lp-data/abigen/pair"
	"dddx-lp-data/abigen/token"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Step 01: Initialize BSC Client
	client, err := ethclient.Dial("https://bsc-dataseed4.ninicoin.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection")

	_ = client

	// Step 02: Initialize gin router
	r := gin.Default()
	r.GET("/api/:usrAddr", func(c *gin.Context) {
		// Get data off body req
		walletAddr := common.HexToAddress(c.Param("usrAddr"))

		/**********************************************************************/
		/* Reading contract */
		pairAddress := common.HexToAddress("0x322Ba943c19f9ec1EF8ceAD8260b30E789Ca1846") // WBNB/BUSD
		instance, err := pair.NewPair(pairAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		name, err := instance.Name(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		symbol, err := instance.Symbol(&bind.CallOpts{})
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
		token0Instance, err := token.NewToken(common.HexToAddress(token0Addr.Hex()), client)
		if err != nil {
			log.Fatal(err)
		}

		token0Bal, err := token0Instance.BalanceOf(&bind.CallOpts{}, pairAddress)
		if err != nil {
			log.Fatal(err)
		}

		token0BalOfUser := new(big.Float).Mul(new(big.Float).SetInt(token0Bal), lpRatio)

		// Get pool's token1 balance
		// Ex: WBNB
		// Get token0's address
		token1Addr, err := instance.Token1(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		// Get pool's balance
		token1Instance, err := token.NewToken(common.HexToAddress(token1Addr.Hex()), client)
		if err != nil {
			log.Fatal(err)
		}

		token1Bal, err := token1Instance.BalanceOf(&bind.CallOpts{}, pairAddress)
		if err != nil {
			log.Fatal(err)
		}

		token1BalOfUser := new(big.Float).Mul(new(big.Float).SetInt(token1Bal), lpRatio)

		/* End of reading contract*/
		/**********************************************************************/

		c.JSON(http.StatusOK, gin.H{
			"message":        "successful",
			"userAddress":    walletAddr,
			"protocol":       "DDDX",
			"tag":            "LP",
			"pairName":       name,
			"pairSymbol":     symbol,
			"token0Bal":      fmt.Sprintf("%.18f", WeiToEth(token0BalOfUser)),
			"token1Bal":      fmt.Sprintf("%.18f", WeiToEth(token1BalOfUser)),
			"token0BalInUSD": "",
			"token1BalInUSD": "",
		})
	})
	r.Run()
}

func WeiToEth(balance *big.Float) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	bnbValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return bnbValue
}
