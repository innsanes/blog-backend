package library_test

import (
	"blog-backend/library"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// newTestBucket 创建测试用的令牌桶
func newTestBucket(capacity, rate int32) *library.TokenBucket {
	return library.NewTokenBucket(library.TokenBucketConfig{
		Capacity: capacity,
		Rate:     rate,
		Retries:  20,
	})
}

// TestTokenBucketBasic 基础功能测试
func TestTokenBucketBasic(t *testing.T) {
	tb := newTestBucket(10, 5) // 容量10，速率5个/秒

	// 测试初始状态
	for i := 0; i < 10; i++ {
		assert.True(t, tb.Allow(1), "初始状态应该允许10个请求，第%d个失败", i+1)
	}

	// 第11个应该被拒绝
	assert.False(t, tb.Allow(1), "桶已空，应该拒绝请求")

	// 等待1秒，应该有5个新token
	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		assert.True(t, tb.Allow(1), "1秒后应该有5个token，第%d个失败", i+1)
	}

	// 第6个应该被拒绝
	assert.False(t, tb.Allow(1), "1秒后只有5个token，第6个应该被拒绝")
}

// TestTokenBucketConcurrent 并发测试
func TestTokenBucketConcurrent(t *testing.T) {
	tb := newTestBucket(100, 10) // 容量100，速率10个/秒
	var wg sync.WaitGroup
	successCount := int32(0)
	failCount := int32(0)
	var mu sync.Mutex

	// 启动100个goroutine并发请求
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if tb.Allow(1) {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else {
				mu.Lock()
				failCount++
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	// 验证结果
	assert.Equal(t, int32(100), successCount, "期望成功100次")
	assert.Equal(t, int32(0), failCount, "期望失败0次")
}

// TestTokenBucketHighLoad 高负载测试
func TestTokenBucketHighLoad(t *testing.T) {
	tb := newTestBucket(50, 5) // 容量50，速率5个/秒
	var wg sync.WaitGroup
	successCount := int32(0)
	failCount := int32(0)
	var mu sync.Mutex

	// 启动200个goroutine，超过桶容量
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if tb.Allow(1) {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else {
				mu.Lock()
				failCount++
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	// 验证结果：应该成功50次，失败150次
	assert.Equal(t, int32(50), successCount, "期望成功50次")
	assert.Equal(t, int32(150), failCount, "期望失败150次")
}

// TestTokenBucketContinuous 持续请求测试
func TestTokenBucketContinuous(t *testing.T) {
	tb := newTestBucket(20, 10) // 容量20，速率10个/秒
	var wg sync.WaitGroup
	successCount := int32(0)
	failCount := int32(0)
	var mu sync.Mutex

	// 持续请求3秒
	duration := 3 * time.Second
	start := time.Now()

	for time.Since(start) < duration {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tb.Allow(1) {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else {
				mu.Lock()
				failCount++
				mu.Unlock()
			}
		}()
		time.Sleep(time.Millisecond * 10) // 每10ms发起一个请求
	}

	wg.Wait()

	// 验证结果：3秒内应该成功约20+30=50次（初始20 + 3秒*10/秒）
	expectedSuccess := int32(50)
	tolerance := int32(10) // 允许10次误差
	assert.GreaterOrEqual(t, successCount, expectedSuccess-tolerance,
		"成功次数应该大于等于%d", expectedSuccess-tolerance)
	assert.LessOrEqual(t, successCount, expectedSuccess+tolerance,
		"成功次数应该小于等于%d", expectedSuccess+tolerance)

	t.Logf("持续请求测试结果：成功%d次，失败%d次", successCount, failCount)
}

// TestTokenBucketBurst 突发流量测试
func TestTokenBucketBurst(t *testing.T) {
	tb := newTestBucket(30, 5) // 容量30，速率5个/秒
	var wg sync.WaitGroup
	successCount := int32(0)
	failCount := int32(0)
	var mu sync.Mutex

	// 第一波：30个并发请求（应该全部成功）
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tb.Allow(1) {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else {
				mu.Lock()
				failCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// 验证第一波结果
	assert.Equal(t, int32(30), successCount, "第一波期望成功30次")

	// 第二波：立即发起20个请求（应该全部失败）
	successCount = 0
	failCount = 0
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if tb.Allow(1) {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else {
				mu.Lock()
				failCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// 验证第二波结果
	assert.Equal(t, int32(0), successCount, "第二波期望成功0次")
	assert.Equal(t, int32(20), failCount, "第二波期望失败20次")
}

// TestTokenBucketMultiToken 多token请求测试
func TestTokenBucketMultiToken(t *testing.T) {
	tb := newTestBucket(100, 10) // 容量100，速率10个/秒

	// 测试请求多个token
	assert.True(t, tb.Allow(50), "应该允许请求50个token")
	assert.True(t, tb.Allow(30), "应该允许请求30个token")
	assert.True(t, tb.Allow(20), "应该允许请求20个token")

	// 请求超过剩余token数量
	assert.False(t, tb.Allow(1), "应该拒绝超过剩余token的请求")
}

// TestTokenBucketConcurrentMultiToken 并发多token测试
func TestTokenBucketConcurrentMultiToken(t *testing.T) {
	tb := newTestBucket(200, 20) // 容量200，速率20个/秒
	var wg sync.WaitGroup
	successCount := int32(0)
	failCount := int32(0)
	var mu sync.Mutex

	// 并发请求不同数量的token
	tokenRequests := []int32{10, 20, 30, 40, 50}
	for _, tokens := range tokenRequests {
		wg.Add(1)
		go func(reqTokens int32) {
			defer wg.Done()
			if tb.Allow(reqTokens) {
				mu.Lock()
				successCount++
				mu.Unlock()
			} else {
				mu.Lock()
				failCount++
				mu.Unlock()
			}
		}(tokens)
	}

	wg.Wait()

	// 验证结果
	assert.Greater(t, successCount, int32(0), "应该至少有一些成功的请求")
	t.Logf("并发多token测试结果：成功%d次，失败%d次", successCount, failCount)
}

// TestTokenBucketWait 测试Wait方法
func TestTokenBucketWait(t *testing.T) {
	tb := library.NewTokenBucket(library.TokenBucketConfig{
		Capacity: 10,
		Rate:     5,
		Retries:  20,
	})

	// 消耗所有token
	for i := 0; i < 10; i++ {
		tb.Allow(1)
	}

	// 测试Wait方法
	start := time.Now()
	result := tb.Wait(1, 3*time.Second)
	duration := time.Since(start)

	assert.True(t, result, "Wait方法应该成功获取token")

	// 验证等待时间（由于优化了sleep间隔，等待时间会更短）
	assert.GreaterOrEqual(t, duration, time.Millisecond*100,
		"等待时间应该大于等于100ms，实际：%v", duration)
	assert.LessOrEqual(t, duration, 3*time.Second,
		"等待时间应该小于等于3秒，实际：%v", duration)

	t.Logf("Wait方法测试：等待时间%v", duration)
}

// TestTokenBucketEdgeCases 边界情况测试
func TestTokenBucketEdgeCases(t *testing.T) {
	tb := newTestBucket(10, 5)

	// 测试请求0个token
	assert.True(t, tb.Allow(0), "请求0个token应该总是成功")

	// 测试请求负数token
	assert.True(t, tb.Allow(-1), "请求负数token应该总是成功")

	// 测试请求超过容量的token
	assert.False(t, tb.Allow(11), "请求超过容量的token应该被拒绝")

	// 测试请求等于容量的token
	assert.True(t, tb.Allow(10), "请求等于容量的token应该成功")
	assert.False(t, tb.Allow(1), "桶已空，应该拒绝后续请求")
}

// TestTokenBucketRateCalculation 速率计算测试
func TestTokenBucketRateCalculation(t *testing.T) {
	tb := newTestBucket(10, 10) // 容量10，速率10个/秒

	// 消耗所有token
	for i := 0; i < 10; i++ {
		tb.Allow(1)
	}

	// 等待1秒，应该有10个新token
	time.Sleep(time.Second)
	successCount := int32(0)
	for i := 0; i < 15; i++ {
		if tb.Allow(1) {
			successCount++
		}
	}

	// 应该成功约10次（允许±2的误差）
	assert.GreaterOrEqual(t, successCount, int32(8),
		"1秒后应该至少有8个token，实际：%d", successCount)
	assert.LessOrEqual(t, successCount, int32(12),
		"1秒后应该最多有12个token，实际：%d", successCount)
}

// BenchmarkTokenBucket 性能基准测试
func BenchmarkTokenBucket(b *testing.B) {
	tb := newTestBucket(1000, 100)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tb.Allow(1)
		}
	})
}

// BenchmarkTokenBucketConcurrent 并发性能基准测试
func BenchmarkTokenBucketConcurrent(b *testing.B) {
	tb := newTestBucket(1000, 100)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tb.Allow(1)
		}
	})
}
