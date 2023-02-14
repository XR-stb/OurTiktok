package comment

import (
	"OutTiktok/apps/comment/comment"
	"OutTiktok/apps/gateway/internal/svc"
	"OutTiktok/apps/gateway/internal/types"
	"OutTiktok/apps/gateway/pkg/jwt"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (resp *types.CommentListRes, err error) {
	resp = &types.CommentListRes{}

	// 验证Token
	var UserId int64
	claims, err := jwt.VerifyToken(req.Token)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return
	}
	UserId = claims.UserId

	r, err := l.svcCtx.CommentClient.List(context.Background(), &comment.ListReq{
		UserId:  UserId,
		VideoId: req.VideoId,
	})

	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		return resp, nil
	}

	if r.Status != 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "评论列表获取失败"
		return
	}

	resp.StatusMsg = "成功"
	resp.CommentList = make([]types.Comment, len(r.CommentList))
	for i := 0; i < len(r.CommentList); i++ {
		_ = copier.Copy(&resp.CommentList[i], r.CommentList[i])
	}

	return
}
