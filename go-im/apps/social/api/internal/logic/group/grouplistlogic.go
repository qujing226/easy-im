package group

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"

	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGroupListLogic 用户申群列表
func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupListLogic) GroupList(req *types.GroupListRep) (resp *types.GroupListResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	groups, err := l.svcCtx.Social.GroupList(l.ctx, &social.GroupListReq{
		UserId: uid,
	})
	if err != nil{
		return
	}
	if len(groups.List) == 0{
		return &types.GroupListResp{}, nil
	}
	resp = &types.GroupListResp{}
	var list []*types.Groups
	for _, group := range groups.List {
		list = append(list, &types.Groups{
			Id:              group.Id,
			Name:            group.Name,
			Icon:            group.Icon,
			Status:          int64(group.Status),
			GroupType:       int64(group.GroupType),
			IsVerify:        group.IsVerify,
			Notification:    group.Notification,
			NotificationUid: group.NotificationUid,
		})
	}
	resp.List = list
	return
}
