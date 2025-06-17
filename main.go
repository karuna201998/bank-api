package main

import (
	"go-crud-poc/config"
	"go-crud-poc/models"
	"go-crud-poc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Account{})

	r := gin.Default()

	// Register API endpoints
	r.POST("/accounts", routes.CreateAccount)
	r.GET("/accounts", routes.GetAllAccounts)
	r.PUT("/accounts/:id", routes.UpdateAccount)
	r.DELETE("/accounts/:id", routes.DeleteAccount)

	for _, route := range r.Routes() {
		println("Route:", route.Method, route.Path)
	}

	r.Run(":8088") // Start server on port 8081
}
