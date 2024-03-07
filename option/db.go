package option

type DatabaseOption struct {
	Driver   string `default:"mysql"`
	Port     int    `default:"3306"`
	Host     string `required:"true"`
	Name     string `required:"true"`
	Username string `required:"true"`
	Password string `required:"true"`
}
