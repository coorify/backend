package model

import (
	"github.com/coorify/backend/crypto"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type Account struct {
	Model

	UUID     uuid.UUID `json:"uuid" gorm:"index;unique;comment:用户UUID"`
	Username string    `json:"username" gorm:"size:16;index;unique;comment:用户账号"`
	Password string    `json:"-" gorm:"size:64;comment:用户密码"`
	Nickname string    `json:"nickname" gorm:"size:16;comment:用户昵称"`
	Phone    string    `json:"phone" gorm:"size:16;comment:用户手机号"`
	Status   int       `json:"status" gorm:"default:1;comment:用户状态"`
	Roles    []Role    `json:"roles" gorm:"many2many:sys_account_roles"`
}

func (Account) TableName() string {
	return "sys_accounts"
}

func (a *Account) BeforeCreate(db *gorm.DB) (err error) {
	var rls *[]Role
	if err = db.Model(&Role{}).Where("status&2 != 0").Find(&rls).Error; err != nil {
		return err
	}

	a.UUID = uuid.Must(uuid.NewV4())
	a.Password = crypto.EncodePassword(a.Username, a.Password)
	a.Roles = *rls
	return nil
}

func (a *Account) Verify(pswd string) bool {
	if a.Password == "" {
		return false
	}

	return crypto.EncodePassword(a.Username, pswd) == a.Password
}
