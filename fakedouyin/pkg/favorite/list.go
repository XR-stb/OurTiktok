package favorite

import "github.com/gin-gonic/gin"

type ListResponse struct {
	Status
	VideoList []Favorite `json:"video_list"`
}

func List(c *gin.Context) {
	userId := c.Query("user_id")

	var favoriteList []Favorite
	db.Where("user_id=? AND status=?", userId, 1).Find(&favoriteList)

	c.JSON(200, ListResponse{
		VideoList: favoriteList,
	})
}
