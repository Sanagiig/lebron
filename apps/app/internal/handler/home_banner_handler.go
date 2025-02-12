package handler

import (
	"net/http"

	"github.com/Sanagiig/lebron/apps/app/internal/logic"
	"github.com/Sanagiig/lebron/apps/app/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 首页Banner
func HomeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
