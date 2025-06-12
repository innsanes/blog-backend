package service

import (
	"blog-backend/core"
	"blog-backend/global"
	"blog-backend/library/vodka"
)

var (
	BlogServer = core.NewVodka(
		global.Config,
		"blog",
		vodka.WithLogHandler(core.VodkaLogger(global.Log)),
	)
	InternalServer = core.NewVodka(
		global.Config,
		"internal",
		vodka.WithLogHandler(core.VodkaLogger(global.Log)),
	)
)

func init() {

}
