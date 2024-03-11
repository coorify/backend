package account

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/jwt"
	"github.com/coorify/backend/reply"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	cas := c.MustGet(field.SYS_JWTCLAMIS).(*jwt.Clamis)
	act, err := cas.ToAccount(c)

	if err != nil {
		reply.FailWithMessage(err.Error(), c)
		return
	}

	reply.OkWithPayload(act, c)
}
