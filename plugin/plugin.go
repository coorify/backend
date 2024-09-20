package plugin

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Engin() *gin.Engine
	Option() interface{}

	Set(key string, value interface{})
}

func Setup(s Server) {
	e := s.Engin()

	e.Use(gin.Recovery())

	e.Use(Option(s))
	e.Use(Logger(s))
	e.Use(Redis(s))
	e.Use(DB(s))
	e.Use(Cors(s))
	e.Use(Signature(s))
	e.Use(Jwt(s))
	e.Use(Perm(s))
}
