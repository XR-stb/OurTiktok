package logic

import (
	"OutTiktok/apps/user/user"
	"OutTiktok/dao"
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"strconv"
	"time"

	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/internal/svc"

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

func (l *ActionLogic) Action(in *comment.ActionReq) (*comment.ActionRes, error) {
	actionType := in.ActionType

	if actionType == 1 { //发布评论
		nowUnix := time.Now().Unix()
		newComment := &dao.Comment{
			VideoId:    in.VideoId,
			UserId:     in.UserId,
			CreateTime: nowUnix,
			Content:    in.Content,
		}

		// 写入数据库
		if err := l.svcCtx.DB.Create(&newComment).Error; err != nil {
			return &comment.ActionRes{Status: -1}, err
		}

		// 查询用户信息
		userinfo := &comment.UserInfo{}
		r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &user.GetUsersReq{
			UserIds: []int64{in.UserId},
			ThisId:  in.UserId,
		})
		if err == nil {
			_ = copier.Copy(userinfo, r.Users[0])
		}

		// 写入缓存
		// 更新评论信息
		key := fmt.Sprintf("cinfo_%d", newComment.Id)
		val := fmt.Sprintf("%d_%s", newComment.UserId, newComment.Content)
		_ = l.svcCtx.Redis.Setex(key, val, 86400)

		// 更新视频评论ID
		key2 := fmt.Sprintf("cids_%d", newComment.VideoId)
		_, _ = l.svcCtx.Redis.Zadd(key2, newComment.CreateTime, strconv.FormatInt(newComment.Id, 10))

		return &comment.ActionRes{
			CommentInfo: &comment.CommentInfo{
				Id:         newComment.Id,
				User:       userinfo,
				Content:    newComment.Content,
				CreateTime: nowUnix,
			},
		}, nil
	} else { // 删除评论
		if err := l.svcCtx.DB.Delete(&dao.Comment{
			Id:     in.CommentId,
			UserId: in.UserId,
		}).Error; err != nil {
			return &comment.ActionRes{
				Status: -1,
			}, err
		}

		// 更新缓存
		// 更新评论信息
		key := fmt.Sprintf("cinfo_%d", in.CommentId)
		_, _ = l.svcCtx.Redis.Del(key)
		// 更新视频评论ID
		key2 := fmt.Sprintf("cids_%d", in.VideoId)
		_, _ = l.svcCtx.Redis.Zrem(key2, in.CommentId)

		return &comment.ActionRes{}, nil
	}
}
