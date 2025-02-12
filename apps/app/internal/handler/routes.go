// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	order "github.com/Sanagiig/lebron/apps/app/internal/handler/order"
	user "github.com/Sanagiig/lebron/apps/app/internal/handler/user"
	"github.com/Sanagiig/lebron/apps/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 购物车列表
				Method:  http.MethodGet,
				Path:    "/v1/cart/list",
				Handler: CartListHandler(serverCtx),
			},
			{
				// 分类商品列表
				Method:  http.MethodGet,
				Path:    "/v1/category/list",
				Handler: CategoryListHandler(serverCtx),
			},
			{
				// 限时抢购
				Method:  http.MethodGet,
				Path:    "/v1/flashsale",
				Handler: FlashSaleHandler(serverCtx),
			},
			{
				// 首页Banner
				Method:  http.MethodGet,
				Path:    "/v1/home/banner",
				Handler: HomeBannerHandler(serverCtx),
			},
			{
				// 订单列表
				Method:  http.MethodGet,
				Path:    "/v1/order/list",
				Handler: OrderListHandler(serverCtx),
			},
			{
				// 商品评论列表
				Method:  http.MethodGet,
				Path:    "/v1/product/comment",
				Handler: ProductCommentHandler(serverCtx),
			},
			{
				// 商品详情
				Method:  http.MethodGet,
				Path:    "/v1/product/detail",
				Handler: ProductDetailHandler(serverCtx),
			},
			{
				// 推荐商品列表
				Method:  http.MethodGet,
				Path:    "/v1/recommend",
				Handler: RecommendHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// add order
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: order.AddOrderHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/order"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// login
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// add user collection
				Method:  http.MethodPost,
				Path:    "/addCollection",
				Handler: user.UserCollectionAddHandler(serverCtx),
			},
			{
				// add user receiveAddress
				Method:  http.MethodPost,
				Path:    "/addReceiveAddress",
				Handler: user.AddReceiveAddressHandler(serverCtx),
			},
			{
				// del user collection
				Method:  http.MethodPost,
				Path:    "/delCollection",
				Handler: user.UserCollectionDelHandler(serverCtx),
			},
			{
				// del user receiveAddress list
				Method:  http.MethodPost,
				Path:    "/delReceiveAddress",
				Handler: user.DelReceiveAddressHandler(serverCtx),
			},
			{
				// edit user receiveAddress
				Method:  http.MethodPost,
				Path:    "/editReceiveAddress",
				Handler: user.EditReceiveAddressHandler(serverCtx),
			},
			{
				// get user collection list
				Method:  http.MethodGet,
				Path:    "/getCollectionList",
				Handler: user.UserCollectionListHandler(serverCtx),
			},
			{
				// get user receiveAddress list
				Method:  http.MethodGet,
				Path:    "/getReceiveAddressList",
				Handler: user.UserReceiveAddressListHandler(serverCtx),
			},
			{
				// get user info
				Method:  http.MethodPost,
				Path:    "/info",
				Handler: user.DetailHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/user"),
	)
}
