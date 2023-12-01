package middlewares

import (
	"fmt"

	"github.com/anbangisak/gunaanban_Challenge/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		fmt.Println(config)
		// handles next pending handlers
		c.Next()
	}
}
