package friend

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"

	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFriendPutInListLogic 好友申请列表
func NewFriendPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInListLogic {
	return &FriendPutInListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInListLogic) FriendPutInList(req *types.FriendPutInListReq) (resp *types.FriendPutInListResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)

	friendsReqList, err := l.svcCtx.Social.FriendPutInList(l.ctx, &social.FriendPutInListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	if len(friendsReqList.List) == 0 {
		return &types.FriendPutInListResp{}, nil
	}
	var members []*types.FriendRequests
	for _, requests := range friendsReqList.List {
		members = append(members, &types.FriendRequests{
			Id:           requests.Id,
			UserId:       requests.UserId,
			ReqUid:       requests.ReqUid,
			ReqMsg:       requests.ReqMsg,
			ReqTime:      requests.ReqTime,
			HandleResult: int(requests.HandleResult),
			HandleMsg:    requests.HandleMsg,
		})
	}
	resp = &types.FriendPutInListResp{
		List: members,
	}
	return
}
