package logic

import (
	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/comment/internal/svc"
	"OutTiktok/apps/user/user"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
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
	// 根据videoId 返回 评论列表
	var commentList []*comment.CommentInfo
	rows := int(l.svcCtx.DB.Table("comments").Where("video_id = ?", in.VideoId).Find(&commentList).RowsAffected)
	if rows < 1 {
		return &comment.ListRes{
			Status: -1,
		}, nil
	}

	// 获取用户信息
	userIds := make([]int64, rows)
	for i := 0; i < rows; i++ {
		userIds[i] = commentList[i].UserId
	}

	r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &user.GetUsersReq{
		UserIds: userIds,
		ThisId:  in.UserId,
	})
	if err == nil {
		for _, info := range r.Users {
			l.svcCtx.UserCache[info.Id] = info
		}
		for i := 0; i < rows; i++ {
			commentList[i].UserInfo = &comment.UserInfo{}
			_ = copier.Copy(commentList[i].UserInfo, l.svcCtx.UserCache[commentList[i].UserId])
		}
	}

	return &comment.ListRes{
		Status:      0,
		CommentList: commentList,
	}, nil
}
