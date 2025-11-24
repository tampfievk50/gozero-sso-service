package svc

import (
	"context"
	"gozero-sso-service/application/config"
	"gozero-sso-service/application/middleware"
	"gozero-sso-service/dataaccess/adapter/repository"
	"gozero-sso-service/dataaccess/adapter/repository/role"
	"gozero-sso-service/dataaccess/adapter/repository/user"
	"gozero-sso-service/domain/domain-application/adapter/service"
	"gozero-sso-service/domain/domain-application/adapter/service/auth"
	roleSvc "gozero-sso-service/domain/domain-application/adapter/service/role"
	userSvc "gozero-sso-service/domain/domain-application/adapter/service/user"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"github.com/tampfievk50/gozero-core-api/casbinx"
	"github.com/tampfievk50/gozero-core-api/ormx"
	"github.com/tampfievk50/gozero-core-api/servicecontext"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

func InitRepository(database *ormx.Database) *repository.Repository {
	return &repository.Repository{
		UserRepository: user.NewUserRepository(database),
		RoleRepository: role.NewRoleRepository(database),
	}
}

func InitService(rp *repository.Repository) *service.Service {
	return &service.Service{
		AuthService: auth.NewAuthService(rp.UserRepository),
		UserService: userSvc.NewUserService(rp.UserRepository),
		RoleService: roleSvc.NewRoleService(rp.RoleRepository),
	}
}

type ServiceContext struct {
	Config                 config.Config
	AuthMiddleware         rest.Middleware
	RefreshTokenMiddleware rest.Middleware
	RecoverMiddleware      rest.Middleware

	RedisCli *redis.Client
	Enforcer *casbin.SyncedEnforcer
	Pushers  map[string]*kq.Pusher
	DB       *ormx.Database

	Repo *repository.Repository
	Svc  *service.Service
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

	repo := InitRepository(database)
	svc := InitService(repo)

	return &ServiceContext{
		Config:                 c,
		AuthMiddleware:         middleware.NewAuthMiddleware(rbacEnforcer, svc).Handle,
		RefreshTokenMiddleware: middleware.NewRefreshTokenMiddleware().Handle,
		RecoverMiddleware:      middleware.NewRecoverMiddleware().Handle,
		RedisCli:               rdb,
		Enforcer:               rbacEnforcer,
		Pushers:                pushers,
		DB:                     database,
		Repo:                   repo,
		Svc:                    svc,
	}
}

func (svc *ServiceContext) GetLogger(ctx context.Context) logx.Logger {
	return logx.WithContext(ctx)
}
