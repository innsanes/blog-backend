package core

import (
	"os"
	"path/filepath"

	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
)

type Image struct {
	*serv.Service
	Config *ImageConfig
}

type ImageConfig struct {
	Path      string `conf:"path,default=./images"`
	Quality   int    `conf:"quality,default=90,usage=image_quality"`
	MaxWidth  int    `conf:"max_width,default=1920,usage=image_max_width"`
	MaxHeight int    `conf:"max_height,default=1080,usage=image_max_height"`
}

func NewImage() *Image {
	return &Image{
		Config: &ImageConfig{},
	}
}

func (s *Image) BeforeServe() (err error) {
	conf.RegisterConfWithName("image", s.Config)
	return
}

func (s *Image) Serve() (err error) {
	// 创建目录
	if _, err := os.Stat(s.Config.Path); os.IsNotExist(err) {
		err = os.MkdirAll(s.Config.Path, 0755)
		if err != nil {
			return err
		}
	}
	// 创建origin和compressed目录
	if _, err := os.Stat(filepath.Join(s.Config.Path, "origin")); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Join(s.Config.Path, "origin"), 0755)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(filepath.Join(s.Config.Path, "compressed")); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Join(s.Config.Path, "compressed"), 0755)
		if err != nil {
			return err
		}
	}
	return
}
