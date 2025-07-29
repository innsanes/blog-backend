package library

import (
	"sync/atomic"
	"time"
)

const (
	TokenBucketMinCapacity     = 1
	TokenBucketDefaultCapacity = 1024
	TokenBucketMinRate         = 1
	TokenBucketDefaultRate     = 30
	TokenBucketMinRetries      = 20
	TokenBucketMaxRetries      = 100
	TokenBucketDefaultRetries  = 20
)

// TokenBucket 令牌桶限流器
type TokenBucket struct {
	state    int64
	capacity int32 // 令牌桶容量
	rate     int32 // 令牌产生速率（个/秒）
	retries  int32 // 重试次数
}

// NewTokenBucket 创建新的令牌桶限流器
func NewTokenBucket(config TokenBucketConfig) *TokenBucket {
	tb := &TokenBucket{
		capacity: max(TokenBucketMinCapacity, config.Capacity),
		rate:     max(TokenBucketMinRate, config.Rate),
		retries:  min(TokenBucketMaxRetries, max(TokenBucketMinRetries, config.Retries)),
	}

	// 初始化状态：桶满，时间为当前时间
	tb.state = tb.pack(time.Now(), config.Capacity)

	return tb
}

type TokenBucketConfig struct {
	Capacity int32
	Rate     int32
	Retries  int32
}

func (t *TokenBucket) pack(refillTime time.Time, currentTokens int32) int64 {
	return (int64(currentTokens) << 32) | refillTime.Unix()
}

func (t *TokenBucket) unpack(state int64) (refillTime time.Time, currentTokens int32) {
	currentTokens = int32(state >> 32)
	refillTime = time.Unix(state&0xFFFFFFFF, 0)
	return
}

func (t *TokenBucket) Allow(n int32) bool {
	for retries := 0; retries < int(t.retries); retries++ {
		state := atomic.LoadInt64(&t.state)
		lastRefill, currentTokens := t.unpack(state)

		nowSec := time.Now().Unix()
		timePassed := nowSec - lastRefill.Unix()
		tokensToAdd := int32(timePassed) * t.rate
		availableTokens := min(t.capacity, currentTokens+tokensToAdd)

		if availableTokens < n {
			return false
		}

		newState := (int64(availableTokens-n) << 32) | nowSec
		if atomic.CompareAndSwapInt64(&t.state, state, newState) {
			return true
		}
	}

	return false
}

func (t *TokenBucket) Wait(n int32, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		if t.Allow(n) {
			return true
		}
		now := time.Now()
		nextSec := now.Truncate(time.Second).Add(time.Second)
		sleepDuration := nextSec.Sub(now)
		remaining := deadline.Sub(now)
		if sleepDuration > remaining {
			sleepDuration = remaining
		}
		time.Sleep(sleepDuration)
	}
	return false
}
