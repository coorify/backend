package dialector

import (
	"github.com/coorify/backend/logger"
	"github.com/coorify/backend/option"
	"gorm.io/gorm"
)

func NewDialector(opt *option.DatabaseOption) gorm.Dialector {
	drv := opt.Driver

	logger.Infof("database dialector: %s", drv)

	switch drv {
	default:
		return mysqlDialector(opt)
	}
}
