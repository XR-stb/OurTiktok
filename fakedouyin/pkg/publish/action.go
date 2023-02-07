package publish

import (
	"fakedouyin/pkg/config"
	_ "fakedouyin/pkg/logger"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type ActionResponse struct {
	Status
}

var (
	uploadFail = ActionResponse{
		Status{
			StatusCode: -1,
			StatusMsg:  "上传失败",
		},
	}
	uploadSuccess = ActionResponse{}
)

func Action(c *gin.Context) {
	// 验证Token
	token := c.PostForm("token")
	claims, err := verifyToken(token)
	if err != nil {
		c.JSON(200, uploadFail)
		return
	}

	// 获取文件
	file, err := c.FormFile("data")
	title := c.PostForm("title")
	if err != nil || file == nil || title == "" {
		c.JSON(200, uploadFail)
		return
	}

	// 写入本地
	filename := newId()
	suffix := strings.Split(file.Filename, ".")
	dst := "/static/video/" + filename + "." + suffix[len(suffix)-1]
	err = c.SaveUploadedFile(file, "."+dst)
	if err != nil {
		c.JSON(200, uploadFail)
		return
	}

	// 写入数据库
	video := Video{
		AuthorId:   claims.UserId,
		UploadTime: time.Now().UnixMilli(),
		PlayUrl:    "http://" + config.C.Server.IP + config.C.Server.Port + dst,
		CoverUrl:   "",
		Title:      title,
	}
	db.Create(&video)

	// 返回
	c.JSON(200, uploadSuccess)
}
