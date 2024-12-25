package group

import (
	"context"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/imclient"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/svc"
	"github.com/peninsula12/easy-im/go-im/apps/social/api/internal/types"
	"github.com/peninsula12/easy-im/go-im/apps/social/rpc/social"
	"github.com/peninsula12/easy-im/go-im/pkg/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewCreateGroupLogic 创群
func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.GroupCreateReq) (resp *types.GroupCreateResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctxdata.GetUId(l.ctx)
	res, err := l.svcCtx.Social.GroupCreate(l.ctx, &social.GroupCreateReq{
		Name:       req.Name,
		Icon:       req.Icon,
		Status:     1,
		CreatorUid: uid,
	})
	if err != nil{
		return
	}
	if res.GroupId == ""{
		return
	}
	// 建立会话
	_,err = l.svcCtx.CreateGroupConversation(l.ctx,&imclient.CreateGroupConversationReq{
		GroupId:  res.GroupId,
		CreateId: uid,
	})
	return
}
