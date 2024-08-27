package core

import (
	"blog-backend/library/martini"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/innsanes/serv"
	"strconv"
)

type Martini struct {
	*serv.Service
	*martini.M
	conf   Confer
	config *GinConfig
	router []func(*gin.RouterGroup)
	bfs    []martini.BuildFunc
}

type GinConfig struct {
	IP   string `conf:"ip,default=0.0.0.0,usage=gin_serve_ip"`
	Port int    `conf:"port,default=8000,usage=gin_serve_port"`
}

func NewMartini(conf Confer, bfs ...martini.BuildFunc) *Martini {
	return &Martini{
		conf:   conf,
		config: &GinConfig{},
		bfs:    bfs,
	}
}

func (s *Martini) RegisterRouter(router func(*gin.RouterGroup)) {
	s.router = append(s.router, router)
}

func (s *Martini) BeforeServe() (err error) {
	s.conf.RegisterConfWithName("gin", s.config)
	return
}

func (s *Martini) Serve() (err error) {
	gin.SetMode(gin.ReleaseMode)
	s.M = martini.New(s.bfs...)
	s.Engine.Use(s.Logger(), gin.Recovery())

	// Register router
	routerGroup := s.Engine.Group("")
	for _, r := range s.router {
		r(routerGroup)
	}

	go func() {
		var errRun error
		if errRun = s.Engine.Run(s.config.IP + ":" + strconv.Itoa(s.config.Port)); errRun != nil {
			panic(fmt.Sprintf("gin run error: %v", errRun))
		}
	}()
	return
}

func MartiniLogger(log *Logger) martini.LogHandler {
	return func(param gin.LogFormatterParams) {
		log.Info("%s %s %s %d %s %s\n",
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}
}
