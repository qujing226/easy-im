package ws

import (
	"easy-chat/pkg/status"
)

type (
	Msg struct {
		status.MType `mapstructure:"mType"`
		Content      string `mapstructure:"content"`
	}
)

type (
	Chat struct {
		ConversationId  string `mapstructure:"conversationId"`
		status.ChatType `mapstructure:"chatType"`
		SendId          string `mapstructure:"sendId"`
		RecvId          string `json:"recvId"`
		SendTime        int64  `mapstructure:"sendTime"`
		Msg             `mapstructure:"msg"`
	}
)
