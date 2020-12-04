package user

import (
	"github.com/gin-gonic/gin"
)

func UserFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"controller": "user",
	})
}
