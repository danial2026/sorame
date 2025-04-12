package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp  string `json:"timestamp"`
	Level      string `json:"level"`
	Message    string `json:"message"`
	Method     string `json:"method,omitempty"`
	Path       string `json:"path,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
	Duration   string `json:"duration,omitempty"`
}

// JsonLogging middleware to log request and response details in JSON format
func JsonLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip logging for /status endpoint
		if r.URL.Path == "/status" {
			next.ServeHTTP(w, r)
			return
		}

		start := time.Now() // Capture request start time

		// Wrap ResponseWriter to track the status code
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// Log the request details - only non-identifying information
		requestLogEntry := LogEntry{
			Timestamp: time.Now().Format(time.RFC3339),
			Level:     "info",
			Message:   "Incoming request",
			Method:    r.Method,
			Path:      r.URL.Path,
		}

		jsonRequestLog, err := json.Marshal(requestLogEntry)
		if err == nil {
			// Save the request log to a file
			file, fileErr := os.OpenFile("logs/sorame_request_logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if fileErr != nil {
				log.Printf("Error opening log file: %v", fileErr)
			} else {
				defer file.Close()
				_, writeErr := file.Write(jsonRequestLog)
				if writeErr != nil {
					log.Printf("Error writing log to file: %v", writeErr)
				}
				file.WriteString("\n") // Add a newline for readability
			}
		} else {
			log.Printf("Error marshaling request log to JSON: %v", err)
		}

		next.ServeHTTP(wrapped, r) // Proceed to the next handler

		// Log the response details - minimal information
		responseLogEntry := LogEntry{
			Timestamp:  time.Now().Format(time.RFC3339),
			Level:      "info",
			Message:    "Outgoing response",
			StatusCode: wrapped.statusCode,
			Duration:   time.Since(start).String(),
		}
		jsonResponseLog, err := json.Marshal(responseLogEntry)
		if err == nil {
			// Save the response log to a file
			file, fileErr := os.OpenFile("logs/sorame_response_logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if fileErr != nil {
				log.Printf("Error opening response log file: %v", fileErr)
			} else {
				defer file.Close()
				_, writeErr := file.Write(jsonResponseLog)
				if writeErr != nil {
					log.Printf("Error writing response log to file: %v", writeErr)
				}
				file.WriteString("\n") // Add a newline for readability
			}
		} else {
			log.Printf("Error marshaling response log to JSON: %v", err)
		}
	})
}
