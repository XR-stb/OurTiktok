package user

import (
	"crypto/md5"
	"fmt"
)

type Status struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 加密
func passwordEncrypt(pw string) string {
	hash := md5.Sum([]byte(pw))
	return fmt.Sprintf("%x", hash)
}
