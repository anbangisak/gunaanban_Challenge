package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"

	"github.com/anbangisak/gunaanban_Challenge/config"

	"github.com/anbangisak/gunaanban_Challenge/controllers"
	"github.com/anbangisak/gunaanban_Challenge/middlewares"
)

func NewRouter() *gin.Engine {
	config := config.GetConfig()
	sslHostName := fmt.Sprintf("localhost:%v", config.GetString("server.https_port"))
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			secureMiddleware := secure.New(secure.Options{
				SSLRedirect: true,
				SSLHost:     sslHostName,
			})
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				return
			}

			c.Next()
		}
	}()

	// router := gin.New()
	router := gin.Default()
	router.Use(secureFunc)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("./templates/*")
	router.GET("/", func(c *gin.Context) {
		c.String(200, "X-Frame-Options header is now `DENY`.")
	})

	dashboard := new(controllers.DashboardController)

	router.GET("/dashboard", dashboard.Display)
	router.GET("/health", dashboard.Health)
	router.Use(middlewares.AuthMiddleware())

	return router
}
