package main

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware 记录每个请求的详细日志
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Wrap response writer to capture status code
		lw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next(lw, r)
		log.Printf("[%s] %s %s %d %s %s",
			RepoName, r.Method, r.URL.Path, lw.statusCode,
			time.Since(start), r.RemoteAddr,
		)
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lw *loggingResponseWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.ResponseWriter.WriteHeader(code)
}
