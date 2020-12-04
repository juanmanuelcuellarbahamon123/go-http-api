package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"controller": "auth",
	})
}
