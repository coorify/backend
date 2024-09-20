package router

import (
	"github.com/coorify/backend/middle"
	"github.com/coorify/backend/perm"
	"github.com/coorify/backend/router/account"
	"github.com/coorify/backend/router/admin"
	"github.com/coorify/backend/router/auth"
	"github.com/coorify/backend/router/oauth"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Engin() *gin.Engine
	Group(relativePath string) *gin.RouterGroup
}

func Setup(s Server) {
	s.Engin().GET("version", Version)

	{
		o := s.Engin().Group("/system/oauth")
		o.GET("profileByJwt", oauth.ProfileByJwt)
	}

	sys := s.Group("system")
	{
		a := sys.Group("auth")
		a.POST("sigin", auth.Sigin)
	}
	{
		a := sys.Group("account", middle.Jwt)
		a.GET("profile", account.Profile)
		a.GET("permission", account.AccountPermission)
		a.GET("menu", account.AccountMenu)

		p := a.Group("password")
		p.PUT("update", account.PasswordUpdate)
	}
	{
		adm := sys.Group("admin", middle.Jwt, middle.WithPerm(perm.NewAdminPerm("站点管理员", "超级权限", "admin_all")))

		perm := adm.Group("permission")
		perm.POST("create", admin.PermissionCreate)
		perm.GET("group", admin.PermissionGroup)
		perm.GET("find", admin.PermissionFind)
		perm.DELETE("delete", admin.PermissionDelete)
		perm.GET("system", admin.PermissionSystem)

		role := adm.Group("role")
		role.POST("create", admin.RoleCreate)
		role.GET("find", admin.RoleFind)
		role.DELETE("delete", admin.RoleDelete)
		role.PUT("status/update", admin.RoleStatusUpdate)
		role.GET("permissions", admin.RolePermissions)
		role.PUT("permission/update", admin.RolePermissionUpdate)

		act := adm.Group("account")
		act.POST("create", admin.AccountCreate)
		act.GET("find", admin.AccountFind)
		act.PUT("status/update", admin.AccountStatusUpdate)
		act.GET("roles", admin.AccountRoles)
		act.PUT("role/update", admin.AccountRoleUpdate)

		menu := adm.Group("menu")
		menu.POST("create", admin.MenuCreate)
		menu.GET("find", admin.MenuFind)
		menu.DELETE("delete", admin.MenuDelete)
		menu.PUT("status/update", admin.MenuStatusUpdate)
	}

}
