package user

import (
	"github.com/gin-gonic/gin"
	"time"
)

type RegResponse struct {
	Status
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}

var (
	queryEmpty = RegResponse{
		Status{
			StatusCode: -1,
			StatusMsg:  "用户名或密码为空",
		}, -1, "",
	}
	usernameExist = RegResponse{
		Status{
			StatusCode: -1,
			StatusMsg:  "用户名已被占用",
		}, -1, "",
	}
)

func Reg(c *gin.Context) {
	// 获取用户名和密码
	user := User{}
	_ = c.ShouldBindQuery(&user)
	if user.Username == "" || user.Password == "" {
		c.JSON(200, queryEmpty)
		return
	}

	// 将用户写入数据库
	user.Password = passwordEncrypt(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		c.JSON(200, usernameExist)
		return
	}

	// 生成Token
	claims := &JWTClaims{
		UserId:   user.Id,
		Username: user.Username,
	}
	claims.IssuedAt = time.Now().Unix()                  //当前时间
	claims.ExpiresAt = time.Now().Add(ExpireTime).Unix() //过期时间
	token := getToken(claims)

	// 返回
	c.JSON(200, RegResponse{Status{}, user.Id, token})
}
