package handlers

import (
	"akshu-21/finance-manager/model"
	"akshu-21/finance-manager/services"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

// GetAllCategories : to get all configured Categories
func GetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := services.GetAllCategories()
		if err != nil {
			log.Error("GetAllCategories: unable_to_fetch_Categories")
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

// GetCategory : to get all configured categories
func GetCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqst := model.Category{}
		c.Bind(&rqst)
		data, err := services.GetCategory(&rqst)
		if err != nil {
			log.Error("GetCategory: unable_to_fetch_category")
			return
		}
		c.JSON(http.StatusOK, data)
	}
}

// AddNewCategory : to add a new Category
func AddNewCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Category{}
		c.Bind(&rqstData)
		err := services.AddNewCategory(&rqstData)
		if err != nil {
			log.Error("AddNewCategory: unable_to_add_new_Category")
			return
		}
		c.JSON(http.StatusOK, "Category added successfully !!")
	}
}

// UpdateCategory :  to add a new Category
func UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Category{}
		c.Bind(&rqstData)
		err := services.UpdateCategory(&rqstData)
		if err != nil {
			log.Error("UpdateCategory: unable_to_add_new_Category")
			return
		}
		c.JSON(http.StatusOK, "Category updated successfully !!")
	}
}

// DeleteCategory : to add a new Category
func DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		rqstData := model.Category{}
		c.Bind(&rqstData)
		err := services.DeleteCategory(&rqstData)
		if err != nil {
			log.Error("DeleteCategory: unable_to_delete_Category")
			return
		}
		c.JSON(http.StatusOK, "Category deleted successfully !!")
	}
}
