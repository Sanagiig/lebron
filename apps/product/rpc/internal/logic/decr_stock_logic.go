package logic

import (
	"context"

	"github.com/Sanagiig/lebron/apps/product/rpc/internal/svc"
	"github.com/Sanagiig/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockLogic {
	return &DecrStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockLogic) DecrStock(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
	// todo: add your logic here and delete this line

	return &product.DecrStockResponse{}, nil
}
