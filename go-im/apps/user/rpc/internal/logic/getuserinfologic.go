package logic

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/models"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/user"
	"github.com/peninsula12/easy-im/go-im/pkg/xerr"
	"github.com/pkg/errors"

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
	u := make([]models.User, 0, 1)
	s := make([]string, 1)
	s[0] = in.User
	var userEntity user.UserEntity
	err := l.svcCtx.CSvc.GetUserByIds(&u, s)
	ur := u[0]
	if err != nil {
		if ur.ID == "" {
			return nil, errors.WithStack(xerr.IdNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find api by id "+
			" err %v req %v", err, in.User)
	}

	userEntity = user.UserEntity{
		Id:       ur.ID,
		Avatar:   ur.Avatar,
		Nickname: ur.Nickname,
		Phone:    ur.Phone,
		Status:   int32(*ur.Status),
		Sex:      int32(*ur.Sex),
	}
	return &user.GetUserInfoResp{User: &userEntity}, nil
}
