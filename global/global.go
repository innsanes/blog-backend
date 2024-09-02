package global

import (
	"blog-backend/core"
	"blog-backend/library/martini"
)

var (
	Log    = core.NewLog()
	Config = core.NewConfig()
	Gin    = core.NewMartini(
		Config,
		martini.WithLogHandler(core.MartiniLogger(Log)),
	)
	MySQL = core.NewMOrm(Config)
	Token = core.NewTokenPool()
)
