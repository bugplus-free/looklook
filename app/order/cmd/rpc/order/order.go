// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package order

import (
	"context"

	"looklook/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateHomestayOrderReq            = pb.CreateHomestayOrderReq
	CreateHomestayOrderResp           = pb.CreateHomestayOrderResp
	HomestayOrder                     = pb.HomestayOrder
	HomestayOrderDetailReq            = pb.HomestayOrderDetailReq
	HomestayOrderDetailResp           = pb.HomestayOrderDetailResp
	UpdateHomestayOrderTradeStateReq  = pb.UpdateHomestayOrderTradeStateReq
	UpdateHomestayOrderTradeStateResp = pb.UpdateHomestayOrderTradeStateResp
	UserHomestayOrderListReq          = pb.UserHomestayOrderListReq
	UserHomestayOrderListResp         = pb.UserHomestayOrderListResp

	Order interface {
		// 民宿下订单
		CreateHomestayOrder(ctx context.Context, in *CreateHomestayOrderReq, opts ...grpc.CallOption) (*CreateHomestayOrderResp, error)
		// 民宿订单详情
		HomestayOrderDetail(ctx context.Context, in *HomestayOrderDetailReq, opts ...grpc.CallOption) (*HomestayOrderDetailResp, error)
		// 更新民宿订单状态
		UpdateHomestayOrderTradeState(ctx context.Context, in *UpdateHomestayOrderTradeStateReq, opts ...grpc.CallOption) (*UpdateHomestayOrderTradeStateResp, error)
		// 用户民宿订单
		UserHomestayOrderList(ctx context.Context, in *UserHomestayOrderListReq, opts ...grpc.CallOption) (*UserHomestayOrderListResp, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

// 民宿下订单
func (m *defaultOrder) CreateHomestayOrder(ctx context.Context, in *CreateHomestayOrderReq, opts ...grpc.CallOption) (*CreateHomestayOrderResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.CreateHomestayOrder(ctx, in, opts...)
}

// 民宿订单详情
func (m *defaultOrder) HomestayOrderDetail(ctx context.Context, in *HomestayOrderDetailReq, opts ...grpc.CallOption) (*HomestayOrderDetailResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.HomestayOrderDetail(ctx, in, opts...)
}

// 更新民宿订单状态
func (m *defaultOrder) UpdateHomestayOrderTradeState(ctx context.Context, in *UpdateHomestayOrderTradeStateReq, opts ...grpc.CallOption) (*UpdateHomestayOrderTradeStateResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.UpdateHomestayOrderTradeState(ctx, in, opts...)
}

// 用户民宿订单
func (m *defaultOrder) UserHomestayOrderList(ctx context.Context, in *UserHomestayOrderListReq, opts ...grpc.CallOption) (*UserHomestayOrderListResp, error) {
	client := pb.NewOrderClient(m.cli.Conn())
	return client.UserHomestayOrderList(ctx, in, opts...)
}
