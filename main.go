package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	limit      int
	window     time.Duration
	clientData map[string][]time.Time
	mu         sync.Mutex
}

func New(limit int, window time.Duration) *Limiter {
	return &Limiter{
		limit:      limit,
		window:     window,
		clientData: make(map[string][]time.Time),
	}
}

func (l *Limiter) isRateLimited(fingerprint string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	requests, exists := l.clientData[fingerprint]
	if !exists {
		return false
	}

	var validRequests []time.Time
	for _, t := range requests {
		if now.Sub(t) <= l.window {
			validRequests = append(validRequests, t)
		}
	}

	l.clientData[fingerprint] = validRequests

	return len(validRequests) >= l.limit
}

func (l *Limiter) recordRequest(fingerprint string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	l.clientData[fingerprint] = append(l.clientData[fingerprint], now)
}
