package plugin

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Engin() *gin.Engine
	Option() interface{}
}

func Setup(s Server) {
	e := s.Engin()
	o := s.Option()

	e.Use(gin.Recovery())
	e.Use(Option(o))
	e.Use(Logger(o))
	e.Use(redisPlugin(o))
	e.Use(dbPlugin(o))
	e.Use(Cors(o))
	e.Use(Signature(o))
	e.Use(Jwt(o))
	e.Use(Perm(o))
}
