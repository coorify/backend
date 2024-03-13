package middle

import (
	"net/http"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
)

func Jwt(c *gin.Context) {
	name := "token"
	tk := c.GetHeader(name)
	if len(tk) == 0 {
		tk = c.Query(name)
	}
	if len(tk) == 0 {
		tk, _ = c.Cookie(name)
	}

	jwt := c.MustGet(field.SYS_JWT).(jwt.JwtEncoding)
	cas, ok := jwt.Decode(tk)
	if !ok {
		c.Abort()
		reply.NewReply(http.StatusUnauthorized, nil, "令牌过期", c)
	}

	c.Set(field.SYS_JWTCLAMIS, cas)
}
