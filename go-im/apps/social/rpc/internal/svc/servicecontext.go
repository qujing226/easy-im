package svc

import (
	"easy-chat/apps/social/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	CSvc   *CacheService // 缓存与数据读写
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		CSvc:   NewCacheService(c),
	}
}
