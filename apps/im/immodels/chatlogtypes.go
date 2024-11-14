package immodels

import (
	"easy-chat/pkg/status"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatLog struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationId string          `bson:"conversationId"`
	SendId         string          `bson:"sendId"`
	RecvId         string          `bson:"recvId"`
	MsgFrom        int             `bson:"msgFrom"`
	ChatType       status.ChatType `bson:"chatType"`
	MsgType        status.MType    `bson:"msgType"`
	MsgContent     string          `bson:"msgContent"`
	SendTime       int64           `bson:"sendTime"`
	Status         int             `bson:"status"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
