package conversation

import (
	"github.com/mitchellh/mapstructure"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/websocket"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/ws"
	"github.com/peninsula12/easy-im/go-im/apps/task/mq/mq"
	"github.com/peninsula12/easy-im/go-im/pkg/status"
	"github.com/peninsula12/easy-im/go-im/pkg/suid"
	"time"
)

func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			_ = srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		if data.ConversationId == "" {
			switch data.ChatType {
			case status.SingleChatType:
				data.ConversationId = suid.CombineId(conn.Uid, data.RecvId)
			case status.GroupChatType:
				data.ConversationId = data.RecvId
			default:
			}
		}
		err := svc.MsgChatTransferClient.Push(&mq.MsgChatTransfer{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			SendTime:       time.Now().UnixMilli(),
			MType:          data.MType,
			Content:        data.Content,
			MsgId:          msg.Id,
		})
		if err != nil {
			_ = srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
	}
}

func MarkRead(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.MarkRead
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			_ = srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		err := svc.MsgReadTransferClient.Push(&mq.MsgMarkRead{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MsgIds:         data.MsgIds,
		})
		if err != nil {
			_ = srv.Send(websocket.NewErrMessage(err), conn)
			return
		}
	}
}
