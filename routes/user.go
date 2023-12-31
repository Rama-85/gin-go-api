package routes

import (
	"gin-go-api/controller"

	"github.com/gin-gonic/gin"
)

func CalRoute(router *gin.Engine) {
	router.GET("/", controller.GetDays)
	router.POST("/", controller.CreateMonth)
	router.DELETE("/:id", controller.DeleteMonth)
	router.PUT("/:id", controller.UpdateMonth)
}
