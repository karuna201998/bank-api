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
	r.GET("/accounts/:id", routes.GetAccount)
	r.PUT("/accounts/:id", routes.UpdateAccount)
	r.DELETE("/accounts/:id", routes.DeleteAccount)

	r.Run(":8081") // Start server on port 8081
}
