package option

type DatabaseOption struct {
	Driver string `default:"mysql"`

	// mysql
	Port     int    `default:"0"`
	Host     string `default:""`
	Name     string `default:""`
	Username string `default:""`
	Password string `default:""`

	// sqlite
	DSN string `default:""`
}
