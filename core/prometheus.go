package core

import (
	"github.com/innsanes/serv"
)

type Prometheus struct {
	*serv.Service
}

type PrometheusConfig struct{}

func NewPrometheus() *Prometheus {
	return &Prometheus{}
}
