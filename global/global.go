package g

import (
	"blog-backend/core"
)

var (
	Log    = core.NewLog()
	Config = core.NewConfig()
	MySQL  = core.NewMOrm()
)
