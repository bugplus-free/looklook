package logic

import (
	"context"

	"looklook/app/payment/cmd/rpc/internal/svc"
	"looklook/app/payment/cmd/rpc/pb"
	"looklook/app/payment/model"
	"looklook/common/xerr"
	"looklook/common/uniqueid"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建微信支付预处理订单
func (l *CreatePaymentLogic) CreatePayment(in *pb.CreatePaymentReq) (*pb.CreatePaymentResp, error) {
	// todo: add your logic here and delete this line
	data := new(model.ThirdPayment)
	data.Sn = uniqueid.GenSn(uniqueid.SN_PREFIX_THIRD_PAYMENT)
	data.UserId = in.UserId
	data.PayMode = in.PayModel
	data.PayTotal = in.PayTotal
	data.OrderSn = in.OrderSn
	data.ServiceType = model.ThirdPaymentServiceTypeHomestayOrder

	_, err := l.svcCtx.ThirdPaymentModel.Insert(l.ctx, nil, data)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create wechat pay prepayorder db insert fail , err:%v ,data : %+v  ", err, data)
	}

	return &pb.CreatePaymentResp{
		Sn: data.Sn,
	}, nil
}
