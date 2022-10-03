package routers

import (
	"dddx-lp-data/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() http.Handler {
	r := gin.Default()

	// Write the log to gin.DefaultWriter
	r.Use(gin.Logger())

	// Recovery middleware
	// Recovers from any panics any writes a 500 if there was one
	r.Use(gin.Recovery())

	// My APIs
	// Test
	r.GET("/ping", controllers.Ping)
	// Utils API
	// Get pair ids
	r.GET("/api/v1/pair_ids", controllers.GetAllPairIds)

	// Get user's (all) information
	// r.GET("/api/v1/user", controllers.GetUserInfo)

	/**********************************************************************/
	/* LP Data APIs*/

	// r.GET("/api/v1/user/lps/all", controllers.GetAllLp) // Get LP data (multiple protocols from multiple chains)
	// r.GET("/api/v1/user/lps/:chain", controllers.GetAllLpOfChain) // Get LP data (multiple protocols from one chain)
	// r.GET("/api/v1/user/lps/:protocol", controllers.GetAllLPOfProtocol) // Get LP data (multiple protocols from one chain)
	r.GET("/api/v1/user/lps", controllers.GetLpOfProtocol) // Get user's specific LP pair data from specific chain

	/* End of LP Data APIs*/
	/**********************************************************************/

	/**********************************************************************/
	/* Farming Data APIs*/

	// r.GET("/api/v1/user/farms/all", controllers.GetAllLp) // Get Farming data (multiple protocols from multiple chains)
	// r.GET("/api/v1/user/farms/:chain", controllers.GetAllLpOfChain) // Get Farming data (multiple protocols from one chain)
	// r.GET("/api/v1/user/farms/:protocol", controllers.GetAllLPOfProtocol) // Get Farming data (multiple protocols from one chain)
	r.GET("/api/v1/user/farms", controllers.GetFarmOfProtocol) // Get user's specific Farming pair data from specific chain

	/* End of Farming Data APIs*/
	/**********************************************************************/

	/**********************************************************************/
	/* Staked Data APIs*/

	// r.GET("/api/v1/user/staked/all", controllers.GetAllLp) // Get Farming data (multiple protocols from multiple chains)
	// r.GET("/api/v1/user/staked/:chain", controllers.GetAllLpOfChain) // Get Farming data (multiple protocols from one chain)
	// r.GET("/api/v1/user/staked/:protocol", controllers.GetAllLPOfProtocol) // Get Farming data (multiple protocols from one chain)
	r.GET("/api/v1/user/staked", controllers.GetStakedOfProtocol) // Get user's specific Farming pair data from specific chain

	/* End of Staked Data APIs*/
	/**********************************************************************/

	return r
}
