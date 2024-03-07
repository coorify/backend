package model

type Perm struct {
	Model

	Value string `json:"value" gorm:"size:64;index;unique;comment:权限值"`
	Group string `json:"group" gorm:"size:32;comment:权限组"`
	Desc  string `json:"desc" gorm:"size:32;comment:权限描述"`
}

func (Perm) TableName() string {
	return "sys_perms"
}
