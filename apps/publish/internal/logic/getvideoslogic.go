package logic

import (
	"OutTiktok/apps/user/userclient"
	"context"
	"github.com/jinzhu/copier"

	"OutTiktok/apps/publish/internal/svc"
	"OutTiktok/apps/publish/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideosLogic {
	return &GetVideosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideosLogic) GetVideos(in *publish.GetVideosReq) (*publish.GetVideosRes, error) {
	videoIds := in.VideoIds

	// 查询数据库
	var videoList []*publish.Video
	rows := int(l.svcCtx.DB.Where("id IN ?", videoIds).Find(&videoList).RowsAffected)

	// 查询作者信息
	AuthorIds := make([]int64, rows)
	for i := 0; i < rows; i++ {
		AuthorIds[i] = videoList[i].AuthorId
	}
	if r, err := l.svcCtx.UserClient.GetUsers(context.Background(), &userclient.GetUsersReq{
		ThisId:  in.UserId,
		UserIds: AuthorIds,
	}); err == nil {
		for _, user := range r.Users {
			l.svcCtx.UserCache[user.Id] = user
		}
		for i := 0; i < rows; i++ {
			userinfo := l.svcCtx.UserCache[videoList[i].AuthorId]
			videoList[i].Author = &publish.UserInfo{}
			copier.Copy(videoList[i].Author, userinfo)
		}
	}

	return &publish.GetVideosRes{
		VideoList: videoList,
	}, nil
}
