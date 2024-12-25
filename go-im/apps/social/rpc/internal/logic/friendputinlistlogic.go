package logic

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/models"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/pkg/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInListLogic {
	return &FriendPutInListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInListLogic) FriendPutInList(in *social.FriendPutInListReq) (*social.FriendPutInListResp, error) {
	// todo: add your logic here and delete this line
	// 获取好友请求列表
	var friendReqList []models.FriendRequest
	result := l.svcCtx.CSvc.DB.Where("user_id = ?", in.UserId).Find(&friendReqList)
	if result.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friend request list by req_uid %v err %v", in.UserId, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, errors.WithStack(xerr.FriendReqListNotFound)
	}
	var friends []*social.FriendRequests

	for _, friend := range friendReqList {
		friends = append(friends, &social.FriendRequests{
			Id:           friend.ID,
			UserId:       friend.UserID,
			ReqUid:       friend.ReqUID,
			ReqMsg:       friend.ReqMsg,
			ReqTime:      friend.ReqTime.Unix(),
			HandleResult: int32(friend.HandleResult),
			HandleMsg: friend.HandleMsg,
		})
	}

	return &social.FriendPutInListResp{
		List: friends,
	}, nil
}
