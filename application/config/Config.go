package config

import (
	"github.com/tampfievk50/gozero-core-api/casbinx"
	"github.com/tampfievk50/gozero-core-api/ormx"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	Rest            RestConf
	Auth            AuthConf
	Db              ormx.Conf
	Redis           RedisConf
	Casx            casbinx.CasxConf
	EnableMessaging bool
	KqPusherConf    KqPusherConf
	KqConsumerConf  kq.KqConf
}

type ServiceUrlConf struct {
	QuotaServiceUrl    string
	ContractServiceUrl string
}

type RestConf struct {
	rest.RestConf
	CorsOrigins []string
}

type AppConf struct {
	EmailSenderAddress string
	Endpoint           string `json:",optional"`
}

type CaptchaConf struct {
	Enable bool `json:",default=true"`

	Version   int     `json:",default=2,options=2|3"`
	Action    string  `json:",default=auth"`
	Score     float64 `json:",default=0.5"`
	SiteKey   string  `json:",optional"`
	SecretKey string  `json:",optional"`

	ImageCaptchaEnable         bool  `json:",default=false"`
	ImageCaptchaTtl            int64 `json:",optional,default=300,comment:second"`
	ImageCaptchaLoginFailCount int64 `json:",optional,default=5"`
	ImageCaptchaLoginFailTtl   int64 `json:",optional,default=3600,comment:second"`
}

type AuthConf struct {
	AccessSecret string
	AccessExpire int64
}

type OtpConf struct {
	OtpLength            int32 `json:",optional,default=6"`
	SmsResendDelay       int32 `json:",optional,comment:second"`
	MaxOTPResendAttempts int32 `json:",optional,comment:second"`
	OtpTtl               int32 `json:",optional,comment:second"`
	PaddingCacheTime     int32 `json:",optional"`
	VerificationAttempts int32 `json:",optional,comment:second"`
}

type KqPusherConf struct {
	Brokers []string
	Topics  []string
}

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}

type Job struct {
	UploadPath   string
	VT8x79Table  string
	VTOtherTable string
}

type FtpInfo struct {
	Hostname string
	Port     string
	Username string
	Password string
	BasePath string
}
