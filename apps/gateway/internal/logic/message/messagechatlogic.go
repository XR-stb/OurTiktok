package message

import (
	"OutTiktok/apps/gateway/pkg/jwt"
	"OutTiktok/apps/message/message"
	"context"
	"github.com/jinzhu/copier"

	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageChatLogic) MessageChat(req *types.MessageChatReq) (resp *types.MessageChatRes, err error) {
	resp = &types.MessageChatRes{}

	// 验证Token
	var UserId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	UserId = claims.UserId

	r, err := l.svcCtx.MessageClient.Chat(context.Background(), &message.MessageChatReq{
		FromUserID: UserId,
		ToUserId:   req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	if r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "获取失败"
		return
	}

	resp.MessageList = make([]types.Message, len(r.MessageList))
	for i, info := range r.MessageList {
		resp.MessageList[i] = types.Message{}
		_ = copier.Copy(&resp.MessageList[i], info)
	}

	return
}
