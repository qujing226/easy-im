package logic

import (
	"context"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/models"
	"easy-chat/pkg/status"
	"easy-chat/pkg/suid"
	"easy-chat/pkg/utils"
	"easy-chat/pkg/xerr"
	"github.com/pkg/errors"
	"time"

	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GroupCreate 群要求
func (l *GroupCreateLogic) GroupCreate(in *social.GroupCreateReq) (*social.GroupCreateResp, error) {
	// todo: add your logic here and delete this line
	var group = models.Group{
		ID:         suid.GenerateID(),
		Name:       in.Name,
		Icon:       in.Icon,
		GroupType:  utils.ConvertToInt8(int8(status.GroupTypeNormal)),
		Status:     utils.ConvertToInt8(in.Status),
		CreatorUID: in.CreatorUid,
		IsVerify:   true,
	}

	var groupMember = models.GroupMember{
		ID:        suid.GenerateID(),
		GroupID:   group.ID,
		UserID:    group.CreatorUID,
		RoleLevel: status.GroupMemberRoleLevelOwner,
		JoinTime:  time.Now(),
	}

	// 开启事务
	tx := l.svcCtx.CSvc.DB.Begin()
	err := tx.Create(&group).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(xerr.NewDBErr(), "create group %v failed err %v", group.ID, err)
	}
	err = tx.Create(&groupMember).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.Wrapf(xerr.NewDBErr(), "create group member %v failed err %v", groupMember.ID, err)
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "commit tx failed err %v", err)
	}

	return &social.GroupCreateResp{
		GroupId: group.ID,
	}, nil
}
