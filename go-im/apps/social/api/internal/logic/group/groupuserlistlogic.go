package group

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/types"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/user"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGroupUserListLogic 成员列表列表
func NewGroupUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserListLogic {
	return &GroupUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUserListLogic) GroupUserList(req *types.GroupUserListReq) (resp *types.GroupUserListResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.GroupUserListResp{}
	_ = ctxdata.GetUId(l.ctx)
	members, err := l.svcCtx.Social.GroupUsers(l.ctx, &social.GroupUsersReq{GroupId: req.GroupId})
	if err != nil {
		return
	}
	if len(members.List) == 0 {
		return
	}
	ids := make([]string, 0, len(members.List))
	for _, member := range members.List {
		ids = append(ids, member.UserId)
	}

	users, err := l.svcCtx.User.FindUser(l.ctx, &user.FindUserReq{
		Ids: ids,
	})
	if err != nil {
		return
	}
	if len(users.Users) == 0 {
		return
	}
	memberEntities := make(map[string]*user.UserEntity, len(ids))
	for _, entity := range users.Users {
		memberEntities[entity.Id] = entity
	}

	var list []*types.GroupMembers
	for i, member := range members.List {
		ids[i] = member.UserId
		list = append(list, &types.GroupMembers{
			Id:            member.Id,
			GroupId:       member.GroupId,
			UserId:        member.UserId,
			Nickname:      memberEntities[member.UserId].Nickname,
			UserAvatarUrl: memberEntities[member.UserId].Avatar,
			RoleLevel:     int(member.RoleLevel),
			InviterUid:    member.InviterUid,
			OperatorUid:   member.OperatorUid,
		})
	}
	resp.List = list
	return
}
