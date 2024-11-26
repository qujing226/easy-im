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

type GroupPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInHandleLogic {
	return &GroupPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutInHandleLogic) GroupPutInHandle(in *social.GroupPutInHandleReq) (*social.GroupPutInHandleResp, error) {
	// todo: add your logic here and delete this line
	var groupReq models.GroupRequest
	err := l.svcCtx.CSvc.DB.Where("id = ? and handle_result = ?", in.GroupReqId, status.PendingHandlerResult).First(&groupReq).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(xerr.GroupPutInNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group request by req_id %v err %v", in.GroupReqId, err)
	}
	switch status.HandlerResult(in.HandleResult) {
	case status.RefuseHandlerResult:
		groupReq.HandleResult = status.RefuseHandlerResult
		groupReq.HandleTime = time.Now()
		groupReq.HandleUserID = in.HandleUid
		err = l.svcCtx.CSvc.DB.Updates(&groupReq).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "refuse group request by id %v err %v", in.GroupReqId, err)
		}
	case status.CancelHandlerResult:
		// 我不确定取消的过程是谁来完成的,这个逻辑应该是不会走的
		groupReq.HandleResult = status.CancelHandlerResult
		groupReq.HandleTime = time.Now()
		groupReq.HandleUserID = in.HandleUid
		err = l.svcCtx.CSvc.DB.Updates(&groupReq).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "cancle group request by id %v err %v", in.GroupReqId, err)
		}
	case status.PassHandlerResult:
		// 更新请求
		groupReq.HandleResult = status.PassHandlerResult
		groupReq.HandleTime = time.Now()
		groupReq.HandleUserID = in.HandleUid

		tx := l.svcCtx.CSvc.DB.Begin()
		err = tx.Updates(&groupReq).Error
		if err != nil {
			tx.Rollback()
			return nil, errors.Wrapf(xerr.NewDBErr(), "update group request by id %v err %v", in.GroupReqId, err)
		}

		// 添加 member 数据
		// 通过nickname查id
		var groupMember = models.GroupMember{
			ID:          suid.GenerateID(),
			GroupID:     groupReq.GroupID,
			UserID:      groupReq.ReqID, // 不确定这俩爹username是指ID还是nickname
			RoleLevel:   status.GroupMemberRoleLevelMember,
			JoinTime:    time.Now(),
			JoinSource:  groupReq.JoinSource,
			InviterUID:  groupReq.InviterUserID,
			OperatorUID: groupReq.HandleUserID,
		}
		err = tx.Create(&groupMember).Error
		if err != nil {
			tx.Rollback()
			return nil, errors.Wrapf(xerr.NewDBErr(), "create group member by id %v err %v", in.GroupReqId, err)
		}
		err = tx.Commit().Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "commit tx failed err %v", err)
		}
	default:
		return nil, errors.WithStack(xerr.ParamError)
	}
	if groupReq.HandleResult == status.PassHandlerResult {
		return &social.GroupPutInHandleResp{
			GroupId: in.GroupId,
		}, nil
	}
	return &social.GroupPutInHandleResp{}, nil
}
