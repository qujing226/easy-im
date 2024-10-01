package logic

import (
	"context"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/models"
	"easy-chat/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"easy-chat/apps/social/rpc/social"

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
	err := l.svcCtx.CSvc.DB.Where("user_id = ?", in.UserId).Find(&friendList).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.WithStack(xerr.FrinndListNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friend list by user_id %v err %v", in.UserId, err)
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
