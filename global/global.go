package global

import (
	"blog-backend/core"
	"blog-backend/library/vodka"
)

var (
	Log        = core.NewLog()
	Config     = core.NewConfig()
	BlogServer = core.NewVodka(
		Config,
		vodka.WithLogHandler(core.VodkaLogger(Log)),
	)
	MySQL = core.NewMOrm(Config)
	Token = core.NewTokenPool()
)
