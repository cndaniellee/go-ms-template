package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtAuthConf struct {
	AccessSecret string
	AccessExpire int64
}

const JwtUserIdKey = "UserId"

type JwtClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

func GenerateUserToken(config JwtAuthConf, userId int64) (string, error) {
	claims := &JwtClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(config.AccessExpire))),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.AccessSecret))
}
