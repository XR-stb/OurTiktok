package dao

import "time"

type User struct {
	Id       int64  `gorm:"primaryKey;autoincrement"`
	Username string `gorm:"size:32;uniqueIndex;notnull"`
	Password string `gorm:"size:32;notnull"`
}

type Video struct {
	Id         int64  `gorm:"primaryKey;autoincrement"`
	AuthorId   int64  `gorm:"notnull"`
	UploadTime int64  `gorm:"notnull;index"`
	PlayUrl    string `gorm:"notnull;size:128" `
	CoverUrl   string `gorm:"notnull;size:128" `
	Title      string `gorm:"notnull;size:128" `
}

type Favorite struct {
	Id      int64 `gorm:"primaryKey;autoincrement"`
	VideoId int64 `gorm:"notnull;uniqueIndex:idx_vid_uid"`
	UserId  int64 `gorm:"notnull;uniqueIndex:idx_vid_uid"`
	Status  int32 `gorm:"notnull;type:enum('1','2')"` // 1-点赞 2-未点赞
}

type Comment struct {
	Id         int64     `gorm:"primaryKey;autoincrement"`
	VideoId    int64     `gorm:"notnull;index"`
	UserId     int64     `gorm:"notnull"`
	CreateDate time.Time `gorm:"notnull"`
	Content    string    `gorm:"notnull"`
}

type Relation struct {
	Id         int64 `gorm:"primaryKey;autoincrement"`
	FollowedId int64 `gorm:"notnull;uniqueIndex:idx_vid_uid"`       // 被关注
	FollowerId int64 `gorm:"notnull;uniqueIndex:idx_vid_uid;index"` // 关注者
	Status     int32 `gorm:"notnull;type:enum('1','2')"`            // 1-关注 2-未关注
}
