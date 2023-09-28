package router

import (
	"ecommerce/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./")
	router.GET("/getalldata", controller.GetAllCustomers)
	router.GET("/getbyid",controller.GetById)
	router.GET("/viewappointment",controller.ViewAppointment)
	router.GET("/viewallfeedback",controller.ViewFeedback)
	router.POST("/create", controller.CreateProfile)
	router.POST("/appoitnment", controller.Appoitment)
	router.POST("/login", controller.Login)
	router.POST("/feedback", controller.Feedback)
	router.POST("/createadmin", controller.Createadmin)
	router.POST("/adminlogin", controller.Adminlogin)
	router.POST("/updatePatientHealth",controller.Updatepatient)
	router.DELETE("/deletebyid",controller.Deletebyid)
	router.GET("/getChartData",controller.Getgraphdata)


	return router
}
