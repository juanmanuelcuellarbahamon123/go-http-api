package routes

import (
	"go-http-api/controllers/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {

	userGroup := route.Group("/user")
	{
		userGroup.GET("/", user.UserFunction)
	}

}
