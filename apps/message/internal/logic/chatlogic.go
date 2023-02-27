package logic

import (
	"OutTiktok/apps/message/internal/svc"
	"OutTiktok/apps/message/message"
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatLogic) Chat(in *message.MessageChatReq) (*message.MessageChatRes, error) {
	key := fmt.Sprintf("chat:%d:%d", in.FromUserID, in.ToUserId)
	chatKey, err := l.svcCtx.Redis.Get(key)
	if err != nil {
		return &message.MessageChatRes{Status: -1}, nil
	}
	if chatKey == "" {
		return &message.MessageChatRes{}, nil
	}

	// 每次获取十条
	start := in.PreMsgTime + 1
	pairs, err := l.svcCtx.Redis.ZrangebyscoreWithScoresAndLimit(chatKey, start, math.MaxInt64, 0, 10)
	msgList := make([]*message.MessageInfo, len(pairs))
	for i, pair := range pairs {
		msgList[i] = parseToMsg(pair.Score, pair.Key)
	}

	return &message.MessageChatRes{
		MessageList: msgList,
	}, nil
}

func parseToMsg(unix int64, str string) *message.MessageInfo {
	splits := strings.Split(str, "_")
	id, _ := strconv.ParseInt(splits[0], 10, 64)
	fromUserId, _ := strconv.ParseInt(splits[1], 10, 64)
	toUserId, _ := strconv.ParseInt(splits[2], 10, 64)
	return &message.MessageInfo{
		Id:         id,
		ToUserId:   toUserId,
		FromUserId: fromUserId,
		Content:    splits[3],
		CreateTime: unix,
	}
}
