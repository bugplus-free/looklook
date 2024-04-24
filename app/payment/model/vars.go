package model

import (
    "errors"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrNoRowsUpdate = errors.New("update db no rows change")

// 支付业务类型
var ThirdPaymentServiceTypeHomestayOrder string ="homestayOrder"

// 支付方式
var ThirdPaymentPayModelWechatPay ="WECHAT_PAY"

// 支付状态
var ThirdPaymentPayTradeStateFAIL int64 =-1
var ThirdPaymentPayTradeStateWait int64=0
var ThirdPaymentPayTradeStateSuccess int64=1
var ThirdPaymentPayTradeStateRefund int64=2

