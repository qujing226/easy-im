package msgTransfer

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/websocket"
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/ws"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/socialclient"
	"github.com/peninsula12/easy-im/go-im/apps/task/mq/internal/svc"
	"github.com/peninsula12/easy-im/go-im/pkg/status"
	"github.com/zeromicro/go-zero/core/logx"
)

type baseMsgTransfer struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseMsgTransfer(svc *svc.ServiceContext) *baseMsgTransfer {
	return &baseMsgTransfer{
		svcCtx: svc,
		Logger: logx.WithContext(context.Background()),
	}
}

func (m *baseMsgTransfer) Transfer(ctx context.Context, data *ws.Push) error {
	var err error
	switch data.ChatType {
	case status.SingleChatType:
		err = m.single(ctx, data)
	case status.GroupChatType:
		err = m.group(ctx, data)
	}
	return err
}

func (m *baseMsgTransfer) single(ctx context.Context, data *ws.Push) error {
	return m.svcCtx.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push",
		// todo: 此处可能需要更改
		UserId: "",
		FormId: status.SYSTEM_ROOT_UID,
		Data:   data,
	})
}

func (m *baseMsgTransfer) group(ctx context.Context, data *ws.Push) error {
	// 查询群的用户
	users, err := m.svcCtx.Social.GroupUsers(ctx, &socialclient.GroupUsersReq{
		GroupId: data.RecvId,
	})
	if err != nil {
		return err
	}
	data.RecvIds = make([]string, 0, len(users.List))
	//fmt.Printf("group user: %+v", users.List)
	for _, members := range users.List {
		if members.UserId == data.SendId {
			continue
		}
		data.RecvIds = append(data.RecvIds, members.UserId)
	}
	return m.svcCtx.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push",
		// todo: 此处可能需要更改
		UserId: "",
		FormId: status.SYSTEM_ROOT_UID,
		Data:   data,
	})

}
