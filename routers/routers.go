package routers

import (
	"dddx-lp-data/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() http.Handler {
	r := gin.Default()

	// My API
	r.GET("/api/:usrAddr", controllers.GetUserInfo)

	return r
}
