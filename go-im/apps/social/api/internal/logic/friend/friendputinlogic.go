package friend

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/types"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFriendPutInLogic 好友申请
func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInLogic) FriendPutIn(req *types.FriendPutInReq) (resp *types.FriendPutInResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	_, err = l.svcCtx.Social.FriendPutIn(l.ctx, &social.FriendPutInReq{
		UserId:  uid,
		ReqUid:  req.UserId,
		ReqMsg:  req.ReqMsg,
		ReqTime: req.ReqTime,
	})
	return
}
