package feed

import (
	"fakedouyin/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Author struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Video struct {
	Id            int    `json:"id"`
	AuthorId      int    `json:"-"`
	Author        Author `json:"author"`
	UploadTime    int64  `json:"-"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

func init() {
	var err error
	// 连接Mysql
	if db, err = gorm.Open(mysql.Open(config.C.Mysql.Dsn)); err != nil {
		log.Fatalln(err)
	}
}
