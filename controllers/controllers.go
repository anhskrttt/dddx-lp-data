package controllers

import (
	"dddx-lp-data/models"
	"dddx-lp-data/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is a handler func for testing purpose.
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "pong",
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
