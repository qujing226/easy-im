package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/config"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/im"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/imclient"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/socialclient"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/user"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/userclient"
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
