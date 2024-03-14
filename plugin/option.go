package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/gin-gonic/gin"
)

func Option(opt interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(field.SYS_OPTION, opt)
	}
}
