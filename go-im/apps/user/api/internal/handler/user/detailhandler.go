package user

import (
	"net/http"

	"github.com/peninsula12/easy-im/go-im/apps/user/api/internal/logic/user"
	"github.com/peninsula12/easy-im/go-im/apps/user/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// DetailHandler 获取用户信息
func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
