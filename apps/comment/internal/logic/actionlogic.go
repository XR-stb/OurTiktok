package logic

import (
	"OutTiktok/dao"
	"context"
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
	newComment := &dao.Comment{
		VideoId:    in.VideoId,
		UserId:     in.UserId,
		CreateDate: time.Now(),
		Content:    in.Content,
	}

	if err := l.svcCtx.DB.Create(newComment).Error; err != nil {
		return &comment.ActionRes{
			Status: -1,
		}, nil
	}

	return &comment.ActionRes{
		CommentInfo: &comment.CommentInfo{
			Id:         newComment.Id,
			UserInfo:   nil,
			Content:    in.Content,
			CreateDate: newComment.CreateDate.Format("01-02"),
		},
	}, nil
}
