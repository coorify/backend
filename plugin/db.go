package plugin

import (
	"github.com/coorify/backend/dialector"
	"github.com/coorify/backend/field"
	"github.com/coorify/backend/model"
	"github.com/coorify/backend/option"
	"github.com/coorify/go-value"
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

var _db *gorm.DB

func GetDB() *gorm.DB {
	return _db
}

func dbPlugin(opt interface{}) gin.HandlerFunc {
	var err error

	o := value.MustGet(opt, "DB").(option.DatabaseOption)
	a := value.MustGet(opt, "Admin").(option.AdminOption)

	drv := dialector.NewDialector(&o)

	_db, err = gorm.Open(drv, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	if err = initMigrate(_db); err != nil {
		panic(err)
	}

	if err = initAdmin(_db, &a); err != nil {
		panic(err)
	}

	return func(ctx *gin.Context) {
		ctx.Set(field.SYS_DB, _db)
	}
}
