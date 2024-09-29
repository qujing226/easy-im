package logic

import (
	"context"
	"easy-chat/apps/user/common/models"
	"easy-chat/pkg/ctxdata"
	"easy-chat/pkg/encrypy"
	"easy-chat/pkg/suid"
	"github.com/pkg/errors"
	"time"

	"easy-chat/apps/user/rpc/internal/svc"
	"easy-chat/apps/user/rpc/user"

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
	var u models.User
	// 1.检查用户是否存在(phone)
	l.svcCtx.DB.Where("phone = ?", in.Phone).First(&u)
	if u.ID != "" {
		l.Logger.Error(errors.New("user exists"))
		return nil, errors.New("user exists")
	}

	// 2.定义新增用户
	u = models.User{
		ID:        suid.GenerateID(),
		Avatar:    in.Avatar,
		Nickname:  in.Nickname,
		Phone:     in.Phone,
		Status:    0,
		Sex:       int8(in.Sex),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	if in.Password != "" {
		pass, err := encrypy.GenPasswordHash([]byte(in.Password))
		if err != nil {
			l.Logger.Error(err)
			return nil, err
		}
		u.Password = string(pass)
	}
	// 3.保存用户
	err := l.svcCtx.DB.Create(&u).Error
	if err != nil {
		l.Logger.Error(err)
		return nil, errors.New("save user failed")
	}

	// 4. 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, u.ID)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil

}
