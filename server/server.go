package server

import (
	"fmt"

	"github.com/anbangisak/gunaanban_Challenge/config"
)

func Init() {
	config := config.GetConfig()
	route := NewRouter()
	// httpRouter := gin.New()
	// httpRouter.GET("/*path", func(c *gin.Context) {
	// 	fmt.Println(c.Param("variable"))
	// 	c.Redirect(302, "https://localhost:"+config.GetString("server.https_port")+"/"+c.Param("variable"))
	// })
	go route.RunTLS(fmt.Sprintf(":%v", config.GetString("server.https_port")), "localhost.crt", "localhost.key")
	route.Run(fmt.Sprintf(":%v", config.GetString("server.http_port")))

	// route.Run(":8080")
}
