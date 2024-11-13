package websocket

type FrameType uint8

const (
	FrameData  FrameType = 0x0
	FramePing  FrameType = 0x1
	FrameError FrameType = 0x9
)

type Message struct {
	FrameType FrameType `json:"frameType,omitempty"`
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
