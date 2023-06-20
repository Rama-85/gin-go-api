package main

import (
	"gin-go-api/config"
	"gin-go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.CalRoute(router)
	router.Run(":8080")
}
