package xerr

// user rpc
var (
	PhoneNotFound = New(ServerCommonError, "api not found")
	IdNotFound    = New(ServerCommonError, "api id not found")
	UserPwdErr    = New(ServerCommonError, "password is wrong")
	ParamError    = New(RequestParamError, "params error")
)

// social rpc
var (
	FriendAlreadyExists    = New(ServerCommonError, "friend already exists")
	FriendRequestOnPending = New(ServerCommonError, "friend request on pending")
	FriendRequestRefused   = New(ServerCommonError, "friend request refused")
	FrinndListNotFound     = New(ServerCommonError, "friend list not found")
	FindFriendByIdErr      = New(ServerCommonError, "find friend by id error")
)
