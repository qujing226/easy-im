package logic

import (
	"context"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/models"
	"easy-chat/apps/social/rpc/social"
	"easy-chat/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinListLogic {
	return &GroupPutinListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutinListLogic) GroupPutinList(in *social.GroupPutinListReq) (*social.GroupPutinListResp, error) {
	// todo: add your logic here and delete this line
	var groupRequests []models.GroupRequest
	result := l.svcCtx.CSvc.DB.Where("group_id = ? ", in.GroupId).Find(&groupRequests)
	if result.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group put in list by group_id %v err %v", in.GroupId, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, errors.WithStack(xerr.GroupPutInNotFound)
	}
	var groupRequestsResp []*social.GroupRequests
	for _, request := range groupRequests {
		groupRequestsResp = append(groupRequestsResp, &social.GroupRequests{
			Id:           request.ID,
			GroupId:      request.GroupID,
			ReqId:        request.ReqID,
			ReqMsg:       request.ReqMsg,
			ReqTime:      request.ReqTime.Unix(),
			JoinSource:   int32(*request.JoinSource),
			InviterUid:   request.InviterUserID,
			HandleUid:    request.HandleUserID,
			HandleResult: int32(request.HandleResult),
		})
	}
	return &social.GroupPutinListResp{
		List: groupRequestsResp,
	}, nil
}
