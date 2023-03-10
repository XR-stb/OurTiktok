// Code generated by goctl. DO NOT EDIT.
package types

type Status struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type User struct {
	Id              int64  `json:"id"`
	Username        string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type Video struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type Comment struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type FriendUser struct {
	User
	Message string `json:"message"`
	MsgType int32  `json:"msg_type"`
}

type Message struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type FeedReq struct {
	LatestTime int64  `form:"latest_time,default=0"`
	Token      string `form:"token"`
}

type FeedRes struct {
	Status
	NextTime  int64   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}

type RegisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterRes struct {
	Status
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginRes struct {
	Status
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type UserRes struct {
	Status
	User User `json:"user"`
}

type PublishActionReq struct {
	Data  []byte `json:"data"`
	Token string `json:"token"`
	Title string `form:"title"`
}

type PublishActionRes struct {
	Status
}

type PublishListReq struct {
	Token  string `form:"token"`
	UserId int64  `form:"user_id"`
}

type PublishListRes struct {
	Status
	VideoList []Video `json:"video_list"`
}

type FavoriteActionReq struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int32  `form:"action_type"`
}

type FavoriteActionRes struct {
	Status
}

type FavoriteListReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type FavoriteListRes struct {
	Status
	VideoList []Video `json:"video_list"`
}

type CommentActionReq struct {
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

type CommentActionRes struct {
	Status
	Comment Comment `json:"comment"`
}

type CommentListReq struct {
	Token   string `form:"token"`
	VideoId int64  `form:"video_id"`
}

type CommentListRes struct {
	Status
	CommentList []Comment `json:"comment_list"`
}

type RelationActionReq struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"`
}

type RelationActionRes struct {
	Status
}

type RelationFollowListReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type RelationFollowListRes struct {
	Status
	UserList []User `json:"user_list"`
}

type RelationFollowerListReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type RelationFollowerListRes struct {
	Status
	UserList []User `json:"user_list"`
}

type RelationFriendListReq struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type RelationFriendListRes struct {
	Status
	UserList []FriendUser `json:"user_list"`
}

type MessageActionReq struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"`
	Content    string `form:"content"`
}

type MessageActionRes struct {
	Status
}

type MessageChatReq struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	PreMsgTime int64  `form:"pre_msg_time"`
}

type MessageChatRes struct {
	Status
	MessageList []Message `json:"message_list"`
}
