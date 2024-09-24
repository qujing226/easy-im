package svc

import (
	"context"
	"easy-chat/apps/user/common/utils"
	"easy-chat/apps/user/rpc/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     utils.InitDB(c.Mysql.DataSource),
		RDB:    utils.InitRDB(c.Redis.Host, c.Redis.Pass),
	}
}

func (s *ServiceContext) SetCache(key string, value interface{}, expiration time.Duration) error {
	return s.RDB.Set(context.Background(), key, value, expiration).Err()
}

func (s *ServiceContext) GetCache(key string) (string, error) {
	return s.RDB.Get(context.Background(), key).Result()
}
