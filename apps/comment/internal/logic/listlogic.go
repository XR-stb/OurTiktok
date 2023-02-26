package logic

import (
	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/internal/svc"
	"OutTiktok/apps/user/user"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"math"
	"strconv"
	"strings"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *comment.ListReq) (*comment.ListRes, error) {
	// 获取评论ID
	key := fmt.Sprintf("cids_%d", in.VideoId)
	pairs, _ := l.svcCtx.Redis.ZrevrangebyscoreWithScores(key, 0, math.MaxInt64)
	if len(pairs) == 0 {
		return &comment.ListRes{}, nil
	}

	commentList := make([]*comment.CommentInfo, len(pairs))
	nonCacheList := make([]int64, 0, len(pairs)) // 未命中列表
	commentIds := make([]int64, len(pairs))
	// 查询缓存
	for i, pair := range pairs {
		id, _ := strconv.ParseInt(pair.Key, 10, 64)
		commentIds[i] = id
		key := fmt.Sprintf("cinfo_%d", id)
		str, err := l.svcCtx.Redis.Get(key)
		if err != nil {
			nonCacheList = append(nonCacheList, id)
			continue
		}
		// 刷新过期时间
		_ = l.svcCtx.Redis.Expire(key, 86400)

		l.svcCtx.CommentCache[id] = parseToComment(id, pair.Score, str)
	}

	// 查询数据库
	if len(nonCacheList) > 0 {
		var queryCommentList []*comment.CommentInfo
		l.svcCtx.DB.Table("comments").Where("id IN ?", nonCacheList)

		for _, info := range queryCommentList {
			l.svcCtx.CommentCache[info.Id] = info
		}

		// 写回缓存
		for _, info := range queryCommentList {
			key := fmt.Sprintf("cinfo_%d", info.Id)
			val := fmt.Sprintf("%d_%s", info.UserId, info.Content)
			_ = l.svcCtx.Redis.Setex(key, val, 86400)
		}
	}

	// 写入结果并获取评论用户ID
	userIds := make([]int64, len(commentList))
	for i, id := range commentIds {
		commentList[i] = l.svcCtx.CommentCache[id]
		userIds[i] = l.svcCtx.CommentCache[id].UserId
	}

	// 获取用户信息
	r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &user.GetUsersReq{
		UserIds: userIds,
		ThisId:  in.UserId,
	})
	if err == nil {
		for _, info := range r.Users {
			l.svcCtx.UserCache[info.Id] = info
		}
		for i := 0; i < len(commentList); i++ {
			commentList[i].User = &comment.UserInfo{}
			_ = copier.Copy(commentList[i].User, l.svcCtx.UserCache[commentList[i].UserId])
		}
	}

	return &comment.ListRes{
		CommentList: commentList,
	}, nil
}

func parseToComment(id int64, createTime int64, str string) *comment.CommentInfo {
	splits := strings.Split(str, "_")
	uid, _ := strconv.ParseInt(splits[0], 10, 64)
	return &comment.CommentInfo{
		Id:         id,
		UserId:     uid,
		Content:    splits[1],
		CreateTime: createTime,
	}
}
