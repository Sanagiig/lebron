package svc

import (
	"github.com/Sanagiig/lebron/apps/product/rpc/internal/config"
	"github.com/Sanagiig/lebron/apps/product/rpc/model"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
	"time"
)

const localCacheExpire = time.Duration(time.Second * 60)

type ServiceContext struct {
	Config         config.Config
	ProductModel   model.ProductModel
	CategoryModel  model.CategoryModel
	OperationModel model.ProductOperationModel
	SingleGroup    singleflight.Group
	LocalCache     *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:         c,
		LocalCache:     localCache,
		ProductModel:   model.NewProductModel(conn),
		CategoryModel:  model.NewCategoryModel(conn),
		OperationModel: model.NewProductOperationModel(conn),
	}
}
