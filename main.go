package main

import (
	"blog-backend/global"
	"blog-backend/router"
	"github.com/innsanes/serv"
)

func main() {
	router.RegisterRouter()
	serv.Serve(
		global.Log,
		global.Config,
		global.Gin,
		global.MySQL,
	)
}
