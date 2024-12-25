package msgTransfer

import (
	"github.com/peninsula12/easy-im/go-im/apps/im/ws/ws"
	"github.com/peninsula12/easy-im/go-im/pkg/status"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

type groupMsgRead struct {
	mu sync.Mutex

	conversationId string

	push   *ws.Push
	pushCh chan *ws.Push

	count int
	// 上次推送时间
	pushTime time.Time
	done     chan struct{}
}

func newGroupMsgRead(push *ws.Push, pushCh chan *ws.Push) *groupMsgRead {
	m := &groupMsgRead{
		conversationId: push.ConversationId,
		push:           push,
		pushCh:         pushCh,
		count:          1,
		pushTime:       time.Now(),
		done:           make(chan struct{}),
	}
	go m.transfer()
	return m
}

// 合并
func (m *groupMsgRead) mergePush(push *ws.Push) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.count++
	// 只需要进行记录，因为消息发送有别处完成W
	//fmt.Printf("push %+v\n", m.push)
	for msgId, read := range push.ReadRecords {
		m.push.ReadRecords[msgId] = read
	}
}

func (m *groupMsgRead) transfer() {
	// 1. 超市发送
	// 2。 超量发送
	timer := time.NewTimer(GroupMsgReadRecordDelayTime / 2)
	defer timer.Stop()
	for {
		select {
		case <-m.done:
			return
		case <-timer.C:
			m.mu.Lock()
			val := GroupMsgReadRecordDelayTime*2 - time.Since(m.pushTime)
			push := m.push

			if val > 0 && m.count < GroupMsgReadRecordDelayCount || push == nil {
				// 未达标
				timer.Reset(GroupMsgReadRecordDelayTime / 2)
				m.mu.Unlock()
				continue
			}
			// 达标了，进行推送
			m.pushTime = time.Now()
			m.push = nil
			m.count = 0
			timer.Reset(val)
			m.mu.Unlock()

			logx.Infof("push time condition satified,start pushing %+v", push)
			m.pushCh <- push
		default:
			m.mu.Lock()
			if m.count >= GroupMsgReadRecordDelayCount {
				push := m.push
				m.push = nil
				m.count = 0
				m.mu.Unlock()
				logx.Infof("push count condition satified,start pushing %+v", push)
				m.pushCh <- push
				continue
			}
			if m.isIdle() {
				m.mu.Unlock()
				// 使用msgReadTransfer释放
				m.pushCh <- &ws.Push{
					ConversationId: m.conversationId,
					ChatType:       status.GroupChatType,
				}
				continue
			}
			m.mu.Unlock()

			tempDelay := GroupMsgReadRecordDelayTime / 4
			if tempDelay > time.Second {
				tempDelay = time.Second
			}
			time.Sleep(tempDelay)

		}
	}

}

// IsIdle 检查是否是活跃状态
func (m *groupMsgRead) IsIdle() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.isIdle()
}

func (m *groupMsgRead) isIdle() bool {
	pushTime := m.pushTime
	val := GroupMsgReadRecordDelayTime*2 - time.Since(pushTime)

	if val <= 0 || (m.push == nil && m.count == 0) {
		return true
	}
	return false
}

func (m *groupMsgRead) clear() {
	select {
	case <-m.done:
	default:
		close(m.done)
	}
	m.push = nil
}
