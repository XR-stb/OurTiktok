package publish

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

var (
	secret = []byte("kfcvme50")
)

// 验证Token
func verifyToken(token string) (*JWTClaims, error) {
	tokened, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, errors.New("token已失效")
	}
	claims, ok := tokened.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("token已失效")
	}
	if err := tokened.Claims.Valid(); err != nil {
		return nil, errors.New("token已失效")
	}
	return claims, nil
}
