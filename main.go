package main

import (
	"blog-backend/global"
	"blog-backend/global/service"
	"blog-backend/router"
	"github.com/innsanes/serv"
)

func main() {
	service.BlogServer.RegisterRouter(router.RegisterBlog)
	service.InternalServer.RegisterRouter(router.RegisterBlogAuth)

	serv.Serve(
		global.Log,
		global.Config,
		service.BlogServer,
		service.InternalServer,
		global.MySQL,
	)
}
