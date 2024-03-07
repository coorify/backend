package option

type AdminOption struct {
	ID       uint   `default:"1"`
	Nickname string `default:"管理员"`
	Username string `default:"admin"`
	Password string `required:"true"`
}
