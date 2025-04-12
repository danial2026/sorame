package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
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

		// Read and log the request body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			return
		}
		// Restore the request body for further use
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		if strings.Contains(fmt.Sprintf("%v", r.Header), "uptimerobot") == false && r.URL.Path != "/status" {
			// Log the request details and response status
			log.Println(
				"➡️",
				"handled request",
				slog.Int("statusCode", wrapped.statusCode),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Any("duration", time.Since(start)),
				slog.String("xff", r.Header.Get("X-Forwarded-For")),
			)

			// Log the request details
			log.Println(
				"📥",
				"request details",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("headers", fmt.Sprintf("%v", r.Header)),
				slog.String("body", string(bodyBytes)),
				slog.String("queryParams", r.URL.RawQuery),
				slog.String("xff", r.Header.Get("X-Forwarded-For")),
			)

			// Log the response details
			log.Println(
				"📤",
				"response details",
				slog.Int("statusCode", wrapped.statusCode),
				slog.String("headers", fmt.Sprintf("%v", wrapped.Header())),
				slog.Any("duration", time.Since(start)),
				// Assuming the response body can be accessed via wrappedWriter
				slog.String("responseBody", "(response body tracking not implemented)"),
			)
		}
		next.ServeHTTP(wrapped, r) // Proceed to the next handler
	})
}
