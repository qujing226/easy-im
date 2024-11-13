package svc

import (
	"easy-chat/apps/im/immodels"
	"easy-chat/apps/im/ws/internal/config"
)

type ServiceContext struct {
	Config config.Config

	immodels.ChatLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
