package homestayBussiness

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/travel/cmd/api/internal/logic/homestayBussiness"
	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"
)

func GoodBossHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodBossReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBussiness.NewGoodBossLogic(r.Context(), svcCtx)
		resp, err := l.GoodBoss(&req)
		result.HttpResult(r, w, resp, err)
	}
}
