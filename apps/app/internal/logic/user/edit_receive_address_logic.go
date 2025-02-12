package user

import (
	"context"

	"github.com/Sanagiig/lebron/apps/app/internal/svc"
	"github.com/Sanagiig/lebron/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditReceiveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// edit user receiveAddress
func NewEditReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditReceiveAddressLogic {
	return &EditReceiveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditReceiveAddressLogic) EditReceiveAddress(req *types.UserReceiveAddressEditReq) (resp *types.UserReceiveAddressEditRes, err error) {
	// todo: add your logic here and delete this line

	return
}
