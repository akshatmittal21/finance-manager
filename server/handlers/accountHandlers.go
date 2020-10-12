package handlers

import (
	"akshu-21/finance-manager/model"
	"akshu-21/finance-manager/services"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

// GetAllAccounts : to get all configured accounts
func GetAllAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := services.GetAllAccounts()
		if err != nil {
			log.Error("GetAllAccounts: unable_to_fetch_accounts")
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

//GetAccount : to get all configured accounts
func GetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqst := model.Account{}
		c.Bind(&rqst)
		data, err := services.GetAccount(&rqst)
		if err != nil {
			log.Error("GetAccount: unable_to_fetch_account")
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

// AddNewAccount : to add a new account
func AddNewAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Account{}
		c.Bind(&rqstData)
		err := services.AddNewAccount(&rqstData)
		if err != nil {
			log.Error("AddNewAccount: unable_to_add_new_account")
			return
		}
		c.JSON(http.StatusOK, "Account added successfully !!")
	}
}

// UpdateAccount : to add a new account
func UpdateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Account{}
		c.Bind(&rqstData)
		err := services.UpdateAccount(&rqstData)
		if err != nil {
			log.Error("UpdateAccount: unable_to_add_new_account")
			return
		}
		c.JSON(http.StatusOK, "Account updated successfully !!")
	}
}

//DeleteAccount : to add a new account
func DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Account{}
		c.Bind(&rqstData)
		err := services.DeleteAccount(&rqstData)
		if err != nil {
			log.Error("DeleteAccount: unable_to_delete_account")
			return
		}
		c.JSON(http.StatusOK, "Account deleted successfully !!")
	}
}
