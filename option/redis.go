package option

type RedisOption struct {
	Enable   bool   `default:"false"`
	Host     string `default:"localhost"`
	Password string `default:""`
	DB       int    `default:"0"`
	Port     int    `default:"6379"`
}
