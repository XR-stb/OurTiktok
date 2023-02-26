package svc

import (
	"OutTiktok/apps/message/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis: redis.New(c.Redis.Host, redis.Option(func(r *redis.Redis) {

		})),
	}
}
