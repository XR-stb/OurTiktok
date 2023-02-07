package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(dsn string) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	if err := DB.AutoMigrate(&User{}, &Video{}, &Favorite{}, &Comment{}, &Relation{}); err != nil {
		panic(err)
	}
	return DB
}
