package backend

import (
	"context"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/coorify/backend/logger"
	"github.com/coorify/backend/middle"
	"github.com/coorify/backend/plugin"
	"github.com/coorify/backend/router"
	"github.com/coorify/go-value"
	"github.com/gin-gonic/gin"
)

type Server struct {
	keys map[string]interface{}
	opt  interface{}
	eng  *gin.Engine
	svr  *http.Server
	exit chan error
	mu   sync.RWMutex
}

type SetupPlugin interface {
	Plugin(s *Server) error
}

type SetupRouter interface {
	Router(s *Server) error
}

func NewServer(opt interface{}) *Server {
	logger.SetLevel(value.MustGet(opt, "Logger.Level").(string))
	eng := gin.New()

	svr := &Server{
		keys: make(map[string]interface{}),
		eng:  eng,
		opt:  opt,
		exit: make(chan error),
	}

	return svr
}

func (s *Server) Engin() *gin.Engine {
	return s.eng
}

func (s *Server) Option() interface{} {
	return s.opt
}

func (s *Server) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.keys[key] = value
}

func (s *Server) Get(key string) (value interface{}, exists bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exists = s.keys[key]
	return
}

func (s *Server) MustGet(key string) interface{} {
	if value, exists := s.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

func (s *Server) Group(relativePath string) *gin.RouterGroup {
	prefix := value.MustGet(s.opt, "Router.Prefix").(string)
	if prefix != "" {
		return s.Engin().Group(prefix).Group(relativePath, middle.Signature)
	}
	return s.Engin().Group(relativePath, middle.Signature)
}

func (s *Server) Frontend(fe fs.FS) {
	e := s.Engin()

	// only in release mode
	if gin.Mode() == gin.ReleaseMode {
		e.StaticFS("/fe/", http.FS(fe))
		e.NoRoute(func(c *gin.Context) {
			if c.Request.RequestURI == "/" && c.Request.Method == "GET" {
				c.Redirect(http.StatusFound, "/fe")
			}
		})
	}
}

func (s *Server) Start() error {
	if s.svr != nil {
		return nil
	}

	plugin.Setup(s)
	if v, ok := s.opt.(SetupPlugin); ok {
		v.Plugin(s)
	}

	router.Setup(s)
	if v, ok := s.opt.(SetupRouter); ok {
		v.Router(s)
	}

	host := value.MustGet(s.opt, "Server.Host").(string)
	port := value.MustGet(s.opt, "Server.Port").(int)

	addr := fmt.Sprintf("%s:%d", host, port)
	s.svr = &http.Server{
		Handler: s.eng,
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		err := s.svr.Serve(ln)
		if err != nil {
			logger.Info("Server closed")
			err = nil
		}

		s.exit <- err
	}()

	logger.Infof("Listen on %s", addr)
	return err
}

func (s *Server) Stop(grace bool) error {
	if s.svr == nil {
		return nil
	}

	var err error
	if grace {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		err = s.svr.Shutdown(ctx)
	} else {
		err = s.svr.Close()
	}

	if err != nil {
		return err
	}

	return <-s.exit
}
