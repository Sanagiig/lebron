package user

import (
	"context"

	"github.com/Sanagiig/lebron/apps/app/internal/svc"
	"github.com/Sanagiig/lebron/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollectionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// add user collection
func NewUserCollectionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionAddLogic {
	return &UserCollectionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionAddLogic) UserCollectionAdd(req *types.UserCollectionAddReq) (resp *types.UserCollectionAddRes, err error) {
	// todo: add your logic here and delete this line

	return
}
