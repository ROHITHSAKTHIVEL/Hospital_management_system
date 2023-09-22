package controller

import (
	"ecommerce/models"
	"ecommerce/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getalldata(c *gin.Context) {
	alltransaction := service.Getalldata()
	c.JSON(http.StatusOK, alltransaction)
}


func CreateProfile(c *gin.Context) {
	fmt.Println("Creating Profile")
	var profile models.Customer
	fmt.Println("Created")
	if err := c.BindJSON(&profile); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println("json binded")
	fmt.Println(profile)
	service.Insert(profile)
	c.JSON(http.StatusOK, profile)
}

func Appoitment(c *gin.Context) {
	fmt.Println("Creating Profile")
	var profile models.Appoitment
	fmt.Println("Created")
	if err := c.BindJSON(&profile); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println("json binded")
	fmt.Println(profile)
	service.Appoitment(profile)
	c.JSON(http.StatusOK, profile)
}
