package logic

import (
	"OutTiktok/apps/favorite/favorite"
	"OutTiktok/apps/favorite/internal/svc"
	"OutTiktok/apps/publish/publishclient"
	"OutTiktok/dao"
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

func (l *ListLogic) List(in *favorite.ListReq) (*favorite.ListRes, error) {
	userId := in.UserId
	thisId := in.ThisId

	// 查询数据库
	var favoriteList []dao.Favorite
	rows := int(l.svcCtx.DB.Where("user_id=? AND status=?", userId, 1).Order("video_id").Find(&favoriteList).RowsAffected)
	if rows == 0 {
		return &favorite.ListRes{}, nil
	}
	videoList := make([]*favorite.Video, rows)

	// 查询视频信息
	videoIds := make([]int64, rows)
	for i := 0; i < rows; i++ {
		videoIds[i] = favoriteList[i].VideoId
	}
	if r, err := l.svcCtx.PublishClient.GetVideos(context.Background(), &publishclient.GetVideosReq{
		UserId:   thisId,
		VideoIds: videoIds,
	}); err == nil {
		for i := 0; i < rows; i++ {
			videoList[i] = &favorite.Video{}
			_ = copier.Copy(videoList[i], r.VideoList[i])
		}
	}

	return &favorite.ListRes{
		VideoList: videoList,
	}, nil
}
