package model

import "github.com/gofrs/uuid/v5"

type RolePerm struct {
	Model

	RoleId uuid.UUID `json:"uuid" gorm:"index;comment:角色UUID"`
	Perm   string    `json:"perm" gorm:"size:64;comment:权限值"`
}

func (RolePerm) TableName() string {
	return "sys_role_perms"
}
