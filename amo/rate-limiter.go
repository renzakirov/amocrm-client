package amo

import "time"

// RateLimiter limits requests rate to AmoCRM API endpoint
type RateLimiter interface {
	// WaitForRequest waits while the client can made next request
	WaitForRequest()
}

var defaultRTLimiter = newDefaultRateLimiter()

type empty struct{}

type defaultRateLimiter struct {
	limiter chan empty
}

func newDefaultRateLimiter() *defaultRateLimiter {
	limiter := &defaultRateLimiter{}
	limiter.limiter = make(chan empty)
	go func() {
		limiter.limiter <- empty{}
		time.Sleep(time.Second)
	}()
	return limiter
}

func (rl *defaultRateLimiter) WaitForRequest() {
	<-rl.limiter
}
