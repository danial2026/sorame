package middleware

import (
	"log"
	"log/slog"
	"net/http"
	"time"
)

// wrappedWriter wraps ResponseWriter to capture the status code
type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

// Overrides WriteHeader to capture the status code
func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// Logging middleware to log request details
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Capture request start time

		// Wrap ResponseWriter to track the status code
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		if r.URL.Path != "/status" {
			// Log only non-identifying information
			log.Println(
				"➡️",
				"request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Any("time", time.Now().Format(time.RFC3339)),
			)
		}

		next.ServeHTTP(wrapped, r) // Proceed to the next handler

		if r.URL.Path != "/status" {
			// Log only minimal response information
			log.Println(
				"📤",
				"response",
				slog.Int("statusCode", wrapped.statusCode),
				slog.Any("duration", time.Since(start)),
			)
		}
	})
}
