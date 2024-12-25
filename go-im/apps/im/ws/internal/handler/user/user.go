package user

import (
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/websocket"
)

func OnLine(srv *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		ids := srv.GetUsers()
		u := srv.GetUsers(conn)
		err := srv.Send(websocket.NewMessage(u[0], ids), conn)
		srv.Info("err", err)
	}
}
