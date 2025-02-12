package logic

import (
	"context"

	"github.com/Sanagiig/lebron/apps/order/rpc/internal/svc"
	"github.com/Sanagiig/lebron/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *order.OrdersRequest) (*order.OrdersResponse, error) {
	uid := int64(123)
	if in.UserId == uid {
		orders := []*order.OrderItem{
			{
				OrderId:  "20220609123456",
				UserId:   uid,
				Quantity: 1,
			},
		}
		return &order.OrdersResponse{Orders: orders}, nil
	}
	return &order.OrdersResponse{}, nil
}
