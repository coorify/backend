package dialector

import (
	"errors"

	"github.com/coorify/backend/logger"
	"github.com/coorify/backend/option"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func sqliteDialector(opt *option.DatabaseOption) gorm.Dialector {
	if opt.DSN == "" {
		panic(errors.New("sqlite option error"))
	}

	logger.Infof("gorm remote: %s", opt.DSN)

	return sqlite.Open(opt.DSN)
}
