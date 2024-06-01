package middleware

import (
	"strings"
	"time"

	"github.com/bhumong/dbo-test/model"
	"github.com/bhumong/dbo-test/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")

		splitter := strings.Split(authorization, " ")
		if len(splitter) < 2 {
			ctx.Error(model.ErrAuthJwt)
			ctx.Abort()
			return
		}
		tokenString := splitter[1]

		jwtAuth := utils.GetJwt()
		claims, err := jwtAuth.Validate(tokenString)
		if err != nil {
			log.Debug().AnErr("jwtAuth.Validate", err).Send()
			ctx.Error(model.ErrAuthJwt.SetCause(err))
			ctx.Abort()
			return
		}

		ctx.Set("jwt.claims", claims)
	}
}
