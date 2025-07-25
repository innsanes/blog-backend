package main

import (
	"blog-backend/core"
	"blog-backend/global"
	"blog-backend/router"

	"github.com/innsanes/serv"
)

func main() {
	var blogServer = core.NewVodka("blog")
	blogServer.RegisterRouter(router.RegisterBlog)
	var internalServer = core.NewVodka("internal")
	internalServer.RegisterRouter(router.RegisterBlogAuth)
	internalServer.RegisterRouter(router.RegisterEcho)
	internalServer.RegisterRouter(router.RegisterPrometheus)
	internalServer.RegisterRouter(router.RegisterPProf)

	serv.Serve(
		global.Config,
		global.MySQL,
		blogServer,
		internalServer,
	)
}
