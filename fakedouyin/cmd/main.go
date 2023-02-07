package main

import (
	"fakedouyin/pkg/config"
	"fakedouyin/pkg/favorite"
	"fakedouyin/pkg/feed"
	_ "fakedouyin/pkg/logger"
	"fakedouyin/pkg/publish"
	"fakedouyin/pkg/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)
	r.Run(config.C.Server.Port)
}

func initRouter(r *gin.Engine) {
	r.Static("/static/video/", "./static/video")

	r.GET("/douyin/feed", feed.Feed)

	r.POST("/douyin/user/register/", user.Reg)
	r.POST("/douyin/user/login/", user.Log)
	r.GET("/douyin/user/", user.GetUser)

	r.POST("/douyin/publish/action/", publish.Action)
	r.GET("/douyin/publish/list/", publish.List)

	r.POST("/douyin/favorite/action/", favorite.Action)
	r.POST("/douyin/favorite/list/", favorite.List)
}
