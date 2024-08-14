package global

import (
	"blog-backend/core"
	"blog-backend/library/ginx"
	"github.com/innsanes/conf"
)

var (
	Log    = core.NewLog()
	Config = core.NewConfig(
		conf.WithResultLogger(Log),
	)
	Gin = core.NewGin(
		Config,
		ginx.WithLogger(ginx.DefaultLogger(Log)),
	)
)
