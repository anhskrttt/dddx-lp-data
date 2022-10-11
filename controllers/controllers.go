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

// GetAllLiquidityPoolsOfProtocol returns a list of available pools of protocol.
// Note: ~8s
// Refactor 01: ~16s - save factory contract address to local storage
func GetAllLiquidityPoolsOfProtocol(c *gin.Context) {
	// Get data off url query: protocol_id
	switch protocolId := c.Query("protocol_id"); protocolId {
	case "dddx":
		// Query all pool addresses of dddx
		len, pools := utils.GetAllLiquidityPools(utils.PAIR_FACTORY_CONTRACT_ADDRESS)

		response := models.AllPoolsOfProtocolResponse{
			ProtocolId: protocolId,
			Length:     len,
			Pools:      pools,
		}

		c.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	case "aave":
		c.JSON(http.StatusOK, gin.H{
			"response": "coming soon",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"response": "unsupported protocol",
		})
	}
}

// Original query time: ~40s
// Refactor 01: ~40s
// Refactor 02: ~32.28s - Get rid of multiple call to create contract instance
// Refactor 03: ~32.87s - Save available pools to local storage
// Refactor 04: ~34.27s - Get rid of multiple call to create contract instance
// Refactor 05: ~44.88s - Iterate through poolAddresses to find both LP and farm bal
// Refactor 06: ~16.51s - Save LP pool addresses & gauge addresses to local storage
// Refactor 07: ~15.02s - Using global contract instances
// Refactor 08: ~18.83s
func GetAllBalOfUser(c *gin.Context) {
	// Get data off url query: user_address
	userAddress := c.Query("user_address")

	// Get data off url query: protocol_id
	switch protocolId := c.Query("protocol_id"); protocolId {
	case "dddx":

		// Get LP data
		lpBals := []models.AllLPFarmBalResponse{}
		for _, poolAddress := range utils.LPPools {
			poolInstance := utils.GetPoolInstance(poolAddress)
			userBal := utils.GetLPBalFromPoolInstance(userAddress, poolInstance)

			// If bal != 0, add to []responses
			if len(userBal.Bits()) != 0 {
				// Read data from pool address and calculate LP balance
				token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUserFrom(userBal, poolAddress, poolInstance)

				// Create response & Attach data to response struct
				sampleResponse := models.AllLPFarmBalResponse{
					Token0: token0BalOfUser,
					Token1: token1BalOfUser,
					Pool:   utils.GetPoolFromInstance(poolAddress, poolInstance), // Generate a simple pool model for general pool info
				}

				lpBals = append(lpBals, sampleResponse)
			}
		}

		// Get Farming data
		// Check bal of each gauge
		farmBals := []models.AllLPFarmBalResponse{}
		for _, gaugeAddress := range utils.Gauges {
			gauge := utils.GetGaugeInstance(gaugeAddress)

			// Get gauge's bal of user
			bal := utils.GetFarmBalFromGauge(userAddress, gauge)

			if len(bal.Bits()) != 0 {
				// Create pair instance
				pool := utils.GetPoolInstanceFromGauge(gauge)
				poolAddress := utils.GetPoolAddressFromGauge(gauge)
				fmt.Println(poolAddress)

				// Read data from pool address and calculate LP balance
				token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUserFrom(bal, poolAddress, pool)

				// Create response & Attach data to response struct
				sampleResponse := models.AllLPFarmBalResponse{
					Token0: token0BalOfUser,
					Token1: token1BalOfUser,
					Pool:   utils.GetPoolFromInstance(poolAddress, pool), // Generate a simple pool model for general pool info
				}

				farmBals = append(farmBals, sampleResponse)
			}

		}

		// Get Staked data
		staked_balances := utils.GetStakedBalancesOf(userAddress, utils.DDDXVeNFT)

		// Create response & Attach data to response struct
		stakedBalance := models.AllStakeResponse{
			TokenSymbol:    utils.GetTokenSymbolFromVeInstance(utils.DDDXVeNFT),
			Staked_balance: staked_balances,
		}

		response := models.AllBalOfUserResponse{
			UserAddress:   userAddress,
			ProtocolId:    protocolId,
			LPBalance:     lpBals,
			FarmBalance:   farmBals,
			StakedBalance: stakedBalance,
		}

		c.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	case "aave":
		c.JSON(http.StatusOK, gin.H{
			"response": "coming soon",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"response": "unsupported protocol",
		})
	}
}

