package router

import (
	"ecommerce/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./frontend")
	router.GET("/getalldata", controller.Getalldata)
	router.POST("/create", controller.CreateProfile)

	return router
}
