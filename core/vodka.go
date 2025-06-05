package core

import (
	"blog-backend/handler/middleware"
	"blog-backend/library/vodka"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/innsanes/serv"
)

type Vodka struct {
	*serv.Service
	*vodka.Vodka
	conf   Confer
	config *VodkaConfig
	router []func(*gin.RouterGroup)
	bfs    []vodka.BuildFunc
}

type VodkaConfig struct {
	Addr     string `conf:"addr,default=0.0.0.0:8000,usage=gin_serve_ip"`
	BasePath string `conf:"base_path,default=,usage=gin_serve_base_path"`
	CORS     bool   `conf:"cors,default=false,usage=gin_serve_cors"`
	GinMode  string `conf:"mode,default=release,usage=gin_mode(debug,release,test)"`
}

func NewVodka(conf Confer, bfs ...vodka.BuildFunc) *Vodka {
	return &Vodka{
		conf:   conf,
		config: &VodkaConfig{},
		bfs:    bfs,
	}
}

func (s *Vodka) RegisterRouter(router func(*gin.RouterGroup)) {
	s.router = append(s.router, router)
}

func (s *Vodka) BeforeServe() (err error) {
	s.conf.RegisterConfWithName("gin", s.config)
	return
}

func (s *Vodka) Serve() (err error) {
	gin.SetMode(s.config.GinMode)
	s.Vodka = vodka.New(s.bfs...)
	s.Engine.Use(s.Logger(), gin.Recovery())

	if s.config.CORS {
		s.Engine.Use(middleware.CORS())
	}

	routerGroup := s.Engine.Group(s.config.BasePath)
	for _, r := range s.router {
		r(routerGroup)
	}

	go func() {
		var errRun error
		if errRun = s.Engine.Run(s.config.Addr); errRun != nil {
			panic(fmt.Sprintf("gin run error: %v", errRun))
		}
	}()
	return
}

func VodkaLogger(log *Logger) vodka.LogHandler {
	return func(param gin.LogFormatterParams) {
		log.Info("%s %s %s %d %s %s",
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}
}
