package global

import (
	"blog-backend/core"
)

var (
	Log       = core.NewLog()
	Config    = core.NewConfig()
	MySQL     = core.NewMOrm()
	Pyroscope = core.NewPyroscope()
	Image     = core.NewImage()
)