// GetAllLpOfProtocol returns all pools user's currently providing liquidity.
// Default pool is DDDX.
// Time response: 40.17s
// Refactor 01: 29.31s - Avoid re-create contract instances
// Refactor 01: 40.52s - Add switch case
func GetAllLpBalOfProtocol(c *gin.Context) {
	// Get data off url query: user_address
	userAddress := c.Query("user_address")

	// Get data off url query: protocol_id
	switch protocolId := c.Query("protocol_id"); protocolId {
	case "dddx":
		var responses []models.AllLPFarmBalResponse

		for _, poolAddress := range utils.LPPools {
			bal := utils.GetLPBalFromPool(userAddress, poolAddress)

			// If bal != 0, add to []responses
			if len(bal.Bits()) != 0 {
				// Get pool instance
				poolInstance := utils.GetPoolInstance(poolAddress)

				// Read data from pool address and calculate LP balance
				token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUserFrom(bal, poolAddress, poolInstance)

				// Create response & Attach data to response struct
				response := models.AllLPFarmBalResponse{
					Token0: token0BalOfUser,
					Token1: token1BalOfUser,
					Pool:   utils.GetPoolFromInstance(poolAddress, poolInstance), // Generate a simple pool model for general pool info
				}

				responses = append(responses, response)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"user_address": userAddress,
			"protocol_id":  protocolId,
			"response":     responses,
		})
	case "aave":
		c.JSON(http.StatusOK, gin.H{
			"response": "coming soon",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"response": "unsupported protocol",
		})
	}
}

// Note: Original time : 19.38s
// Refactor 01: 21.98s - Save gauges to local storage
func GetAllFarmBalOfProtocol(c *gin.Context) {
	var responses []models.AllLPFarmBalResponse

	// Get data off url query: user_address
	userAddress := c.Query("user_address")

	// Get data off url query: protocol_id
	switch protocolId := c.Query("protocol_id"); protocolId {
	case "dddx":
		// check bal of each gauge
		for _, gaugeAddress := range utils.Gauges {
			bal := utils.GetFarmBalFromGaugeAddress(userAddress, gaugeAddress)

			if len(bal.Bits()) != 0 {
				// Read data from pool address and calculate LP balance
				token0BalOfUser, token1BalOfUser := utils.GetTokenPairBalOfUser(userAddress, gaugeAddress, true)

				// Create response & Attach data to response struct
				response := models.AllLPFarmBalResponse{
					Token0: token0BalOfUser,
					Token1: token1BalOfUser,
					Pool:   utils.GetPoolFromGaugeAddress(gaugeAddress), // Generate a simple pool model for general pool info
				}

				responses = append(responses, response)
			}

		}

		c.JSON(http.StatusOK, gin.H{
			"user_address": userAddress,
			"protocol_id":  protocolId,
			"response":     responses,
		})

	case "aave":
		c.JSON(http.StatusOK, gin.H{
			"response": "coming soon",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"response": "unsupported protocol",
		})
	}
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
		Pool:        utils.GetPoolFromGaugeAddress(data.GaugeAddress), // Generate a simple pool model for general pool info
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
		Pool:        utils.GetPoolFromGaugeAddress(data.GaugeAddress), // Generate a simple pool model for general pool info
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
	// Get data off url query: user_address
	userAddress := c.Query("user_address")

	// Get data off url query: protocol_id
	switch protocolId := c.Query("protocol_id"); protocolId {
	case "dddx":
		staked_balances := utils.GetStakedBalances(userAddress, utils.VE_NFT_CONTRACT_ADDRESS)

		// Create response & Attach data to response struct
		response := models.StakeResponse{
			UserAddress:    userAddress,
			TokenSymbol:    utils.GetTokenSymbolFromVeAddress(utils.VE_NFT_CONTRACT_ADDRESS),
			Staked_balance: staked_balances,
		}

		// Respond with defined data
		c.JSON(http.StatusOK, gin.H{
			"staked_response": response,
		})

	case "aave":
		c.JSON(http.StatusOK, gin.H{
			"response": "coming soon",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"response": "unsupported protocol",
		})
	}

}
