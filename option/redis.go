package option

import "fmt"

type RedisOption struct {
	Enable   bool   `default:"false"`
	Host     string `default:"localhost"`
	Password string `default:""`
	DB       int    `default:"0"`
	Port     int    `default:"6379"`
}

func (r RedisOption) Addr() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
