package perm

import "strings"

type PermValue struct {
	Value string `json:"value"`
}

type Perm struct {
	PermValue

	Group   string `json:"group"`
	Desc    string `json:"desc"`
	IsAdmin bool   `json:"-"`
}

var perms []*Perm

func init() {
	perms = make([]*Perm, 0)
}

func AllPerm(full bool) []*Perm {
	if full {
		return perms
	}

	pms := make([]*Perm, 0)
	for _, p := range perms {
		if !p.IsAdmin {
			pms = append(pms, p)
		}
	}

	return pms
}

func AddPerm(pm *Perm) {
	perms = append(perms, pm)
}

func FindPerm(pv *PermValue) *Perm {
	for _, p := range perms {
		if p.Value == pv.Value {
			return p
		}
	}
	return nil
}

func GetGroups(prefix string) []string {
	gps := make([]string, 0)

	for _, gp := range perms {
		if strings.HasPrefix(gp.Group, prefix) {
			gps = append(gps, gp.Group)
		}
	}

	return gps
}

func NewPerm(grp, desc, val string) *Perm {
	return &Perm{
		Group:   grp,
		Desc:    desc,
		IsAdmin: false,
		PermValue: PermValue{
			Value: val,
		},
	}
}

func NewAdminPerm(grp, desc, val string) *Perm {
	return &Perm{
		Group:   grp,
		Desc:    desc,
		IsAdmin: true,
		PermValue: PermValue{
			Value: val,
		},
	}
}

func NewPermValue(val string) *Perm {
	return &Perm{
		Group:   "",
		Desc:    "",
		IsAdmin: false,
		PermValue: PermValue{
			Value: val,
		},
	}
}
