package favorite

import (
	"fakedouyin/pkg/config"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var r redis.Conn

type Favorite struct {
	Id      int `gorm:"primaryKey;autoincrement"`
	VideoId int `gorm:"notnull;uniqueindex:idx_vid_uid"`
	UserId  int `gorm:"notnull;uniqueindex:idx_vid_uid;index"`
	Status  int `gorm:"notnull;type:enum('1','2')"` // 1-点赞 2-未点赞
}

func init() {
	var err error
	// 连接Mysql
	if db, err = gorm.Open(mysql.Open(config.C.Mysql.Dsn)); err != nil {
		log.Fatalln(err)
	}
	// 格式化数据库
	if err := db.AutoMigrate(&Favorite{}); err != nil {
		log.Fatalln(err)
	}
	// 连接Redis
	if r, err = redis.Dial("tcp", config.C.Redis.IP+config.C.Redis.Port); err != nil {
		log.Fatalln(err)
	}
}
