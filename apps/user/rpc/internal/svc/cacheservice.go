package svc

import (
	"context"
	"easy-chat/apps/user/rpc/internal/config"
	"easy-chat/apps/user/rpc/models"
	"easy-chat/pkg/utils"
	"encoding/json"
	"fmt"
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

func (s *CacheService) GetUserCache(user *models.User, key string) error {
	data, err := s.RDB.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), user)
}

func (s *CacheService) SetUserCache(user *models.User, key string) error {
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return s.SetCache(key, userData, time.Hour*24)
}

// GetUserByPhone 业务逻辑层
func (s *CacheService) GetUserByPhone(user *models.User, phone string) error {
	cacheKey := "user_phone:" + phone
	// 处理缓存
	err := s.GetUserCache(user, cacheKey)
	if err == nil {
		return nil
	}

	err = s.DB.Where("phone = (?)", phone).Find(user).Error
	if err != nil {
		return err
	}
	// 更新缓存
	err = s.SetUserCache(user, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

func (s *CacheService) GetUserByIds(users *[]models.User, ids []string) error {
	for _, id := range ids {
		cacheKey := "user_id:" + id
		var user models.User
		// 处理缓存
		err := s.GetUserCache(&user, cacheKey)
		if err == nil {
			*users = append(*users, user)
			continue
		}
		fmt.Println(ids, id)
		err = s.DB.Where("id = ?", id).First(&user).Error
		if err != nil {
			return err
		}
		err = s.SetUserCache(&user, cacheKey)
		if err != nil {
			return err
		}
		*users = append(*users, user)
	}
	return nil
}

func (s *CacheService) GetUserByName(user *models.User, name string) error {
	cacheKey := "user_name:" + name
	err := s.GetUserCache(user, cacheKey)
	if err == nil {
		return nil
	}
	err = s.DB.Where("nickname = (?)", name).Find(user).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = s.SetUserCache(user, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// CreateUser 保存
func (s *CacheService) CreateUser(user *models.User) error {
	cacheKey := "user_phone:" + user.Phone
	err := s.DB.Create(user).Error
	if err != nil {
		return err
	}
	s.RDB.Del(context.Background(), cacheKey)
	// 延时双删策略
	time.Sleep(1 * time.Second)
	s.RDB.Del(context.Background(), cacheKey)
	return nil
}
