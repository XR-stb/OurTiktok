package user

import "github.com/gin-gonic/gin"

type GetUserResponse struct {
	Status
	User Info `json:"user"`
}

type Info struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

var (
	notExist = GetUserResponse{
		Status: Status{
			StatusCode: -1,
			StatusMsg:  "未找到用户",
		},
	}
)

func GetUser(c *gin.Context) {
	userId := c.Query("user_id")

	// 查询
	user := User{}
	result := db.Where("Id=?", userId).First(&user)
	if result.Error != nil || result.RowsAffected < 1 {
		c.JSON(200, notExist)
		return
	}

	// 获取关注/粉丝数量 redis/跨微服务

	// 返回
	c.JSON(200, GetUserResponse{
		Status{}, Info{
			Id:   user.Id,
			Name: user.Username,
		},
	})
}
