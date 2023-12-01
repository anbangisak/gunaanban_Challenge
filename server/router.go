package server

import (
	"github.com/gin-gonic/gin"

	"github.com/anbangisak/gunaanban_Challenge/controllers"
	"github.com/anbangisak/gunaanban_Challenge/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("./templates/*")

	dashboard := new(controllers.DashboardController)

	router.GET("/dashboard", dashboard.Display)
	router.GET("/health", dashboard.Health)
	router.Use(middlewares.AuthMiddleware())

	return router
}
