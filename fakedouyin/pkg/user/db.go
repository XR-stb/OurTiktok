package user

import (
	"fakedouyin/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type User struct {
	Id       int    `gorm:"primaryKey;autoincrement"`
	Username string `gorm:"size:32;uniqueIndex;notnull" form:"username"`
	Password string `gorm:"size:32;notnull" form:"password"`
}

func init() {
	var err error
	// 连接Mysql
	if db, err = gorm.Open(mysql.Open(config.C.Mysql.Dsn)); err != nil {
		log.Fatalln(err)
	}
	// 格式化数据库
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatalln(err)
	}
}
