package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfzzf/gorder-v2/common/config"
	"github.com/mfzzf/gorder-v2/common/server"
	"github.com/mfzzf/gorder-v2/order/ports"
	"github.com/spf13/viper"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serviceName := viper.GetString("order.service-name")
	server.RunHTTPServer(serviceName, func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router, HTTPServer{}, ports.GinServerOptions{
			BaseURL:      "/api",
			Middlewares:  nil,
			ErrorHandler: nil,
		})
	})
}
