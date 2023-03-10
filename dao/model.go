package dao

type User struct {
	Id              int64  `gorm:"primaryKey;autoincrement"`
	Username        string `gorm:"notnull;size:32;uniqueIndex"`
	Password        string `gorm:"notnull;size:32"`
	Avatar          string `gorm:"notnull;size:128"`
	BackgroundImage string `gorm:"notnull;size:128"`
	Signature       string `gorm:"notnull;size:128"`
}

type Video struct {
	Id         int64  `gorm:"primaryKey;autoincrement"`
	AuthorId   int64  `gorm:"notnull;index"`
	UploadTime int64  `gorm:"notnull"`
	PlayUrl    string `gorm:"notnull;size:128"`
	CoverUrl   string `gorm:"notnull;size:128"`
	Title      string `gorm:"notnull;size:128"`
}

type Favorite struct {
	Id      int64 `gorm:"primaryKey;autoincrement"`
	VideoId int64 `gorm:"notnull;uniqueIndex:idx_vid_uid"`
	UserId  int64 `gorm:"notnull;uniqueIndex:idx_vid_uid"`
	Status  int32 `gorm:"notnull;type:enum('1','2')"` // 1-点赞 2-未点赞
}

type Comment struct {
	Id         int64  `gorm:"primaryKey;autoincrement"`
	VideoId    int64  `gorm:"notnull;index"`
	UserId     int64  `gorm:"notnull"`
	CreateTime int64  `gorm:"notnull"`
	Content    string `gorm:"notnull"`
}

type Relation struct {
	Id         int64 `gorm:"primaryKey;autoincrement"`
	FollowedId int64 `gorm:"notnull;uniqueIndex:idx_vid_uid"`       // 被关注
	FollowerId int64 `gorm:"notnull;uniqueIndex:idx_vid_uid;index"` // 关注者
	Status     int32 `gorm:"notnull;type:enum('1','2')"`            // 1-关注 2-未关注
}
