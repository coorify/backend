package dialector

import (
	"github.com/coorify/backend/logger"
	"github.com/coorify/backend/option"
	"gorm.io/gorm"
)

func NewDialector(opt *option.DatabaseOption) gorm.Dialector {
	drv := opt.Driver

	logger.Infof("gorm dialector: %s", drv)

	switch drv {
	case "sqlite":
		return sqliteDialector(opt)
	default:
		return mysqlDialector(opt)
	}
}
