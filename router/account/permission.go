package account

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/perm"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
)

func AccountPermission(c *gin.Context) {
	cas := c.MustGet(field.SYS_JWTCLAMIS).(*jwt.Clamis)
	pc := c.MustGet(field.SYS_PERMCACHE).(perm.PermCache)

	pvs := pc.Get(cas.UUID, c)
	rep := make([]string, 0)
	for _, p := range pvs {
		rep = append(rep, p.Value)
	}

	reply.OkWithPayload(rep, c)
}
