package logic

import (
	"context"
	"github.com/Sanagiig/lebron/apps/product/rpc/model"
	"github.com/zeromicro/go-zero/core/mr"
	"strconv"
	"strings"

	"github.com/Sanagiig/lebron/apps/product/rpc/internal/svc"
	"github.com/Sanagiig/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {
	products := make(map[int64]*product.ProductItem)
	pdis := strings.Split(in.ProductIds, ",")
	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pdis {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer[any], cancel func(error)) {
		pidStr := item.(string)
		pid, err := strconv.ParseInt(pidStr, 10, 64)
		if err != nil {
			return
		}
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(pid))
		if err != nil {
			return
		}
		writer.Write(p)
	}, func(pipe <-chan interface{}, writer mr.Writer[any], cancel func(error)) {
		var r []*model.Product
		for p := range pipe {
			r = append(r, p.(*model.Product))
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	for _, p := range ps.([]*model.Product) {
		products[int64(p.Id)] = &product.ProductItem{
			ProductId: int64(p.Id),
			Name:      p.Name,
		}
	}
	return &product.ProductResponse{Products: products}, nil
}
