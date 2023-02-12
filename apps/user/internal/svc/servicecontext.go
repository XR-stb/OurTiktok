package svc

import (
	"OutTiktok/apps/user/internal/config"
	"OutTiktok/dao"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     dao.NewGorm(c.MysqlDsn),
	}
}
