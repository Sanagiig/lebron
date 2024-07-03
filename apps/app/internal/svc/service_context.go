package svc

import (
	"github.com/Sanagiig/lebron/apps/app/internal/config"
	"github.com/Sanagiig/lebron/apps/order/rpc/order_client"
	"github.com/Sanagiig/lebron/apps/product/rpc/product_client"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order_client.Order
	ProductRPC product_client.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order_client.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product_client.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
