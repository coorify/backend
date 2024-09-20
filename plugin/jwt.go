package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/go-value"
	"github.com/gin-gonic/gin"
)

func Jwt(s Server) gin.HandlerFunc {
	o := s.Option()

	secret := value.MustGet(o, "Jwt.Secret").(string)
	expire := value.MustGet(o, "Jwt.Expire").(int)

	jwt := jwt.NewJwt([]byte(secret), expire)

	return func(c *gin.Context) {
		c.Set(field.SYS_JWT, jwt)
	}
}
