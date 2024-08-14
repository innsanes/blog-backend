package core

import (
	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
)

type Config struct {
	*serv.Service
	*conf.X
}

func NewConfig(bfs ...conf.BuildFunc) *Config {
	return &Config{
		X: conf.New(bfs...),
	}
}

func (s *Config) BeforeServe() (err error) {
	flag := conf.NewFlag(s.X)
	s.X.RegisterSource(flag)
	yaml := conf.NewYaml(s.X)
	s.X.RegisterSource(yaml)
	s.X.RegisterConfWithName("yaml", yaml.YamlConf)
	return
}

func (s *Config) Serve() (err error) {
	s.X.Parse()
	return
}

func (s *Config) AfterServe() (err error) {
	s.X.PrintResult()
	return
}
