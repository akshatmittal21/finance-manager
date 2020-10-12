package main

import (
	"akshu-21/finance-manager/datamdl"
	"akshu-21/finance-manager/model"
	"akshu-21/finance-manager/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	_, configReadError := datamdl.InitConfig(model.GetConfigFilePath(), &model.Config)
	if configReadError != nil {
		log.Error("configReadError: ", configReadError)
		return
	}
	mongoInitError := datamdl.InitMongo()
	if mongoInitError != nil {
		log.Error("mongoInitError : ", mongoInitError)
		return
	}
	initServer()

}
func initServer() {
	r := gin.Default()
	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	r.Use(cors.New(md))

	routes.InitRoutes(r)
	s := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	s.ListenAndServe()
}
