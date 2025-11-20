package consumer

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentSuccess struct{}

func NewPaymentSuccess() *PaymentSuccess {
	return &PaymentSuccess{}
}

func (l *PaymentSuccess) Consume(ctx context.Context, key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)
	return nil
}
