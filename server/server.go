package server

import (
	"fmt"

	"github.com/anbangisak/gunaanban_Challenge/config"
)

func Init() {
	config := config.GetConfig()
	route := NewRouter()
	route.Run(fmt.Sprintf(":%v", config.GetString("server.port")))
	// route.Run(":8080")
}
