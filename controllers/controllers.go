package controllers

import (
	"dddx-lp-data/models"
	"dddx-lp-data/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a handler func for testing purpose.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "pong",
	})
}

func GetAllLiquidityPoolsOfProtocol(c *gin.Context) {
	var pools []string

	// Get data off url query: user_address
	protocolId := c.Query("protocol_id")

	if protocolId != "dddx" {
		c.JSON(http.StatusOK, gin.H{
			"response": "unsupported protocol",
		})
		return
	}

	// Query all pool addresses of dddx
	poolLength, pools := utils.GetAllLiquidityPools("0xb5737A06c330c22056C77a4205D16fFD1436c81b")

	c.JSON(http.StatusOK, gin.H{
		"protocol_id": protocolId,
		"length":      poolLength,
		"pools":       pools,
	})
}

// GetAllLpOfProtocol returns all pools user's currently providing liquidity.
// Default pool is DDDX.
func GetAllLpBalOfProtocol(c *gin.Context) {
	// Get data off url query: user_address, protocol_id
	userAddress := c.Query("user_address")
	protocolId := c.Query("protocol_id")

	if protocolId != "dddx" {
		c.JSON(http.StatusOK, gin.H{
			"response": "unsupported protocol",
		})
		return
	}

	// Get all liquidity pools
	_, pools := utils.GetAllLiquidityPools("0xb5737A06c330c22056C77a4205D16fFD1436c81b")

	var responses []models.AllLPFarmBalResponse

	for _, poolAddress := range pools {
		bal := utils.GetLPBalFromPool(userAddress, poolAddress)
		fmt.Println(bal)

		// If bal != 0, add to []responses
		if len(bal.Bits()) != 0 {
			// Read data from pool address and calculate LP balance
			token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUserFromPoolAddress(userAddress, poolAddress)

			// Create response & Attach data to response struct
			response := models.AllLPFarmBalResponse{
				Token0: token0BalOfUser,
				Token1: token1BalOfUser,
				Pool:   utils.GetPoolFromAddress(poolAddress), // Generate a simple pool model for general pool info
			}

			responses = append(responses, response)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"user_address": userAddress,
		"protocol_id":  protocolId,
		"response":     responses,
	})
}

func GetAllFarmBalOfProtocol(c *gin.Context) {
	var responses []models.AllLPFarmBalResponse

	// Get data off url query: user_address, protocol_id
	userAddress := c.Query("user_address")
	protocolId := c.Query("protocol_id")

	if protocolId != "dddx" {
		c.JSON(http.StatusOK, gin.H{
			"response": "unsupported protocol",
		})
		return
	}

	// Get all liquidity pools
	// _, pools := utils.GetAllLiquidityPools("0xb5737A06c330c22056C77a4205D16fFD1436c81b")

	// Get all gauges address
	_, gauges := utils.GetAllGauges("0xAd8Ab2C2270Ab0603CFC674d28fd545495369f31", "0xb5737A06c330c22056C77a4205D16fFD1436c81b")

	// check bal of each gauge
	for _, gaugeAddress := range gauges {
		bal := utils.GetFarmBalFromGauge(userAddress, gaugeAddress)

		if len(bal.Bits()) != 0 {
			// Read data from pool address and calculate LP balance
			token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUser(userAddress, gaugeAddress, true)

			// Create response & Attach data to response struct
			response := models.AllLPFarmBalResponse{
				Token0: token0BalOfUser,
				Token1: token1BalOfUser,
				Pool:   utils.GetPoolFromGauge(gaugeAddress), // Generate a simple pool model for general pool info
			}

			responses = append(responses, response)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"user_address": userAddress,
		"protocol_id":  protocolId,
		"response":     responses,
	})
}

// TestAddress: 0x043220ac21c0ce367689d93822ad70fe95ea8d2e
// GetLpOfProtocol get LP/Farming data from one pool.
// Default pool is DDDX.
func GetLpOfProtocol(c *gin.Context) {
	// Declare request model to use
	var data models.LPRequest

	// Get data off header: user_address, protocol_id, pool_address
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	// Read data from pool address and calculate LP balance
	token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUser(data.UserAddress, data.GaugeAddress, false)

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
		"lp_response": response,
	})

}

// TestAddress: 0x972b0f9cde1266e860e546ac92e783741769400f
// GetFarmOfProtocol get Farming data from one pool.
// Default pool is DDDX.
func GetFarmOfProtocol(c *gin.Context) {
	// Declare request model to use
	var data models.LPRequest

	// Get data off header: user_address, protocol_id, pool_address
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	// Read data from pool address and calculate LP balance
	token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUser(data.UserAddress, data.GaugeAddress, true)

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

// TestAddr: 0x94fac6b9634f00801b122e2c3dfe1c29b44cda25
// GetStakedOfProtocol get staked data from one pool.
// Default pool is DDDX.
func GetStakedOfProtocol(c *gin.Context) {
	// Declare request model to use
	var data models.GetStakedOfProtocolRequest

	// Get data off header: user_address, protocol_id, pool_address
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	staked_balances := utils.GetStakedBalances(data.UserAddress, data.VeAddress)

	// Create response & Attach data to response struct
	response := models.StakeResponse{
		UserAddress:    data.UserAddress,
		TokenSymbol:    utils.GetTokenSymbolFromVeAddress(data.VeAddress),
		Staked_balance: staked_balances,
	}

	// Respond with defined data
	c.JSON(http.StatusOK, gin.H{
		"staked_response": response,
	})

}
