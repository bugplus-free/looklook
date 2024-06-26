// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	homestay "looklook/app/travel/cmd/api/internal/handler/homestay"
	homestayBussiness "looklook/app/travel/cmd/api/internal/handler/homestayBussiness"
	homestayComment "looklook/app/travel/cmd/api/internal/handler/homestayComment"
	"looklook/app/travel/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 房东所有民宿列表
				Method:  http.MethodPost,
				Path:    "/homestay/businessList",
				Handler: homestay.BusinessListHandler(serverCtx),
			},
			{
				// 猜你喜欢民宿列表
				Method:  http.MethodPost,
				Path:    "/homestay/guessList",
				Handler: homestay.GuessListHandler(serverCtx),
			},
			{
				// 民宿详情
				Method:  http.MethodPost,
				Path:    "/homestay/homestayDetail",
				Handler: homestay.HomestayDetailHandler(serverCtx),
			},
			{
				// 民宿列表（为你优选）
				Method:  http.MethodPost,
				Path:    "/homestay/homestayList",
				Handler: homestay.HomestayListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 最佳房东
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/goodBoss",
				Handler: homestayBussiness.GoodBossHandler(serverCtx),
			},
			{
				// 房东信息
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/homestayBussinessDetail",
				Handler: homestayBussiness.HomestayBussinessDetailHandler(serverCtx),
			},
			{
				// 店铺列表
				Method:  http.MethodPost,
				Path:    "/homestayBussiness/homestayBussinessList",
				Handler: homestayBussiness.HomestayBussinessListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 民宿评论列表
				Method:  http.MethodPost,
				Path:    "/homestayComment/commentList",
				Handler: homestayComment.CommentListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/travel/v1"),
	)
}
