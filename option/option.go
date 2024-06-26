package option

type Option struct {
	Logger    LoggerOption
	Server    ServerOption
	DB        DatabaseOption
	Admin     AdminOption
	Jwt       JwtOption
	Router    RouterOption
	Signature SignatureOption
	Redis     RedisOption
}
