package main

import (
	"blog-backend/global"
	"github.com/innsanes/serv"
)

func main() {
	serv.Serve(
		global.Log,
		global.Config,
		global.Gin,
	)
}
