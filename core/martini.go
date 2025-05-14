package core

import (
	"blog-backend/library/martini"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/innsanes/serv"
	"net/http"
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
	s.Engine.Use(s.Logger(), gin.Recovery(), func(context *gin.Context) {
		method := context.Request.Method
		// 1. [必须]接受指定域的请求，可以使用*不加以限制，但不安全
		//context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
		fmt.Println(context.GetHeader("Origin"))
		// 2. [必须]设置服务器支持的所有跨域请求的方法
		context.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		// 3. [可选]服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
		context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, X-Token")
		// 4. [可选]设置XMLHttpRequest的响应对象能拿到的额外字段
		context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, X-Token")
		// 5. [可选]是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
		context.Header("Access-Control-Allow-Credentials", "true")
		// 6. 放行所有OPTIONS方法
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
			return
		}
		context.Next()
	})

	// Register router
	routerGroup := s.Engine.Group("/api")
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
