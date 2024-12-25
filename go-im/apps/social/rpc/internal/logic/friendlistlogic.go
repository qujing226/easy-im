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

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *social.FriendListReq) (*social.FriendListResp, error) {
	// todo: add your logic here and delete this line
	// 查询 friend 列表
	var friendList []models.Friend
	result := l.svcCtx.CSvc.DB.Where("user_id = ?", in.UserId).Find(&friendList)
	if result.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friend list by user_id %v err %v", in.UserId, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, errors.WithStack(xerr.FriendListNotFound)
	}

	var friends []*social.Friends

	for _, friend := range friendList {
		friends = append(friends, &social.Friends{
			FriendUid: friend.FriendUID,
			Remark:    friend.Remark,
		})
	}

	return &social.FriendListResp{
		List: friends,
	}, nil
}
