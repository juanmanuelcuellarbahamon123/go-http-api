package routes

import (
	"go-http-api/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {

	authGroup := route.Group("/auth")
	{
		authGroup.GET("/", auth.AuthFunction)
	}

}
