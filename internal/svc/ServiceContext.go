package svc

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/tampfievk50/gozero-core-api/casbinx"
	"github.com/tampfievk50/gozero-core-api/ormx"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"gozero-sso-service/internal/config"
	"gozero-sso-service/internal/middleware"
)

type ServiceContext struct {
	Config                 config.Config
	AuthMiddleware         rest.Middleware
	RefreshTokenMiddleware rest.Middleware
	RecoverMiddleware      rest.Middleware

	RedisCli *redis.Client
	Enforcer *casbin.SyncedEnforcer
	Pushers  map[string]*kq.Pusher
	DB       *ormx.Database
}

func (svc *ServiceContext) GetLogger(ctx context.Context) logx.Logger {
	return logx.WithContext(ctx)
}

func NewServiceContext(c config.Config) servicecontext.ServiceContextInterface {
	database := ormx.InitDB(&c.Db, ormx.WithMigration(map[string]string{}, []interface{}{}))

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	rbacEnforcer := casbinx.InitRBAC(database, c.Casx)

	pushers := map[string]*kq.Pusher{}
	for _, topic := range c.KqPusherConf.Topics {
		pushers[topic] = kq.NewPusher(c.KqPusherConf.Brokers, topic)
	}

	return &ServiceContext{
		Config:                 c,
		AuthMiddleware:         middleware.NewAuthMiddleware(rbacEnforcer).Handle,
		RefreshTokenMiddleware: middleware.NewRefreshTokenMiddleware().Handle,
		RecoverMiddleware:      middleware.NewRecoverMiddleware().Handle,
		RedisCli:               rdb,
		Enforcer:               rbacEnforcer,
		Pushers:                pushers,
		DB:                     database,
	}
}
