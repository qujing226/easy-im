package logic

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/im/model"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/websocket"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/ws"
	"github.com/peninsula12/easy-im/go-im/pkg/suid"
	"time"
)

type Conversation struct {
	ctx context.Context
	srv *websocket.Server
	svc *svc.ServiceContext
}

func NewConversation(ctx context.Context, srv *websocket.Server, svc *svc.ServiceContext) *Conversation {
	return &Conversation{
		ctx: ctx,
		srv: srv,
		svc: svc,
	}
}

func (l *Conversation) SingleChat(data *ws.Chat, userId string) error {
	if data.ConversationId == "" {
		data.ConversationId = suid.CombineId(userId, data.RecvId)
	}

	// 记录消息
	var chatLog = immodels.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         userId,
		RecvId:         data.RecvId,
		MsgFrom:        0,
		MsgType:        data.MType,
		MsgContent:     data.Content,
		SendTime:       time.Now().UnixNano(),
		Status:         0,
	}
	err := l.svc.ChatLogModel.Insert(l.ctx, &chatLog)
	return err
}
