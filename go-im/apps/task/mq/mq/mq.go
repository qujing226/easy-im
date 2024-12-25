package mq

import "github.com/peninsula12/easy-im/go-im/pkg/status"

type MsgChatTransfer struct {
	MsgId string `mapstructure:"msgId"`

	ConversationId  string `json:"conversationId"`
	status.ChatType `json:"chatType"`
	SendId          string   `json:"sendId"`
	RecvId          string   `json:"recvId"`
	RecvIds         []string `json:"recvIds"`
	SendTime        int64    `json:"sendTime"`

	status.MType `json:"status.mType"`
	Content      string `json:"content"`
}

type MsgMarkRead struct {
	ConversationId  string `json:"conversationId"`
	status.ChatType `json:"chatType"`
	SendId          string   `json:"sendId"`
	RecvId          string   `json:"recvId"`
	MsgIds          []string `json:"msgIds"`
}
