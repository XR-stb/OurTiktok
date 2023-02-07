package publish

type Status struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type Author struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	FollowCount   int    `json:"follow_count,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
