package config

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}
