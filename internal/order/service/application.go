package service

import (
	"context"
	"github.com/mfzzf/gorder-v2/order/adapters"
	"github.com/mfzzf/gorder-v2/order/app"
)

func NewApplication(context context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	return app.Application{}
}
