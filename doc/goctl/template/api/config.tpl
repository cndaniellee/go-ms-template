package config

import (
    {{.authImport}}
    "goms/common/auth"
)

type Config struct {
	rest.RestConf

	JwtAuth auth.JwtAuthConf
}
