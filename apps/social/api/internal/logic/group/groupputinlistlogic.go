package group

import (
	"context"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/pkg/ctxdata"

	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutInListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGroupPutInListLogic 申请进群列表
func NewGroupPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInListLogic {
	return &GroupPutInListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutInListLogic) GroupPutInList(req *types.GroupPutInListRep) (resp *types.GroupPutInListResp, err error) {
	// todo: add your logic here and delete this line
	_ = ctxdata.GetUId(l.ctx)
	groupPutinList, err := l.svcCtx.Social.GroupPutinList(l.ctx, &social.GroupPutinListReq{
		GroupId: req.GroupId,
	})
	if err != nil{
		return
	}
	if len(groupPutinList.List) == 0{
		return &types.GroupPutInListResp{}, nil
	}
	resp = &types.GroupPutInListResp{}
	var list []*types.GroupRequests
	for _, requests := range groupPutinList.List {
		list = append(list, &types.GroupRequests{
			Id:            requests.Id,
			UserId:        requests.ReqId,
			GroupId:       requests.GroupId,
			ReqMsg:        requests.ReqMsg,
			ReqTime:       requests.ReqTime,
			JoinSource:    int64(requests.JoinSource),
			InviterUserId: requests.InviterUid,
			HandleUserId:  requests.HandleUid,
			HandleTime:    requests.HandleTime,
			HandleResult:  int64(requests.HandleResult),
		})
	}
	resp.List = list
	return
}
