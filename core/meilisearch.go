package core

import (
	"blog-backend/services/search"

	"github.com/innsanes/conf"
	"github.com/innsanes/serv"
	"github.com/meilisearch/meilisearch-go"
)

type Meilisearch struct {
	*serv.Service
	config         *MeilisearchConfig
	ServiceManager meilisearch.ServiceManager
}

type MeilisearchConfig struct {
	Host string `conf:"server,default=localhost:7700,usage=host"`
}

func NewMeilisearch() *Meilisearch {
	return &Meilisearch{
		config: &MeilisearchConfig{},
	}
}

func (s *Meilisearch) BeforeServe() (err error) {
	conf.RegisterConfWithName("meilisearch", s.config)
	return
}

func (s *Meilisearch) Serve() (err error) {
	s.ServiceManager = meilisearch.New("http://" + s.config.Host)
	search.CreateBlogIndex(s.ServiceManager)
	return
}
