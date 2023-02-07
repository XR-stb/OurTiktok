package user

import (
	"github.com/gin-gonic/gin"
	"time"
)

type LogResponse struct {
	Status
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}

var (
	wrongNorP = LogResponse{
		Status{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误",
		}, -1, "",
	}
)

func Log(c *gin.Context) {
	// 获取用户名和密码
	user := User{}
	c.ShouldBindQuery(&user)

	// 查询账号并验证密码
	user.Password = passwordEncrypt(user.Password)
	if err := db.Select("Id").Where("username=? AND password=?", user.Username, user.Password).First(&user).Error; err != nil {
		c.JSON(200, wrongNorP)
		return
	}

	//生成Token
	claims := &JWTClaims{
		UserId:   user.Id,
		Username: user.Username,
	}
	claims.IssuedAt = time.Now().Unix()                  //当前时间
	claims.ExpiresAt = time.Now().Add(ExpireTime).Unix() //过期时间
	token := getToken(claims)

	//返回
	c.JSON(200, LogResponse{Status{}, user.Id, token})
}
