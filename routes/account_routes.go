package routes

import (
	"net/http"
	"strconv"

	"go-crud-poc/config"
	"go-crud-poc/models"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&account)
	c.JSON(http.StatusCreated, account)
}

func GetAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var account models.Account
	result := config.DB.First(&account, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	c.JSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var account models.Account
	result := config.DB.First(&account, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	var updated models.Account
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account.AccountNumber = updated.AccountNumber
	account.AccountHolder = updated.AccountHolder
	account.Balance = updated.Balance

	config.DB.Save(&account)
	c.JSON(http.StatusOK, account)
}

func DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var account models.Account
	result := config.DB.First(&account, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	config.DB.Delete(&account)
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
}
