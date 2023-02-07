package publish

import (
	"context"
	"errors"
	"fakedouyin/pkg/service"
	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	Status
	VideoList []Video `json:"video_list"`
}

func List(c *gin.Context) {
	// 验证Token
	token := c.Query("token")
	claims, err := verifyToken(token)
	userId := 0
	if err == nil {
		userId = claims.UserId
	}

	authorId := c.Query("user_id")

	// 查询数据库
	var videoList []Video
	db.Where("author_id=?", authorId).Find(&videoList)

	if len(videoList) == 0 {
		c.JSON(200, ListResponse{})
		return
	}

	// 查询作者信息
	if err := getAuthor(videoList); err != nil {
		c.JSON(200, ListResponse{
			Status: Status{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	// 查询点赞
	if err := getFavorite(videoList, userId); err != nil {
		c.JSON(200, ListResponse{
			Status: Status{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(200, ListResponse{
		Status{},
		videoList,
	})
}

func getAuthor(videoList []Video) error {
	// 获取作者id
	AuthorId := videoList[0].AuthorId

	// rpc调用
	res, err := userclient.GetUserInfo(context.Background(), &service.GetUserInfoReq{
		UsersId: []int64{int64(AuthorId)},
	})
	if err != nil {
		errors.New("获取作者信息失败")
	}

	// 写入videolist
	authorName := res.UsersInfo[0].Name
	for i := 0; i < len(videoList); i++ {
		videoList[i].Author.Id = AuthorId
		videoList[i].Author.Name = authorName
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
