package core

import (
	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
)

type Conf struct {
	*serv.Service
}

func NewConfig() *Conf {
	return &Conf{}
}

func (s *Conf) BeforeServe() (err error) {
	flag := conf.NewFlag(conf.GetConf())
	conf.RegisterSource(flag)
	yaml := conf.NewYaml(conf.GetConf())
	conf.RegisterSource(yaml)
	conf.RegisterConfWithName("yaml", yaml.YamlConf)
	return
}

func (s *Conf) Serve() (err error) {
	conf.Parse()
	return
}
