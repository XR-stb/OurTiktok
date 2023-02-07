package favorite

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type FavoriteResponse struct {
	Status
}

var (
	favoriteSuccess = FavoriteResponse{}
	favortieFail    = FavoriteResponse{
		Status{-1, "点赞失败"},
	}
)

func Action(c *gin.Context) {
	// 验证Token
	token := c.Query("token")
	claims, err := verifyToken(token)
	if err != nil {
		c.JSON(200, favortieFail)
		return
	}
	userId := claims.UserId

	// 处理数据
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(200, favortieFail)
		return
	}
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		c.JSON(200, favortieFail)
		return
	}

	favorite := Favorite{
		0, videoId, userId, actionType,
	}
	// 查询缓存
	key := "f_" + c.Query("video_id")
	res, err := r.Do("get", key)
	if err != nil {
		c.JSON(200, favortieFail)
		return
	}

	if res != nil {
		if actionType == 1 {
			r.Do("incr", key)
			r.Do("expire", key, 86400)
			if err := db.Create(&favorite).Error; err != nil {
				db.Model(&favorite).Where("video_id=? AND user_id=?", videoId, userId).Update("status", actionType)
			}
		} else {
			r.Do("decr", key)
			r.Do("expire", key, 86400)
			db.Model(&favorite).Where("video_id=? AND user_id=?", videoId, userId).Update("status", actionType)
		}
	} else {
		if actionType == 1 {
			if err := db.Create(&favorite).Error; err != nil {
				db.Model(&favorite).Where("video_id=? AND user_id=?", videoId, userId).Update("status", actionType)
			}
		} else {
			db.Model(&favorite).Where("video_id=? AND user_id=?", videoId, userId).Update("status", actionType)
		}
		var count int64
		db.Model(&favorite).Where("video_id=?", videoId).Count(&count)
		r.Do("setex", key, 86400, count)
	}

	// 返回
	c.JSON(200, favoriteSuccess)
}
