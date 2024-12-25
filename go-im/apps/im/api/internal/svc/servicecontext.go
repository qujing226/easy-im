package svc

import (
	"easy-chat/apps/im/api/internal/config"
	"easy-chat/apps/im/rpc/im"
	"easy-chat/apps/im/rpc/imclient"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/apps/social/rpc/socialclient"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	im.ImClient
	social.SocialClient
	user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ImClient:     imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
		SocialClient: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		UserClient:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
