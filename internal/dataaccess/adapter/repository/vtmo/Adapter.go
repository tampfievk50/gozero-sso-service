package vtmo

import (
	"context"
	"github.com/tampfievk50/gozero-core-api/app"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gozero-sso-service/internal/config"
	op "gozero-sso-service/internal/logic/ports/output/repository"
)

type repository struct {
	logx.Logger
	db     *gorm.DB
	config *config.Config
}

func NewVTMORepository() op.VTMORepository {
	return &repository{}
}

func (r *repository) SetState(ctx context.Context, db *gorm.DB, config *config.Config) op.VTMORepository {
	r.Logger = logx.WithContext(ctx)
	r.db = db
	r.config = config
	return r
}

func init() {
	app.Register("vTMORepository", func() op.VTMORepository {
		return NewVTMORepository()
	})
}
