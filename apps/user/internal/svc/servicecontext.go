package svc

import (
	"OutTiktok/apps/user/internal/config"
	"OutTiktok/dao"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     dao.NewGorm(c.MysqlDsn),
		Redis:  redis.New(c.Redis.Host),
	}
}
