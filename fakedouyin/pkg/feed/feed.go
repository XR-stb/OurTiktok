package feed

import (
	"context"
	"errors"
	"fakedouyin/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var authorCache map[int]Author = make(map[int]Author)

type Status struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type FeedResponse struct {
	Status
	NextTime  int64   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}

func Feed(c *gin.Context) {
	// 验证Token
	token := c.Query("token")
	claims, err := verifyToken(token)
	userId := 0
	if err == nil {
		userId = claims.UserId
	}

	latestTime := c.Query("latest_time")
	var now int64
	if latestTime == "" {
		now := time.Now().UnixMilli()
		latestTime = fmt.Sprintf("%d", now)
	}

	// 查询数据库
	var videolist []Video
	rows := db.Where("upload_time<?", latestTime).Order("upload_time desc").Order("id desc").Limit(30).Find(&videolist).RowsAffected
	if rows == 0 {
		c.JSON(200, FeedResponse{
			Status{}, now, nil,
		})
		return
	}

	// 获取作者信息
	if err := getAuthor(videolist); err != nil {
		c.JSON(200, FeedResponse{
			Status: Status{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 获取点赞数量以及是否点赞
	if err := getFavorite(videolist, userId); err != nil {
		c.JSON(200, FeedResponse{
			Status: Status{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(200, FeedResponse{
		Status{}, videolist[rows-1].UploadTime, videolist,
	})
}

func getAuthor(videoList []Video) error {
	// 获取作者id
	usersId := make([]int64, 0, len(videoList))
	for _, v := range videoList {
		if _, ok := authorCache[v.AuthorId]; !ok {
			usersId = append(usersId, int64(v.AuthorId))
		}
	}

	// rpc调用
	res, err := userclient.GetUserInfo(context.Background(), &service.GetUserInfoReq{
		UsersId: usersId,
	})
	if err != nil {
		errors.New("获取作者信息失败")
	}

	// 写入缓存
	for _, v := range res.UsersInfo {
		authorCache[int(v.UserId)] = Author{
			Id:   int(v.UserId),
			Name: v.Name,
		}
	}

	// 写入videolist
	for i := 0; i < len(videoList); i++ {
		videoList[i].Author = authorCache[videoList[i].AuthorId]
	}

	return nil
}

func getFavorite(videoList []Video, userId int) error {
	// 获取视频Id
	videosId := make([]int64, 0, len(videoList))
	for _, v := range videoList {
		videosId = append(videosId, int64(v.Id))
	}

	// rpc调用
	res, err := favoriteclient.GetFavorite(context.Background(), &service.GetFavoriteReq{
		VideosId: videosId,
		UsersId:  int64(userId),
	})
	if err != nil {
		return errors.New("获取点赞信息失败")
	}

	// 写入viedolist
	for i, v := range res.Favorites {
		videoList[i].FavoriteCount = int(v.Count)
		videoList[i].IsFavorite = v.IsFavorite
	}

	return nil
}
