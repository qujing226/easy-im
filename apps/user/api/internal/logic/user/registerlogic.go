package user

import (
	"context"
	"easy-chat/apps/user/rpc/user"

	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewRegisterLogic 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.User.Register(l.ctx, &user.RegisterReq{
		Phone:    req.Phone,
		Password: req.Password,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Sex:      int32(req.Sex),
	})
	if err != nil {
		return nil, err
	}

	var res types.RegisterResp
	err = copier.Copy(&res, registerResp)
	if err != nil {
		return nil, err
	}
	resp = &res
	return
}
