package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/peninsula12/easy-im/go-im/apps/im/model"
	"github.com/peninsula12/easy-im/go-im/pkg/xerr"

	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/im"
	"github.com/peninsula12/easy-im/go-im/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConversationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetConversations 获取会话
func (l *GetConversationsLogic) GetConversations(in *im.GetConversationsReq) (*im.GetConversationsResp, error) {
	// todo: add your logic here and delete this line
	// 根据用户查询用户会话列表
	data, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		if errors.Is(err, immodels.ErrNotFound) {
			return &im.GetConversationsResp{}, nil
		}
		return nil, errors.Wrapf(xerr.NewDBErr(), "ConversationsModel.FindByUserId err %v,req %v", err, in.UserId)
	}
	var res im.GetConversationsResp

	err = copier.Copy(&res, &data)
	if err != nil {
		return &im.GetConversationsResp{}, errors.Wrapf(xerr.NewDBErr(), "copier err %v", err)
	}

	// 根据会话列表，查询具体的会话
	ids := make([]string, 0, len(data.ConversationList))
	for _, v := range data.ConversationList {
		ids = append(ids, v.ConversationId)
	}
	conversations, err := l.svcCtx.ConversationModel.ListByConversationIds(l.ctx, ids)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "conversationModel.ListByConversationIds err %v,req %v", err, ids)
	}
	// 计算是否包含未读消息
	for _, v := range conversations {
		if _, ok := res.ConversationList[v.ConversationId]; !ok {
			continue
		}
		// 用户读取的消息量
		total := res.ConversationList[v.ConversationId].Total
		if total < int32(v.Total) {
			// 有新消息
			res.ConversationList[v.ConversationId].Total = int32(v.Total)
			// 有多少是未读的
			res.ConversationList[v.ConversationId].ToRead = int32(v.Total) - total
			// 更改当前会话为显示状态
			res.ConversationList[v.ConversationId].IsShow = true
		}
	}

	return &res, nil
}
