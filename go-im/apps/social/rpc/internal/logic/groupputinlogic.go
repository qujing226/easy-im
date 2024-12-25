package logic

import (
	"context"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/models"
	"easy-chat/pkg/status"
	"easy-chat/pkg/suid"
	. "easy-chat/pkg/utils"
	"easy-chat/pkg/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinLogic {
	return &GroupPutinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutinLogic) GroupPutin(in *social.GroupPutinReq) (*social.GroupPutinResp, error) {
	// todo: add your logic here and delete this line
	//  1. 普通用户申请 ： 如果群无验证直接进入
	//  2. 群成员邀请： 如果群无验证直接进入
	//  3. 群管理员/群创建者邀请：直接进入群
	var groupReq = models.GroupRequest{
		ID:            suid.GenerateID(),
		ReqID:         in.ReqId,
		GroupID:       in.GroupId,
		ReqMsg:        in.ReqMsg,
		ReqTime:       time.Unix(in.ReqTime, 0),
		JoinSource:    ConvertToInt8(in.JoinSource),
		InviterUserID: in.InviterUid,
		HandleResult:  status.PendingHandlerResult,
		HandleTime:    time.Now(),
	}

	// 群是否开启了验证功能
	// 检查群是否有验证
	var group models.Group
	err := l.svcCtx.CSvc.DB.Where("id = ?", in.GroupId).Find(&group).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(xerr.GroupNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "get group %v failed err %v", in.GroupId, err)
	}
	if group.IsVerify == false {
		// 直接加入
		groupReq.HandleResult = status.PassHandlerResult
		tx := l.svcCtx.CSvc.DB.Begin()
		err = tx.Create(&models.GroupMember{
			ID:         suid.GenerateID(),
			UserID:     in.ReqId,
			GroupID:    in.GroupId,
			JoinSource: ConvertToInt8(in.JoinSource),
			JoinTime:   time.Now(),
			RoleLevel:  status.GroupMemberRoleLevelMember,
			InviterUID: in.InviterUid,
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, errors.Wrapf(xerr.NewDBErr(), "create group member by id %v err %v", in.GroupId, err)
		}
		err = tx.Create(&groupReq).Error
		if err != nil {
			tx.Rollback()
			return nil, errors.Wrapf(xerr.NewDBErr(), "create group putin %v failed err %v", groupReq.ID, err)
		}
		err = tx.Commit().Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "commit tx failed err %v", err)
		}
		return &social.GroupPutinResp{}, nil
	}

	// 检查是否是管理者邀请入群
	if in.InviterUid != "" {
		var groupMember models.GroupMember
		err := l.svcCtx.CSvc.DB.Where("user_id = ?", in.InviterUid).Find(&groupMember).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.WithStack(xerr.GroupInviterNotFound)
			}
			return nil, errors.Wrapf(xerr.NewDBErr(), "get group %v failed err %v", in.GroupId, err)
		}
		if groupMember.RoleLevel == status.GroupMemberRoleLevelAdmin || groupMember.RoleLevel == status.GroupMemberRoleLevelOwner {
			groupReq.HandleResult = status.PassHandlerResult
			groupReq.HandleUserID = groupMember.UserID
			groupReq.ReqTime = time.Now()
			tx := l.svcCtx.CSvc.DB.Begin()
			err = tx.Create(&groupReq).Error
			if err != nil {
				tx.Rollback()
				return nil, errors.Wrapf(xerr.NewDBErr(), "create group putin %v failed err %v", groupReq.ID, err)
			}
			err = tx.Create(&models.GroupMember{
				ID:          suid.GenerateID(),
				GroupID:     in.GroupId,
				UserID:      in.ReqId,
				RoleLevel:   status.GroupMemberRoleLevelMember,
				JoinTime:    time.Now(),
				JoinSource:  groupReq.JoinSource,
				InviterUID:  groupReq.InviterUserID,
				OperatorUID: groupReq.InviterUserID,
			}).Error
			if err != nil {
				tx.Rollback()
				return nil, errors.Wrapf(xerr.NewDBErr(), "create group member by id %v err %v", in.GroupId, err)
			}
			err = tx.Commit().Error
			if err != nil {
				return nil, errors.Wrapf(xerr.NewDBErr(), "commit tx failed err %v", err)
			}
			return &social.GroupPutinResp{}, nil
		}
	}
	// 普通申请
	err = l.svcCtx.CSvc.DB.Create(&groupReq).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "create group putin %v failed err %v", groupReq.ID, err)
	}
	return &social.GroupPutinResp{
		GroupId: groupReq.GroupID,
	}, nil
}
