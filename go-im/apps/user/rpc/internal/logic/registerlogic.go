package logic

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/models"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"
	"github.com/peninsula12/easy-im/go-im/pkg/encrypy"
	"github.com/peninsula12/easy-im/go-im/pkg/suid"
	"github.com/peninsula12/easy-im/go-im/pkg/utils"
	"github.com/peninsula12/easy-im/go-im/pkg/xerr"
	"github.com/pkg/errors"
	"time"

	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	u := models.User{}
	var err error
	// 1.检查用户是否存在(phone)
	err = l.svcCtx.CSvc.GetUserByPhone(&u, in.Phone)
	if err != nil {
		if u.ID == "" {
			return nil, errors.WithStack(xerr.PhoneNotFound)
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "find api by phone "+
			" err %v req %v", err, in.Phone)
	}

	// 2.定义新增用户
	U := &models.User{
		ID:       suid.GenerateID(),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Status:   utils.ConvertToInt8(0),
		Sex:      utils.ConvertToInt8(in.Sex),
	}

	if in.Password != "" {
		pass, err := encrypy.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, errors.Wrapf(xerr.NewServerCommonErr(), "passwordHash gen err %v", err)
		}
		U.Password = string(pass)
	}
	// 3.保存用户
	err = l.svcCtx.CSvc.CreateUser(U)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "save api %v failed ,err %v", in, err)
	}

	// 4. 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, u.ID)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "extdata get jwt token"+
			" err %v", in.Phone)
	}
	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil

}
