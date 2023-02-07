package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

var (
	secret     = []byte("kfcvme50")
	ExpireTime = 7 * 24 * time.Hour
)

func getToken(claims *JWTClaims) string {
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return token
}

func VerifyToken(token string) (*JWTClaims, error) {
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
