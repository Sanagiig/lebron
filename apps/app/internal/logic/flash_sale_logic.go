package logic

import (
	"context"

	"github.com/Sanagiig/lebron/apps/app/internal/svc"
	"github.com/Sanagiig/lebron/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashSaleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 限时抢购
func NewFlashSaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashSaleLogic {
	return &FlashSaleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashSaleLogic) FlashSale() (resp *types.FlashSaleResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
