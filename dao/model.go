package dao

import "time"

type User struct {
	Id       int    `gorm:"primaryKey;autoincrement"`
	Username string `gorm:"size:32;uniqueIndex;notnull"`
	Password string `gorm:"size:32;notnull"`
}

type Video struct {
	Id         int    `gorm:"primaryKey;autoincrement"`
	AuthorId   int    `gorm:"notnull"`
	UploadTime int64  `gorm:"notnull;index"`
	PlayUrl    string `gorm:"notnull;size:128" `
	CoverUrl   string `gorm:"notnull;size:128" `
	Title      string `gorm:"notnull;size:128" `
}

type Favorite struct {
	Id      int `gorm:"primaryKey;autoincrement"`
	VideoId int `gorm:"notnull;unique index:idx_vid_uid"`
	UserId  int `gorm:"notnull;unique index:idx_vid_uid;index"`
	Status  int `gorm:"notnull;type:enum('1','2')"` // 1-点赞 2-未点赞
}

type Comment struct {
	Id         int       `gorm:"primaryKey;autoincrement"`
	VideoId    int       `gorm:"notnull;index"`
	UserId     int       `gorm:"notnull;"`
	CreateTime time.Time `gorm:"notnull"`
	Content    string    `gorm:"notnull"`
}

type Relation struct {
	Id         int `gorm:"primaryKey;autoincrement"`
	FollowedId int `gorm:"notnull;unique index:idx_vid_uid"`       // 被关注
	FollowerId int `gorm:"notnull;unique index:idx_vid_uid;index"` // 关注者
	Status     int `gorm:"notnull;type:enum('1','2')"`             // 1-关注 2-未关注
}
