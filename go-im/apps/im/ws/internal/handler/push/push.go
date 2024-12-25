package push

import (
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/im/ws/ws"
	"easy-chat/pkg/status"
	"github.com/mitchellh/mapstructure"
)

func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			_ = srv.Send(websocket.NewErrMessage(err))
			return
		}
		// 发送目标
		// todo: 这里的recv是一个切片
		switch data.ChatType {
		case status.SingleChatType:
			err := single(srv, &data, data.RecvId)
			if err != nil {
				srv.Error(err)
			}
		case status.GroupChatType:
			err := group(srv, &data)
			if err != nil {
				srv.Error(err)
			}
		default:

		}

	}

}

func single(srv *websocket.Server, data *ws.Push, recvId string) error {
	recvConn := srv.GetConn(recvId)
	if recvConn == nil {
		// todo: 目标离线
		return nil
	}
	srv.Infof("push msg %v", data)

	return srv.Send(websocket.NewMessage(data.SendId, &ws.Chat{
		ConversationId: data.ConversationId,
		ChatType:       data.ChatType,
		SendTime:       data.SendTime,
		Msg: ws.Msg{
			MsgId:       data.MsgId,
			MType:       data.MType,
			Content:     data.Content,
			ReadRecords: data.ReadRecords,
		},
	}), recvConn[0])

}

func group(srv *websocket.Server, data *ws.Push) (err error) {
	//fmt.Println("group push")
	for _, id := range data.RecvIds {
		func(id string) {
			srv.Schedule(func() {
				err = single(srv, data, id)
			})
		}(id)
		//fmt.Println(id)
	}
	return
}
