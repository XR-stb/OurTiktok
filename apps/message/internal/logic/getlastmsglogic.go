package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"OutTiktok/apps/message/internal/svc"
	"OutTiktok/apps/message/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLastMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLastMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLastMsgLogic {
	return &GetLastMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLastMsgLogic) GetLastMsg(in *message.GetLastMsgReq) (*message.GetLastMsgRes, error) {
	lastMsg := make([]*message.LastMsg, len(in.ToUserId))
	fromUserId := in.FromUserId
	for i, toUserId := range in.ToUserId {
		// 获取ChatKey
		key := fmt.Sprintf("chat:%d:%d", fromUserId, toUserId)
		chatKey, err := l.svcCtx.Redis.Get(key)
		if err != nil || chatKey == "" {
			continue
		}

		// 获取最后一条消息
		msg, err := l.svcCtx.Redis.Zrevrange(chatKey, 0, 0)
		if err != nil || len(msg) == 0 {
			continue
		}

		// 解析消息
		splits := strings.Split(msg[0], "_")
		msgFromUserId, _ := strconv.ParseInt(splits[1], 10, 64)
		lastMsg[i] = &message.LastMsg{
			Message: splits[3],
		}
		if fromUserId == msgFromUserId {
			lastMsg[i].MsgType = 1
		}
	}

	return &message.GetLastMsgRes{
		LastMsg: lastMsg,
	}, nil
}
