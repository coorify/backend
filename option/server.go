package option

type ServerOption struct {
	Port int `default:"3080"`
	Host string
}
