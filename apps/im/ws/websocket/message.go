package websocket

import "time"

type FrameType uint8

const (
	FrameData  FrameType = 0x0
	FramePing  FrameType = 0x1
	FrameAck   FrameType = 0x2
	FrameNoAck FrameType = 0x3
	FrameError FrameType = 0x9
)

type Message struct {
	Id        string    `json:"id,omitempty"`
	FrameType FrameType `json:"frameType,omitempty"`
	AckSeq    int       `json:"ackSeq"`
	ackTime   time.Time `json:"ackTime"`
	errCount  int       `json:"errCount"`
	Method    string    `json:"method,omitempty"`
	UserId    string    `json:"userId,omitempty"`
	FormId    string    `json:"formId,omitempty"`
	Data      any       `json:"data,omitempty"`
}

func NewMessage(formId string, data any) *Message {
	return &Message{
		FrameType: FrameData,
		FormId:    formId,
		Data:      data,
	}
}

func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameError,
		Data:      err.Error(),
	}
}
