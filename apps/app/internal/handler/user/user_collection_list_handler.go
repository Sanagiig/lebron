package user

import (
	"net/http"

	"github.com/Sanagiig/lebron/apps/app/internal/logic/user"
	"github.com/Sanagiig/lebron/apps/app/internal/svc"
	"github.com/Sanagiig/lebron/apps/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get user collection list
func UserCollectionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCollectionListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserCollectionListLogic(r.Context(), svcCtx)
		resp, err := l.UserCollectionList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
