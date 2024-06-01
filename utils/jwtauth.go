package utils

import (
	"time"

	"github.com/bhumong/dbo-test/config"
	"github.com/bhumong/dbo-test/pkg/auth"
	"github.com/golang-jwt/jwt/v5"
)

var jwtAuthInstance *auth.AuthJwt

func Init() {
	cfg := config.GetConfig()
	expired := time.Second * time.Duration(cfg.GetInt("jwt.expired"))

	jwtAuthInstance = auth.New(cfg.GetString("jwt.secret"), expired, jwt.SigningMethodHS256, nil)
	if jwtAuthInstance == nil {
		panic("cannot create new instance of AuthJwt")
	}
}

func GetJwt() *auth.AuthJwt {
	return jwtAuthInstance
}
