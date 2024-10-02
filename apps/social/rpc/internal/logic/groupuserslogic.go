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

type GroupUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUsersLogic {
	return &GroupUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupUsersLogic) GroupUsers(in *social.GroupUsersReq) (*social.GroupUsersResp, error) {
	// todo: add your logic here and delete this line
	var groupMembers []models.GroupMember
	result := l.svcCtx.CSvc.DB.Where("group_id = ?", in.GroupId).Find(&groupMembers)
	if result.Error != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find group members by group_id %v err %v", in.GroupId, result.Error)
	}
	if result.RowsAffected == 0 {
		return nil, errors.WithStack(xerr.GroupNotFound)
	}
	var users []*social.GroupMembers
	for _, member := range groupMembers {
		var joinSource int32
		if member.JoinSource != nil {
			joinSource = int32(*member.JoinSource)
		} else {
			joinSource = 0 // 默认值
		}
		users = append(users, &social.GroupMembers{
			Id:          member.ID,
			GroupId:     member.GroupID,
			UserId:      member.UserID,
			RoleLevel:   int32(member.RoleLevel),
			JoinTime:    member.JoinTime.Unix(),
			JoinSource:  joinSource,
			InviterUid:  member.InviterUID,
			OperatorUid: member.OperatorUID,
		})
	}
	return &social.GroupUsersResp{
		List: users,
	}, nil
}
