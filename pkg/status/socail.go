package status

// HandlerResult 1. pending 2. processed 3. rejected
type HandlerResult int8

const (
	PendingHandlerResult HandlerResult = 1 + iota
	PassHandlerResul
	RefuseHandlerResul
	CancelHandlerResul
)
