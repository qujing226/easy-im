package mq

import "easy-chat/pkg/status"

type MsgChatTransfer struct {
	ConversationId  string `json:"conversationId"`
	status.ChatType `json:"charType"`
	SendId          string `json:"sendId"`
	RecvId          string `json:"recvId"`
	SendTime        int64  `json:"sendTime"`

	status.MType `json:"status.mType"`
	Content      string `json:"content"`
}
