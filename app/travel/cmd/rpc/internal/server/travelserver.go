// Code generated by goctl. DO NOT EDIT.
// Source: travel.proto

package server

import (
	"context"

	"looklook/app/travel/cmd/rpc/internal/logic"
	"looklook/app/travel/cmd/rpc/internal/svc"
	"looklook/app/travel/cmd/rpc/pb"
)

type TravelServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTravelServer
}

func NewTravelServer(svcCtx *svc.ServiceContext) *TravelServer {
	return &TravelServer{
		svcCtx: svcCtx,
	}
}

// homestayDetail
func (s *TravelServer) HomestayDetail(ctx context.Context, in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	l := logic.NewHomestayDetailLogic(ctx, s.svcCtx)
	return l.HomestayDetail(in)
}
