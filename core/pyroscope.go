package core

import (
	"fmt"
	"github.com/grafana/pyroscope-go"
	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
)

type Pyroscope struct {
	*serv.Service
	config *PyroscopeConfig
}

type PyroscopeConfig struct {
	Host string `conf:"server,default=localhost:4040,usage=host"`
}

func NewPyroscope() *Pyroscope {
	return &Pyroscope{
		config: &PyroscopeConfig{},
	}
}

func (s *Pyroscope) BeforeServe() (err error) {
	conf.RegisterConfWithName("pyroscope", s.config)
	return
}

func (s *Pyroscope) Serve() (err error) {
	_, err = pyroscope.Start(pyroscope.Config{
		ApplicationName: "blog-backend",
		ServerAddress:   fmt.Sprintf("http://%s", s.config.Host),
		Tags:            map[string]string{},
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
		},
	})
	return
}
