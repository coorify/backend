package plugin

import (
	"github.com/coorify/backend/dialector"
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/option"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Account{}, &model.Perm{}, &model.Role{}, &model.RolePerm{}, &model.Menu{})
}

func initAdmin(db *gorm.DB, adm *option.AdminOption) error {
	nt := int64(0)
	md := &model.Account{}

	if err := db.Model(md).Count(&nt).Error; err != nil {
		return err
	}

	if nt > 0 {
		return nil
	}

	return db.Model(md).Create(&model.Account{
		Model: model.Model{
			ID: adm.ID,
		},
		Username: adm.Username,
		Password: adm.Password,
		Nickname: adm.Nickname,
	}).Error
}

func Database(opt *option.Option) gin.HandlerFunc {
	drv := dialector.NewDialector(&opt.DB)

	db, err := gorm.Open(drv, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	if err = initMigrate(db); err != nil {
		panic(err)
	}

	if err = initAdmin(db, &opt.Admin); err != nil {
		panic(err)
	}

	return func(ctx *gin.Context) {
		ctx.Set(field.SYS_DB, db)
	}
}
