package status

// HandlerResult 1. pending 2. processed 3. rejected
type HandlerResult int8

const (
	PendingHandlerResult HandlerResult = 1 + iota
	PassHandlerResult
	RefuseHandlerResult
	CancelHandlerResult
)

type GroupMemberRoleLevel int8

const (
	GroupMemberRoleLevelOwner GroupMemberRoleLevel = 1 + iota
	GroupMemberRoleLevelAdmin
	GroupMemberRoleLevelMember
)
type GroupType int8

const (
	GroupTypeNormal GroupType = 1 + iota
	GroupTypePrivate
	GroupTypePublic
)
