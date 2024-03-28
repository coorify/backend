package middle

import (
	"net/http"

	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/perm"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
)

func WithPerm(pm *perm.Perm) gin.HandlerFunc {
	if perm.FindPerm(&pm.PermValue) == nil {
		perm.AddPerm(pm)
	}

	return func(c *gin.Context) {
		cas := c.MustGet(field.SYS_JWTCLAMIS).(*jwt.Clamis)
		pc := c.MustGet(field.SYS_PERMCACHE).(perm.PermCache)

		if !pc.Has(cas.UUID, &pm.PermValue, c) {
			c.Abort()
			reply.NewReply(http.StatusUnauthorized, nil, "invalid permission", c)
		}
	}
}
