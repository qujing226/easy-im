package logic

import (
	"context"
	"easy-chat/apps/social/rpc/internal/svc"
	"easy-chat/apps/social/rpc/models"
	"easy-chat/pkg/xerr"
	"github.com/pkg/errors"

	"easy-chat/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupListLogic) GroupList(in *social.GroupListReq) (*social.GroupListResp, error) {
	// todo: add your logic here and delete this line
	var groupMember []models.GroupMember
	result := l.svcCtx.CSvc.DB.Where("user_id = ?", in.UserId).Find(&groupMember)
	if result.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group list by user_id %v err %v", in.UserId, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, errors.WithStack(xerr.FindGroupByIdErr)
	}

	var groups []*social.Groups
	for _, member := range groupMember {
		var group models.Group
		err := l.svcCtx.CSvc.DB.Where("id = ?", member.GroupID).Find(&group).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewDBErr(), "find group by id %v err %v", member.GroupID, err)
		}
		groups = append(groups, &social.Groups{
			Id:              group.ID,
			Name:            group.Name,
			Icon:            group.Icon,
			Status:          int32(*group.Status),
			CreatorUid:      group.CreatorUID,
			GroupType:       int32(*group.GroupType),
			IsVerify:        group.IsVerify,
			Notification:    group.Notification,
			NotificationUid: group.NotificationUID,
		})
	}

	return &social.GroupListResp{
		List: groups,
	}, nil
}
