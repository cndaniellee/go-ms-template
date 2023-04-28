package config

import (
    {{.authImport}}
    "goms/common/auth"
)

type Config struct {
	rest.RestConf
}
