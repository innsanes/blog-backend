package service

import (
	"blog-backend/structs/req"
)

var Echo IEcho = &EchoService{}

type EchoService struct{}

type IEcho interface {
	Echo(in *req.Echo) (string, error)
}

func (s *EchoService) Echo(in *req.Echo) (msg string, err error) {
	return in.Message, nil
}
