package plugin

import (
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/option"
	"github.com/gin-gonic/gin"
)

func Option(opt *option.Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(field.SYS_OPTION, opt)
	}
}
