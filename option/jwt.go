package option

type JwtOption struct {
	Secret string `default:"jwt_secret"`
	Expire int    `default:"7200"`
}
