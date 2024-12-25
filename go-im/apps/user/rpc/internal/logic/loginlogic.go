package logic

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/models"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"
	"github.com/peninsula12/easy-im/go-im/pkg/encrypy"
	"github.com/peninsula12/easy-im/go-im/pkg/xerr"
	"github.com/pkg/errors"
	"time"

	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/user"

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
	u := &models.User{}
	var err error
	// 1.检查用户是否存在(phone)
	err = l.svcCtx.CSvc.GetUserByPhone(u, in.Phone)
	if err != nil {
		if u.ID == "" {
			return nil, errors.WithStack(xerr.PhoneNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find api by phone "+
			"%v err %v ", in.Phone, err)
	}

	// 2. 密码验证
	if !encrypy.ValidatePasswordHash([]byte(u.Password), []byte(in.Password)) {
		return nil, errors.WithStack(xerr.UserPwdErr)
	}
	// 3. 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, u.ID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "etxdata get jwt token"+
			" err %v ", in.Phone)
	}
	return &user.LoginResp{
		Id:     u.ID,
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
