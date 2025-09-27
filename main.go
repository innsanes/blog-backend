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
	blogServer.RegisterRouter(router.RegisterImage)
	var internalServer = core.NewVodka("internal")
	internalServer.RegisterRouter(router.RegisterBlogAuth)
	internalServer.RegisterRouter(router.RegisterImageAuth)
	internalServer.RegisterRouter(router.RegisterEcho)
	internalServer.RegisterRouter(router.RegisterPrometheus)
	internalServer.RegisterRouter(router.RegisterPProf)

	serv.Serve(
		global.Config,
		global.Image,
		global.MySQL,
		global.Pyroscope,
		global.Meilisearch,
		blogServer,
		internalServer,
	)
}
