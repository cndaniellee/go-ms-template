package jwtauth

import (
	"context"
	"fmt"
	"github.com/cndaniellee/go-utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

type JwtAuthConf struct {
	AccessSecret string
	AccessExpire int64
}

const (
	JwtUserIdKey      = "UserId"
	JwtExtraCodeKey   = "ExtraCode"
	CacheExtraUserKey = "extra:user:%d"
)

type JwtClaims struct {
	UserId    int64
	ExtraCode string
	jwt.RegisteredClaims
}

type TokenGenerator struct {
	rds           *redis.Redis
	accessSecret  string
	accessExpire  int64
	refreshSecret string
	refreshExpire int64
}

func NewTokenGenerator(rds *redis.Redis, authConf JwtAuthConf, refreshConf JwtAuthConf) *TokenGenerator {
	return &TokenGenerator{
		rds:           rds,
		accessSecret:  authConf.AccessSecret,
		accessExpire:  authConf.AccessExpire,
		refreshSecret: refreshConf.AccessSecret,
		refreshExpire: refreshConf.AccessExpire,
	}
}

func (g *TokenGenerator) ExecuteUser(ctx context.Context, userId int64) (userToken string, refreshToken string, err error) {

	// 生成扩展码，用于Token唯一性、单点登录
	extraCode := utils.RandStr(16)

	// 签发登录Token
	tokenClaims := &JwtClaims{
		UserId:    userId,
		ExtraCode: extraCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(g.accessExpire))),
		},
	}
	userToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString([]byte(g.accessSecret))
	if err != nil {
		return
	}

	// 签发刷新Token
	refreshClaims := &JwtClaims{
		UserId:    userId,
		ExtraCode: extraCode,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(g.refreshExpire))),
		},
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(g.refreshSecret))
	if err != nil {
		return
	}

	// 缓存ExtraCode
	err = g.rds.SetexCtx(ctx, fmt.Sprintf(CacheExtraUserKey, userId), extraCode, int(g.refreshExpire))
	return
}
