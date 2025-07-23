package core

import (
	"blog-backend/service/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
)

type Vodka struct {
	*serv.Service
	Engine *gin.Engine
	config *VodkaConfig
	logger *VodkaLog
	router []func(*gin.RouterGroup)
	name   string
}

type VodkaConfig struct {
	Addr     string `conf:"addr,default=0.0.0.0:8000,usage=gin_serve_ip"`
	BasePath string `conf:"base_path,default=,usage=gin_serve_base_path"`
	CORS     bool   `conf:"cors,default=false,usage=gin_serve_cors"`
	GinMode  string `conf:"mode,default=release,usage=gin_mode(debug/release/test)"`
}

func NewVodka(name string) *Vodka {
	return &Vodka{
		config: &VodkaConfig{},
		logger: NewVodkaLog(name),
		name:   name,
	}
}

func (s *Vodka) RegisterRouter(router func(*gin.RouterGroup)) {
	s.router = append(s.router, router)
}

func (s *Vodka) BeforeServe() (err error) {
	conf.RegisterConfWithName(s.name, s.config)
	return
}

func (s *Vodka) Serve() (err error) {
	gin.SetMode(s.config.GinMode)
	s.Engine = gin.New()
	s.Engine.Use(s.logger.Logger(), gin.Recovery())

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
