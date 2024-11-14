package push

import (
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/im/ws/ws"
	"github.com/mitchellh/mapstructure"
)

func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err))
			return
		}
		// 发送目标
		// todo: 这里的recv是一个切片
		recvConn := srv.GetConn(data.RecvId)
		if recvConn == nil {
			// todo: 目标离线
			return
		}
		srv.Infof("push msg %v", data)

		err := srv.Send(websocket.NewMessage(data.SendId, &ws.Chat{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendTime:       data.SendTime,
			Msg: ws.Msg{
				MType:   data.MType,
				Content: data.Content,
			},
		}), recvConn[0])
		if err != nil {
			srv.Error(err)
		}
	}

}
