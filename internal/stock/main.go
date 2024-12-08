package main

import (
	"github.com/mfzzf/gorder-v2/common/genproto/stockpb"
	"github.com/mfzzf/gorder-v2/common/server"
	"github.com/spf13/viper"
)

func main() {
	serviceName := viper.GetString("stock.service-name")

	server.RunGRPCServer(serviceName, stockpb.RegisterStockServiceServer())
}
