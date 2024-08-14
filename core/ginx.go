package core

import (
	"blog-backend/library/ginx"
	"blog-backend/router"
	"github.com/gin-gonic/gin"
	"github.com/innsanes/serv"
	"strconv"
)

type Gin struct {
	*serv.Service
	*ginx.Gin
	conf   Confer
	config *GinConfig
}

type GinConfig struct {
	IP   string `conf:"ip,default=0.0.0.0"`
	Port int    `conf:"port,default=8000"`
}

func NewGin(conf Confer, bfs ...ginx.BuildFunc) *Gin {
	return &Gin{
		Gin:    ginx.New(bfs...),
		conf:   conf,
		config: &GinConfig{},
	}
}

func (s *Gin) BeforeServe() (err error) {
	s.conf.RegisterConfWithName("gin", s.config)
	return
}

func (s *Gin) Serve() (err error) {
	s.Engine.Use(gin.Recovery())
	if s.GetLogger() == nil {
		s.Engine.Use(gin.Logger())
	} else {
		s.Engine.Use(s.GetLogger())
	}
	router.RegisterRouter(s.Engine)
	err = s.Engine.Run(s.config.IP + ":" + strconv.Itoa(s.config.Port))
	return
}
