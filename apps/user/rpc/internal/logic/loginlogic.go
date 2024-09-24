package logic

import (
	"context"
	"easy-chat/apps/user/common/models"
	"easy-chat/pkg/ctxdata"
	"easy-chat/pkg/encrypy"
	"errors"
	"time"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	var u models.User
	// 1.检查用户是否存在(phone)
	l.svcCtx.DB.Where("phone = ?", in.Phone).First(&u)
	if u.ID == "" {
		return nil, errors.New("user doesn't exists")
	}

	// 2. 密码验证
	if !encrypy.ValidatePasswordHash([]byte(u.Password), []byte(in.Password)) {
		return nil, errors.New("password is wrong")
	}
	// 3. 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, u.ID)
	if err != nil {
		return nil, err
	}
	return &user.LoginResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
