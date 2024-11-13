package status

type MType int8

const (
	TextMType MType = iota
)

type ChatType int

const (
	GroupChatType ChatType = iota
	SingleChatType
)
