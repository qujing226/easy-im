package user

import (
	"context"
	"easy-chat/apps/user/api/internal/svc"
	"easy-chat/apps/user/api/internal/types"
	"easy-chat/apps/user/rpc/user"
	"easy-chat/pkg/status"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewLoginLogic 用户登入
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.User.Login(l.ctx, &user.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var res types.LoginResp
	err = copier.Copy(&res, loginResp)
	if err != nil {
		return nil, err
	}
	resp = &res

	// 处理登入的业务
	err = l.svcCtx.Redis.HsetCtx(l.ctx, status.REDIS_ONLINE_USER, loginResp.Id, "1")

	// 为每个用户的在线状态单独设置过期时间
	//expireKey := fmt.Sprintf("%s:%s", status.REDIS_ONLINE_USER, loginResp.Id)
	//expireErr := l.svcCtx.Redis.ExpireCtx(l.ctx, expireKey, 30*60)
	//if expireErr != nil {
	//	return nil, expireErr
	//}

	return
}
