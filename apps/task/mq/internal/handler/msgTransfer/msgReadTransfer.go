package msgTransfer

import (
	"context"
	"easy-chat/apps/im/ws/ws"
	"easy-chat/apps/task/mq/internal/svc"
	"easy-chat/apps/task/mq/mq"
	"easy-chat/pkg/bitmap"
	"easy-chat/pkg/status"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"sync"
	"time"
)

var (
	GroupMsgReadRecordDelayTime  = time.Second
	GroupMsgReadRecordDelayCount = 10
)

const (
	GroupMsgReadHandlerAtTransfer = iota
	GroupMsgReadHandlerDelayTransfer
)

type MsgReadTransfer struct {
	*baseMsgTransfer

	cache.Cache

	mu sync.Mutex

	groupMsgs map[string]*groupMsgRead
	push      chan *ws.Push
}

func NewMsgReadTransfer(svc *svc.ServiceContext) kq.ConsumeHandler {
	m := &MsgReadTransfer{
		baseMsgTransfer: NewBaseMsgTransfer(svc),
		groupMsgs:       make(map[string]*groupMsgRead, 1),
		push:            make(chan *ws.Push, 1),
	}

	if svc.Config.MsgReadHandler.GroupMsgReadHandler != GroupMsgReadHandlerAtTransfer {
		if svc.Config.MsgReadHandler.GroupMsgReadRecordDelayCount > 0 {
			GroupMsgReadRecordDelayCount = svc.Config.MsgReadHandler.GroupMsgReadRecordDelayCount

		}
		if svc.Config.MsgReadHandler.GroupMsgReadRecordDelayTime > 0 {
			GroupMsgReadRecordDelayTime = time.Duration(svc.Config.MsgReadHandler.GroupMsgReadRecordDelayTime) *
				time.Second
		}
	}
	go m.transfer()

	return m
}

func (m *MsgReadTransfer) Consume(ctx context.Context, key, value string) error {
	m.Infof("MsgReadTransfer ", value)
	var data mq.MsgMarkRead

	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	// 业务处理 ： 更新已读未读
	readRecords, err := m.UpdateChatLogRead(ctx, &data)
	if err != nil {
		return err
	}

	// map[消息ID] 已读记录  map[string]string
	push := &ws.Push{
		ConversationId: data.ConversationId,
		ChatType:       data.ChatType,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ContentType:    status.ContentMarkRead,
		ReadRecords:    readRecords,
	}
	switch data.ChatType {
	case status.SingleChatType:
		//  直接推送
		m.push <- push
	case status.GroupChatType:
		if m.svcCtx.Config.MsgReadHandler.GroupMsgReadHandler == GroupMsgReadHandlerAtTransfer {
			m.push <- push
		}
		// 合并推送
		m.mu.Lock()
		defer m.mu.Unlock()

		push.SendId = ""
		if _, ok := m.groupMsgs[push.ConversationId]; ok {
			// 合并
			m.Infof("merge push %v", push.ConversationId)
			m.groupMsgs[push.ConversationId].mergePush(push)
		} else {
			m.Infof("new read merge push %v", push.ConversationId)
			m.groupMsgs[push.ConversationId] = newGroupMsgRead(push, m.push)
		}
	}
	return nil
}

func (m *MsgReadTransfer) UpdateChatLogRead(ctx context.Context, data *mq.MsgMarkRead) (map[string]string, error) {
	res := make(map[string]string)

	chatLogs, err := m.svcCtx.ChatLogModel.ListByMsgIds(ctx, data.MsgIds)
	if err != nil {
		return nil, err
	}
	// 处理已读
	for _, chatLog := range chatLogs {
		switch chatLog.ChatType {
		case status.SingleChatType:
			chatLog.ReadRecords = []byte{1}
		case status.GroupChatType:
			readRecords := bitmap.Load(chatLog.ReadRecords)
			readRecords.Set(data.SendId)
			chatLog.ReadRecords = readRecords.Export()
		}

		res[chatLog.ID.Hex()] = base64.StdEncoding.EncodeToString(chatLog.ReadRecords)
		err = m.svcCtx.ChatLogModel.UpdateMarkRead(ctx, chatLog.ID, chatLog.ReadRecords)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// 异步任务的写成处理发送
func (m *MsgReadTransfer) transfer() {
	for push := range m.push {
		if push.RecvId != "" || len(push.RecvIds) > 0 {
			if err := m.Transfer(context.Background(), push); err != nil {
				m.Errorf("m transfer err %v push %v", err, push)
			}
		}

		if push.ChatType == status.SingleChatType {
			continue
		}

		if m.svcCtx.Config.MsgReadHandler.GroupMsgReadHandler == GroupMsgReadHandlerAtTransfer {
			continue
		}
		// 清空数据
		m.mu.Lock()
		if _, ok := m.groupMsgs[push.ConversationId]; ok && m.groupMsgs[push.ConversationId].IsIdle() {
			fmt.Printf("clear groupMsg: %+v\n", m.groupMsgs[push.ConversationId])
			m.groupMsgs[push.ConversationId].clear()
			delete(m.groupMsgs, push.ConversationId)
		}
		m.mu.Unlock()

	}
}
