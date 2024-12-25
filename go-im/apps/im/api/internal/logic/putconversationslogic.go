package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/imclient"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"

	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewPutConversationsLogic 更新会话
func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutConversationsLogic) PutConversations(req *types.PutConversationsReq) (resp *types.PutConversationsResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	var conversationList map[string]*imclient.Conversation
	err = copier.Copy(&conversationList, req.ConversationList)
	if err != nil {
		return
	}

	_, err = l.svcCtx.PutConversations(l.ctx, &imclient.PutConversationsReq{
		UserId:           uid,
		ConversationList: conversationList,
	})
	return
}
