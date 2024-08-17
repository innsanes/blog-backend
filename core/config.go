package core

import (
	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
)

type Conf struct {
	*serv.Service
	*conf.X
}

func NewConfig(bfs ...conf.BuildFunc) *Conf {
	return &Conf{
		X: conf.New(bfs...),
	}
}

func (s *Conf) BeforeServe() (err error) {
	flag := conf.NewFlag(s.X)
	s.X.RegisterSource(flag)
	yaml := conf.NewYaml(s.X)
	s.X.RegisterSource(yaml)
	s.X.RegisterConfWithName("yaml", yaml.YamlConf)
	return
}

func (s *Conf) Serve() (err error) {
	s.X.Parse()
	return
}

func (s *Conf) AfterServe() (err error) {
	s.X.PrintResult()
	return
}
