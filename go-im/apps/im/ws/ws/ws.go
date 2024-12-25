package ws

import (
	"github.com/peninsula12/easy-im/go-im/pkg/status"
)

type (
	Msg struct {
		MsgId        string `mapstructure:"msgId"`
		status.MType `mapstructure:"mType"`
		Content      string            `mapstructure:"content"`
		ReadRecords  map[string]string `mapstructure:"readRecords"`
	}
)

type (
	Chat struct {
		ConversationId  string `mapstructure:"conversationId"`
		status.ChatType `mapstructure:"chatType"`
		SendId          string `mapstructure:"sendId"`
		RecvId          string `mapstructure:"recvId"`

		SendTime int64 `mapstructure:"sendTime"`
		Msg      `mapstructure:"msg"`
	}
)

type (
	Push struct {
		ConversationId  string `mapstructure:"conversationId"`
		status.ChatType `mapstructure:"chatType"`
		SendId          string   `mapstructure:"sendId"`
		RecvId          string   `mapstructure:"recvId"`
		RecvIds         []string `mapstructure:"recvIds"`
		SendTime        int64    `mapstructure:"sendTime"`

		MsgId       string             `mapstructure:"msgId"`
		ReadRecords map[string]string  `mapstructure:"readRecords"`
		ContentType status.ContentType `mapstructure:"contentType"`

		status.MType `mapstructure:"mType"`
		Content      string `mapstructure:"content"`
	}

	MarkRead struct {
		status.ChatType `mapstructure:"chatType"`
		ConversationId  string   `mapstructure:"conversationId"`
		RecvId          string   `mapstructure:"recvId"`
		MsgIds          []string `mapstructure:"msgIds"`
	}
)
