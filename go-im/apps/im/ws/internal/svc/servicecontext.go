package svc

import (
	"github.com/peninsula12/easy-im/go-im/apps/im/model"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/internal/config"
	"github.com/peninsula12/easy-im/go-im/apps/task/mq/mqclient"
)

type ServiceContext struct {
	Config config.Config
	mqclient.MsgChatTransferClient
	mqclient.MsgReadTransferClient
	immodels.ChatLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MsgChatTransferClient: mqclient.NewMsgChatTransferClient(
			c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
		MsgReadTransferClient: mqclient.NewMsgReadTransferClient(
			c.MsgReadTransfer.Addrs, c.MsgReadTransfer.Topic),
		ChatLogModel: immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}
}
