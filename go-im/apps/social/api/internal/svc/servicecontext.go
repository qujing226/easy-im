package svc

import (
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/imclient"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/config"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/socialclient"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/userclient"
	"github.com/peninsula12/easy-im/go-im/pkg/interceptor"
	"github.com/peninsula12/easy-im/go-im/pkg/middleware"
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
