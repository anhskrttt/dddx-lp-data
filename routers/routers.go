package routers

import (
	"dddx-lp-data/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() http.Handler {
	r := gin.Default()

	// My API
	r.GET("/ping", controllers.Ping)
	r.GET("/api/lp/:usrAddr", controllers.GetUserInfo)
	r.GET("/api/farm/:usrAddr", controllers.GetUserInfoFarming)
	r.GET("/api/lock/:usrAddr", controllers.GetUserInfoLocked)

	return r
}
