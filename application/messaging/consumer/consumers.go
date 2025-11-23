package consumer

import (
	"gozero-sso-service/application/config"
	"gozero-sso-service/messaging/consumer"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqConsumerConf, consumer.NewPaymentSuccess()),
	}
}
