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

func GetAllAccounts(c *gin.Context) {
	var accounts []models.Account
	if err := config.DB.Find(&accounts).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch accounts"})
		return
	}
	c.JSON(200, accounts)
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
