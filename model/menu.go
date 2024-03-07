package model

import (
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type Menu struct {
	Model

	UUID   uuid.UUID `json:"uuid" gorm:"index;unique;comment:UUID"`
	Link   string    `json:"link" gorm:"size:256;comment:链接"`
	Perm   string    `json:"perm" gorm:"size:64;comment:权限值"`
	Name   string    `json:"name" gorm:"size:16;comment:名称"`
	Group  string    `json:"group" gorm:"size:32;comment:分组"`
	Pos    string    `json:"pos" gorm:"size:16;comment:位置"`
	Icon   string    `json:"icon" gorm:"size:32;comment:图标"`
	Auth   int       `json:"auth" gorm:"default:0;comment:认证"`
	Status int       `json:"status" gorm:"default:1;comment:状态"`
}

func (Menu) TableName() string {
	return "sys_menus"
}

func (m *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	m.UUID = uuid.Must(uuid.NewV4())
	return nil
}
