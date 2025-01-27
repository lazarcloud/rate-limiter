package limiter

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
)

func (l *Limiter) generateFingerprint(r *http.Request) string {
	// Uses IP, User-Agent, and X-Forwarded-For
	clientIP := r.RemoteAddr
	userAgent := r.UserAgent()
	forwardedFor := r.Header.Get("X-Forwarded-For")

	data := strings.Join([]string{clientIP, userAgent, forwardedFor}, "|")

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
