package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	ListenOn string

	SocialRpc zrpc.RpcClientConf

	MsgChatTransfer kq.KqConf
	MsgReadTransfer kq.KqConf

	MsgReadHandler struct {
		GroupMsgReadHandler          int
		GroupMsgReadRecordDelayTime  int64
		GroupMsgReadRecordDelayCount int
	}

	Redisx redis.RedisConf
	Mongo  struct {
		Url string
		Db  string
	}

	Ws struct {
		Host string
	}
}
