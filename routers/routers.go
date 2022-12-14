package routers

import (
	"dddx-lp-data/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Routers includes all api routes.
func Routers() http.Handler {
	r := gin.Default()

	// Write the log to gin.DefaultWriter
	r.Use(gin.Logger())

	// Recovery middleware
	// Recovers from any panics any writes a 500 if there was one
	r.Use(gin.Recovery())

	r.GET("/ping", controllers.Ping) // Test

	// Core APIs
	// Protocol APIs
	r.GET("/api/v1/all_pools", controllers.GetAllLiquidityPoolsOfProtocol) // Get all Liquidity pools of protocol

	// User APIs
	// [User APIs] All
	r.GET("/api/v1/user", controllers.GetAllBalOfUser) // Get all LP balances of users

	// [User APIs] LP
	r.GET("/api/v1/user/all_pools", controllers.GetAllLpBalOfProtocol) // Get all LP balances of users
	r.GET("/api/v1/user/lps", controllers.GetLpOfProtocol)             // Get user's specific LP pair data from a specific protocol

	// [User APIs] Farming
	r.GET("/api/v1/user/all_farms", controllers.GetAllFarmBalOfProtocol) // Get all LP balances of users
	r.GET("/api/v1/user/farms", controllers.GetFarmOfProtocol)           // Get user's specific Farming pair data from a specific protocol

	// [User APIs] Staked
	r.GET("/api/v1/user/staked", controllers.GetStakedOfProtocol) // Get user's specific Farming pair data from a specific protocol

	return r
}
