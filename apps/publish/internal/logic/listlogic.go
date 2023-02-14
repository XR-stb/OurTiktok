package logic

import (
	"OutTiktok/apps/user/userclient"
	"context"
	"github.com/jinzhu/copier"

	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

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

func (l *ListLogic) List(in *publish.ListReq) (*publish.ListRes, error) {
	authorId := in.UserId

	// 查询数据库
	var videoList []*publish.Video
	rows := int(l.svcCtx.DB.Where("author_id=?", authorId).Find(&videoList).RowsAffected)

	// 查询作者信息
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &userclient.GetUsersReq{
		ThisId:  in.ThisId,
		UserIds: []int64{authorId},
	}); err == nil {
		userinfo := &publish.UserInfo{
			Id: authorId,
		}
		copier.Copy(userinfo, r.Users[0])
		for i := 0; i < rows; i++ {
			videoList[i].Author = userinfo
		}
	}
	return &publish.ListRes{
		VideoList: videoList,
	}, nil
}
