package router

import (
	"ecommerce/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	//router.Static("/static", ".index.html")
	router.Static("/static", "./")
	router.GET("/getalldata", controller.Getalldata)
	router.POST("/create", controller.CreateProfile)
	router.POST("/appoitnment", controller.Appoitment)
	router.POST("/login",controller.Login)


	return router
}
