package thirdPayment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/payment/cmd/api/internal/logic/thirdPayment"
	"looklook/app/payment/cmd/api/internal/svc"
	"looklook/app/payment/cmd/api/internal/types"
)

// third payment：wechat pay
func ThirdPaymentwxPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThirdPaymentWxPayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := thirdPayment.NewThirdPaymentwxPayLogic(r.Context(), svcCtx)
		resp, err := l.ThirdPaymentwxPay(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
