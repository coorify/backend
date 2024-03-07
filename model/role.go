package model

import (
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type Role struct {
	Model

	UUID   uuid.UUID `json:"uuid" gorm:"index;unique;comment:角色UUID"`
	Name   string    `json:"name" gorm:"size:16;unique;comment:角色名称"`
	Desc   string    `json:"desc" gorm:"size:256;comment:角色描述"`
	Status int       `json:"status" gorm:"default:1;comment:角色状态"`

	Perms []RolePerm `json:"perms" gorm:"foreignKey:RoleId;references:UUID"`
}

func (Role) TableName() string {
	return "sys_roles"
}

func (a *Role) BeforeCreate(tx *gorm.DB) (err error) {
	a.UUID = uuid.Must(uuid.NewV4())
	return nil
}
