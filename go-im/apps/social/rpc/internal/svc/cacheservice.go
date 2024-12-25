package svc

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/config"
	"github.com/peninsula12/easy-im/go-im/pkg/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type CacheService struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewCacheService(c config.Config) *CacheService {
	return &CacheService{
		DB:  utils.InitDB(c.Mysql.DataSource),
		RDB: utils.InitRDB(c.Cache[0].Host, c.Cache[0].Pass),
	}
}

// SetCache 定义缓存读写服务
func (s *CacheService) SetCache(key string, value interface{}, expiration time.Duration) error {
	return s.RDB.Set(context.Background(), key, value, expiration).Err()
}

func (s *CacheService) GetCache(key string) (string, error) {
	return s.RDB.Get(context.Background(), key).Result()
}
