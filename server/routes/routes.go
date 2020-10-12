package routes

import (
	"akshu-21/finance-manager/handlers"
	"github.com/gin-gonic/gin"
)

// InitRoutes : To initialize routes
func InitRoutes(r *gin.Engine) {

	// accounts routes
	r.GET("/getAllAccounts", handlers.GetAllAccounts())
	r.POST("/getAccount", handlers.GetAccount())
	r.POST("/addNewAccount", handlers.AddNewAccount())
	r.POST("/updateAccount", handlers.UpdateAccount())
	r.POST("/deleteAccount", handlers.DeleteAccount())

	// transactions routes
	r.GET("/getAllTransactions", handlers.GetAllTransactions())
	r.POST("/getTransaction", handlers.GetTransaction())
	r.POST("/addNewTransaction", handlers.AddNewTransaction())
	r.POST("/updateTransaction", handlers.UpdateTransaction())
	r.POST("/deleteTransaction", handlers.DeleteTransaction())

	//category routes
	r.GET("/getAllCategories", handlers.GetAllCategories())
	r.POST("/getCategory", handlers.GetCategory())
	r.POST("/addNewCategory", handlers.AddNewCategory())
	r.POST("/updateCategory", handlers.UpdateCategory())
	r.POST("/deleteCategory", handlers.DeleteCategory())

	//cdn handler
	r.POST("/uploadFile", handlers.UploadFileHandler())
	r.Static("/files/", handlers.UploadPath)

}
