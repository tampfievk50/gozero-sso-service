package consumer

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"gozero-sso-service/application/config"
	"gozero-sso-service/messaging/consumer"
)

func Consumers(c config.Config) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqConsumerConf, consumer.NewPaymentSuccess()),
	}
}
