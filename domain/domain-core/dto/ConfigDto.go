package dto

import (
	"github.com/tampfievk50/gozero-core-api/casbinx"
	"github.com/tampfievk50/gozero-core-api/ormx"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
)

type ConfigDto struct {
	Rest            RestConfDto
	Auth            AuthConfDto
	Db              ormx.Conf
	Redis           RedisConfDto
	Casx            casbinx.CasxConf
	EnableMessaging bool
	KqPusherConf    KqPusherConfDto
	KqConsumerConf  kq.KqConf
}

type ServiceUrlConfDto struct {
	QuotaServiceUrl    string
	ContractServiceUrl string
}

type RestConfDto struct {
	rest.RestConf
	CorsOrigins []string
}

type AppConfDto struct {
	EmailSenderAddress string
	Endpoint           string `json:",optional"`
}

type AuthConfDto struct {
	AccessSecret  string
	AccessExpire  int64
	RefreshSecret string
	RefreshExpire int64
}

type OtpConfDto struct {
	OtpLength            int32 `json:",optional,default=6"`
	SmsResendDelay       int32 `json:",optional,comment:second"`
	MaxOTPResendAttempts int32 `json:",optional,comment:second"`
	OtpTtl               int32 `json:",optional,comment:second"`
	PaddingCacheTime     int32 `json:",optional"`
	VerificationAttempts int32 `json:",optional,comment:second"`
}

type KqPusherConfDto struct {
	Brokers []string
	Topics  []string
}

type RedisConfDto struct {
	Addr     string
	Password string
	DB       int
}
