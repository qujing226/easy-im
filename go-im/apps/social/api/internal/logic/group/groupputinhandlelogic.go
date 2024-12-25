package group

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/imclient"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"
	"github.com/peninsula12/easy-im/go-im/pkg/status"

	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutInHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGroupPutInHandleLogic 申请进群处理
func NewGroupPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInHandleLogic {
	return &GroupPutInHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutInHandleLogic) GroupPutInHandle(req *types.GroupPutInHandleRep) (resp *types.GroupPutInHandleResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	res, err := l.svcCtx.Social.GroupPutInHandle(l.ctx, &social.GroupPutInHandleReq{
		GroupReqId:    req.GroupReqId,
		GroupId:       req.GroupId,
		HandleUid:     uid,
		HandleResult:  req.HandleResult,
		Username:      "",
		UserAvatarUrl: "",
	})
	if err != nil {
		return
	}
	if res.GroupId == "" {
		return
	}
	if status.HandlerResult(req.HandleResult) != status.PassHandlerResult {
		return
	}


	_, err = l.svcCtx.SetUpUserConversation(l.ctx, &imclient.SetUpUserConversationReq{
		SendId:   uid,
		RecvId:   res.GroupId,
		ChatType: int32(status.GroupChatType),
	})

	return
}
