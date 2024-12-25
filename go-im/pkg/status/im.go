package status

type MType int8

const (
	TextMType MType = iota
)

type ChatType int

const (
	SingleChatType ChatType = iota + 1
	GroupChatType
)

type ContentType int

const (
	ContentChatMsg ContentType = iota + 1
	ContentMarkRead
)
