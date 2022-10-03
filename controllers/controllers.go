package controllers

import (
	"dddx-lp-data/abigen/gauge"
	"dddx-lp-data/abigen/pair"
	"dddx-lp-data/abigen/token"
	"dddx-lp-data/abigen/ve"
	"dddx-lp-data/initializers"
	"dddx-lp-data/models"
	"dddx-lp-data/utils"
	"fmt"
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

/**********************************************************************/
/* New APIs*/

// Get LP data (multiple protocols from multiple chains)
func GetAllLP(c *gin.Context) {

}

// Get LP data (multiple protocols from one chain)
func GetAllLPsOfChain(c *gin.Context) {

}

// Get LP data (multiple protocols from one chain)
func GetAllLPsOfProtocol(c *gin.Context) {

}

// Get LP data from one pool
func GetLpOfProtocol(c *gin.Context) {
	// Declare request model to use
	var data models.GetLpOfProtocolRequest

	// Get data off header: user_address, protocol_id, pool_address
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	// Read data from pool address and calculate LP balance
	token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUser(data.UserAddress, data.PoolAddress)

	// Create response & Attach data to response struct
	response := models.LPFarmResponse{
		UserAddress: data.UserAddress,
		Token0:      token0BalOfUser,
		Token1:      token1BalOfUser,
		ProtocolId:  data.ProtocolId,
		Pool:        utils.GetPoolFromAddress(data.PoolAddress), // Generate a simple pool model for general pool info
	}

	// Respond with defined data
	c.JSON(http.StatusOK, gin.H{
		"lp_response": response,
	})

}

// Get LP data from one pool
func GetFarmOfProtocol(c *gin.Context) {
	// Declare request model to use
	var data models.GetFarmOfProtocolRequest

	// Get data off header: user_address, protocol_id, pool_address
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	// Read data from pool address and calculate LP balance
	token0BalOfUser, token1BalOfUser := utils.GetFarmTokenPairBalOfUser(data.UserAddress, data.GaugeAddress)

	// Create response & Attach data to response struct
	response := models.LPFarmResponse{
		UserAddress: data.UserAddress,
		Token0:      token0BalOfUser,
		Token1:      token1BalOfUser,
		ProtocolId:  data.ProtocolId,
		Pool:        utils.GetPoolFromGauge(data.GaugeAddress), // Generate a simple pool model for general pool info
	}

	// Respond with defined data
	c.JSON(http.StatusOK, gin.H{
		"farm_response": response,
	})

}

func getParamFromURLQuery(c *gin.Context, param_name string) string {
	query := c.Request.URL.Query()
	return query.Get("user_address")
}

// Utils
func GetAllChainIds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pair_ids": models.PairIds,
	})
}

func GetAllProtocolIds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pair_ids": models.PairIds,
	})
}

func GetAllPairIds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pair_ids": models.PairIds,
	})
}

/* End of new APIs*/
/**********************************************************************/

// TestAddress: 0x043220ac21c0ce367689d93822ad70fe95ea8d2e
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

	user.Token0Bal = result0.String()
	// user.Token0Bal = int(result0.Int64())

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

	user.Token1Bal = result1.String()
	// user.Token1Bal = int(result1.Int64())

	// user.Token1Bal = new(big.Float).Mul(new(big.Float).SetInt(token1Bal), lpRatio)

	/* End of reading contract*/
	/**********************************************************************/

	/**********************************************************************/
	/* Query data from coingecko */

	// WBNB to USD
	price0, price1 := utils.GetUsdPriceInPair("wbnb", "busd")
	user.Token0BalInUsd = utils.WeiToEth(result0) * price0

	// BUSD to USD

	user.Token1BalInUsd = utils.WeiToEth(result1) * price1

	/* End of querying data from coingecko*/
	/**********************************************************************/

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
		"user":    user,
	})
}

