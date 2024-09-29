package logic

import (
	"context"
	"easy-chat/apps/user/common/models"
	"errors"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	var u models.User
	var userEntity user.UserEntity
	l.svcCtx.DB.Where("id = ?", in.User).Find(&u)
	if u.ID == "" {
		l.Logger.Error(errors.New("user not exists"))
		return nil, errors.New("user not exists")
	}

	userEntity = user.UserEntity{
		Id:       u.ID,
		Avatar:   u.Avatar,
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Status:   int32(u.Status),
		Sex:      int32(u.Sex),
	}

	return &user.GetUserInfoResp{User: &userEntity}, nil
}
