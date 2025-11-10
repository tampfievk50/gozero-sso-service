package consumer

import (
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"gozero-sso-service/internal/config"
	"gozero-sso-service/internal/messaging/consumer"
	"gozero-sso-service/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqConsumerConf, consumer.NewPaymentSuccess(ctx, svcContext)),
	}
}