// TestAddress: 0x972b0f9cde1266e860e546ac92e783741769400f
func GetUserInfoFarming(c *gin.Context) {
	// Declare a model to response
	var user models.UserInformation
	user.Protocol = "DDDX" // Default
	user.Tag = "Farming"   // Default

	// Get data off body req
	walletAddr := common.HexToAddress(c.Param("usrAddr"))
	user.Address = walletAddr.Hex()

	/**********************************************************************/
	/* Reading contract */
	pairGaugeAddress := common.HexToAddress("0x928180c1BdEdb6eb2ED9625C66F45757E4aBc623") // gauge-vWBNB/BUSD
	instance, err := gauge.NewGauge(pairGaugeAddress, initializers.Client)
	if err != nil {
		log.Fatal(err)
	}

	// Get pair contract address from gauge contract
	pairAddr, err := instance.Stake(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	instancePair, err := pair.NewPair(pairAddr, initializers.Client)
	if err != nil {
		log.Fatal(err)
	}

	user.PairName, err = instancePair.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	user.PairSymbol, err = instancePair.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get wallet's LP token balance
	bal, err := instance.BalanceOf(&bind.CallOpts{}, walletAddr)
	if err != nil {
		log.Fatal(err)
	}

	// After getting wallet's LP token balance, calculate each token in pair as normal liquidity provide
	// Get total supply of LP tokens
	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get the wallet's LP ratio
	lpRatio := new(big.Float).Quo(new(big.Float).SetInt(bal), new(big.Float).SetInt(totalSupply))
	fmt.Println(lpRatio)

	// Get pool's token0 balance
	// Ex: WBNB

	// Get token0's address
	token0Addr, err := instancePair.Token0(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get pool's balance
	token0Instance, err := token.NewToken(common.HexToAddress(token0Addr.Hex()), initializers.Client)
	if err != nil {
		log.Fatal(err)
	}

	token0Bal, err := token0Instance.BalanceOf(&bind.CallOpts{}, pairAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token0Bal)

	token0BalRatio := new(big.Float).Mul(new(big.Float).SetInt(token0Bal), lpRatio)
	fmt.Println(token0BalRatio)
	result0 := new(big.Int)
	token0BalRatio.Int(result0)

	user.Token0Bal = result0.String()

	// Get pool's token1 balance
	// Ex: WBNB
	// Get token0's address
	token1Addr, err := instancePair.Token1(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Get pool's balance
	token1Instance, err := token.NewToken(common.HexToAddress(token1Addr.Hex()), initializers.Client)
	if err != nil {
		log.Fatal(err)
	}

	token1Bal, err := token1Instance.BalanceOf(&bind.CallOpts{}, pairAddr)
	if err != nil {
		log.Fatal(err)
	}

	_ = token1Bal

	token1BalRatio := new(big.Float).Mul(new(big.Float).SetInt(token1Bal), lpRatio)
	result1 := new(big.Int)
	token1BalRatio.Int(result1)

	user.Token1Bal = result1.String()

	// user.Token1Bal = new(big.Float).Mul(new(big.Float).SetInt(token1Bal), lpRatio)

	/* End of reading contract*/
	/**********************************************************************/

	/**********************************************************************/
	/* Query data from coingecko */

	// WBNB to USD
	price0, price1 := utils.GetUsdPriceInPair("wbnb", "busd")
	user.Token0BalInUsd = utils.WeiToEth(result0) * price0

	// BUSD to USD

	user.Token1BalInUsd = utils.WeiToEth(result1) * price1

	/* End of querying data from coingecko*/
	/**********************************************************************/

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
		"user":    user,
	})
}

type UserInformationLocked struct {
	Amount *big.Int
	End    *big.Int
}

// TestAddr: 0x94fac6b9634f00801b122e2c3dfe1c29b44cda25
func GetUserInfoLocked(c *gin.Context) {
	var result []UserInformationLocked

	// Declare a model to response
	var user models.UserInformation
	user.Protocol = "DDDX"         // Default
	user.Tag = "Liquidity Provide" // Default

	// Get data off body req
	walletAddr := common.HexToAddress(c.Param("usrAddr"))
	user.Address = walletAddr.Hex()

	veNFTAddress := common.HexToAddress("0xFe9e21e78089094E1443169c4c74bBBBcBb13DE0") // gauge-vWBNB/BUSD
	instance, err := ve.NewVe(veNFTAddress, initializers.Client)
	if err != nil {
		log.Fatal(err)
	}

	// Get user's total veNFT
	userBal, err := instance.BalanceOf(&bind.CallOpts{}, walletAddr)
	if err != nil {
		log.Fatal(err)
	}

	num := int(userBal.Int64())

	for i := 0; i < num; i++ {
		set_i := big.NewInt(int64(i))
		var test UserInformationLocked
		tokenId, err := instance.TokenOfOwnerByIndex(&bind.CallOpts{}, walletAddr, set_i)
		if err != nil {
			log.Fatal(err)
		}

		test, err = instance.Locked(&bind.CallOpts{}, tokenId)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, test)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "successful",
		"wallet":   walletAddr,
		"protocol": "DDDX",
		"tag":      "Locked",
		"token":    "DDDX",
		"details":  result,
	})

}
