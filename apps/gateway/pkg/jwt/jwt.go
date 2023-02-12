package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

var (
	secret     = []byte("kfcvme50")
	ExpireTime = 7 * 24 * time.Hour
)

func GetToken(claims *JWTClaims) string {
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(ExpireTime))
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
	// 过期
	if err := tokened.Claims.Valid(); err != nil {
		return nil, errors.New("token已失效")
	}
	return claims, nil
}
