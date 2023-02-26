package logic

import (
	"context"
	"fmt"
	"time"

	"OutTiktok/apps/message/internal/svc"
	"OutTiktok/apps/message/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActionLogic) Action(in *message.MessageActionReq) (*message.MessageActionRes, error) {
	now := time.Now().Unix()
	first := false
	// 获取ChatKey
	key1 := fmt.Sprintf("chat:%d:%d", in.FromUserID, in.ToUserId)
	chatKey, err := l.svcCtx.Redis.Get(key1)
	if err != nil {
		return &message.MessageActionRes{Status: -1}, nil
	}
	if chatKey == "" {
		// 分配ChatKey
		chatKey = fmt.Sprintf("msglist:%d:%d", in.FromUserID, in.ToUserId)
		key2 := fmt.Sprintf("chat:%d:%d", in.ToUserId, in.FromUserID)
		_ = l.svcCtx.Redis.Setex(key1, chatKey, 604800)
		_ = l.svcCtx.Redis.Setex(key2, chatKey, 604800)
		first = true
	}

	msgId, err := l.svcCtx.Redis.Incr("autoIncr")
	if err != nil {
		return &message.MessageActionRes{Status: -1}, nil
	}
	msg := fmt.Sprintf("%d_%d_%d_%s", msgId, in.FromUserID, in.ToUserId, in.Content)
	_, _ = l.svcCtx.Redis.Zadd(chatKey, now, msg)
	if first {
		_ = l.svcCtx.Redis.Expire(chatKey, 604800)
	}

	return &message.MessageActionRes{}, nil
}
