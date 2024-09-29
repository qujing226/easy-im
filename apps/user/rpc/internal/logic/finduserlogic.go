package logic

import (
	"context"
	"easy-chat/apps/user/common/models"
	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	// todo: add your logic here and delete this line
	var users []models.User
	var userEntities []*user.UserEntity

	if in.Phone != "" {
		err := l.svcCtx.DB.Where("phone = ?", in.Phone).Find(&users).Error
		if err != nil {
			l.Logger.Error(err)
			return nil, err
		}
	} else if len(in.Ids) > 0 {
		err := l.svcCtx.DB.Where("id IN (?)", in.Ids).Find(&users).Error
		if err != nil {
			l.Logger.Error(err)
			return nil, err
		}
	} else if in.Name != "" {
		err := l.svcCtx.DB.Where("nickname LIKE ?", "%"+in.Name+"%").Find(&users).Error
		if err != nil {
			l.Logger.Error(err)
			return nil, err
		}
	} else {
		l.Logger.Error(errors.New("params error"))
		return nil, errors.New("params error")
	}

	userEntities = make([]*user.UserEntity, len(users))

	for index, u := range users {
		userEntities[index] = &user.UserEntity{
			Id:       u.ID,
			Avatar:   u.Avatar,
			Nickname: u.Nickname,
			Phone:    u.Phone,
			Status:   int32(u.Status),
			Sex:      int32(u.Sex),
		}
	}

	return &user.FindUserResp{Users: userEntities}, nil
}
