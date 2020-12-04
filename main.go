package main

import (
	"go-http-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.UserRoutes(router)

	router.Run(":4000")

}
