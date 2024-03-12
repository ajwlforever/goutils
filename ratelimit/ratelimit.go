package ratelimit

type LimitOption func()

type Limiter interface {
	TryAcquire() LimitResult
}

type LimitResult struct {
	ok bool
}
