package dialector

import (
	"errors"
	"fmt"

	"github.com/coorify/backend/logger"
	"github.com/coorify/backend/option"
	cfg "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mysqlDialector(opt *option.DatabaseOption) gorm.Dialector {
	if opt.Port == 0 {
		opt.Port = 3306
	}

	if opt.Host == "" || opt.Name == "" || opt.Username == "" || opt.Password == "" {
		panic(errors.New("mysql option error"))
	}

	logger.Infof("gorm remote: %s:%d", opt.Host, opt.Port)

	cf := mysql.Config{
		DSNConfig: &cfg.Config{
			User:   opt.Username,
			Passwd: opt.Password,
			DBName: opt.Name,
			Net:    "tcp",
			Addr:   fmt.Sprintf("%s:%d", opt.Host, opt.Port),
			Params: map[string]string{
				"charset":   "utf8mb4",
				"parseTime": "true",
				"loc":       "Local",
			},
			AllowNativePasswords: true,
		},
		DefaultStringSize:         256,
		SkipInitializeWithVersion: false,
	}

	return mysql.New(cf)
}
