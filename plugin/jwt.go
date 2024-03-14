package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/option"
	"github.com/coorify/go-value"
	"github.com/gin-gonic/gin"
)

func Jwt(opt interface{}) gin.HandlerFunc {
	o := value.MustGet(opt, "Jwt").(option.JwtOption)
	jwt := jwt.NewJwt([]byte(o.Secret), o.Expire)

	return func(c *gin.Context) {
		c.Set(field.SYS_JWT, jwt)
	}
}
