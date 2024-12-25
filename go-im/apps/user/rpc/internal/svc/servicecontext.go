package svc

import (
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/internal/config"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"
	"github.com/peninsula12/easy-im/go-im/pkg/status"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

type ServiceContext struct {
	Config config.Config
	*redis.Redis
	CSvc *CacheService // 缓存与数据读写
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  redis.MustNewRedis(c.Redisx),
		CSvc:   NewCacheService(c),
	}
}

func (svc *ServiceContext) SetRootToken() error {
	// 生成jkt 再写入到redis
	systemToken, err := ctxdata.GetJwtToken(svc.Config.Jwt.AccessSecret, time.Now().Unix(), 999999999, status.SYSTEM_ROOT_UID)
	if err != nil {
		return err
	}

	return svc.Redis.Set(status.REDIS_SYSTEM_ROOT_TOKEN, systemToken)

}
