package svc

import (
	"easy-chat/apps/im/rpc/imclient"
	"easy-chat/apps/social/api/internal/config"
	"easy-chat/apps/social/rpc/socialclient"
	"easy-chat/apps/user/rpc/userclient"
	"easy-chat/pkg/interceptor"
	"easy-chat/pkg/middleware"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	IdempotenceMiddleware rest.Middleware
	socialclient.Social
	userclient.User
	imclient.Im

	*redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		IdempotenceMiddleware: middleware.NewIdempotenceMiddleware().Handler,
		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc,
			// 重试 + 幂等
			//zrpc.WithDialOption(grpc.WithDefaultServiceConfig()),
			zrpc.WithUnaryClientInterceptor(interceptor.DefaultIdempotentClient),
		)),
		User:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Im:    imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
		Redis: redis.MustNewRedis(c.Redisx),
	}
}
