package limiter

import "net/http"

// Middleware is a https middleware that limits the endpoint
func (l *Limiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientFingerprint := l.generateFingerprint(r)

		if l.isRateLimited(clientFingerprint) {
			http.Error(w, "Too Many Requests. Please try again later.", http.StatusTooManyRequests)
			return
		}

		l.recordRequest(clientFingerprint)
		next.ServeHTTP(w, r)
	})
}
