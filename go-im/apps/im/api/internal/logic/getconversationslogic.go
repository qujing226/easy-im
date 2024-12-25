package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/im"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"

	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetConversationsLogic 获取会话
func NewGetConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConversationsLogic) GetConversations(req *types.GetConversationsReq) (resp *types.GetConversationsResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	data, err := l.svcCtx.GetConversations(l.ctx, &im.GetConversationsReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	var res types.GetConversationsResp
	err = copier.Copy(&res, data)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
