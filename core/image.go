package core

import (
	"os"
	"path/filepath"

	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
)

type Image struct {
	*serv.Service
	Path    string
	Quality int
	config  *ImageConfig
}

type ImageConfig struct {
	Path    string `conf:"path,default=./images"`
	Quality int    `conf:"quality,default=75,usage=image_quality"`
}

func NewImage() *Image {
	return &Image{
		config: &ImageConfig{},
	}
}

func (s *Image) BeforeServe() (err error) {
	conf.RegisterConf(s.config)
	return
}

func (s *Image) Serve() (err error) {
	s.Path = s.config.Path
	s.Quality = s.config.Quality
	// 创建目录
	if _, err := os.Stat(s.Path); os.IsNotExist(err) {
		err = os.MkdirAll(s.Path, 0755)
		if err != nil {
			return err
		}
	}
	// 创建origin和compressed目录
	if _, err := os.Stat(filepath.Join(s.Path, "origin")); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Join(s.Path, "origin"), 0755)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(filepath.Join(s.Path, "compressed")); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Join(s.Path, "compressed"), 0755)
		if err != nil {
			return err
		}
	}
	return
}
