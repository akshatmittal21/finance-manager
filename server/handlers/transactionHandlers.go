package handlers

import (
	"akshu-21/finance-manager/model"
	"akshu-21/finance-manager/services"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

//GetAllTransactions : To get all configured Transactions
func GetAllTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := services.GetAllTransactions()
		if err != nil {
			log.Error("GetAllTransactions: unable_to_fetch_Transactions")
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

// GetTransaction : to get all configured Transactions
func GetTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqst := model.Transaction{}
		c.Bind(&rqst)
		data, err := services.GetTransaction(&rqst)
		if err != nil {
			log.Error("GetTransaction: unable_to_fetch_Transaction")
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

// AddNewTransaction : to add a new Transaction
func AddNewTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Transaction{}
		c.Bind(&rqstData)
		transactionID, err := services.AddNewTransaction(&rqstData)
		if err != nil {
			log.Error("AddNewTransaction: unable_to_add_new_Transaction")
			return
		}
		c.JSON(http.StatusOK, transactionID)
	}
}

// UpdateTransaction : to add a new Transaction
func UpdateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Transaction{}
		c.Bind(&rqstData)
		err := services.UpdateTransaction(&rqstData)
		if err != nil {
			log.Error("UpdateTransaction: unable_to_add_new_Transaction")
			return
		}
		c.JSON(http.StatusOK, "Transaction updated successfully !!")
	}
}

// DeleteTransaction : to add a new Transaction
func DeleteTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Transaction{}
		c.Bind(&rqstData)
		err := services.DeleteTransaction(&rqstData)
		if err != nil {
			log.Error("DeleteTransaction: unable_to_delete_Transaction")
			return
		}
		c.JSON(http.StatusOK, "Transaction deleted successfully !!")
	}
}
