package friend

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/types"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/user"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/userclient"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFriendListLogic 好友列表
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)

	friends, err := l.svcCtx.Social.FriendList(l.ctx, &social.FriendListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	if len(friends.List) == 0 {
		return &types.FriendListResp{}, nil
	}

	// 获取好友的个人信息
	ids := make([]string, 0, len(friends.List))
	for _, f := range friends.List {
		ids = append(ids, f.FriendUid)
	}
	users, err := l.svcCtx.User.FindUser(l.ctx, &user.FindUserReq{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}

	userRecords := make(map[string]*userclient.UserEntity, len(users.Users))
	for i, _ := range users.Users {
		userRecords[users.Users[i].Id] = users.Users[i]
	}

	respList := make([]*types.Friends, len(friends.List))
	for index, f := range friends.List {
		respList[index] = &types.Friends{
			Id:        f.Id,
			FriendUid: f.FriendUid,
			Remark:    f.Remark,
			Avatar:    userRecords[f.FriendUid].Avatar,
			Nickname:  userRecords[f.FriendUid].Nickname,
		}
	}
	return &types.FriendListResp{List: respList}, err
}
