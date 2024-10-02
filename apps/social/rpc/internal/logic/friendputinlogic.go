package logic

import (
	"context"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/models"
	"easy-chat/pkg/status"
	"easy-chat/pkg/suid"
	"easy-chat/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// todo: add your logic here and delete this line
	var friend models.Friend
	// 1. 申请人是否与目标是好友关系 此操作不需要使用缓存
	err := l.svcCtx.CSvc.DB.Where("user_id = ? and friend_uid = ?", in.UserId, in.ReqUid).First(&friend).Error
	if err == nil {
		return nil, errors.WithStack(xerr.FriendAlreadyExists)
	}
	if err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friend by user_id %v and friend_uid %v err %v", in.UserId, in.ReqUid, err)
	}
	// 2. 是否已经有过申请，请申请尚未通过
	var friendReq models.FriendRequest
	err = l.svcCtx.CSvc.DB.Where("req_uid = ? and user_id = ? and handle_result != ?", in.ReqUid, in.UserId, status.PassHandlerResult).First(&friendReq).Error
	if err == nil {
		return nil, errors.WithStack(xerr.FriendRequestOnPending)
	}
	if err != gorm.ErrRecordNotFound {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friend request refused by req_uid %v and user_id %v err %v", in.ReqUid, in.UserId, err)
	}
	// 3. 创建申请记录
	err = l.svcCtx.CSvc.DB.Debug().Create(&models.FriendRequest{
		ID:           suid.GenerateID(),
		UserID:       in.UserId,
		ReqUID:       in.ReqUid,
		ReqMsg:       in.ReqMsg,
		ReqTime:      time.Unix(in.ReqTime, 0),
		HandleResult: status.PendingHandlerResult,
		HandledAt:    time.Now(),
	}).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "create friend request by user_id %v and req_uid %v err %v", in.UserId, in.ReqUid, err)
	}
	return &social.FriendPutInResp{}, nil
}
