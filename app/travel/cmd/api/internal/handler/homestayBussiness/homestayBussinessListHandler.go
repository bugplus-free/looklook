package homestayBussiness

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/travel/cmd/api/internal/logic/homestayBussiness"
	"looklook/app/travel/cmd/api/internal/svc"
	"looklook/app/travel/cmd/api/internal/types"
)

func HomestayBussinessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBussiness.NewHomestayBussinessListLogic(r.Context(), svcCtx)
		resp, err := l.HomestayBussinessList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
