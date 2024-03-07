package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/option"
	"github.com/gin-gonic/gin"
)

func Jwt(opt *option.Option) gin.HandlerFunc {
	o := &opt.Jwt
	jwt := jwt.NewJwt([]byte(o.Secret), o.Expire)

	return func(c *gin.Context) {
		c.Set(field.SYS_JWT, jwt)
	}
}
