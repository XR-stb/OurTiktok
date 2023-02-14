package logic

import (
	"OutTiktok/apps/user/user"
	"OutTiktok/dao"
	"context"
	"github.com/jinzhu/copier"
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
		newComment := &dao.Comment{
			VideoId:    in.VideoId,
			UserId:     in.UserId,
			CreateDate: time.Now(),
			Content:    in.Content,
		}

		if err := l.svcCtx.DB.Create(&newComment).Error; err != nil {
			return &comment.ActionRes{
				Status: -1,
			}, err
		}

		var userinfo *comment.UserInfo
		r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &user.GetUsersReq{
			UserIds: []int64{in.UserId},
			ThisId:  in.UserId,
		})
		if err == nil {
			_ = copier.Copy(userinfo, r.Users[0])
		}

		return &comment.ActionRes{
			CommentInfo: &comment.CommentInfo{
				Id:         newComment.Id,
				UserInfo:   userinfo,
				Content:    newComment.Content,
				CreateDate: newComment.CreateDate.Format("01-02"),
			},
		}, nil
	} else { // 删除评论
		if err := l.svcCtx.DB.Delete(&dao.Comment{}).Where("id = ? AND user_id = ?", in.CommentId, in.UserId).Error; err != nil {
			return &comment.ActionRes{
				Status: -1,
			}, err
		}

		return &comment.ActionRes{}, nil
	}
}
