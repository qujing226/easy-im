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

type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	// todo: add your logic here and delete this line
	// 1. 获取好友申请记录
	var friendReq models.FriendRequest
	err := l.svcCtx.CSvc.DB.Where("id = ? and handle_result = ?", in.FriendReqId, status.PendingHandlerResult).First(&friendReq).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.WithStack(xerr.FindFriendByIdErr)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find friend by id %v err %v", in.FriendReqId, err)
	}
	// 处理操作的状态
	switch status.HandlerResult(in.HandleResult) {
	case status.RefuseHandlerResult:
		{
			// 拒绝好友申请
			friendReq.HandleResult = status.RefuseHandlerResult
			friendReq.HandleMsg = "申请被拒绝"
			friendReq.HandledAt = time.Now()
			err = l.svcCtx.CSvc.DB.Updates(&friendReq).Error
			if err != nil {
				return nil, errors.Wrapf(xerr.NewDBErr(), "update friend request by id %v err %v", in.FriendReqId, err)
			}
		}
	case status.CancelHandlerResult:
		{
			// 将好友申请状态置为取消
			friendReq.HandleResult = status.CancelHandlerResult
			friendReq.HandleMsg = "申请已取消"
			friendReq.HandledAt = time.Now()
			err = l.svcCtx.CSvc.DB.Updates(&friendReq).Error
			if err != nil {
				return nil, errors.Wrapf(xerr.NewDBErr(), "update friend request by id %v err %v", in.FriendReqId, err)
			}
		}

	case status.PassHandlerResult:
		{
			// 2. 更新好友申请记录
			friendReq.HandleResult = status.PassHandlerResult
			friendReq.HandleMsg = "申请通过"
			friendReq.HandledAt = time.Now()
			// 3. 更新请求记录 + 建立两条好友关系记录  -->  事务
			tx := l.svcCtx.CSvc.DB.Begin()
			err = tx.Updates(&friendReq).Error
			if err != nil {
				tx.Rollback()
				return nil, errors.Wrapf(xerr.NewDBErr(), "update friend request by id %v err %v", in.FriendReqId, err)
			}
			friends := []models.Friend{
				{
					ID:        suid.GenerateID(),
					UserID:    friendReq.UserID,
					FriendUID: friendReq.ReqUID,
				},
				{
					ID:        suid.GenerateID(),
					UserID:    friendReq.ReqUID,
					FriendUID: friendReq.UserID,
				},
			}
			err = tx.Create(&friends).Error
			if err != nil {
				tx.Rollback()
				return nil, errors.Wrapf(xerr.NewDBErr(), "create friend by user_id %v and friend_uid %v err %v", friendReq.UserID, friendReq.ReqUID, err)
			}

			// commit
			err = tx.Commit().Error
			if err != nil {
				return nil, errors.Wrapf(xerr.NewDBErr(), "commit tx err %v", err)
			}
		}
	default:
		return nil, errors.WithStack(xerr.ParamError)
	}
	return &social.FriendPutInHandleResp{}, nil
}
