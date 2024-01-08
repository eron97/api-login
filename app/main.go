package main

import (
	"github.com/eron97/api-login.git/controllers/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	router.Run(":8080")
}
