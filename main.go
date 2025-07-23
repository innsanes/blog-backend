package main

import (
	"blog-backend/global"
	"blog-backend/router"
	"github.com/innsanes/serv"
)

func main() {
	g.BlogServer.RegisterRouter(router.RegisterBlog)
	g.InternalServer.RegisterRouter(router.RegisterBlogAuth)

	serv.Serve(
		g.Config,
		g.MySQL,
		g.BlogServer,
		g.InternalServer,
	)
}
