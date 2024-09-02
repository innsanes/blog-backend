package core

import "sync"

// 暂时使用 sync.Map 后面改为 redis

type TokenPool struct {
	tokens sync.Map
}

func NewTokenPool() *TokenPool {
	return &TokenPool{
		tokens: sync.Map{},
	}
}

func (tp *TokenPool) AddToken(token string) {
	tp.tokens.Store(token, struct{}{})
}

func (tp *TokenPool) RemoveToken(token string) {
	tp.tokens.Delete(token)
}

func (tp *TokenPool) CheckToken(token string) bool {
	_, ok := tp.tokens.Load(token)
	return ok
}
