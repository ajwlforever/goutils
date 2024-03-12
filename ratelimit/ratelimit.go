package ratelimit

import "time"

type LimiterOption func() Limiter

type Limiter interface {
	TryAcquire() LimitResult
}

type LimitResult struct {
	Ok bool
}

func WithFixedWindowLimiter(unitTime time.Duration, count int) LimiterOption {
	return func() Limiter {
		limiter := NewFixedWindowLimiter(unitTime, count)
		return limiter
	}
}

func NewLimiter(option LimiterOption) Limiter {
	return option()
}

// Record 限流情况全部记录下来。
// todo
type Record struct {
}
