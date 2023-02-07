package publish

import (
	"fakedouyin/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Video struct {
	Id            int    `gorm:"primaryKey;autoincrement" json:"id,omitempty"`
	Author        Author `gorm:"-" json:"author"`
	AuthorId      int    `gorm:"notnull" json:"-"`
	UploadTime    int64  `gorm:"index;notnull" json:"-"`
	PlayUrl       string `gorm:"size:128;notnull" json:"play-url,omitempty"`
	CoverUrl      string `gorm:"size:128;notnull" json:"cover-url,omitempty"`
	FavoriteCount int    `gorm:"-" json:"favorite_count"`
	CommentCount  int    `gorm:"-" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"-" json:"is_favorite"`
	Title         string `gorm:"size:128;notnull" json:"title,omitempty"`
}

func init() {
	var err error
	// 连接Mysql
	if db, err = gorm.Open(mysql.Open(config.C.Mysql.Dsn)); err != nil {
		log.Fatalln(err)
	}
	// 格式化数据库
	if err := db.AutoMigrate(&Video{}); err != nil {
		log.Fatalln(err)
	}
}
